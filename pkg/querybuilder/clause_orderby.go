package querybuilder

import (
	"strings"

	"github.com/nikunjratanpara/cmms/pkg/linq"
)

type OrderByClause struct {
	*AbstractClause
	Columns []string
	IsDesc  bool
}

func makeOrderBy(abstractClause *AbstractClause, isDescending bool, columns ...string) OrderByClause {
	return OrderByClause{
		AbstractClause: withOrderByType(abstractClause),
		Columns:        columns,
		IsDesc:         isDescending,
	}
}

func (clause OrderByClause) GetSql(queryContext QueryContext) string {
	ascDesc := ""
	if clause.IsDesc {
		ascDesc = " DESC"
	}
	columns := linq.Map(
		clause.Columns,
		func(column string) string {
			return queryContext.prepareIdentifier(column) + ascDesc
		})

	return strings.Join(columns, ", ")
}
