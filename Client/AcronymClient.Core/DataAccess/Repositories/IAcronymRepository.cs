using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using AcronymClient.Core.DataAccess.Entities;
using AcronymClient.Core.Utils;
using SQLite;

namespace AcronymClient.Core.DataAccess.Repository
{
	public interface IAcronymRepository
	{
		[Utils.NotNull]
		Task<IEnumerable<AcronymModel>> FindAllAsync([Utils.NotNull] string acronym);

		[Utils.NotNull]
		Task UpdateAllAsync([Utils.NotNull] IEnumerable<AcronymModel> acronyms);
	}
	
}
