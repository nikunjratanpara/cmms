package querybuilder

import "github.com/nikunjratanpara/cmms/pkg/linq"

/*
Add DatePart Where conditions
Add Union Support
Add Aggregation Support
*/

type AggregateFunction string

const (
	CountFn AggregateFunction = "COUNT"
	SumFn   AggregateFunction = "SUM"
	AvgFn   AggregateFunction = "AVG"
	MinFn   AggregateFunction = "MIN"
	MaxFn   AggregateFunction = "MAX"
)

type AggregateClause struct {
	Fn      AggregateFunction
	Args    []string
	Columns []string
	Alias   string
}

func (a AggregateClause) Clause() {
	// Implement clause interface
}

//TODO: Add Support for DateParts

type DatePart string

const (
	YearDatePart   DatePart = "YEAR"
	MonthDatePart  DatePart = "MONTH"
	DayDatePart    DatePart = "DAY"
	HourDatePart   DatePart = "HOUR"
	MinuteDatePart DatePart = "MINUTE"
	SecondDatePart DatePart = "SECOND"
)

type DatePartClause struct {
	Part  DatePart
	Expr  string
	Alias string
}

func (d DatePartClause) Clause() {}

type Query struct {
	whereClauses   []Clause
	havingClauses  []Clause
	orderByClauses []Clause
	groupByClauses []Clause
	selectClauses  []Clause
	clauses        []Clause
	queryType      string
	isDistinct     bool
}

func NewQuery() *Query {
	return &Query{
		whereClauses:   []Clause{},
		havingClauses:  []Clause{},
		orderByClauses: []Clause{},
		groupByClauses: []Clause{},
		selectClauses:  []Clause{},
		clauses:        []Clause{},
		queryType:      "Select",
	}
}

func NewQueryFromTable(tableName string) *Query {
	return NewQuery().From(tableName)
}

func NewQueryFromSubQuery(fromQuery Query, alias string) *Query {
	return NewQuery().FromSubQuery(fromQuery, alias)
}

func NewQueryFromRawSql(rawSql string) *Query {
	return NewQuery().FromRaw(rawSql)
}

func (query *Query) FromRaw(rawSql string) *Query {
	query.addClause(makeRawSqlClause(makeAbstractClause(withFromComponent, withRawSqlType), rawSql))
	return query
}

func (query *Query) From(tableName string) *Query {
	query.addClause(makeFromClause(tableName))
	return query
}

func (query *Query) FromSubQuery(subQuery Query, alias string) *Query {
	subQuery.As(alias)
	query.addClause(makeSubQueryClause(makeAbstractClause(withFromComponent), subQuery))
	return query
}

func (query *Query) As(alias string) *Query {
	if alias == "" {
		return query
	}

	clauses := linq.Filter(query.clauses, func(clause Clause) bool {
		return isAsType(clause) && isFromComponent(clause)
	})

	if len(clauses) > 0 {
		asClause := clauses[0].(*AsClause)
		asClause.Alias = alias
		return query
	}
	return query.addClause(makeAsClause(makeAbstractClause(withFromComponent, withAsType), alias))
}

func (query *Query) GetAllClauses() []Clause {
	return append(
		append(
			append(query.whereClauses, query.havingClauses...), query.orderByClauses...), query.groupByClauses...)
}

func (query *Query) Distinct() *Query {
	query.isDistinct = true
	return query
}

func (query *Query) addWhereClause(clause Clause) *Query {
	query.whereClauses = append(query.whereClauses, clause)
	return query
}

func (query *Query) addHavingClause(clause Clause) *Query {
	query.havingClauses = append(query.havingClauses, clause)
	return query
}

func (query *Query) addOrderByClause(clause Clause) *Query {
	query.orderByClauses = append(query.orderByClauses, clause)
	return query
}

func (query *Query) addGroupByClause(clause Clause) *Query {
	query.groupByClauses = append(query.groupByClauses, clause)
	return query
}

func (query *Query) addSelectClause(clause Clause) *Query {
	query.selectClauses = append(query.selectClauses, clause)
	return query
}

func (query *Query) addClause(clause Clause) *Query {
	query.clauses = append(query.clauses, clause)
	return query
}
