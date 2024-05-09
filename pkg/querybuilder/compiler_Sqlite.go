package querybuilder

/*
 strftime('%Y-%m-%d', "CreatedAt") = cast('2018-04-01' as text)
  AND strftime('%H:%M:%S', "CreatedAt") > cast('16:30' as text)
  AND strftime('%m', "CreatedAt") = cast(2 as text)
*/

type SqliteCompiler struct {
	*BaseCompiler
}

func makeSqliteCompiler() BaseCompiler {
	compiler := &SqliteCompiler{
		BaseCompiler: &BaseCompiler{
			openingIdentifier: "\"",
			closingIdentifier: "\"",
			defaultNamespace:  "public",
			booleanAsBit:      false,
			lastId:            "last_insert_rowid() as [Id]",
			columnAsKeyword:   " AS ",
			tableAsKeyword:    " AS ",
			engineName:        SqliteEngineScope,
		},
	}
	compiler.Compiler = compiler
	return BaseCompiler{Compiler: compiler}
}

func (compiler SqliteCompiler) CompileTrue() string {
	return "1"
}

func (compiler SqliteCompiler) CompileFalse() string {
	return "0"
}
