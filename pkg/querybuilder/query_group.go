package querybuilder

func (query *Query) GroupBy(columns ...string) *Query {
	return query.addGroupByClause(makeGroupByClause(columns...))
}

func (query *Query) GroupByRaw(groupByRaw string) *Query {
	return query.addGroupByClause(makeRawSqlClause(makeAbstractClause(withGroupByComponent), groupByRaw))
}
