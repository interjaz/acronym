using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using ModernHttpClient;
using System.Net.Http;
using Newtonsoft.Json;

namespace AcronymClient.Core
{
	using EitherAcronyms = Either<IEnumerable<AcronymModel>>;

	public interface IAcronymProvider
	{
		[NotNull]
		Task<EitherAcronyms> FindAsync([NotNull] string acronym);
	}
}