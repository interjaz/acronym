using System;
namespace AcronymClient.Core.Utils
{
	[AttributeUsage(AttributeTargets.Parameter | AttributeTargets.ReturnValue | AttributeTargets.Field | AttributeTargets.Method)]
	public class NotNullAttribute : Attribute { }

	[AttributeUsage(AttributeTargets.Parameter | AttributeTargets.ReturnValue | AttributeTargets.Field | AttributeTargets.Method)]
	public class CanBeNull : Attribute { }
}


