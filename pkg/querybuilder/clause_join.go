package querybuilder

const (
	InnerJoinType = "INNER"
	LeftJoinType  = "LEFT"
	RightJoinType = "RIGHT"
	CrossJoinType = "CROSS"
)

// Join clauses
type JoinClause struct {
	*AbstractClause
	JoinType       string
	TableName      string
	JoinConditions []Clause
}

type JoinQueryClause struct {
	*AbstractClause
	JoinType       string
	Query          Query
	JoinConditions []Clause
}

func makeJoinClause(
	abstractClause *AbstractClause,
	joinType string,
	tableName string,
	joinConditions ...Clause,
) *JoinClause {
	return &JoinClause{
		AbstractClause: withJoinComponent(abstractClause),
		TableName:      tableName,
		JoinType:       joinType,
		JoinConditions: joinConditions,
	}
}

func makeJoinQueryClause(
	abstractClause *AbstractClause,
	joinType string,
	query Query,
	joinConditions ...Clause,
) *JoinQueryClause {
	return &JoinQueryClause{
		AbstractClause: withJoinType(abstractClause),
		JoinType:       joinType,
		Query:          query,
		JoinConditions: joinConditions,
	}
}

func (clause *JoinQueryClause) GetSql(ctx QueryContext) string {
	subQuerySql := ctx.compileSubQuery(clause.Query)
	joinConditions := ctx.compileClauses(" AND ", clause.JoinConditions)
	if joinConditions != "" {
		joinConditions = " ON " + joinConditions
	}
	return clause.JoinType + " JOIN " + subQuerySql + joinConditions
}

func (clause *JoinClause) GetSql(ctx QueryContext) string {
	conditions := ctx.compileClauses(" AND ", clause.JoinConditions)
	if conditions != "" {
		conditions = " ON " + conditions
	}
	return clause.JoinType + " JOIN " + ctx.prepareTableName(clause.TableName) + conditions
}
