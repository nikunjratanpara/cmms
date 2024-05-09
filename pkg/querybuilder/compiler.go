package querybuilder

import (
	"fmt"
	"strings"
)

type Compiler interface {
	Compile(Query) QueryResult
	CompileTrue() string
	CompileFalse() string
	CompileLimit(int) string
	CompileOffset(int) string

	GetOpeningIdentifier() string
	GetClosingIdentifier() string
	GetDefaultNamespace() string
	GetBooleanAsBit() bool
	GetLastId() string
	GetColumnAsKeyword() string
	GetTableAsKeyword() string
	GetEngineScope() string

	prepareIdentifier(string) string
	prepareTableName(string) string
	wrapToken(string) string

	compileWhere(Query, QueryContext) string
	compileSelect(Query, QueryContext) string
	compileFrom(Query, QueryContext) string
	compileJoin(Query, QueryContext) string
	compileGroupBy(Query, QueryContext) string
	compileHaving(Query, QueryContext) string
	compileOrderBy(Query, QueryContext) string
	compilePagination(Query, QueryContext) string
	compileUnion(Query, QueryContext) string

	compileSubQuery(Query, QueryContext) string
	compileSelectQuery(Query, QueryContext) string
}

type QueryResult struct {
	Sql        string
	Parameters map[int]interface{}
	query      *Query
}

type BaseCompiler struct {
	openingIdentifier string
	closingIdentifier string
	defaultNamespace  string
	booleanAsBit      bool
	lastId            string
	columnAsKeyword   string
	tableAsKeyword    string
	Compiler          Compiler
	engineName        string
}

// compileSubQuery compiles a subquery.
//
// It takes a query and a query context as parameters and returns a string.
func (compiler BaseCompiler) compileSubQuery(query Query, ctx QueryContext) string {
	queryStr := ctx.compileSelectQuery(query)
	alias := ctx.compileClauses("", query.clauses, isAsType, isFromComponent)
	if alias != "" {
		return "(" + queryStr + ") " + alias
	}
	return queryStr
}

// compileSelectQuery compiles a select query based on the provided query and context.
//
// Parameters:
// - query: the query to be compiled
// - ctx: the context for compiling the query
// Return type: string
func (compiler BaseCompiler) compileSelectQuery(query Query, ctx QueryContext) string {
	queryBuilder := strings.Builder{}

	queryBuilder.WriteString(ctx.compileSelect(query))
	queryBuilder.WriteString(ctx.compileFrom(query))
	queryBuilder.WriteString(ctx.compileJoin(query))
	queryBuilder.WriteString(ctx.compileWhere(query))
	queryBuilder.WriteString(ctx.compileGroupBy(query))
	queryBuilder.WriteString(ctx.compileHaving(query))
	queryBuilder.WriteString(ctx.compileOrderBy(query))
	queryBuilder.WriteString(ctx.compilePagination(query))
	queryBuilder.WriteString(ctx.compileUnion(query))
	return queryBuilder.String()
}

// TODO: Add function to handle dateparts for various databases

// GetEngineScope returns the engine scope of the BaseCompiler.
//
// It does not take any parameters.
// It returns a string.
func (compiler BaseCompiler) GetEngineScope() string {
	return compiler.engineName
}

func (compiler BaseCompiler) GetOpeningIdentifier() string {
	return compiler.openingIdentifier
}

func (compiler BaseCompiler) GetClosingIdentifier() string {
	return compiler.closingIdentifier
}

func (compiler BaseCompiler) GetDefaultNamespace() string {
	return compiler.defaultNamespace
}

func (compiler BaseCompiler) GetBooleanAsBit() bool {
	return compiler.booleanAsBit
}

func (compiler BaseCompiler) GetLastId() string {
	return compiler.lastId
}

// GetColumnAsKeyword returns the columnAsKeyword field of the BaseCompiler struct.
//
// This function does not take any parameters.
// It returns a string representing the columnAsKeyword field.
func (compiler BaseCompiler) GetColumnAsKeyword() string {
	return compiler.columnAsKeyword
}

// GetTableAsKeyword returns the tableAsKeyword field of the BaseCompiler struct.
//
// It does not take any parameters.
// It returns a string.
func (compiler BaseCompiler) GetTableAsKeyword() string {
	return compiler.tableAsKeyword
}

