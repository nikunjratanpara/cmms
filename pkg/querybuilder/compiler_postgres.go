package querybuilder

/*
.WhereDate("CreatedAt", "2018-04-01")
.WhereTime("CreatedAt", ">", "16:30")
.WhereDatePart("month", "CreatedAt", 2);

"CreatedAt" :: date = '2018-04-01'
  AND "CreatedAt" :: time > '16:30'
  AND DATE_PART('MONTH', "CreatedAt") = 2
*/

type PostgresCompiler struct {
	*BaseCompiler
}

func makePostgresCompiler() BaseCompiler {
	compiler := &PostgresCompiler{
		BaseCompiler: &BaseCompiler{
			openingIdentifier: "\"",
			closingIdentifier: "\"",
			defaultNamespace:  "public",
			booleanAsBit:      false,
			lastId:            "lastval() as [Id]",
			columnAsKeyword:   " AS ",
			tableAsKeyword:    " AS ",
			engineName:        PostgresEngineScope,
		},
	}
	compiler.Compiler = compiler
	return BaseCompiler{Compiler: compiler}
}
