using System;
using SQLite;

namespace AcronymClient.Core
{
	[Table("Acronyms")]
	public class AcronymModel
	{
		public long AcronymId { get; set; }
		public string Acronym { get; set; }
		public string Language { get; set; }
		public string Definition { get; set; }
		public string Url { get; set; }
		public DateTime CreatedAt { get; set; }
		public DateTime ModifiedAt { get; set; }
	}
}