func (compiler BaseCompiler) GetCompiler() Compiler {
	return compiler.Compiler
}

// CompileTrue returns the string "true".
//
// This function does not take any parameters and returns a string.
func (compiler BaseCompiler) CompileTrue() string {
	return "true"
}

// CompileFalse generates the Go code for the "false" value.
//
// It does not take any parameters.
// It returns a string representing the Go code for the "false" value.
func (compiler BaseCompiler) CompileFalse() string {
	return "false"
}

// prepareIdentifier prepares the identifier for the given column name.
//
// It takes a string parameter `columnName` and returns a string in proper identifier format.
func (compiler BaseCompiler) prepareIdentifier(columnName string) string {
	columnName = compiler.prepareTableName(columnName)
	tokenizedColumnName := strings.ReplaceAll(columnName, ".", compiler.closingIdentifier+"."+compiler.openingIdentifier)
	return compiler.wrapToken(tokenizedColumnName)
}

func (compiler BaseCompiler) prepareTableName(tableName string) string {
	aliasIndex := strings.Index(strings.ToLower(tableName), " as ")
	if aliasIndex > -1 {
		table := strings.TrimSpace(tableName[0:aliasIndex])
		alias := strings.TrimSpace(tableName[aliasIndex+3:])
		return fmt.Sprintf(
			"%s%s%s",
			compiler.prepareIdentifier(table),
			compiler.GetTableAsKeyword(),
			compiler.prepareIdentifier(alias),
		)
	}
	from := compiler.prepareIdentifier(tableName)
	return from
}

// wrapToken is a Go function that takes a string token and returns a wrapped string.
//
// It takes a string token as a parameter and returns a string.
func (compiler BaseCompiler) wrapToken(token string) string {
	var buf strings.Builder
	buf.WriteString(compiler.openingIdentifier)
	buf.WriteString(token)
	buf.WriteString(compiler.closingIdentifier)
	return buf.String()
}

// Compile compiles the query using the BaseCompiler.
//
// query Query - the query to compile.
// Return QueryResult - the result of the compilation.
func (compiler BaseCompiler) Compile(query Query) QueryResult {
	var c Compiler
	c = compiler
	if compiler.Compiler != nil {
		c = compiler.Compiler
	}

	params := map[int]interface{}{}
	ctx := QueryContext{
		Compiler:     c,
		AddParameter: getParameterAccumlator(params),
	}
	sql := ctx.compileSelectQuery(query)
	return QueryResult{
		Sql:        sql,
		Parameters: params,
		query:      &query,
	}
}

// appendWhenExists appends the given clause to the SQL string when the SQL string is not empty.
//
// Parameters:
// - sqlStr: The SQL string to append the clause to.
// - clause: The clause to append to the SQL string.
//
// Return:
// - The resulting SQL string with the clause appended, or an empty string if the SQL string is empty.
func appendWhenExists(sqlStr string, clause string) string {
	if sqlStr != "" {
		return clause + sqlStr
	}
	return ""
}

type ParameterAccumulator func(value interface{}) string

// getParameterAccumlator returns a ParameterAccumulator function that takes a value and adds it to the params map,
// returning the parameter placeholder string in the format "$X", where X is the index of the added value in the map.
//
// params: a map of integers to interface{} values representing the parameters.
// Returns: a function that takes a value and returns the parameter placeholder string.
func getParameterAccumlator(params map[int]interface{}) ParameterAccumulator {
	return func(value interface{}) string {
		paramNo := len(params)
		params[paramNo] = value
		return fmt.Sprintf("$%d", paramNo)
	}
}

/*

pagination with Order by

Oracle -  ORDER BY "Name" OFFSET 30 ROWS FETCH NEXT 10 ROWS ONLY

  MySql - ORDER BY `Name` LIMIT 10 OFFSET 30

  Postgres - ORDER BY "Name" LIMIT 10 OFFSET 30

  Sqlite - ORDER BY "Name" LIMIT 10 OFFSET 30
*/

// CompileLimit generates a SQL LIMIT clause based on the given limit value.
//
// Parameters:
// - limit: an integer representing the maximum number of records to return.
//
// Returns:
// - a string containing the SQL LIMIT clause.
func (compiler BaseCompiler) CompileLimit(limit int) string {
	return fmt.Sprintf("LIMIT %d", limit)
}

