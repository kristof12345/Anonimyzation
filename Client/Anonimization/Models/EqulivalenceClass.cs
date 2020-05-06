using System.Collections.Generic;

namespace Anonimization.Models
{
    public class EqulivalenceClass
    {
        public int Id { get; set; }
        public Dictionary<string, string> CategoricAttributes { get; set; } = new Dictionary<string, string>();
        public Dictionary<string, NumericRange> IntervalAttributes { get; set; } = new Dictionary<string, NumericRange>();
        public int Count { get; set; }
        public int IntentCount { get; set; } = 1;
        public bool Active { get; set; } = true;

        public override string ToString()
        {
            string attributes = "";

            foreach (var attr in CategoricAttributes)
            {
                attributes = attributes + attr.Key + " : " + attr.Value + "; ";
            }

            foreach (var attr in IntervalAttributes)
            {
                attributes = attributes + attr.Key + " : " + attr.Value.ToString() + "; ";
            }

            return attributes;
        }
    }
}