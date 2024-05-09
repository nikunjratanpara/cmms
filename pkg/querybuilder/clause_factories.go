package querybuilder

import "github.com/nikunjratanpara/cmms/pkg/linq"

const (
	AsType        = "As"
	FromType      = "From"
	DistinctType  = "Distinct"
	UnionType     = "Union"
	UnionAllType  = "UnionAll"
	BetweenType   = "Between"
	ContainsType  = "Contains"
	StartWithType = "StartWith"
	EndWithType   = "EndWith"
	LikeType      = "Like"
	AndType       = "And"
	OrType        = "Or"
	InType        = "In"
	NullType      = "Null"
	BooleanType   = "IsBoolean"
	ColumnsType   = "Columns"
	SubQueryType  = "SubQuery"
	RawSqlType    = "RawSql"

	JoinType = "Join"
	OnType   = "On"

	GroupByType = "GroupBy"
	OrderByType = "OrderBy"
	LimitType   = "Limit"
	OffsetType  = "Offset"

	FromComponent       = "From"
	WhereComponent      = "Where"
	GroupByComponent    = "GroupBy"
	HavingComponent     = "Having"
	OrderByComponent    = "OrderBy"
	SelectComponent     = "Select"
	PaginationComponent = "Pagination"
	OnComponent         = "On"
	JoinComponent       = "Join"
	UnionComponent      = "Union"
	SetComponent        = "Set"
	InsertComponent     = "Insert"
	UpdateComponent     = "Update"
	DeleteComponent     = "Delete"

	SqlServerEngineScope = "SqlServer"
	PostgresEngineScope  = "Postgres"
	OracleEngineScope    = "Oracle"
	MySqlEngineScope     = "MySql"
	SqliteEngineScope    = "Sqlite"
)

func makeAbstractClause(optionFuncs ...OptionsBuilder[AbstractClause]) *AbstractClause {
	abstractClause := &AbstractClause{}
	linq.ForEach(optionFuncs, func(optionFunc OptionsBuilder[AbstractClause]) {
		optionFunc(abstractClause)
	})
	return abstractClause
}

func makeAbstractColumnClause(abstractClause *AbstractClause, columnName string,
	isNot bool) *AbstractColumnClause {
	return &AbstractColumnClause{
		AbstractClause: abstractClause,
		ColumnName:     columnName,
		IsNot:          isNot,
	}
}

func makeCoditionClause(abstractColumnClause *AbstractColumnClause, value interface{}, operator Operator) *ConditionClause {
	return &ConditionClause{
		AbstractColumnClause: abstractColumnClause,
		Value:                value,
		Operator:             operator,
	}
}

func withComponent(component string) OptionsBuilder[AbstractClause] {
	return func(abstractClause *AbstractClause) *AbstractClause {
		abstractClause.Component = component
		return abstractClause
	}
}

func isComponent(component string) IsClause {
	return func(clause Clause) bool { return clause.GetComponent() == component }
}

func withEngineScope(engineScope string) OptionsBuilder[AbstractClause] {
	optionsBuilder := func(abstractClause *AbstractClause) *AbstractClause {
		abstractClause.EngineScope = engineScope
		return abstractClause
	}

	return optionsBuilder
}

func isEngineScope(engineScope string) IsClause {
	return func(clause Clause) bool {
		return clause.GetEngineScope() == "" || clause.GetEngineScope() == engineScope
	}
}

func withType(t string) OptionsBuilder[AbstractClause] {
	return func(abstractClause *AbstractClause) *AbstractClause {
		abstractClause.Type = t
		return abstractClause
	}
}
func isType(t string) IsClause {
	return func(clause Clause) bool { return clause.GetType() == t }
}

type OptionsBuilder[T interface{}] func(*T) *T
type IsClause func(clause Clause) bool
type ClauseFactory[T Clause, R interface{}] func() *T

// ---------------------------------Engine Scopes -------------------------
var isSqlServerEngineScope = isEngineScope(SqlServerEngineScope)
var isOracleEngineScope = isEngineScope(OracleEngineScope)

// --------------------------------Types-------------------------------------
var withUnionType = withType(UnionType)
var withUnionAllType = withType(UnionAllType)

var withAsType = withType(AsType)
var isAsType = isType(AsType)

var withFromType = withType(FromType)
var isFromType = isType(FromType)

var withJoinType = withType(JoinType)
var withOnType = withType(OnType)

var withLimitType = withType(LimitType)
var isLimitType = isType(LimitType)

var withOffsetType = withType(OffsetType)
var isOffsetType = isType(OffsetType)

var withOrderByType = withType(OrderByType)

var withGroupByType = withType(GroupByType)

var withAndType = withType(AndType)
var withOrType = withType(OrType)

var withNullType = withType(NullType)

var withInType = withType(InType)

var withContainsType = withType(ContainsType)
var withStartWithType = withType(StartWithType)
var withEndWithType = withType(EndWithType)
var withLikeType = withType(LikeType)

var withBetweenType = withType(BetweenType)

var withSubQueryType = withType(SubQueryType)
var isSubQueryType = isType(SubQueryType)

var withRawSqlType = withType(RawSqlType)
var isRawSqlType = isType(RawSqlType)

var withIsBooleanType = withType(BooleanType)
var withColumnsType = withType(ColumnsType)

//-----------------------------Components--------------------------------

var withSelectComponent = withComponent(SelectComponent)
var isSelectComponent = isComponent(SelectComponent)

var withFromComponent = withComponent(FromComponent)
var isFromComponent IsClause = isComponent(FromComponent)

var withWhereComponent = withComponent(WhereComponent)
var isWhereComponent = isComponent(WhereComponent)

var withHavingComponent = withComponent(HavingComponent)
var isHavingComponent = isComponent(HavingComponent)

var withGroupByComponent = withComponent(GroupByComponent)

var withOrderByComponent = withComponent(OrderByComponent)
var isOrderByComponent = isComponent(OrderByComponent)

var withPaginationComponent = withComponent(PaginationComponent)
var isPaginationComponent = isComponent(PaginationComponent)

var withJoinComponent = withComponent(JoinComponent)
var isJoinComponent = isComponent(JoinComponent)

var withUnionComponent = withComponent(UnionComponent)
var isUnionComponent = isComponent(UnionComponent)
