using System;
using AcronymClient.Core;
using AcronymClient.Core.Controls;
using AcronymClient.iOS;
using AcronymClient.iOS.Renderers;
using UIKit;
using Xamarin.Forms;
using Xamarin.Forms.Platform.iOS;

[assembly: ExportRenderer(typeof(NoHelpEntry), typeof(NoHelpEntryRenderer))]

namespace AcronymClient.iOS.Renderers
{
	public class NoHelpEntryRenderer : EntryRenderer
    {
        protected override void OnElementChanged(ElementChangedEventArgs<Entry> e)
		{
			base.OnElementChanged(e);
			if (Control != null)
			{
				Control.SpellCheckingType = UITextSpellCheckingType.No;             // No Spellchecking
				Control.AutocorrectionType = UITextAutocorrectionType.No;           // No Autocorrection
				Control.AutocapitalizationType = UITextAutocapitalizationType.None; // No Autocapitalization
			}
		}
	}
}

