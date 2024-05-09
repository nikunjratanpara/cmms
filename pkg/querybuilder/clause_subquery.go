package querybuilder

type SubQueryConditionClause struct {
	*AbstractColumnClause
	Query    *Query
	operator Operator
}

type SubQueryClause struct {
	*AbstractClause
	Query *Query
}

func makeSubQueryClause(abstractClause *AbstractClause, query Query) *SubQueryClause {
	return &SubQueryClause{
		AbstractClause: withSubQueryType(abstractClause),
		Query:          &query,
	}
}

func makeSubQueryConditionClause(abstractClause *AbstractClause, columnName string, operator Operator, query Query, alias string, isNot bool) *SubQueryConditionClause {
	(&query).As(alias)
	return &SubQueryConditionClause{
		AbstractColumnClause: makeAbstractColumnClause(withSubQueryType(abstractClause), columnName, isNot),
		Query:                &query,
		operator:             operator,
	}
}

func (c SubQueryClause) GetSql(ctx QueryContext) string {
	return ctx.compileSubQuery(*c.Query)
}

func (c SubQueryConditionClause) GetSql(ctx QueryContext) string {
	subQuerySQL := ctx.compileSubQuery(*c.Query)
	isNot := ""
	if c.IsNot {
		isNot = "NOT"
	}

	return isNot + "(" + ctx.prepareIdentifier(c.ColumnName) + " " + string(c.operator) + " (" + subQuerySQL + ")" + ")"
}
