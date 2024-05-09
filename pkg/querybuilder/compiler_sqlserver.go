package querybuilder

import (
	"fmt"
	"strings"
)

/*
.WhereDate("CreatedAt", "2018-04-01")
.WhereTime("CreatedAt", ">", "16:30")
.WhereDatePart("month", "CreatedAt", 2);


  CAST([CreatedAt] AS DATE) = '2018-04-01'
  AND CAST([CreatedAt] AS TIME) > '16:30'
  AND DATEPART(MONTH, [CreatedAt]) = 2
*/

type SqlServerCompiler struct {
	*BaseCompiler
}

func makeSqlServerCompiler() BaseCompiler {
	compiler := &SqlServerCompiler{
		BaseCompiler: &BaseCompiler{
			openingIdentifier: "[",
			closingIdentifier: "]",
			defaultNamespace:  "dbo",
			booleanAsBit:      true,
			lastId:            "scope_identity() as [Id]",
			columnAsKeyword:   " AS ",
			tableAsKeyword:    " AS ",
			engineName:        SqlServerEngineScope,
		},
	}
	compiler.Compiler = compiler
	return BaseCompiler{Compiler: compiler}
}

func (compiler SqlServerCompiler) CompileTrue() string {
	return "cast(1 as bit)"
}

func (compiler SqlServerCompiler) CompileFalse() string {
	return "cast(0 as bit)"
}

func (compiler SqlServerCompiler) CompileDatePart(part string, column string) string {
	return fmt.Sprintf("DATEPART(%s, %s)", part, column)
}

func (compiler SqlServerCompiler) CompileDate(column string) string {
	return fmt.Sprintf("CAST(%s AS DATE)", column)
}

func (compiler SqlServerCompiler) CompileTime(column string) string {
	return fmt.Sprintf("CAST(%s AS TIME)", column)
}

func (compiler SqlServerCompiler) CompileDateTime(column string) string {
	return fmt.Sprintf("CAST(%s AS DATETIME)", column)
}

/*
 SqlServer -
 All three defined - ORDER BY [Name] OFFSET 30 ROWS FETCH NEXT 10 ROWS ONLY
 Offset only - ORDER BY (SELECT NULL) OFFSET 10 ROWS
 Limit only - SELECT TOP(10) * FROM
 Orderby not defined - ORDER BY (SELECT NULL) OFFSET 10 ROWS FETCH NEXT 10 ROWS ONLY

  SqlServer (legacy) - SELECT
  *
FROM
  (
    SELECT
      *,
      ROW_NUMBER() OVER (
        ORDER BY
          [Name]
      ) AS [row_num]
    FROM
      [Users]
    WHERE
      [Name] = 'John'
      AND (
        [Country] = 'FR'
        OR [Country] = 'UK'
      )
  ) AS [results_wrapper]
WHERE
  [row_num] BETWEEN 31
  AND 40
*/

func (compiler SqlServerCompiler) CompileLimit(limit int) string {
	return fmt.Sprintf(" TOP(%d)", limit)
}

func (compiler SqlServerCompiler) compileSelect(query Query, ctx QueryContext) string {
	limitClause := GetOneComponent[*LimitClause](query.clauses, isLimitType)
	offsetClause := GetOneComponent[*OffsetClause](query.clauses, isOffsetType)

	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("SELECT")

	if query.isDistinct {
		stringBuilder.WriteString(" DISTINCT")
	}

	if limitClause != nil && offsetClause == nil {
		stringBuilder.WriteString(ctx.CompileLimit((*limitClause).Count))
	}

	selectClauses := ctx.compileClauses(", ", query.selectClauses, isSelectComponent)
	if selectClauses == "" {
		selectClauses = " *"
	} else {
		stringBuilder.WriteString(" ")
	}

	stringBuilder.WriteString(selectClauses)
	return stringBuilder.String()
}

func (compiler SqlServerCompiler) compilePagination(query Query, ctx QueryContext) string {
	limitClause := GetOneComponent[*LimitClause](query.clauses, isLimitType, isPaginationComponent, isSqlServerEngineScope)
	offsetClause := GetOneComponent[*OffsetClause](query.clauses, isOffsetType, isPaginationComponent, isSqlServerEngineScope)

	if offsetClause == nil {
		//exit from the function if offset is not defined
		//limit will be handled by select clause for sql server
		return ""
	}

	stringBuilder := strings.Builder{}
	//Handle by order by
	orderbyClause := ctx.compileOrderBy(query)
	if orderbyClause == "" {
		stringBuilder.WriteString(" ORDER BY (SELECT NULL)")
	}

	//limit and offset both defined
	if limitClause != nil {
		stringBuilder.WriteString(fmt.Sprintf(" OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", (*offsetClause).Count, (*limitClause).Count))
		return stringBuilder.String()
	} else {
		//only offset defined
		stringBuilder.WriteString(fmt.Sprintf(" OFFSET %d ROWS", (*offsetClause).Count))
	}
	return stringBuilder.String()
}
