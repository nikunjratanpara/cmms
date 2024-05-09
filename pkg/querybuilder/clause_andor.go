package querybuilder

type OrClause struct {
	*AbstractClause
	Clauses []Clause
}

type AndClause struct {
	*AbstractClause
	Clauses []Clause
}

func makeAndClause(abstractClause *AbstractClause, clauses []Clause) *AndClause {
	return &AndClause{
		AbstractClause: withAndType(abstractClause),
		Clauses:        clauses,
	}
}

func makeOrClause(abstractClause *AbstractClause, clauses []Clause) *OrClause {
	return &OrClause{
		AbstractClause: withOrType(abstractClause),
		Clauses:        clauses,
	}
}

func (clause OrClause) GetSql(context QueryContext) string {
	sql := context.compileClauses(" OR ", clause.Clauses)
	return " ( " + sql + " ) "
}

func (clause AndClause) GetSql(context QueryContext) string {
	sql := context.compileClauses(" AND ", clause.Clauses)
	return " ( " + sql + " ) "
}
