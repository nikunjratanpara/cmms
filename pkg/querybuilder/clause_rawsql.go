package querybuilder

import (
	"errors"
	"strings"
)

type RawSqlClause struct {
	*AbstractClause
	Sql    string
	Params []interface{}
}

func makeRawSqlClause(abstractClause *AbstractClause, sql string, params ...interface{}) *RawSqlClause {
	if sql == "" {
		panic("sql is empty")
	}

	return &RawSqlClause{
		AbstractClause: withRawSqlType(abstractClause),
		Sql:            sql,
		Params:         params,
	}
}

func (clause RawSqlClause) GetSql(context QueryContext) string {
	sql := clause.Sql
	sql = strings.ReplaceAll(sql, "[", context.GetOpeningIdentifier())
	sql = strings.ReplaceAll(sql, "]", context.GetClosingIdentifier())
	paramsCount := strings.Count(sql, "?")
	assertParametersMatch(paramsCount, clause)
	for _, value := range clause.Params {
		sql = strings.Replace(sql, "?", context.AddParameter(value), 1)
	}
	return sql
}

func assertParametersMatch(paramsCount int, clause RawSqlClause) {
	if paramsCount != len(clause.Params) {
		panic(errors.New("mismatch parameters and parameter palceholders"))
	}
}
