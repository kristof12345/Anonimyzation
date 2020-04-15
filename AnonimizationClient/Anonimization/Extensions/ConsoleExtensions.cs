using Anonimization.Models;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;

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
    }
}
