using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using AcronymClient.Core.DataAccess.Entities;
using AcronymClient.Core.Utils;

namespace AcronymClient.Core.Providers
{
	using AcronymClient.Core.DataAccess.Repository;
	using EitherAcronyms = Either<IEnumerable<AcronymModel>>;

	public class CachedProvider : IAcronymProvider
	{
		readonly RestfulAcronymProvider _restfulProvider;
		readonly DatabaseAcronymProvider _databaseProvider;
		readonly IAcronymRepository _acronymRepository;


		public CachedProvider(
			[NotNull] RestfulAcronymProvider resftulProvider,
			[NotNull] DatabaseAcronymProvider databaseProivder,
			[NotNull] IAcronymRepository acronymRepository
		)
		{
			Check.NotNull(resftulProvider, nameof(resftulProvider));
			Check.NotNull(databaseProivder, nameof(databaseProivder));
			Check.NotNull(acronymRepository, nameof(acronymRepository));

			_restfulProvider = resftulProvider;
			_databaseProvider = databaseProivder;
			_acronymRepository = acronymRepository;
		}

		public async Task<EitherAcronyms> FindAsync([NotNull] string acronym)
		{
			Check.NotNull(acronym, nameof(acronym));

			var restfulResult = await _restfulProvider.FindAsync(acronym);
			if (restfulResult.Error != null)
			{
				var databaseResult = await _databaseProvider.FindAsync(acronym);
				return databaseResult;
			}

			await _acronymRepository.UpdateAllAsync(restfulResult.Some);
			return restfulResult;
		}
	}
}

