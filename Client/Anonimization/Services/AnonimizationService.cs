using Anonimization.Models;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;

namespace Anonimization.Services
{
    public class AnonimizationService
    {
        private IAnonimizationApi Api { get; set; }
        static SemaphoreSlim semaphore = new SemaphoreSlim(1, 1);

        public AnonimizationService(IAnonimizationApi api)
        {
            Api = api;
        }

        public async Task<List<Dictionary<string, object>>> GetAnonimDocumentsWithClassIds(Dataset dataset)
        {
            var result = await Api.GetDocuments(dataset.Name);
            return result.Result;
        }

        public async Task<List<(EqulivalenceClass, Dictionary<string, object>)>> GetAnonimDocumentsWithClasses(Dataset dataset)
        {
            var result = new List<(EqulivalenceClass, Dictionary<string, object>)>();

            var documents = await GetAnonimDocumentsWithClassIds(dataset);
            foreach (var document in documents)
            {
                var equlivalenceClass = await Api.GetEqulivalenceClassById(int.Parse(document["classId"].ToString()));
                result.Add((equlivalenceClass, document));
            }

            return result;
        }

        public async Task AnonimizeDocument(string datasetName, Document document)
        {
            await semaphore.WaitAsync();
            var dataset = await Api.GetDatasetByName(datasetName);

            var matchingClasses = await Api.GetMatchingEqulivalenceClasses(document.PublicFields);
            //matchingClasses.Write();

            foreach (var equlivalenceClass in matchingClasses)
            {
                // If a matching equlivalence class was found, we signal our upload intent
                if (IsMatching(equlivalenceClass, document.PublicFields))
                {
                    await Api.RegisterUploadIntent(dataset.Name, equlivalenceClass.Id);
                    semaphore.Release();
                    await CheckCentralTable(dataset.Name, equlivalenceClass.Id, document.PrivateFields, document.Id);
                    return;
                }
            }

            // If no matching equlivalence class was found we create a new one.
            var result = await CreateEqulivalenceClass(dataset, document.PublicFields);
            await Api.RegisterUploadIntent(dataset.Name, result.Id);
            semaphore.Release();
            await CheckCentralTable(dataset.Name, result.Id, document.PrivateFields, document.Id);
        }

        private async Task CheckCentralTable(string dataset, int id, Dictionary<string, object> document, string documentId)
        {
            //Periodically checking central table
            while (true)
            {
                await semaphore.WaitAsync();
                var response = await Api.CheckCentralTable(id);
                if (response != null)
                {
                    Console.WriteLine("Uploading: " + documentId);
                    var request = new UploadSessionRequest
                    {
                        DatasetName = dataset
                    };

                    var session = await Api.StartUploadSession(request);
                    var uploadResponse = await Api.UploadDocument(session.SessionId, id, document);
                    Console.WriteLine("Finished: " + documentId + " " + uploadResponse.Error);
                    semaphore.Release();
                    return;
                }
                semaphore.Release();
                Thread.Sleep(1000);
            }
        }

        private async Task<EqulivalenceClass> CreateEqulivalenceClass(Dataset dataset, Dictionary<string, object> document)
        {
            var newClass = new EqulivalenceClass();

            foreach (var categoricField in dataset.GetCategoricFields())
            {
                newClass.CategoricAttributes.Add(categoricField.Name, document[categoricField.Name] as string);
            }



            foreach (var intervalField in dataset.GetIntervalFields())
            {
                NumericRange range;
                var exactValue = (double)document[intervalField.Name];
                var size = intervalField.PreferedSize;

                if (size != 0)
                {
                    range = new NumericRange(exactValue - size / 2, exactValue + size / 2);
                }
                else
                {
                    range = NumericRange.CreateDefault(exactValue);
                }

                newClass.IntervalAttributes.Add(intervalField.Name, range);
            }

            var result = await Api.CreateEqulivalenceClass(newClass);

            //Check if the equlivalence class was created
            await Api.GetMatchingEqulivalenceClasses(document);
            return result;
        }

        private bool IsMatching(EqulivalenceClass equlivalenceClass, Dictionary<string, object> document)
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
