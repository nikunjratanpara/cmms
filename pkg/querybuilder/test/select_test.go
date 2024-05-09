package querybuilder

import (
	"testing"

	"github.com/nikunjratanpara/cmms/pkg/querybuilder"
)

func Test_Select(t *testing.T) {
	t.Run("check schema with table", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("dbo.Account")
		query.Select("Id", "Name")
		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT \"Id\", \"Name\" FROM \"dbo\".\"Account\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT [Id], [Name] FROM [dbo].[Account]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT \"Id\", \"Name\" FROM \"dbo\".\"Account\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT `Id`, `Name` FROM `dbo`.`Account`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT \"Id\", \"Name\" FROM \"dbo\".\"Account\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Check column alias", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("dbo.Account")
		query.Select("Id", "Name as NameAlias")
		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT \"Id\", \"Name\" AS \"NameAlias\" FROM \"dbo\".\"Account\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT [Id], [Name] AS [NameAlias] FROM [dbo].[Account]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT \"Id\", \"Name\" AS \"NameAlias\" FROM \"dbo\".\"Account\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT `Id`, `Name` AS `NameAlias` FROM `dbo`.`Account`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT \"Id\", \"Name\" AS \"NameAlias\" FROM \"dbo\".\"Account\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})
}
