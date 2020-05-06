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

        public static NumericRange CreateDefault(double value)
        {
            return new NumericRange(value - 10.0, value + 10.0);
        }

        public override string ToString()
        {
            return Min + " - " + Max;
        }
    }
}
