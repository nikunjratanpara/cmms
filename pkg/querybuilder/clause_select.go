package querybuilder

import (
	"strings"

	"github.com/nikunjratanpara/cmms/pkg/linq"
)

type SelectClause struct {
	*AbstractClause
	Columns []string
}

func makeSelectAbstractClause(optionFuncs ...OptionsBuilder[AbstractClause]) *AbstractClause {
	return makeAbstractClause(append(optionFuncs, withSelectComponent)...)
}

var withSelectType = withType("Select")

func makeSelectClause(columns ...string) *SelectClause {
	return &SelectClause{
		AbstractClause: makeSelectAbstractClause(withSelectType),
		Columns:        columns,
	}
}

func makeSelectClauseFromSubQuery(subQuery Query, alias string) *SubQueryClause {
	(&subQuery).As(alias)
	return &SubQueryClause{
		AbstractClause: makeSelectAbstractClause(withSelectType),
		Query:          &subQuery,
	}
}

func makeSelectRawSqlClause(rawSql string) *RawSqlClause {
	return makeRawSqlClause(makeSelectAbstractClause(withRawSqlType), rawSql)
}

func (clause SelectClause) GetSql(context QueryContext) string {
	columns := linq.Map(
		clause.Columns,
		func(column string) string {
			return context.prepareIdentifier(column)
		})

	return strings.Join(columns, ", ")
}
