package querybuilder

func makeUnionAbstractClause() *AbstractClause {
	return makeAbstractClause(withUnionComponent)
}

func (query *Query) Union(unionQuery *Query) *Query {
	return query.addClause(makeUnionClause(makeUnionAbstractClause(), false, *unionQuery))
}

func (query *Query) UnionAll(unionQuery *Query) *Query {
	return query.addClause(makeUnionClause(makeUnionAbstractClause(), true, *unionQuery))
}
