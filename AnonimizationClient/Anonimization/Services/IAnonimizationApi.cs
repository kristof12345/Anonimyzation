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

        [Get("/classes/matching")]
        public Task<List<EqulivalenceClass>> GetMatchingEqulivalenceClasses([Body] Document document);

        [Post("/classes")]
        public Task<EqulivalenceClass> CreateEqulivalenceClass([Body] EqulivalenceClass equlivalenceClass);

        [Put("/classes/{dataset}/{id}")]
        public Task<bool> RegisterUploadIntent(string dataset, int id);

        [Put("/classes/{dataset}/{id}")]
        public Task<bool> UploadDocument(string dataset, int id, [Body] Document document);

        [Get("/central/{id}")]
        public Task<CentralTableItem> CheckCentralTable(int id);
    }
}