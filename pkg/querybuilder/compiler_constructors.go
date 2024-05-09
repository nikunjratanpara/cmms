package querybuilder

import (
	"fmt"
	"os"
)

var (
	compilers map[string]Compiler = make(map[string]Compiler)
)

// GetCompiler returns the SqlCompiler for the given engine.
// The function takes a string parameter engine and returns a SqlCompiler.
// Supported engines are SqlServer, Postgres,
// if engine value is empty, it reads from environment variable "querybuilder.compiler"
func GetCompiler(engine string) Compiler {
	compiler, ok := compilers[engine]
	if ok {
		return compiler
	}
	compilers[engine] = makeCompiler(engine)
	return compilers[engine]
}

// makeCompiler creates a SqlCompiler based on the provided SQL engine.
// It takes a string sqlEngine as a parameter and returns a SqlCompiler.
// If sqlEngine is empty, it falls back to the "querybuilder.compiler" environment variable.
// Supported engines are "SqlServer" and "Postgres".
func makeCompiler(sqlEngine string) Compiler {
	if sqlEngine == "" {
		sqlEngine = os.Getenv("querybuilder.compiler")
	}

	var compiler BaseCompiler
	switch sqlEngine {
	case "SqlServer":
		compiler = makeSqlServerCompiler()
	case "Postgres":
		compiler = makePostgresCompiler()
	case "MySql":
		compiler = makeMysqlCompiler()
	case "Sqlite":
		compiler = makeSqliteCompiler()
	case "Oracle":
		compiler = makeOracleCompiler()
	default:
		panic(fmt.Errorf("unsupported %s engine by query builder", sqlEngine))
	}
	return compiler.Compiler
}
