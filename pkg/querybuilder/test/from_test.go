package querybuilder

import (
	"testing"

	"github.com/nikunjratanpara/cmms/pkg/querybuilder"
)

func Test_From(t *testing.T) {

	t.Run("From Table", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"Account\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [Account]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"Account\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `Account`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"Account\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("From Table with as clause tokenize alias", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account As AccountSrc")
		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"Account\" AS \"AccountSrc\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [Account] AS [AccountSrc]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"Account\" \"AccountSrc\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `Account` AS `AccountSrc`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"Account\" AS \"AccountSrc\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("From SubQuery Respects last defined alias", func(t *testing.T) {
		subquery := querybuilder.NewQueryFromTable("Account").As("Account1")
		query := querybuilder.NewQueryFromSubQuery(*subquery, "AccountSrc")

		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM (SELECT * FROM \"Account\") AS \"AccountSrc\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM (SELECT * FROM [Account]) AS [AccountSrc]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM (SELECT * FROM \"Account\") AS \"AccountSrc\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM (SELECT * FROM `Account`) AS `AccountSrc`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM (SELECT * FROM \"Account\") AS \"AccountSrc\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("From Raw SQL", func(t *testing.T) {
		query := querybuilder.NewQueryFromRawSql("[Comments] TABLESAMPLE SYSTEM (10 PERCENT)")
		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"Comments\" TABLESAMPLE SYSTEM (10 PERCENT)",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [Comments] TABLESAMPLE SYSTEM (10 PERCENT)",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"Comments\" TABLESAMPLE SYSTEM (10 PERCENT)",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `Comments` TABLESAMPLE SYSTEM (10 PERCENT)",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"Comments\" TABLESAMPLE SYSTEM (10 PERCENT)",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})
}
