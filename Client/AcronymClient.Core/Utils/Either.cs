using System;
using System.Collections;
using System.Collections.Generic;

namespace AcronymClient.Core.Utils
{
	public class Either<T> 
	{
		public ErrorMessage Error { get; set;}

		public T Some { get; set;}

		public Either()
		{
		}

		public Either(T some)
		{
			Some = some;
		}

		public Either([NotNull] ErrorMessage error)
		{
			Check.NotNull(error, nameof(error));

			Error = error;
		}

		public Either(ErrorCode errorCode)
		{
			Error = new ErrorMessage
			{
				ErrorCode = errorCode
			};
		}

	}
}

