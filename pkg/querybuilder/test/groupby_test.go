package querybuilder

import (
	"testing"

	"github.com/nikunjratanpara/cmms/pkg/querybuilder"
	"github.com/stretchr/testify/assert"
)

func Test_GroupBy(t *testing.T) {

	t.Run("Group by without having", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		query.GroupBy("Account.Name")

		tests := map[string]querybuilder.QueryResult{
			"Postgres": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\"",
				Parameters: map[int]interface{}{},
			},
			"SqlServer": {
				Sql:        "SELECT * FROM [Account] GROUP BY [Account].[Name]",
				Parameters: map[int]interface{}{},
			},
			"Oracle": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\"",
				Parameters: map[int]interface{}{},
			},
			"MySql": {
				Sql:        "SELECT * FROM `Account` GROUP BY `Account`.`Name`",
				Parameters: map[int]interface{}{},
			},
			"Sqlite": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\"",
				Parameters: map[int]interface{}{},
			},
		}
		for compiler, expectedResult := range tests {
			t.Run(compiler, func(t *testing.T) {
				compileQueryAndMatchResult(t, compiler, query, expectedResult)
			})
		}
	})

	t.Run("Group by with multiple columns", func(t *testing.T) {

		query := querybuilder.NewQueryFromTable("Account")
		query.GroupBy("Account.Name", "Account.ContactNo")

		tests := map[string]querybuilder.QueryResult{
			"Postgres": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\", \"Account\".\"ContactNo\"",
				Parameters: map[int]interface{}{},
			},
			"SqlServer": {
				Sql:        "SELECT * FROM [Account] GROUP BY [Account].[Name], [Account].[ContactNo]",
				Parameters: map[int]interface{}{},
			},
			"Oracle": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\", \"Account\".\"ContactNo\"",
				Parameters: map[int]interface{}{},
			},
			"MySql": {
				Sql:        "SELECT * FROM `Account` GROUP BY `Account`.`Name`, `Account`.`ContactNo`",
				Parameters: map[int]interface{}{},
			},
			"Sqlite": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\", \"Account\".\"ContactNo\"",
				Parameters: map[int]interface{}{},
			},
		}

		runTests(t, query, tests)

	})

	t.Run("Group by with multiple columns each with single groupby clause", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		query.GroupBy("Account.Name")
		query.GroupBy("Account.ContactNo")

		tests := map[string]querybuilder.QueryResult{
			"Postgres": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\", \"Account\".\"ContactNo\"",
				Parameters: map[int]interface{}{},
			},
			"SqlServer": {
				Sql:        "SELECT * FROM [Account] GROUP BY [Account].[Name], [Account].[ContactNo]",
				Parameters: map[int]interface{}{},
			},
			"Oracle": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\", \"Account\".\"ContactNo\"",
				Parameters: map[int]interface{}{},
			},
			"MySql": {
				Sql:        "SELECT * FROM `Account` GROUP BY `Account`.`Name`, `Account`.`ContactNo`",
				Parameters: map[int]interface{}{},
			},
			"Sqlite": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\", \"Account\".\"ContactNo\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Group by with multiple columns multiple groupby Clause", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		query.GroupBy("Account.Name", "Account.ContactNo")
		query.GroupBy("Account.Debit", "Account.Credit")

		tests := map[string]querybuilder.QueryResult{
			"Postgres": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\", \"Account\".\"ContactNo\", \"Account\".\"Debit\", \"Account\".\"Credit\"",
				Parameters: map[int]interface{}{},
			},
			"SqlServer": {
				Sql:        "SELECT * FROM [Account] GROUP BY [Account].[Name], [Account].[ContactNo], [Account].[Debit], [Account].[Credit]",
				Parameters: map[int]interface{}{},
			},
			"Oracle": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\", \"Account\".\"ContactNo\", \"Account\".\"Debit\", \"Account\".\"Credit\"",
				Parameters: map[int]interface{}{},
			},
			"MySql": {
				Sql:        "SELECT * FROM `Account` GROUP BY `Account`.`Name`, `Account`.`ContactNo`, `Account`.`Debit`, `Account`.`Credit`",
				Parameters: map[int]interface{}{},
			},
			"Sqlite": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\", \"Account\".\"ContactNo\", \"Account\".\"Debit\", \"Account\".\"Credit\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Group By raw sql", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		query.GroupByRaw("[Account].[Profit] WITH ROLLUP")

		tests := map[string]querybuilder.QueryResult{
			"Postgres": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Profit\" WITH ROLLUP",
				Parameters: map[int]interface{}{},
			},
			"SqlServer": {
				Sql:        "SELECT * FROM [Account] GROUP BY [Account].[Profit] WITH ROLLUP",
				Parameters: map[int]interface{}{},
			},
			"Oracle": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Profit\" WITH ROLLUP",
				Parameters: map[int]interface{}{},
			},
			"MySql": {
				Sql:        "SELECT * FROM `Account` GROUP BY `Account`.`Profit` WITH ROLLUP",
				Parameters: map[int]interface{}{},
			},
			"Sqlite": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Profit\" WITH ROLLUP",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Group by with having", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account")
		query.GroupBy("Account.Name")
		query.Having("Account.Name", "Asset", "=")

		expectedParams := map[int]interface{}{0: "Asset"}
		tests := map[string]querybuilder.QueryResult{
			"Postgres": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\" HAVING \"Account\".\"Name\" = $0",
				Parameters: expectedParams,
			},
			"SqlServer": {
				Sql:        "SELECT * FROM [Account] GROUP BY [Account].[Name] HAVING [Account].[Name] = $0",
				Parameters: expectedParams,
			},
			"Oracle": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\" HAVING \"Account\".\"Name\" = $0",
				Parameters: expectedParams,
			},
			"MySql": {
				Sql:        "SELECT * FROM `Account` GROUP BY `Account`.`Name` HAVING `Account`.`Name` = $0",
				Parameters: expectedParams,
			},
			"Sqlite": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\" HAVING \"Account\".\"Name\" = $0",
				Parameters: expectedParams,
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Group by with all type of having conditions", func(t *testing.T) {

		query := querybuilder.NewQueryFromTable("Account")
		query.
			GroupBy("Account.Name").
			Having("Account.Name", "Asset", "=").
			HavingNull("Account.AccountId").
			HavingTrue("Account.Active").
			HavingFalse("Account.IsDeleted").
			HavingBetween("Account.Debit", 1500, 2000).
			HavingColumns("Account.Debit", "Account.Credit", "!=").
			HavingStartWith("Account.AccountType", "Asset").
			HavingEndWith("Account.ContactNo", "11111").
			HavingIn("Account.AssetType", []interface{}{1, 2, 3}...)

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
			"Postgres": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\" HAVING \"Account\".\"Name\" = $0 AND \"Account\".\"AccountId\" IS NULL AND \"Account\".\"Active\" = true AND \"Account\".\"IsDeleted\" = false AND \"Account\".\"Debit\" BETWEEN $1 AND $2 AND \"Account\".\"Debit\" != \"Account\".\"Credit\" AND LOWER(\"Account\".\"AccountType\") LIKE $3 AND LOWER(\"Account\".\"ContactNo\") LIKE $4 AND \"Account\".\"AssetType\" IN ( $5,$6,$7 )",
				Parameters: expectedParams,
			},
			"SqlServer": {
				Sql:        "SELECT * FROM [Account] GROUP BY [Account].[Name] HAVING [Account].[Name] = $0 AND [Account].[AccountId] IS NULL AND [Account].[Active] = cast(1 as bit) AND [Account].[IsDeleted] = cast(0 as bit) AND [Account].[Debit] BETWEEN $1 AND $2 AND [Account].[Debit] != [Account].[Credit] AND LOWER([Account].[AccountType]) LIKE $3 AND LOWER([Account].[ContactNo]) LIKE $4 AND [Account].[AssetType] IN ( $5,$6,$7 )",
				Parameters: expectedParams,
			},
			"MySql": {
				Sql:        "SELECT * FROM `Account` GROUP BY `Account`.`Name` HAVING `Account`.`Name` = $0 AND `Account`.`AccountId` IS NULL AND `Account`.`Active` = true AND `Account`.`IsDeleted` = false AND `Account`.`Debit` BETWEEN $1 AND $2 AND `Account`.`Debit` != `Account`.`Credit` AND LOWER(`Account`.`AccountType`) LIKE $3 AND LOWER(`Account`.`ContactNo`) LIKE $4 AND `Account`.`AssetType` IN ( $5,$6,$7 )",
				Parameters: expectedParams,
			},
			"Oracle": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\" HAVING \"Account\".\"Name\" = $0 AND \"Account\".\"AccountId\" IS NULL AND \"Account\".\"Active\" = true AND \"Account\".\"IsDeleted\" = false AND \"Account\".\"Debit\" BETWEEN $1 AND $2 AND \"Account\".\"Debit\" != \"Account\".\"Credit\" AND LOWER(\"Account\".\"AccountType\") LIKE $3 AND LOWER(\"Account\".\"ContactNo\") LIKE $4 AND \"Account\".\"AssetType\" IN ( $5,$6,$7 )",
				Parameters: expectedParams,
			},
			"Sqlite": {
				Sql:        "SELECT * FROM \"Account\" GROUP BY \"Account\".\"Name\" HAVING \"Account\".\"Name\" = $0 AND \"Account\".\"AccountId\" IS NULL AND \"Account\".\"Active\" = 1 AND \"Account\".\"IsDeleted\" = 0 AND \"Account\".\"Debit\" BETWEEN $1 AND $2 AND \"Account\".\"Debit\" != \"Account\".\"Credit\" AND LOWER(\"Account\".\"AccountType\") LIKE $3 AND LOWER(\"Account\".\"ContactNo\") LIKE $4 AND \"Account\".\"AssetType\" IN ( $5,$6,$7 )",
				Parameters: expectedParams,
			},
		}
		runTests(t, query, tests)
	})
}

