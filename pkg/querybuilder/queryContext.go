package querybuilder

import (
	"strings"

	"github.com/nikunjratanpara/cmms/pkg/linq"
)

type QueryContext struct {
	Compiler
	AddParameter ParameterAccumulator
}

func (ctx QueryContext) compileWhere(query Query) string {
	return ctx.Compiler.compileWhere(query, ctx)
}

func (ctx QueryContext) compileSelect(query Query) string {
	return ctx.Compiler.compileSelect(query, ctx)
}

func (ctx QueryContext) compileFrom(query Query) string {
	return ctx.Compiler.compileFrom(query, ctx)
}

func (ctx QueryContext) compileGroupBy(query Query) string {
	return ctx.Compiler.compileGroupBy(query, ctx)
}

func (ctx QueryContext) compileHaving(query Query) string {
	return ctx.Compiler.compileHaving(query, ctx)
}

func (ctx QueryContext) compileOrderBy(query Query) string {
	return ctx.Compiler.compileOrderBy(query, ctx)
}

func (ctx QueryContext) compilePagination(query Query) string {
	return ctx.Compiler.compilePagination(query, ctx)
}

func (ctx QueryContext) compileSubQuery(query Query) string {
	return ctx.Compiler.compileSubQuery(query, ctx)
}

func (ctx QueryContext) compileSelectQuery(query Query) string {
	return ctx.Compiler.compileSelectQuery(query, ctx)
}

func (ctx QueryContext) compileJoin(query Query) string {
	return ctx.Compiler.compileJoin(query, ctx)
}

func (ctx QueryContext) compileUnion(query Query) string {
	return ctx.Compiler.compileUnion(query, ctx)
}

// compileClauses generates the SQL query for a list of clauses based on the given separator and conditions.
//
// Parameters:
// - separator: The separator string used to concatenate the SQL queries of each clause.
// - clauses: The list of clauses to be compiled.
// - isClauses: Optional additional conditions for filtering the clauses.
//
// Returns:
// - string: The generated SQL query for the compiled clauses.
func (ctx QueryContext) compileClauses(separator string, clauses []Clause, isClauses ...IsClause) string {
	if len(clauses) == 0 {
		return ""
	}

	isEngineScopeMatch := isEngineScope(ctx.GetEngineScope())

	clauses = linq.Filter(
		clauses,
		func(clause Clause) bool {
			return clause != nil &&
				isEngineScopeMatch(clause) &&
				linq.AllMatch(isClauses, func(c IsClause) bool { return c(clause) })
		},
	)

	if len(clauses) == 1 {
		return clauses[0].GetSql(ctx)
	}

	var buf strings.Builder
	for _, clause := range clauses {
		buf.WriteString(clause.GetSql(ctx))
		buf.WriteString(separator)
	}

	return strings.TrimSuffix(buf.String(), separator)
}
