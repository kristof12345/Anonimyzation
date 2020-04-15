namespace Anonimization.Models
{
    public class DatasetSettings
    {
        public int K { get; set; }
        public int E { get; set; }
        public int Max { get; set; }
        public string Algorithm { get; set; } = "client-side";
        public string Mode { get; set; } = "continuous";
    }
}