func runTests(t *testing.T, query *querybuilder.Query, tests map[string]querybuilder.QueryResult) {
	for compiler, expectedResult := range tests {
		t.Run(compiler, func(t *testing.T) {
			compileQueryAndMatchResult(t, compiler, query, expectedResult)
		})
	}
}

func compileQueryAndMatchResult(t *testing.T, compiler string, query *querybuilder.Query, expectedResult querybuilder.QueryResult) {
	sqlCompiler := querybuilder.GetCompiler(compiler)
	queryResult := sqlCompiler.Compile(*query)
	assertQueryResultMatch(t, expectedResult, queryResult)
}

func assertQueryResultMatch(t *testing.T, expectedResult, queryResult querybuilder.QueryResult) {
	assert.Equalf(t, expectedResult.Sql, queryResult.Sql, "Generated sql is not matching expected sql, generated sql : %s", queryResult.Sql)
	assertMapMatch(t, expectedResult.Parameters, queryResult.Parameters)
}

func assertStringMatch(t *testing.T, expectedString string, actualString string) {
	if expectedString != actualString {
		t.Fatalf("Generated sql is not matching expected sql, generated sql : %s", actualString)
	}
}

func assertMapMatch[TKey comparable, TVal comparable](t *testing.T, expectedParams map[TKey]TVal, actualParams map[TKey]TVal) {
	for index, value := range actualParams {
		expectedValue := expectedParams[index]
		if expectedValue != value {
			t.Fatalf("Params are not matching at index %v, expected: %v ,actual: %v", index, expectedParams[index], value)
		}
	}
}
