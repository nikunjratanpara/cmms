package querybuilder

import "fmt"

type UnionClause struct {
	*AbstractClause
	IsUnionAll bool
	Query      Query
}

func makeUnionClause(abstractClause *AbstractClause, isUnionAll bool, query Query) *UnionClause {
	return &UnionClause{
		AbstractClause: withUnionType(abstractClause),
		IsUnionAll:     isUnionAll,
		Query:          query,
	}
}

func (unionClause UnionClause) GetSql(ctx QueryContext) string {
	unionType := "UNION"

	if unionClause.IsUnionAll {
		unionType = "UNION ALL"
	}
	return fmt.Sprintf("%s %s", unionType, ctx.compileSelectQuery(unionClause.Query))
}
