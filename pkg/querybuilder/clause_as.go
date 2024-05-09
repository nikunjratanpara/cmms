package querybuilder

type AsClause struct {
	*AbstractClause
	Alias string
}

func makeAsClause(abstractClause *AbstractClause, alias string) *AsClause {
	return &AsClause{
		AbstractClause: withAsType(abstractClause),
		Alias:          alias,
	}
}

func (clause AsClause) GetSql(context QueryContext) string {
	return "AS " + context.wrapToken(clause.Alias)
}
