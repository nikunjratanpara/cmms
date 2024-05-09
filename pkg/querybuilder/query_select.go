package querybuilder

func (query *Query) Select(columns ...string) *Query {
	return query.addSelectClause(makeSelectClause(columns...))
}

func (query *Query) SelectFromSubQuery(subQuery Query, alias string) *Query {
	return query.addSelectClause(makeSelectClauseFromSubQuery(subQuery, alias))
}

func (query *Query) SelectRaw(rawSql string) *Query {
	return query.addSelectClause(makeSelectRawSqlClause(rawSql))
}
