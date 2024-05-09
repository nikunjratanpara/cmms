package querybuilder

func makeOrderByAbstractClause() *AbstractClause {
	return makeAbstractClause(withOrderByComponent)
}

func (query *Query) OrderBy(columns ...string) *Query {
	return query.addOrderByClause(makeOrderBy(makeOrderByAbstractClause(), false, columns...))
}

func (query *Query) OrderByDesc(columns ...string) *Query {
	return query.addOrderByClause(makeOrderBy(makeOrderByAbstractClause(), true, columns...))
}

func (query *Query) OrderByRaw(orderBy string) *Query {
	return query.addOrderByClause(makeRawSqlClause(makeOrderByAbstractClause(), orderBy))
}
