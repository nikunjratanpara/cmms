package querybuilder

type BetweenCondition struct {
	*AbstractColumnClause
	FromValue interface{}
	ToValue   interface{}
}

// makeBetweenClause creates a BetweenCondition with the given column name,
// from value, and to value. The function takes a string for the column name,
// and any interface{} type for the from and to values. It returns a pointer
// to a BetweenCondition.
func makeBetweenClause(abstractClause *AbstractClause, columnName string, fromValue, toValue interface{}, isNot bool) *BetweenCondition {
	return &BetweenCondition{
		AbstractColumnClause: makeAbstractColumnClause(withBetweenType(abstractClause), columnName, isNot),
		FromValue:            fromValue,
		ToValue:              toValue,
	}
}

func makeBetween(abstractClause *AbstractClause, columnName string, fromValue, toValue interface{}) *BetweenCondition {
	return makeBetweenClause(abstractClause, columnName, fromValue, toValue, false)
}

func makeNotBetween(abstractClause *AbstractClause, columnName string, fromValue, toValue interface{}) *BetweenCondition {
	return makeBetweenClause(abstractClause, columnName, fromValue, toValue, true)
}

func (clause BetweenCondition) GetSql(context QueryContext) string {
	if clause.FromValue == nil || clause.ToValue == nil {
		panic("from or to value is nil")
	}

	if clause.FromValue == clause.ToValue {
		return context.prepareIdentifier(clause.ColumnName) + " = " +
			context.AddParameter(clause.FromValue)
	}

	strClause := " BETWEEN "
	if clause.IsNot {
		strClause = " NOT BETWEEN "
	}

	return context.prepareIdentifier(clause.ColumnName) +
		strClause +
		context.AddParameter(clause.FromValue) + " AND " +
		context.AddParameter(clause.ToValue)
}
