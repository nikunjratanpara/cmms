package querybuilder

type ColumnConditionClause struct {
	*AbstractColumnClause
	Operator    Operator
	ColumnName2 string
}

func makeColumnClause(abstractClause *AbstractClause, columnName1 string, columnName2 string, operator Operator, isNot bool) *ColumnConditionClause {
	return &ColumnConditionClause{
		AbstractColumnClause: makeAbstractColumnClause(withColumnsType(abstractClause), columnName1, isNot),
		ColumnName2:          columnName2,
		Operator:             operator,
	}
}

func (clause ColumnConditionClause) GetSql(context QueryContext) string {
	return context.prepareIdentifier(clause.ColumnName) + " " + string(clause.Operator) + " " + context.prepareIdentifier(clause.ColumnName2)
}
