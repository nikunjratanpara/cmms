package querybuilder

import (
	"fmt"
	"strings"
)

/*



Oracle Datepart
 TO_CHAR("CreatedAt", 'YY-MM-DD') = TO_CHAR(TO_DATE('2018-04-01', 'YY-MM-DD'), 'YY-MM-DD')
  AND TO_CHAR("CreatedAt", 'HH24:MI:SS') > TO_CHAR(TO_DATE('16:30', 'HH24:MI'), 'HH24:MI:SS')
  AND EXTRACT(
    MONTH
    FROM
      "CreatedAt"
  ) = 2

*/

type OracleCompiler struct {
	*BaseCompiler
}

func makeOracleCompiler() BaseCompiler {
	compiler := &OracleCompiler{
		BaseCompiler: &BaseCompiler{
			openingIdentifier: "\"",
			closingIdentifier: "\"",
			defaultNamespace:  "public",
			booleanAsBit:      false,
			lastId:            "sys_guid() as [Id]",
			columnAsKeyword:   " ",
			tableAsKeyword:    " ",
			engineName:        OracleEngineScope,
		},
	}
	compiler.Compiler = compiler
	return BaseCompiler{Compiler: compiler}
}

/*
without Order By

ORDER BY

	 (
	   SELECT
	     0
	   FROM
	     DUAL
	 ) OFFSET 30 ROWS FETCH NEXT 10 ROWS ONLY

	Without Order By and offset
	 ORDER BY
	 (
	   SELECT
	     0
	   FROM
	     DUAL
	 ) OFFSET 0 ROWS FETCH NEXT 10 ROWS ONLY

	 without Order By and limit
	 ORDER BY
	 (
	   SELECT
	     0
	   FROM
	     DUAL
	 ) OFFSET 10 ROWS

	 with Order By, limit and offset
	 ORDER BY "Test" OFFSET 0 ROWS FETCH NEXT 10 ROWS ONLY
*/
func (compiler OracleCompiler) compilePagination(query Query, ctx QueryContext) string {
	limitClause := GetOneComponent[*LimitClause](query.clauses, isLimitType, isPaginationComponent, isOracleEngineScope)
	offsetClause := GetOneComponent[*OffsetClause](query.clauses, isOffsetType, isPaginationComponent, isOracleEngineScope)
	offset := 0
	limit := 0

	if limitClause == nil && offsetClause == nil {
		return ""
	}

	if offsetClause != nil {
		offset = (*offsetClause).Count
	}

	if limitClause != nil {
		limit = (*limitClause).Count
	}

	orderbyClause := ctx.compileOrderBy(query)
	stringBuilder := strings.Builder{}

	if orderbyClause == "" {
		stringBuilder.WriteString(" ORDER BY (SELECT 0 FROM DUAL)")
	}

	if limitClause != nil {
		stringBuilder.WriteString(fmt.Sprintf(" OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, limit))
		return stringBuilder.String()
	} else {
		stringBuilder.WriteString(fmt.Sprintf(" OFFSET %d ROWS", offset))
	}
	return stringBuilder.String()
}
