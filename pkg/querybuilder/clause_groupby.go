package querybuilder

import (
	"strings"

	"github.com/nikunjratanpara/cmms/pkg/linq"
)

type GroupByClause struct {
	*AbstractClause
	Columns []string
}

func makeGroupByClause(columns ...string) *GroupByClause {
	return &GroupByClause{
		AbstractClause: makeAbstractClause(
			withGroupByComponent,
			withGroupByType),
		Columns: columns,
	}
}

func (clause GroupByClause) GetSql(context QueryContext) string {
	groupByClauses := linq.Map(
		clause.Columns,
		func(columnName string) string {
			return context.prepareIdentifier(columnName)
		})
	return strings.Join(groupByClauses, ", ")
}
