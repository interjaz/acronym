using System;
using AcronymClient.Core.Providers;
using Xamarin.Forms;
using Xamarin.Forms.Xaml;

[assembly: XamlCompilation(XamlCompilationOptions.Compile)]
namespace AcronymClient.Core
{
	public class App : Application
	{
		public App()
		{
			var acronymProvider = new RestfulAcronymProvider("http://192.168.0.10:8080/api/v1/Acronym");

			var firstPage = new FirstPage(acronymProvider);
			MainPage = new NavigationPage(firstPage);
		}
	}
}

