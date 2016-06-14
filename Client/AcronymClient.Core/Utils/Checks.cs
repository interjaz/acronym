using System;
namespace AcronymClient.Core.Utils
{
	public static class Check
	{
		public static void NotNull<T>(T obj, string name)
			where T : class
		{
			if (obj == null)
			{
				throw new ArgumentNullException(name);
			}
		}
	}
}

