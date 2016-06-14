using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Threading.Tasks;
using System.Windows.Input;
using AcronymClient.Core.DataAccess.Entities;
using AcronymClient.Core.Providers;
using AcronymClient.Core.Utils;
using Xamarin.Forms;

namespace AcronymClient.Core
{
	public partial class FirstPage : ContentPage
	{
		private readonly IAcronymProvider _acronymProvider;

		public FirstPage(
			[NotNull] IAcronymProvider acronymProvider
			)
		{
			Check.NotNull(acronymProvider, nameof(acronymProvider));

			InitializeComponent();
			this.BindingContext = this;

			_acronymProvider = acronymProvider;
			lstAcronyms.ItemTapped += OnItemTapped;

			LstAcronyms = new ObservableCollection<AcronymModel>();
			TxtToFind = "aa";
			SearchAsync(TxtToFind).ConfigureAwait(false);
		}

		private ObservableCollection<AcronymModel> _lstAcronyms;
		public ObservableCollection<AcronymModel> LstAcronyms
		{
			get
			{
				return _lstAcronyms;
			}
			set
			{
				if (_lstAcronyms == value)
				{
					return;
				}

				_lstAcronyms = value;
				OnPropertyChanged();
			}
		}

		private string _txtStatus;
		public string TxtStatus
		{
			get
			{
				return _txtStatus;
			}
			set
			{
				if (_txtStatus == value)
				{
					return;
				}
				_txtStatus = value;
				OnPropertyChanged();
			}
		}

		private string _txtToFind;
		public string TxtToFind
		{
			get
			{
				return _txtToFind;
			}
			set
			{
				if (_txtToFind == value)
				{
					return;
				}
				_txtToFind = value;
				OnPropertyChanged();
			}
		}

		private Command _searchCommand;
		public Command SearchCommand
		{
			get
			{
				_searchCommand = _searchCommand ?? new Command(() => SearchAsync(TxtToFind).ConfigureAwait(false));
				return _searchCommand;
			}
		}

		private async Task SearchAsync(string toFind)
		{
			try
			{
				UpdateStatus("Retriving ...", true);
				var foundAcronyms = await _acronymProvider.FindAsync(toFind);
				if (foundAcronyms.Error != null)
				{
					UpdateStatus("Failed", false);

					System.Diagnostics.Debug.WriteLine(foundAcronyms.Error.ErrorCode);
					System.Diagnostics.Debug.WriteLine(foundAcronyms.Error.Tag);
					return;
				}

				Device.BeginInvokeOnMainThread(() =>
				{
					LstAcronyms = new ObservableCollection<AcronymModel>(foundAcronyms.Some);
					TxtStatus = "Found: " + LstAcronyms.Count;
					lstAcronyms.IsRefreshing = false;
				});
			}
			catch (Exception ex)
			{
				UpdateStatus("Failed", false);
				System.Diagnostics.Debug.WriteLine(ex);
			}
		}

		private void UpdateStatus(string status, bool isRefreshing)
		{
			Device.BeginInvokeOnMainThread(() =>
			{
				TxtStatus = status;
				lstAcronyms.IsRefreshing = isRefreshing;
			});
		}

		private void OnItemTapped(object sender, ItemTappedEventArgs e)
		{
			var acronym = e.Item as AcronymModel;
			if (acronym == null)
			{
				return;
			}

			var uri = new Uri(acronym.Url);
			Device.OpenUri(uri);
		}

	}
}

