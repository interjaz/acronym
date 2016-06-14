using System;

using Android.App;
using Android.Content;
using Android.Content.PM;
using Android.Runtime;
using Android.Views;
using Android.Widget;
using Android.OS;
using AcronymClient.Core;
using Xamarin.Forms.Platform.Android;

namespace AcronymClient.Droid
{
	[Activity(Label = "AcronymClient.Droid", Icon = "@drawable/icon", Theme = "@style/MyTheme", MainLauncher = true, ConfigurationChanges = ConfigChanges.ScreenSize | ConfigChanges.Orientation)]
	public class MainActivity : global::Xamarin.Forms.Platform.Android.FormsAppCompatActivity
	{
		protected override void OnCreate(Bundle bundle)
		{
			// set the layout resources first
			FormsAppCompatActivity.ToolbarResource = Resource.Layout.Toolbar;
			FormsAppCompatActivity.TabLayoutResource = Resource.Layout.Tabbar;

			base.OnCreate(bundle);

			global::Xamarin.Forms.Forms.Init(this, bundle);

			LoadApplication(new App());
		}
	}
}

