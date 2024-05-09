package querybuilder

import "strings"

type InConditionClause struct {
	*AbstractColumnClause
	Value []interface{}
}

func makeInClause(abstractClause *AbstractClause, columnName string, params ...interface{}) *InConditionClause {
	return &InConditionClause{
		AbstractColumnClause: makeAbstractColumnClause(withInType(abstractClause), columnName, false),
		Value:                params,
	}
}

func makeNotInClause(abstractClause *AbstractClause, columnName string, params ...interface{}) *InConditionClause {
	return &InConditionClause{
		AbstractColumnClause: makeAbstractColumnClause(withInType(abstractClause), columnName, true),
		Value:                params,
	}
}

func buildInClauseSql(context QueryContext, columnName string, values []interface{}, clause string) string {
	parametersName := []string{}
	for _, value := range values {
		parametersName = append(parametersName, context.AddParameter(value))
	}
	return context.prepareIdentifier(columnName) + " " + clause + " ( " + strings.Join(parametersName, ",") + " )"
}

func (clause InConditionClause) GetSql(context QueryContext) string {
	strClause := "IN"
	if clause.IsNot {
		strClause = "NOT IN"
	}
	return buildInClauseSql(context, clause.ColumnName, clause.Value, strClause)
}