// CompileOffset generates a string with the specified offset.
//
// offset: an integer representing the offset value.
// string: a string containing the compiled OFFSET statement.
func (compiler BaseCompiler) CompileOffset(offset int) string {
	return fmt.Sprintf("OFFSET %d", offset)
}

func (compiler BaseCompiler) compilePagination(query Query, ctx QueryContext) string {

	limitClause := GetOneComponent[*LimitClause](query.clauses, isLimitType, isPaginationComponent, isEngineScope(ctx.GetEngineScope()))
	offsetClause := GetOneComponent[*OffsetClause](query.clauses, isOffsetType, isPaginationComponent, isEngineScope(ctx.GetEngineScope()))

	stringBuilder := strings.Builder{}

	if limitClause != nil {
		stringBuilder.WriteString(fmt.Sprintf(" LIMIT %d", (*limitClause).Count))
	}
	if offsetClause != nil {
		stringBuilder.WriteString(fmt.Sprintf(" OFFSET %d", (*offsetClause).Count))
	}
	return stringBuilder.String()
}

// compileWhere compiles the WHERE clause of a query based on the given Query and QueryContext.
//
// Parameters:
// - query: The Query object containing the whereClauses to be compiled.
// - ctx: The QueryContext object used for compiling the whereClauses.
//
// Returns:
// - string: The compiled WHERE clause.
func (compiler BaseCompiler) compileWhere(query Query, ctx QueryContext) string {
	return appendWhenExists(ctx.compileClauses(" AND ", query.whereClauses, isWhereComponent), " WHERE ")
}

// compileSelect description of the Go function.
//
// It takes parameters query Query, ctx QueryContext and returns a string.
func (compiler BaseCompiler) compileSelect(query Query, ctx QueryContext) string {
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("SELECT ")
	if query.isDistinct {
		stringBuilder.WriteString("DISTINCT ")
	}

	selectClause := ctx.compileClauses(", ", query.selectClauses, isSelectComponent)
	if selectClause == "" {
		selectClause = "*"
	}
	stringBuilder.WriteString(selectClause)

	return stringBuilder.String()
}

// compileFrom compiles the FROM clause of a SQL query.
//
// It takes a Query object and a QueryContext object as parameters.
// It returns a string representing the compiled FROM clause.
func (compiler BaseCompiler) compileFrom(query Query, ctx QueryContext) string {
	isFrom := func(clause Clause) bool {
		return isFromType(clause) || isSubQueryType(clause) || isRawSqlType(clause)
	}
	return appendWhenExists(ctx.compileClauses(" ", query.clauses, isFromComponent, isFrom), " FROM ")
}

// compileJoin description of the Go function.
//
// It takes a Query and a QueryContext as parameters and returns a string.
func (compiler BaseCompiler) compileJoin(query Query, ctx QueryContext) string {
	return appendWhenExists(ctx.compileClauses(" ", query.clauses, isJoinComponent), " ")
}

// compileGroupBy description of the Go function.
//
// It takes query Query and ctx QueryContext as parameters.
// Returns a string.
func (compiler BaseCompiler) compileGroupBy(query Query, ctx QueryContext) string {
	return appendWhenExists(ctx.compileClauses(", ", query.groupByClauses), " GROUP BY ")
}

// compileHaving compiles the having clauses of the query.
//
// It takes a Query and a QueryContext as parameters and returns a string.
func (compiler BaseCompiler) compileHaving(query Query, ctx QueryContext) string {
	return appendWhenExists(ctx.compileClauses(" AND ", query.havingClauses, isHavingComponent), " HAVING ")
}

// compileOrderBy description of the Go function.
//
// The function takes a Query and a QueryContext as parameters and returns a string.
func (compiler BaseCompiler) compileOrderBy(query Query, ctx QueryContext) string {
	return appendWhenExists(ctx.compileClauses(", ", query.orderByClauses, isOrderByComponent), " ORDER BY ")
}

func (compiler BaseCompiler) compileUnion(query Query, ctx QueryContext) string {
	return appendWhenExists(ctx.compileClauses(" ", query.clauses, isUnionComponent), " ")
}
