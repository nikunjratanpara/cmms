package querybuilder

import (
	"testing"

	"github.com/nikunjratanpara/cmms/pkg/querybuilder"
)

func Test_OrderBy(t *testing.T) {
	compiler := querybuilder.GetCompiler("Postgres")
	t.Run("Order by ascending", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account").
			OrderBy("AccountType", "AccountName")
		sqlResult := compiler.Compile(*query)
		assertStringMatch(t, "SELECT * FROM \"Account\" ORDER BY \"AccountType\", \"AccountName\"", sqlResult.Sql)
	})

	t.Run("Order by descending", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account").
			OrderByDesc("AccountType", "AccountName")
		sqlResult := compiler.Compile(*query)
		assertStringMatch(t, "SELECT * FROM \"Account\" ORDER BY \"AccountType\" DESC, \"AccountName\" DESC", sqlResult.Sql)
	})

	t.Run("Order by ascending and descending ", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("Account").
			OrderByDesc("AccountName").
			OrderBy("AccountType")

		sqlResult := compiler.Compile(*query)
		assertStringMatch(t, "SELECT * FROM \"Account\" ORDER BY \"AccountName\" DESC, \"AccountType\"", sqlResult.Sql)
	})
}
