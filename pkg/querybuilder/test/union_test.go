package querybuilder

import (
	"testing"

	"github.com/nikunjratanpara/cmms/pkg/querybuilder"
)

func TestUnion(t *testing.T) {

	t.Run("Union Test", func(t *testing.T) {
		qb1 := querybuilder.NewQueryFromTable("Customers").
			Select("UserName").
			Where("Catagory", 1, "=")

		qb2 := querybuilder.NewQueryFromTable("Employees").
			Select("UserName").
			Where("Catagory", 2, "=")

		unionQuery := qb1.Union(qb2)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT [UserName] FROM [Customers] WHERE [Catagory] = $0 UNION SELECT [UserName] FROM [Employees] WHERE [Catagory] = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT \"UserName\" FROM \"Customers\" WHERE \"Catagory\" = $0 UNION SELECT \"UserName\" FROM \"Employees\" WHERE \"Catagory\" = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT \"UserName\" FROM \"Customers\" WHERE \"Catagory\" = $0 UNION SELECT \"UserName\" FROM \"Employees\" WHERE \"Catagory\" = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT `UserName` FROM `Customers` WHERE `Catagory` = $0 UNION SELECT `UserName` FROM `Employees` WHERE `Catagory` = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT \"UserName\" FROM \"Customers\" WHERE \"Catagory\" = $0 UNION SELECT \"UserName\" FROM \"Employees\" WHERE \"Catagory\" = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
		}
		runTests(t, unionQuery, tests)
	})

	t.Run("Union All Test", func(t *testing.T) {
		qb1 := querybuilder.NewQueryFromTable("Customers").
			Select("UserName").
			Where("Catagory", 1, "=")

		qb2 := querybuilder.NewQueryFromTable("Employees").
			Select("UserName").
			Where("Catagory", 2, "=")

		unionQuery := qb1.UnionAll(qb2)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT [UserName] FROM [Customers] WHERE [Catagory] = $0 UNION ALL SELECT [UserName] FROM [Employees] WHERE [Catagory] = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT \"UserName\" FROM \"Customers\" WHERE \"Catagory\" = $0 UNION ALL SELECT \"UserName\" FROM \"Employees\" WHERE \"Catagory\" = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT \"UserName\" FROM \"Customers\" WHERE \"Catagory\" = $0 UNION ALL SELECT \"UserName\" FROM \"Employees\" WHERE \"Catagory\" = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT `UserName` FROM `Customers` WHERE `Catagory` = $0 UNION ALL SELECT `UserName` FROM `Employees` WHERE `Catagory` = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT \"UserName\" FROM \"Customers\" WHERE \"Catagory\" = $0 UNION ALL SELECT \"UserName\" FROM \"Employees\" WHERE \"Catagory\" = $1",
				Parameters: map[int]interface{}{0: 1, 1: 2},
			},
		}
		runTests(t, unionQuery, tests)

	})
}
