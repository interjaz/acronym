using System;
namespace AcronymClient.Core
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

