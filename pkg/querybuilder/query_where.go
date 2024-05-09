package querybuilder

func makeWhereAbstractClause(optionsBuilder ...OptionsBuilder[AbstractClause]) *AbstractClause {
	return makeAbstractClause(append(optionsBuilder, withWhereComponent)...)
}

func makeWhereAbstractColumnClause(columnName string, isNot bool) *AbstractColumnClause {
	return makeAbstractColumnClause(makeWhereAbstractClause(), columnName, isNot)
}

func (query *Query) Where(columnName string, value interface{}, operator Operator) *Query {
	if operator == "" {
		operator = Eq
	}
	return query.addWhereClause(makeCoditionClause(makeWhereAbstractColumnClause(columnName, false), value, operator))
}

func (query *Query) NotWhere(columnName string, value interface{}, operator Operator) *Query {
	if operator == "" {
		operator = Eq
	}
	return query.addWhereClause(makeCoditionClause(makeWhereAbstractColumnClause(columnName, true), value, operator))
}

func (query *Query) WhereBetween(columnName string, fromValue interface{}, toValue interface{}) *Query {
	return query.addWhereClause(makeBetween(makeWhereAbstractClause(), columnName, fromValue, toValue))
}

func (query *Query) WhereNotBetween(columnName string, fromValue interface{}, toValue interface{}) *Query {
	return query.addWhereClause(makeNotBetween(makeWhereAbstractClause(), columnName, fromValue, toValue))
}

func (query *Query) WhereContains(columnName string, value string) *Query {
	return query.addWhereClause(makeContains(makeWhereAbstractClause(), columnName, value))
}

func (query *Query) WhereNotContains(columnName string, value string) *Query {
	return query.addWhereClause(makeNotContains(makeWhereAbstractClause(), columnName, value))
}

func (query *Query) WhereLike(columnName string, value string) *Query {
	return query.addWhereClause(makeLike(makeWhereAbstractClause(), columnName, value))
}

func (query *Query) WhereNotLike(columnName string, value string) *Query {
	return query.addWhereClause(makeNotLike(makeWhereAbstractClause(), columnName, value))
}

func (query *Query) WhereStartWith(columnName string, value string) *Query {
	return query.addWhereClause(makeStartWith(makeWhereAbstractClause(), columnName, value))
}

func (query *Query) WhereNotStartWith(columnName string, value string) *Query {
	return query.addWhereClause(makeNotStartWith(makeWhereAbstractClause(), columnName, value))
}

func (query *Query) WhereEndWith(columnName string, value string) *Query {
	return query.addWhereClause(makeEndWith(makeWhereAbstractClause(), columnName, value))
}

func (query *Query) WhereNotEndWith(columnName string, value string) *Query {
	return query.addWhereClause(makeNotEndWith(makeWhereAbstractClause(), columnName, value))
}

func (query *Query) WhereTrue(columnName string) *Query {
	return query.addWhereClause(makeTrue(makeWhereAbstractClause(), columnName))
}

func (query *Query) WhereNotTrue(columnName string) *Query {
	return query.WhereFalse(columnName)
}

func (query *Query) WhereFalse(columnName string) *Query {
	return query.addWhereClause(makeFalse(makeWhereAbstractClause(), columnName))
}

func (query *Query) WhereNotFalse(columnName string) *Query {
	return query.WhereTrue(columnName)
}

func (query *Query) WhereNull(columnName string) *Query {
	return query.addWhereClause(makeNull(makeWhereAbstractClause(), columnName))
}

func (query *Query) WhereNotNull(columnName string) *Query {
	return query.addWhereClause(makeNotNull(makeWhereAbstractClause(), columnName))
}

func (query *Query) WhereColumns(columnName string, columnName2 string, operator Operator) *Query {
	if operator == "" {
		operator = Eq
	}
	return query.addWhereClause(makeColumnClause(makeWhereAbstractClause(), columnName, columnName2, operator, false))
}

func (query *Query) WhereIn(columnName string, params ...interface{}) *Query {
	return query.addWhereClause(makeInClause(makeWhereAbstractClause(), columnName, params...))
}

func (query *Query) WhereNotIn(columnName string, params ...interface{}) *Query {
	return query.addWhereClause(makeNotInClause(makeWhereAbstractClause(), columnName, params...))
}

func (query *Query) WhereRaw(sqlCondition string, params ...interface{}) *Query {
	return query.addWhereClause(makeRawSqlClause(makeWhereAbstractClause(), sqlCondition, params...))
}

func (query *Query) WhereOr(callback func(*Query) *Query) *Query {
	orClause := callback(NewQuery())
	return query.addWhereClause(makeOrClause(makeWhereAbstractClause(), orClause.whereClauses))
}

func (query *Query) WhereAnd(callback func(*Query) *Query) *Query {
	orClause := callback(NewQuery())
	return query.addWhereClause(
		makeAndClause(
			makeWhereAbstractClause(),
			orClause.whereClauses,
		))
}

func (query *Query) WhereQuery(columnName string, operator Operator, subQuery *Query, subQueryAlias string) *Query {
	return query.addWhereClause(makeSubQueryConditionClause(makeWhereAbstractClause(), columnName, operator, *subQuery, subQueryAlias, false))
}

func (query *Query) WhereNotQuery(columnName string, operator Operator, subQuery *Query, subQueryAlias string) *Query {
	return query.addWhereClause(makeSubQueryConditionClause(makeWhereAbstractClause(), columnName, operator, *subQuery, subQueryAlias, true))
}
