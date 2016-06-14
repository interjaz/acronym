using System;
using System.IO;
using AcronymClient.Core.DataAccess.Repository;
using AcronymClient.Core.Providers;
using Xamarin.Forms;
using Xamarin.Forms.Xaml;

[assembly: XamlCompilation(XamlCompilationOptions.Compile)]
namespace AcronymClient.Core
{
	public class App : Application
	{
		public App(string databasePath)
		{
			var sqliteConnection = new SQLite.SQLiteAsyncConnection(databasePath);
			var acronymRepository = new AcronymRepository(sqliteConnection);

			var restfulProvider = new RestfulAcronymProvider("http://192.168.0.10:8080/api/v1/Acronym");
			var databaseProvider = new DatabaseAcronymProvider(acronymRepository);
			var cachedProvider = new CachedProvider(restfulProvider, databaseProvider, acronymRepository);

			var firstPage = new FirstPage(cachedProvider);
			MainPage = new NavigationPage(firstPage);
		}
	}
}

