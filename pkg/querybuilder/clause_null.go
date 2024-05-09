package querybuilder

type NullConditionClause struct {
	*AbstractColumnClause
}

func makeNullClause(abstractClause *AbstractClause, columnName string, isNot bool) *NullConditionClause {
	return &NullConditionClause{
		AbstractColumnClause: makeAbstractColumnClause(withNullType(abstractClause), columnName, isNot),
	}
}

func makeNull(abstractClause *AbstractClause, columnName string) *NullConditionClause {
	return makeNullClause(abstractClause, columnName, false)
}

func makeNotNull(abstractClause *AbstractClause, columnName string) *NullConditionClause {
	return makeNullClause(abstractClause, columnName, true)
}

func (clause NullConditionClause) GetSql(context QueryContext) string {
	strClause := " IS NULL"
	if clause.IsNot {
		strClause = " IS NOT NULL"
	}
	return context.prepareIdentifier(clause.ColumnName) + strClause
}
