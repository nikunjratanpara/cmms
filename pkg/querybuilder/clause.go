package querybuilder

import (
	"time"

	"github.com/google/uuid"
)

type ComparablePremitive interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64 | float32 | float64 | time.Time
}

type Primitive interface {
	bool | string | uuid.UUID | ComparablePremitive
}

type Operator string

const (
	Gt  Operator = ">"
	Lt  Operator = "<"
	Eq  Operator = "="
	NEq Operator = "!="
	GtE Operator = ">="
	LtE Operator = "<="
)

type Clause interface {
	// GetType returns the type of the clause.
	//
	// This method returns a string representing the type of the clause. The
	// returned string could be one of the following: "ComparisonClause",
	// "LogicalOperatorClause", "JoinClause", etc.
	//
	// Each concrete implementation of the `Clause` interface should provide a
	// meaningful type name.
	GetType() string

	// GetSql returns the SQL representation of the clause, given the query context.
	//
	// The `context` parameter is the query context, which provides methods to
	// transform column and table names, add parameters, and escape characters.
	//
	// This method returns a string with the SQL representation of the clause.
	GetSql(ctx QueryContext) string

	GetComponent() string
	GetEngineScope() string
	SetEngineScope(string)
}
