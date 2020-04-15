using System.Collections.Generic;
using System.Linq;

namespace Anonimization.Models
{
    public class Dataset
    {
        public string Name { get; set; }
        public DatasetSettings Settings {get; set;}
        public List<DatasetField> Fields { get; set; }

        public IEnumerable<DatasetField> GetCategoricFields()
        {
            return Fields.Where(f => f.Mode == "cat");
        }

        public IEnumerable<DatasetField> GetIntervalFields()
        {
            return Fields.Where(f => f.Mode == "int");
        }
    }
}
