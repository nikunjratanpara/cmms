package querybuilder

type FromClause struct {
	*AbstractClause
	TableName string
}

func makeFromClause(tableName string) *FromClause {
	return &FromClause{
		AbstractClause: makeAbstractClause(withFromComponent, withFromType),
		TableName:      tableName,
	}
}

func (clause FromClause) GetSql(context QueryContext) string {
	return context.prepareTableName(clause.TableName)
}
