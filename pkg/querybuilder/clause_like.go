package querybuilder

import "strings"

type LikeConditionClause struct {
	*AbstractColumnClause
	Value string
}

func makeLikeClause(abstractClause *AbstractClause, columnName string, value string, isNot bool) *LikeConditionClause {
	return &LikeConditionClause{
		AbstractColumnClause: makeAbstractColumnClause(abstractClause, columnName, isNot),
		Value:                value,
	}
}

func makeContains(abstractClause *AbstractClause, columnName string, value string) *LikeConditionClause {
	return makeLikeClause(withContainsType(abstractClause), columnName, value, false)
}

func makeStartWith(abstractClause *AbstractClause, columnName string, value string) *LikeConditionClause {
	return makeLikeClause(withStartWithType(abstractClause), columnName, value, false)
}

func makeEndWith(abstractClause *AbstractClause, columnName string, value string) *LikeConditionClause {
	return makeLikeClause(withEndWithType(abstractClause), columnName, value, false)
}

func makeLike(abstractClause *AbstractClause, columnName string, value string) *LikeConditionClause {
	return makeLikeClause(withLikeType(abstractClause), columnName, value, false)
}

func makeNotLike(abstractClause *AbstractClause, columnName string, value string) *LikeConditionClause {
	return makeLikeClause(withLikeType(abstractClause), columnName, value, true)
}

func makeNotContains(abstractClause *AbstractClause, columnName string, value string) *LikeConditionClause {
	return makeLikeClause(withContainsType(abstractClause), columnName, value, true)
}

func makeNotStartWith(abstractClause *AbstractClause, columnName string, value string) *LikeConditionClause {
	return makeLikeClause(withStartWithType(abstractClause), columnName, value, true)
}

func makeNotEndWith(abstractClause *AbstractClause, columnName string, value string) *LikeConditionClause {
	return makeLikeClause(withEndWithType(abstractClause), columnName, value, true)
}

func (clause LikeConditionClause) GetSql(context QueryContext) string {
	value := prepareLikeValue(clause)

	strClause := "LIKE "
	if clause.IsNot {
		strClause = "NOT LIKE "
	}

	return "LOWER(" +
		context.prepareIdentifier(clause.ColumnName) +
		") " +
		strClause +
		context.AddParameter(strings.ToLower(value))
}

func prepareLikeValue(clause LikeConditionClause) string {
	value := clause.Value
	if clause.GetType() == "StartWith" {
		value = value + "%"
	} else if clause.GetType() == "EndWith" {
		value = "%" + value
	} else if clause.GetType() == "Contain" {
		value = "%" + value + "%"
	}
	return value
}
