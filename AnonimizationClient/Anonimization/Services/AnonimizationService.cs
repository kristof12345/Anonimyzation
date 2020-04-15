using Anonimization.Extensions;
using Anonimization.Models;
using System;
using System.Threading;
using System.Threading.Tasks;
using Document = System.Collections.Generic.Dictionary<string, object>;

namespace Anonimization.Services
{
    public class AnonimizationService
    {
        private IAnonimizationApi Api { get; set; }

        public AnonimizationService(IAnonimizationApi api)
        {
            Api = api;
        }

        public async Task AnonimizeDocument(object obj)
        {
            var document = ((Request)obj).Document;
            var dataset = ((Request)obj).Dataset;

            var matchingClasses = await Api.GetMatchingEqulivalenceClasses(document);
            //matchingClasses.Write();

            foreach (var equlivalenceClass in matchingClasses)
            {
                // If a matching equlivalence class was found, we signal our upload intent
                if (IsMatching(equlivalenceClass, document))
                {
                    await Api.RegisterUploadIntent(dataset.Name, equlivalenceClass.Id);
                    await CheckCentralTable(dataset.Name, equlivalenceClass.Id, document);
                    return;
                }
            }

            // If no matching equlivalence class was found we create a new one.
            var result = await CreateEqulivalenceClass(dataset, document);
            await Api.RegisterUploadIntent(dataset.Name, result.Id);
            await CheckCentralTable(dataset.Name, result.Id, document);
        }

        private async Task CheckCentralTable(string dataset, int id, Document document)
        {
            //Periodically checking central table
            while (true)
            {
                var response = await Api.CheckCentralTable(id);
                if (response != null)
                {
                    await Api.UploadDocument(dataset, id, document);
                    Console.WriteLine("Uploaded: " + document["private"]);
                    return;
                }
                Thread.Sleep(1000);
            }
        }

        private async Task<EqulivalenceClass> CreateEqulivalenceClass(Dataset dataset, Document document)
        {
            var newClass = new EqulivalenceClass();

            foreach (var categoricField in dataset.GetCategoricFields())
            {
                newClass.CategoricAttributes.Add(categoricField.Name, document[categoricField.Name] as string);
            }

            foreach (var intervalField in dataset.GetIntervalFields())
            {
                newClass.IntervalAttributes.Add(intervalField.Name, NumericRange.Create(document[intervalField.Name] as double?));
            }

            var result = await Api.CreateEqulivalenceClass(newClass);

            //Check if the equlivalence class was created
            await Api.GetMatchingEqulivalenceClasses(document);
            return result;
        }

        private bool IsMatching(EqulivalenceClass equlivalenceClass, Document document)
        {
            foreach (var categoricField in equlivalenceClass.CategoricAttributes)
            {
                var value = document[categoricField.Key] as string;
                if (value != categoricField.Value) return false;
            }

            foreach (var intervalField in equlivalenceClass.IntervalAttributes)
            {
                var value = document[intervalField.Key] as double?;
                if (!intervalField.Value.Contains(value)) return false;
            }

            return true;
        }
    }
}
