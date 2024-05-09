package querybuilder

func makeJoinAbstractClause() *AbstractClause {
	return makeAbstractClause(withJoinComponent)
}

func (query *Query) InnerJoin(tableName string, column string, column1 string, operator Operator) *Query {
	return query.addJoin(tableName, InnerJoinType, column, column1, operator)
}

func (query *Query) InnerJoinWithConditions(tableName string, callback func(*Query) *Query) *Query {
	return query.addJoinWithConditions(InnerJoinType, tableName, callback)
}

func (query *Query) InnerJoinWithQuery(subQuery Query, column string, column1 string, operator Operator) *Query {
	return query.addSubQueryJoin(subQuery, InnerJoinType, column, column1, operator)
}

func (query *Query) InnerJoinWithQueryAndConditions(subQuery Query, callback func(*Query) *Query) *Query {
	return query.addSubQueryJoinWithConditions(InnerJoinType, subQuery, callback)
}

func (query *Query) LeftJoin(tableName string, column string, column1 string, operator Operator) *Query {
	return query.addJoin(tableName, LeftJoinType, column, column1, operator)
}

func (query *Query) LeftJoinWithConditions(tableName string, callback func(*Query) *Query) *Query {
	return query.addJoinWithConditions(LeftJoinType, tableName, callback)
}

func (query *Query) LeftJoinWithQuery(subQuery Query, column string, column1 string, operator Operator) *Query {
	return query.addSubQueryJoin(subQuery, LeftJoinType, column, column1, operator)
}

func (query *Query) LeftJoinWithQueryAndConditions(subQuery Query, callback func(*Query) *Query) *Query {
	return query.addSubQueryJoinWithConditions(LeftJoinType, subQuery, callback)
}

func (query *Query) RightJoin(tableName string, column string, column1 string, operator Operator) *Query {
	return query.addJoin(tableName, RightJoinType, column, column1, operator)
}

func (query *Query) RightJoinWithConditions(tableName string, callback func(*Query) *Query) *Query {
	return query.addJoinWithConditions(RightJoinType, tableName, callback)
}

func (query *Query) RightJoinWithQueryAndConditions(subQuery Query, callback func(*Query) *Query) *Query {
	return query.addSubQueryJoinWithConditions(RightJoinType, subQuery, callback)
}

func (query *Query) RightJoinWithQuery(subQuery Query, column string, column1 string, operator Operator) *Query {
	return query.addSubQueryJoin(subQuery, RightJoinType, column, column1, operator)
}

func (query *Query) CrossJoinWithQuery(subQuery Query) *Query {
	return query.addClause(makeJoinQueryClause(
		makeJoinAbstractClause(),
		CrossJoinType,
		subQuery,
	))
}

func (query *Query) CrossJoin(tableName string) *Query {
	return query.addClause(
		makeJoinClause(
			makeJoinAbstractClause(),
			CrossJoinType,
			tableName,
		))
}

func (query *Query) addJoinWithConditions(joinType string, tableName string, callback func(*Query) *Query) *Query {
	queryWithConditions := callback(NewQuery())
	return query.addClause(makeJoinClause(makeJoinAbstractClause(), joinType, tableName, queryWithConditions.whereClauses...))
}

func (query *Query) addSubQueryJoinWithConditions(joinType string, subQuery Query, callback func(*Query) *Query) *Query {
	queryWithConditions := callback(NewQuery())
	return query.addClause(makeJoinQueryClause(makeJoinAbstractClause(), joinType, subQuery, queryWithConditions.whereClauses...))
}

func (query *Query) addSubQueryJoin(subQueryWithAlias Query, joinType string, column string, column1 string, operator Operator) *Query {
	columnClause := makeColumnClause(
		makeJoinAbstractClause(),
		column,
		column1,
		operator,
		false)
	return query.addClause(makeJoinQueryClause(
		makeJoinAbstractClause(),
		joinType,
		subQueryWithAlias,
		columnClause,
	))
}

func (query *Query) addJoin(tableName string, joinType string, column string, column1 string, operator Operator) *Query {
	columnClause := makeColumnClause(
		makeJoinAbstractClause(),
		column,
		column1,
		operator,
		false)

	return query.addClause(
		makeJoinClause(
			makeJoinAbstractClause(),
			joinType,
			tableName,
			columnClause,
		))
}
