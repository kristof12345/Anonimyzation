using Document = System.Collections.Generic.Dictionary<string, object>;

namespace Anonimization.Models
{
    public class Request
    {
        public Dataset Dataset { get; set; }
        public Document Document { get; set; }
    }
}
