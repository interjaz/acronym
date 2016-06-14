using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Threading.Tasks;
using AcronymClient.Core.DataAccess.Entities;
using AcronymClient.Core.Utils;
using ModernHttpClient;
using Newtonsoft.Json;

namespace AcronymClient.Core.Providers
{
	using EitherAcronyms = Either<IEnumerable<AcronymModel>>;

	public class RestfulAcronymProvider : IAcronymProvider
	{
		private static readonly HttpClient _httpClient;
		private static readonly JsonConverter _jsonConverter;

		readonly string _restfulAcronymEndpoint;

		static RestfulAcronymProvider() 
		{
			_httpClient = new HttpClient(new NativeMessageHandler());
			_jsonConverter = new UnixDateTimeConverter();
		}

		public RestfulAcronymProvider(
			[NotNull] string restfulAcronymEndpoint
			)
		{
			Check.NotNull(restfulAcronymEndpoint, nameof(restfulAcronymEndpoint));

			restfulAcronymEndpoint = restfulAcronymEndpoint.EndsWith("/", StringComparison.OrdinalIgnoreCase) 
                                       ? restfulAcronymEndpoint.Substring(0, restfulAcronymEndpoint.Length - 1)
                                       : restfulAcronymEndpoint;
			
			_restfulAcronymEndpoint = restfulAcronymEndpoint;
		}

		public async Task<EitherAcronyms> FindAsync(string acronym)
		{
			Check.NotNull(acronym, nameof(acronym));

			var requestUri = string.Format("{0}/{1}", _restfulAcronymEndpoint, acronym);
			var response = await _httpClient.GetAsync(requestUri);

			if (!response.IsSuccessStatusCode)
			{
				var errorMessage = new ErrorMessage()
				{
					ErrorCode = ErrorCode.RestfulWrongStatusCode,
					Tag = response.StatusCode
				};
				var error = new EitherAcronyms(errorMessage);
				return error;
			}

			var responesBody = await response.Content.ReadAsStringAsync();
			try
			{
				var acronyms = JsonConvert.DeserializeObject<List<AcronymModel>>(responesBody, _jsonConverter);

				var success = new EitherAcronyms(acronyms);
				return success;
			}
			catch (Exception ex)
			{
				var errorMessage = new ErrorMessage()
				{
					ErrorCode = ErrorCode.RestfulDeserializeError,
					Tag = ex
				};

				var error = new EitherAcronyms(errorMessage);
				return error;
			}
		}

		private class UnixDateTimeConverter : JsonConverter
		{
			public override bool CanConvert(Type objectType)
			{
				return objectType == typeof(DateTime);
			}

			public override object ReadJson(JsonReader reader, Type objectType, object existingValue, JsonSerializer serializer)
			{
				var t = Convert.ToInt64(reader.Value);
				return new DateTime(1970, 1, 1).AddMilliseconds(t);
			}

			public override void WriteJson(JsonWriter writer, object value, JsonSerializer serializer)
			{
				throw new NotImplementedException();
			}
		}
	}
}