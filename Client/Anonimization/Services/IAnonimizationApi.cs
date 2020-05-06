using Anonimization.Models;
using Refit;
using System.Collections.Generic;
using System.Threading.Tasks;
using Document = System.Collections.Generic.Dictionary<string, object>;

namespace Anonimization.Services
{
    public interface IAnonimizationApi
    {
        [Put("/datasets/{id}")]
        public Task<Dataset> CreateDataset(string id, [Body] Dataset dataset);

        [Get("/datasets/{name}")]
        public Task<Dataset> GetDatasetByName(string name);

        [Get("/classes/{id}")]
        public Task<EqulivalenceClass> GetEqulivalenceClassById(int id);

        [Get("/classes/matching")]
        public Task<List<EqulivalenceClass>> GetMatchingEqulivalenceClasses([Body] Document document);

        [Post("/classes")]
        public Task<EqulivalenceClass> CreateEqulivalenceClass([Body] EqulivalenceClass equlivalenceClass);

        [Put("/classes/{dataset}/{id}")]
        public Task<bool> RegisterUploadIntent(string dataset, int id);

        [Get("/central/{id}")]
        public Task<CentralTableItem> CheckCentralTable(int id);

        [Post("/upload")]
        public Task<UploadSessionResponse> StartUploadSession([Body] UploadSessionRequest request);

        [Post("/upload/{sessionId}/{classId}")]
        public Task<UploadResponse> UploadDocument(string sessionId, int classId, [Body] Document document);

        [Get("/data/{dataset}")]
        public Task<AnonimizationResults> GetDocuments(string dataset);
    }
}