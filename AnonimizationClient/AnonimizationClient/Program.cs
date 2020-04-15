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
            Console.WriteLine("Hello World!");

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

            /*
            var thread1 = new Thread(service.AnonimizeDocument);
            thread1.Start(new Request { Dataset = dataset, Document = document1 });

            var thread2 = new Thread(service.AnonimizeDocument);
            thread2.Start(new Request { Dataset = dataset, Document = document2 });

            var thread3 = new Thread(service.AnonimizeDocument);
            thread3.Start(new Request { Dataset = dataset, Document = document3 });

            thread1.Join();
            thread2.Join();
            thread3.Join();
            */
        }

        private static async Task AnonimizeDocument(Dataset dataset, int i)
        {
            var anonimizationApi = RestService.For<IAnonimizationApi>("http://localhost:9137/v1");

            var service = new AnonimizationService(anonimizationApi);

            var document = new Dictionary<string, object>
            {
                { "city", "Budapest" },
                { "age", 20.0 },
                {"private", "secret" + i }
            };

            await service.AnonimizeDocument(new Request { Dataset = dataset, Document = document });
        }

        private static async Task<Dataset> CreateDataset()
        {
            var anonimizationApi = RestService.For<IAnonimizationApi>("http://localhost:9137/v1");

            var dataset = new Dataset
            {
                Settings = new DatasetSettings { E = 1, K = 2, Max = 4 },
                Fields = new List<DatasetField>
                {
                    new DatasetField { Name = "age", Mode = "int", Type = "numeric" },
                    new DatasetField { Name = "city", Mode = "cat", Type = "string" },
                    new DatasetField { Name = "private", Mode = "keep", Type = "string" }
                }
            };

            var id = Guid.NewGuid().ToString();

            return await anonimizationApi.CreateDataset(id, dataset);
        }
    }
}