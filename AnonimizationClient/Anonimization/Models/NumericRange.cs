using System;

namespace Anonimization.Models
{
    public class NumericRange
    {
        public double Min { get; set; }
        public double Max { get; set; }

        public NumericRange(double min, double max)
        {
            Min = min;
            Max = max;
        }

        public bool Contains(double? value)
        {
            if (value == null) return true;
            return value >= Min && value <= Max;
        }

        public static NumericRange Create(double? value)
        {
            if (value == null) throw new Exception();
            return new NumericRange(value.Value - 10.0, value.Value + 10.0); //TODO
        }
    }
}
