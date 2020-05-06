using Anonimization.Extensions;
using Anonimization.Models;
using Anonimization.Services;
using Refit;
using System;
using System.Collections.Generic;
using System.Threading;
using System.Threading.Tasks;

namespace AnonimizationClient
{
    public class Program
    {
        public static async Task Main(string[] args)
        {
            Console.WriteLine("Starting anonimization.");

            var dataset = await CreateDataset();

            var a = AnonimizeDocument(dataset, 1);
            Thread.Sleep(1000);
            var b = AnonimizeDocument(dataset, 2);
            Thread.Sleep(1000);
            var c = AnonimizeDocument(dataset, 3);
            Thread.Sleep(1000);
            var d = AnonimizeDocument(dataset, 4);
            Thread.Sleep(1000);

            await a;
            await b;
            await c;
            await d;

            await WriteResults(dataset);

            await WriteResultsWithClasses(dataset);
        }

        private static async Task WriteResults(Dataset dataset)
        {
            var anonimizationApi = RestService.For<IAnonimizationApi>("http://localhost:9137/v1");

            var service = new AnonimizationService(anonimizationApi);

            var result = await service.GetAnonimDocumentsWithClassIds(dataset);

            foreach (var document in result)
            {
                Console.WriteLine(document.ToMyString());
            }
        }

        private static async Task WriteResultsWithClasses(Dataset dataset)
        {
            var anonimizationApi = RestService.For<IAnonimizationApi>("http://localhost:9137/v1");

            var service = new AnonimizationService(anonimizationApi);

            var result = await service.GetAnonimDocumentsWithClasses(dataset);

            foreach (var (equlivalenceClass, document) in result)
            {
                Console.WriteLine(equlivalenceClass + " " + document.ToMyString());
            }
        }

        private static async Task AnonimizeDocument(Dataset dataset, int i)
        {
            var anonimizationApi = RestService.For<IAnonimizationApi>("http://localhost:9137/v1");

            var service = new AnonimizationService(anonimizationApi);

            var document = new Document
            {
                Id = i.ToString(),
                PublicFields = new Dictionary<string, object>
            {
                { "city", "Budapest" },
                { "age", 20.0 },
            },
                PrivateFields = new Dictionary<string, object>
            {
                { "private", "secret" + i }
            }
            };

            await service.AnonimizeDocument(dataset.Name, document);
        }

        private static async Task<Dataset> CreateDataset()
        {
            var anonimizationApi = RestService.For<IAnonimizationApi>("http://localhost:9137/v1");

            var dataset = new Dataset
            {
                Settings = new DatasetSettings { E = 0, K = 2, Max = 2 },
                Fields = new List<DatasetField>
                {
                    new DatasetField { Name = "age", Mode = "int", Type = "numeric", PreferedSize = 5.0 },
                    new DatasetField { Name = "city", Mode = "cat", Type = "string" },
                    new DatasetField { Name = "private", Mode = "keep", Type = "string" }
                }
            };

            var id = Guid.NewGuid().ToString();

            return await anonimizationApi.CreateDataset(id, dataset);
        }
    }
}