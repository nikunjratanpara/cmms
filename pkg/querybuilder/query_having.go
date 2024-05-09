package querybuilder

func makeHavingAbstractClause(optionsBuilder ...OptionsBuilder[AbstractClause]) *AbstractClause {
	return makeAbstractClause(append(optionsBuilder, withHavingComponent)...)
}

func makeHavingAbstractColumnClause(columnName string, isNot bool) *AbstractColumnClause {
	return makeAbstractColumnClause(makeHavingAbstractClause(), columnName, isNot)
}

func (query *Query) Having(columnName string, value interface{}, operator Operator) *Query {
	if operator == "" {
		operator = Eq
	}
	return query.addHavingClause(makeCoditionClause(makeHavingAbstractColumnClause(columnName, false), value, operator))
}

func (query *Query) NotHaving(columnName string, value interface{}, operator Operator) *Query {
	if operator == "" {
		operator = Eq
	}
	return query.addHavingClause(makeCoditionClause(makeHavingAbstractColumnClause(columnName, true), value, operator))
}

func (query *Query) HavingBetween(columnName string, fromValue interface{}, toValue interface{}) *Query {
	return query.addHavingClause(makeBetween(makeHavingAbstractClause(), columnName, fromValue, toValue))
}

func (query *Query) HavingNotBetween(columnName string, fromValue interface{}, toValue interface{}) *Query {
	return query.addHavingClause(makeNotBetween(makeHavingAbstractClause(), columnName, fromValue, toValue))
}

func (query *Query) HavingContains(columnName string, value string) *Query {
	return query.addHavingClause(makeContains(makeHavingAbstractClause(), columnName, value))
}

func (query *Query) HavingNotContains(columnName string, value string) *Query {
	return query.addHavingClause(makeNotContains(makeHavingAbstractClause(), columnName, value))
}

func (query *Query) HavingLike(columnName string, value string) *Query {
	return query.addHavingClause(makeLike(makeHavingAbstractClause(), columnName, value))
}

func (query *Query) HavingNotLike(columnName string, value string) *Query {
	return query.addHavingClause(makeNotLike(makeHavingAbstractClause(), columnName, value))
}

func (query *Query) HavingStartWith(columnName string, value string) *Query {
	return query.addHavingClause(makeStartWith(makeHavingAbstractClause(), columnName, value))
}

func (query *Query) HavingNotStartWith(columnName string, value string) *Query {
	return query.addHavingClause(makeNotStartWith(makeHavingAbstractClause(), columnName, value))
}

func (query *Query) HavingEndWith(columnName string, value string) *Query {
	return query.addHavingClause(makeEndWith(makeHavingAbstractClause(), columnName, value))
}

func (query *Query) HavingNotEndWith(columnName string, value string) *Query {
	return query.addHavingClause(makeNotEndWith(makeHavingAbstractClause(), columnName, value))
}

func (query *Query) HavingTrue(columnName string) *Query {
	return query.addHavingClause(makeTrue(makeHavingAbstractClause(), columnName))
}

func (query *Query) HavingNotTrue(columnName string) *Query {
	return query.HavingFalse(columnName)
}

func (query *Query) HavingFalse(columnName string) *Query {
	return query.addHavingClause(makeFalse(makeHavingAbstractClause(), columnName))
}

func (query *Query) HavingNotFalse(columnName string) *Query {
	return query.HavingTrue(columnName)
}

func (query *Query) HavingNull(columnName string) *Query {
	return query.addHavingClause(makeNull(makeHavingAbstractClause(), columnName))
}

func (query *Query) HavingNotNull(columnName string) *Query {
	return query.addHavingClause(makeNotNull(makeHavingAbstractClause(), columnName))
}

func (query *Query) HavingColumns(columnName string, columnName2 string, operator Operator) *Query {
	if operator == "" {
		operator = Eq
	}
	return query.addHavingClause(makeColumnClause(makeHavingAbstractClause(), columnName, columnName2, operator, false))
}

func (query *Query) HavingIn(columnName string, params ...interface{}) *Query {
	return query.addHavingClause(makeInClause(makeHavingAbstractClause(), columnName, params...))
}

func (query *Query) HavingNotIn(columnName string, params ...interface{}) *Query {
	return query.addHavingClause(makeNotInClause(makeHavingAbstractClause(), columnName, params...))
}

func (query *Query) HavingRaw(sqlCondition string, params ...interface{}) *Query {
	return query.addHavingClause(makeRawSqlClause(makeHavingAbstractClause(), sqlCondition, params...))
}

func (query *Query) HavingOr(callback func(*Query) *Query) *Query {
	orClause := callback(NewQuery())
	return query.addHavingClause(makeOrClause(makeHavingAbstractClause(), orClause.whereClauses))
}

func (query *Query) HavingAnd(callback func(*Query) *Query) *Query {
	orClause := callback(NewQuery())
	return query.addHavingClause(
		makeAndClause(
			makeHavingAbstractClause(),
			orClause.whereClauses,
		))
}

func (query *Query) HavingQuery(columnName string, operator Operator, subQuery *Query, subQueryAlias string) *Query {
	return query.addHavingClause(makeSubQueryConditionClause(makeHavingAbstractClause(), columnName, operator, *subQuery, subQueryAlias, false))
}

func (query *Query) HavingNotQuery(columnName string, operator Operator, subQuery *Query, subQueryAlias string) *Query {
	return query.addHavingClause(makeSubQueryConditionClause(makeHavingAbstractClause(), columnName, operator, *subQuery, subQueryAlias, true))
}
