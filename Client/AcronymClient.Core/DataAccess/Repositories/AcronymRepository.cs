using System;
using System.Collections.Generic;
using System.Reflection;
using System.Threading.Tasks;
using AcronymClient.Core.DataAccess.Entities;
using AcronymClient.Core.Utils;
using SQLite;

namespace AcronymClient.Core.DataAccess.Repository
{
	public class AcronymRepository : IAcronymRepository
	{
		readonly SQLiteAsyncConnection _connection;
		private readonly Lazy<Task> _initialize;

		public AcronymRepository(
			[Utils.NotNull] SQLiteAsyncConnection connection)
		{
			Check.NotNull(connection, nameof(connection));

			_connection = connection;
			_initialize = new Lazy<Task>(async () =>
			{
				await _connection.CreateTableAsync<AcronymModel>();
			});
		}

		public async Task<IEnumerable<AcronymModel>> FindAllAsync([Utils.NotNull] string acronym)
		{
			Check.NotNull(acronym, nameof(acronym));

			var result = await _connection.Table<AcronymModel>()
										  .Where(s => s.Acronym.StartsWith(acronym))
										  .ToListAsync();

			return result;
		}

		public async Task UpdateAllAsync([Utils.NotNull] IEnumerable<AcronymModel> acronyms)
		{
			Check.NotNull(acronyms, nameof(acronyms));

			await _initialize.Value;

			var tableName = typeof(AcronymModel).GetTypeInfo().GetCustomAttribute<TableAttribute>().Name;
			var columnName = nameof(AcronymModel.Acronym);
			var command = string.Format("DELETE FROM {0} WHERE {1} = ?", tableName, columnName);

			// InsertAllOrReplace is throwing exception here.
			// lets try do it differently
			await _connection.RunInTransactionAsync(conn =>
			{
				foreach (var acronym in acronyms)
				{
					conn.Execute(command, acronym.Acronym);
				}

				conn.InsertAll(acronyms);
			});

		}

	}
}

