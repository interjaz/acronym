using System;
namespace AcronymClient.Core.Utils
{
	public enum ErrorCode
	{
		RestfulWrongStatusCode,
		RestfulDeserializeError,
		DatabaseReadError,
		RestfulNetworkFailure
	}
}

