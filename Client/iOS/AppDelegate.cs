using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using AcronymClient.Core;
using Foundation;
using UIKit;

namespace AcronymClient.iOS
{
	[Register("AppDelegate")]
	public partial class AppDelegate : global::Xamarin.Forms.Platform.iOS.FormsApplicationDelegate
	{
		public override bool FinishedLaunching(UIApplication app, NSDictionary options)
		{
			global::Xamarin.Forms.Forms.Init();

			var dbPath = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.MyDocuments), "Acronyms.db");
			LoadApplication(new App(dbPath));

			return base.FinishedLaunching(app, options);
		}
	}
}

