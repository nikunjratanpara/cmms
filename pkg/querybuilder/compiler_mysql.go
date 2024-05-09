package querybuilder

/*
.WhereDate("CreatedAt", "2018-04-01")
.WhereTime("CreatedAt", ">", "16:30")
.WhereDatePart("month", "CreatedAt", 2);

	DATE(`CreatedAt`) = '2018-04-01'
	AND TIME(`CreatedAt`) > '16:30'
	AND MONTH(`CreatedAt`) = 2
*/
type MysqlCompiler struct {
	*BaseCompiler
}

func makeMysqlCompiler() BaseCompiler {
	compiler := &MysqlCompiler{
		BaseCompiler: &BaseCompiler{
			openingIdentifier: "`",
			closingIdentifier: "`",
			defaultNamespace:  "public",
			booleanAsBit:      false,
			lastId:            "last_insert_id() as [Id]",
			columnAsKeyword:   " AS ",
			tableAsKeyword:    " AS ",
			engineName:        MySqlEngineScope,
		},
	}
	compiler.Compiler = compiler
	return BaseCompiler{Compiler: compiler}
}
