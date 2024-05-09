package querybuilder

func makePaginationAbstractClause() *AbstractClause {
	return makeAbstractClause(withPaginationComponent)
}

func (q *Query) Offset(offset int) *Query {
	return q.addClause(makeOffsetClause(makePaginationAbstractClause(), offset))
}

func (q *Query) Limit(pageSize int) *Query {
	return q.addClause(makeLimitClause(makePaginationAbstractClause(), pageSize))
}

func (q *Query) ForPage(pageNo int) *Query {
	return q.Limit(15).Offset(15 * (pageNo - 1))
}

func (q *Query) ForPageWithOffset(limit int, offset int) *Query {
	return q.Limit(limit).Offset(offset)
}

func (q *Query) Skip(offset int) *Query {
	return q.Offset(offset)
}

func (q *Query) Take(pageSize int) *Query {
	return q.Limit(pageSize)
}
