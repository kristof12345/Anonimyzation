using Anonimization.Models;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using Document = System.Collections.Generic.Dictionary<string, object>;

namespace Anonimization.Extensions
{
    public static class ConsoleExtensions
    {
        public static void Write(this IEnumerable<EqulivalenceClass> list)
        {
            Console.WriteLine(list.Count() + " matching equlivalence classes were found:");

            foreach (var equlivalenceClass in list)
            {
                equlivalenceClass.Write();
            }
        }

        public static void Write(this EqulivalenceClass equlivalenceClass)
        {
            string output = JsonConvert.SerializeObject(equlivalenceClass);
            Console.WriteLine(output);
        }

        public static string ToMyString(this Document document)
        {
            string attributes = "";

            foreach (var key in document.Keys)
            {
                attributes = attributes + key + " : " + document[key] + "; ";
            }

            return attributes;
        }
    }
}
