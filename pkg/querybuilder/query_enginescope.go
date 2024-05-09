package querybuilder

import "github.com/nikunjratanpara/cmms/pkg/linq"

func (query *Query) For(engineScope string, callback func(*Query) *Query) *Query {
	q := callback(NewQuery())
	linq.ForEach(q.GetAllClauses(), func(c Clause) {
		c.SetEngineScope(engineScope)
	})
	return mergeQueries(query, q)
}

func (query *Query) ForMySql(callback func(*Query) *Query) *Query {
	return query.For("MySql", callback)
}

func (query *Query) ForSqlServer(callback func(*Query) *Query) *Query {
	return query.For("SqlServer", callback)
}

func (query *Query) ForPostgres(callback func(*Query) *Query) *Query {
	return query.For("Postgres", callback)
}

func (query *Query) ForSqlite(callback func(*Query) *Query) *Query {
	return query.For("Sqlite", callback)
}

func (query *Query) ForOracle(callback func(*Query) *Query) *Query {
	return query.For("Oracle", callback)
}

func mergeQueries(parentQuery *Query, childQuery *Query) *Query {
	parentQuery.clauses = append(parentQuery.clauses, childQuery.clauses...)
	parentQuery.whereClauses = append(parentQuery.whereClauses, childQuery.whereClauses...)
	parentQuery.havingClauses = append(parentQuery.havingClauses, childQuery.havingClauses...)
	parentQuery.orderByClauses = append(parentQuery.orderByClauses, childQuery.orderByClauses...)
	parentQuery.groupByClauses = append(parentQuery.groupByClauses, childQuery.groupByClauses...)
	parentQuery.selectClauses = append(parentQuery.selectClauses, childQuery.selectClauses...)
	return parentQuery
}
