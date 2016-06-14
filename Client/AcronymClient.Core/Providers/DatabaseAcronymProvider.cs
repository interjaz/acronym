using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using AcronymClient.Core.DataAccess.Entities;
using AcronymClient.Core.Utils;
using System.Linq;
using AcronymClient.Core.DataAccess.Repository;

namespace AcronymClient.Core.Providers
{
	using EitherAcronyms = Either<IEnumerable<AcronymModel>>;

	public class DatabaseAcronymProvider : IAcronymProvider
	{
		readonly IAcronymRepository _acronymRespository;

		public DatabaseAcronymProvider(
			[Utils.NotNull] IAcronymRepository acronymRepository)
		{
			Check.NotNull(acronymRepository, nameof(acronymRepository));

			_acronymRespository = acronymRepository;
		}

		public async Task<EitherAcronyms> FindAsync([Utils.NotNull] string acronym)
		{
			Check.NotNull(acronym, nameof(acronym));


			try
			{
				var result = await _acronymRespository.FindAllAsync(acronym);

				var success = new EitherAcronyms(result.ToList());
				return success;
			}
			catch (Exception ex)
			{
				var errorMessage = new ErrorMessage()
				{
					ErrorCode = ErrorCode.DatabaseReadError,
					Tag = ex
				};

				var error = new EitherAcronyms(errorMessage);
				return error;
			}
		}
	}
}

