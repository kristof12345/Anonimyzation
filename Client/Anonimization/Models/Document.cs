using System.Collections.Generic;

namespace Anonimization.Models
{
    public class Document
    {
        public string Id { get; set; }
        public Dictionary<string, object> PublicFields { get; set; } = new Dictionary<string, object>();

        public Dictionary<string, object> PrivateFields { get; set; } = new Dictionary<string, object>();
    }
}
