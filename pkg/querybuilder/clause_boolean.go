package querybuilder

type BooleanConditionClause struct {
	*ConditionClause
}

func makeBooleanClause(abstractClause *AbstractClause, columnName string, val bool) *BooleanConditionClause {
	abstractColumnClause := makeAbstractColumnClause(withIsBooleanType(abstractClause), columnName, false)
	return &BooleanConditionClause{
		ConditionClause: makeCoditionClause(abstractColumnClause, val, Eq),
	}
}

func makeTrue(abstractClause *AbstractClause, columnName string) *BooleanConditionClause {
	return makeBooleanClause(abstractClause, columnName, true)
}

func makeFalse(abstractClause *AbstractClause, columnName string) *BooleanConditionClause {
	return makeBooleanClause(abstractClause, columnName, false)
}

func (clause BooleanConditionClause) GetSql(context QueryContext) string {
	val, ok := clause.Value.(bool)
	if !ok {
		panic("invalid value for boolean condition clause")
	}
	return getBooleanConditionSql(context, clause.ColumnName, val)
}

func getBooleanConditionSql(context QueryContext, columnName string, bvalue bool) string {
	return context.prepareIdentifier(columnName) + " = " + addBooleanValueParam(context, bvalue)
}

func addBooleanValueParam(context QueryContext, param bool) string {
	if param {
		return context.CompileTrue()
	} else {
		return context.CompileFalse()
	}
}
