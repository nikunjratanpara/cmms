package querybuilder

type AbstractClause struct {
	Type        string
	Component   string
	EngineScope string
}

type AbstractColumnClause struct {
	*AbstractClause
	ColumnName string
	IsNot      bool
}

type ConditionClause struct {
	*AbstractColumnClause
	Value    interface{}
	Operator Operator
}

func (abstractClause AbstractClause) GetType() string {
	return abstractClause.Type
}

func (abstractClause AbstractClause) GetComponent() string {
	return abstractClause.Component
}

func (abstractClause AbstractClause) GetEngineScope() string {
	return abstractClause.EngineScope
}

func (abstractClause *AbstractClause) SetEngineScope(engineScope string) {
	abstractClause.EngineScope = engineScope
}

func (abstractClause AbstractClause) Clone() *AbstractClause {
	return &AbstractClause{
		Type:        abstractClause.Type,
		Component:   abstractClause.Component,
		EngineScope: abstractClause.EngineScope,
	}
}

func (clause ConditionClause) GetSql(context QueryContext) string {
	operator := string(clause.Operator)
	not := ""
	if clause.IsNot {
		not = "NOT " + string(clause.Operator)
	}
	return not + context.prepareIdentifier(clause.ColumnName) + " " + operator + " " + context.AddParameter(clause.Value)
}
