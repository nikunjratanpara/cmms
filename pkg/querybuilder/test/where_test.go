package querybuilder

import (
	"testing"

	"github.com/nikunjratanpara/cmms/pkg/querybuilder"
	"github.com/stretchr/testify/assert"
)

func Test_Where(t *testing.T) {
	t.Run("Where Condition with Operator", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		query.Where("Account.Name", "Asset", "=")
		expectedParams := map[int]interface{}{0: "Asset"}
		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Account\".\"Name\" = $0",
				Parameters: expectedParams,
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [Account] WHERE [Account].[Name] = $0",
				Parameters: expectedParams,
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Account\".\"Name\" = $0",
				Parameters: expectedParams,
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `Account` WHERE `Account`.`Name` = $0",
				Parameters: expectedParams,
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Account\".\"Name\" = $0",
				Parameters: expectedParams,
			},
		}
		runTests(t, query, tests)
	})

	t.Run("raw condition without parameter", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		query.WhereRaw("[Test] <> [Test1]")
		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Test\" <> \"Test1\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [Account] WHERE [Test] <> [Test1]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Test\" <> \"Test1\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `Account` WHERE `Test` <> `Test1`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Test\" <> \"Test1\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("raw condition with parameter", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		query.WhereRaw("[Test] = ? AND [Test1] = ?", true, false)
		expectedParams := map[int]interface{}{0: true, 1: false}

		tests := map[string]querybuilder.QueryResult{
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Test\" = $0 AND \"Test1\" = $1",
				Parameters: expectedParams,
			},
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [Account] WHERE [Test] = $0 AND [Test1] = $1",
				Parameters: expectedParams,
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Test\" = $0 AND \"Test1\" = $1",
				Parameters: expectedParams,
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `Account` WHERE `Test` = $0 AND `Test1` = $1",
				Parameters: expectedParams,
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Test\" = $0 AND \"Test1\" = $1",
				Parameters: expectedParams,
			},
		}
		runTests(t, query, tests)
	})

	t.Run("raw condition with parameter mismatch", func(t *testing.T) {
		compiler := querybuilder.GetCompiler("SqlServer")
		query := querybuilder.NewQueryFromTable("Account")
		query.WhereRaw("[Test] = ?", true, false)
		assert.PanicsWithError(t, "mismatch parameters and parameter palceholders", func() { compiler.Compile(*query) })
	})

	t.Run("raw condition with parameter missing", func(t *testing.T) {
		compiler := querybuilder.GetCompiler("SqlServer")
		query := querybuilder.NewQueryFromTable("Account")
		query.WhereRaw("[Test] = ?")
		assert.PanicsWithError(t, "mismatch parameters and parameter palceholders", func() { compiler.Compile(*query) })
	})

	t.Run("All Type of where Condition", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		query.
			Where("Account.Name", "Asset", "=").
			WhereNull("Account.AccountId").
			WhereTrue("Account.Active").
			WhereFalse("Account.IsDeleted").
			WhereBetween("Account.Debit", 1500, 2000).
			WhereColumns("Account.Debit", "Account.Credit", "!=").
			WhereStartWith("Account.AccountType", "Asset").
			WhereEndWith("Account.ContactNo", "11111").
			WhereIn("Account.AssetType", []interface{}{1, 2, 3}...)

		expectedParams := map[int]interface{}{
			0: "Asset",
			1: 1500,
			2: 2000,
			3: "asset%",
			4: "%11111",
			5: 1,
			6: 2,
			7: 3,
		}
		tests := map[string]querybuilder.QueryResult{
			"SqlServer": {
				Sql:        "SELECT * FROM [Account] WHERE [Account].[Name] = $0 AND [Account].[AccountId] IS NULL AND [Account].[Active] = cast(1 as bit) AND [Account].[IsDeleted] = cast(0 as bit) AND [Account].[Debit] BETWEEN $1 AND $2 AND [Account].[Debit] != [Account].[Credit] AND LOWER([Account].[AccountType]) LIKE $3 AND LOWER([Account].[ContactNo]) LIKE $4 AND [Account].[AssetType] IN ( $5,$6,$7 )",
				Parameters: expectedParams,
			},
			"Postgres": {
				Sql:        "SELECT * FROM \"Account\" WHERE \"Account\".\"Name\" = $0 AND \"Account\".\"AccountId\" IS NULL AND \"Account\".\"Active\" = true AND \"Account\".\"IsDeleted\" = false AND \"Account\".\"Debit\" BETWEEN $1 AND $2 AND \"Account\".\"Debit\" != \"Account\".\"Credit\" AND LOWER(\"Account\".\"AccountType\") LIKE $3 AND LOWER(\"Account\".\"ContactNo\") LIKE $4 AND \"Account\".\"AssetType\" IN ( $5,$6,$7 )",
				Parameters: expectedParams,
			},
		}
		runTests(t, query, tests)
	})
}
