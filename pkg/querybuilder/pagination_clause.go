package querybuilder

type LimitClause struct {
	*AbstractClause
	Count int
}

// Offset clause
type OffsetClause struct {
	*AbstractClause
	Count int
}

func makeLimitClause(abstractClause *AbstractClause, count int) *LimitClause {
	return &LimitClause{
		AbstractClause: withLimitType(abstractClause),
		Count:          count,
	}
}

func makeOffsetClause(abstractClause *AbstractClause, count int) *OffsetClause {
	return &OffsetClause{
		AbstractClause: withOffsetType(abstractClause),
		Count:          count,
	}
}

func (clause LimitClause) GetSql(context QueryContext) string {
	return context.CompileLimit(clause.Count)
}

func (clause OffsetClause) GetSql(context QueryContext) string {
	return context.CompileOffset(clause.Count)
}
