namespace Anonimization.Models
{
    public class UploadResponse
    {
        public bool InsertSuccessful { get; set; }
        public bool FinalizeSuccessful { get; set; }

        public string Error { get; set; }

        public override string ToString()
        {
            return (InsertSuccessful && FinalizeSuccessful).ToString();
        }
    }
}