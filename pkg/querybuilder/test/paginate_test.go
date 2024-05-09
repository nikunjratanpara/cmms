package querybuilder

import (
	"testing"

	"github.com/nikunjratanpara/cmms/pkg/querybuilder"
)

func Test_Paginate(t *testing.T) {
	t.Run("without offset", func(t *testing.T) {

		query := querybuilder.NewQueryFromTable("users").
			OrderBy("BirthDate").
			Limit(20)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY \"BirthDate\" LIMIT 20",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT TOP(20) * FROM [users] ORDER BY [BirthDate]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY \"BirthDate\" OFFSET 0 ROWS FETCH NEXT 20 ROWS ONLY",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` ORDER BY `BirthDate` LIMIT 20",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY \"BirthDate\" LIMIT 20",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Limit, offset and Orderby", func(t *testing.T) {

		query := querybuilder.NewQueryFromTable("users").
			OrderBy("BirthDate").
			Limit(20).
			Offset(20)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY \"BirthDate\" LIMIT 20 OFFSET 20",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] ORDER BY [BirthDate] OFFSET 20 ROWS FETCH NEXT 20 ROWS ONLY",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY \"BirthDate\" OFFSET 20 ROWS FETCH NEXT 20 ROWS ONLY",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` ORDER BY `BirthDate` LIMIT 20 OFFSET 20",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY \"BirthDate\" LIMIT 20 OFFSET 20",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("without limit", func(t *testing.T) {

		query := querybuilder.NewQueryFromTable("users").
			OrderBy("BirthDate").
			Offset(20)

		compiler := querybuilder.GetCompiler(querybuilder.SqlServerEngineScope)
		compiler.Compile(*query)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY \"BirthDate\" OFFSET 20",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] ORDER BY [BirthDate] OFFSET 20 ROWS",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY \"BirthDate\" OFFSET 20 ROWS",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` ORDER BY `BirthDate` OFFSET 20",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY \"BirthDate\" OFFSET 20",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Without Order By", func(t *testing.T) {

		query := querybuilder.NewQueryFromTable("users").
			Limit(20).
			Offset(20)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" LIMIT 20 OFFSET 20",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] ORDER BY (SELECT NULL) OFFSET 20 ROWS FETCH NEXT 20 ROWS ONLY",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" ORDER BY (SELECT 0 FROM DUAL) OFFSET 20 ROWS FETCH NEXT 20 ROWS ONLY",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` LIMIT 20 OFFSET 20",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" LIMIT 20 OFFSET 20",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

}
