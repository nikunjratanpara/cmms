package querybuilder

import (
	"testing"

	"github.com/nikunjratanpara/cmms/pkg/querybuilder"
)

func TestBasicColumnJoin(t *testing.T) {

	t.Run("Inner join", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("users").
			InnerJoin("contacts", "users.id", "contacts.user_id", querybuilder.Eq)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] INNER JOIN [contacts] ON [users].[id] = [contacts].[user_id]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` INNER JOIN `contacts` ON `users`.`id` = `contacts`.`user_id`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" INNER JOIN \"contacts\" ON \"users\".\"id\" = \"contacts\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" INNER JOIN \"contacts\" ON \"users\".\"id\" = \"contacts\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" INNER JOIN \"contacts\" ON \"users\".\"id\" = \"contacts\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Left join", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("users").
			LeftJoin("contacts", "users.id", "contacts.user_id", querybuilder.Eq)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] LEFT JOIN [contacts] ON [users].[id] = [contacts].[user_id]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` LEFT JOIN `contacts` ON `users`.`id` = `contacts`.`user_id`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" LEFT JOIN \"contacts\" ON \"users\".\"id\" = \"contacts\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" LEFT JOIN \"contacts\" ON \"users\".\"id\" = \"contacts\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" LEFT JOIN \"contacts\" ON \"users\".\"id\" = \"contacts\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Right join", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("users").
			RightJoin("contacts", "users.id", "contacts.user_id", querybuilder.Eq)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] RIGHT JOIN [contacts] ON [users].[id] = [contacts].[user_id]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` RIGHT JOIN `contacts` ON `users`.`id` = `contacts`.`user_id`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" RIGHT JOIN \"contacts\" ON \"users\".\"id\" = \"contacts\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" RIGHT JOIN \"contacts\" ON \"users\".\"id\" = \"contacts\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" RIGHT JOIN \"contacts\" ON \"users\".\"id\" = \"contacts\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Cross join", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("users").
			CrossJoin("contacts")

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] CROSS JOIN [contacts]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` CROSS JOIN `contacts`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" CROSS JOIN \"contacts\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" CROSS JOIN \"contacts\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" CROSS JOIN \"contacts\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})
}

func TestSubQueryJoinWithBasicColumnConditions(t *testing.T) {
	t.Run("Inner Join", func(t *testing.T) {
		contactQuery := querybuilder.NewQueryFromTable("contacts").
			Select("id", "user_id", "name", "email").
			WhereNull("user_id").As("C")

		query := querybuilder.NewQueryFromTable("users").
			InnerJoinWithQuery(*contactQuery, "users.id", "C.user_id", querybuilder.Eq)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] INNER JOIN (SELECT [id], [user_id], [name], [email] FROM [contacts] WHERE [user_id] IS NULL) AS [C] ON [users].[id] = [C].[user_id]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` INNER JOIN (SELECT `id`, `user_id`, `name`, `email` FROM `contacts` WHERE `user_id` IS NULL) AS `C` ON `users`.`id` = `C`.`user_id`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" INNER JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"users\".\"id\" = \"C\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" INNER JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"users\".\"id\" = \"C\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" INNER JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"users\".\"id\" = \"C\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Left Join", func(t *testing.T) {
		contactQuery := querybuilder.NewQueryFromTable("contacts").
			Select("id", "user_id", "name", "email").
			WhereNull("user_id").As("C")

		query := querybuilder.NewQueryFromTable("users").
			LeftJoinWithQuery(*contactQuery, "users.id", "C.user_id", querybuilder.Eq)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] LEFT JOIN (SELECT [id], [user_id], [name], [email] FROM [contacts] WHERE [user_id] IS NULL) AS [C] ON [users].[id] = [C].[user_id]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` LEFT JOIN (SELECT `id`, `user_id`, `name`, `email` FROM `contacts` WHERE `user_id` IS NULL) AS `C` ON `users`.`id` = `C`.`user_id`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" LEFT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"users\".\"id\" = \"C\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" LEFT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"users\".\"id\" = \"C\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" LEFT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"users\".\"id\" = \"C\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Right Join", func(t *testing.T) {
		contactQuery := querybuilder.NewQueryFromTable("contacts").
			Select("id", "user_id", "name", "email").
			WhereNull("user_id").As("C")

		query := querybuilder.NewQueryFromTable("users").
			RightJoinWithQuery(*contactQuery, "users.id", "C.user_id", querybuilder.Eq)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] RIGHT JOIN (SELECT [id], [user_id], [name], [email] FROM [contacts] WHERE [user_id] IS NULL) AS [C] ON [users].[id] = [C].[user_id]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` RIGHT JOIN (SELECT `id`, `user_id`, `name`, `email` FROM `contacts` WHERE `user_id` IS NULL) AS `C` ON `users`.`id` = `C`.`user_id`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" RIGHT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"users\".\"id\" = \"C\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" RIGHT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"users\".\"id\" = \"C\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" RIGHT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"users\".\"id\" = \"C\".\"user_id\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})

	t.Run("Cross Join", func(t *testing.T) {
		contactQuery := querybuilder.NewQueryFromTable("contacts").
			Select("id", "user_id", "name", "email").
			WhereNull("user_id").As("C")

		query := querybuilder.NewQueryFromTable("users").
			CrossJoinWithQuery(*contactQuery)

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql:        "SELECT * FROM [users] CROSS JOIN (SELECT [id], [user_id], [name], [email] FROM [contacts] WHERE [user_id] IS NULL) AS [C]",
				Parameters: map[int]interface{}{},
			},
			querybuilder.MySqlEngineScope: {
				Sql:        "SELECT * FROM `users` CROSS JOIN (SELECT `id`, `user_id`, `name`, `email` FROM `contacts` WHERE `user_id` IS NULL) AS `C`",
				Parameters: map[int]interface{}{},
			},
			querybuilder.PostgresEngineScope: {
				Sql:        "SELECT * FROM \"users\" CROSS JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.SqliteEngineScope: {
				Sql:        "SELECT * FROM \"users\" CROSS JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\"",
				Parameters: map[int]interface{}{},
			},
			querybuilder.OracleEngineScope: {
				Sql:        "SELECT * FROM \"users\" CROSS JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\"",
				Parameters: map[int]interface{}{},
			},
		}
		runTests(t, query, tests)
	})
}

func TestSubQueryJoinWithComplexConditions(t *testing.T) {
	t.Run("Inner Join", func(t *testing.T) {
		contactQuery := querybuilder.NewQueryFromTable("contacts").
			Select("id", "user_id", "name", "email").
			WhereNull("user_id").As("C")

		query := querybuilder.NewQueryFromTable("users").
			InnerJoinWithQueryAndConditions(*contactQuery, func(query *querybuilder.Query) *querybuilder.Query {
				return query.WhereNull("C.id").WhereBetween(
					"Users.created_at", "2020-01-01", "2020-01-02",
				)
			})

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql: "SELECT * FROM [users] INNER JOIN (SELECT [id], [user_id], [name], [email] FROM [contacts] WHERE [user_id] IS NULL) AS [C] ON [C].[id] IS NULL AND [Users].[created_at] BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.MySqlEngineScope: {
				Sql: "SELECT * FROM `users` INNER JOIN (SELECT `id`, `user_id`, `name`, `email` FROM `contacts` WHERE `user_id` IS NULL) AS `C` ON `C`.`id` IS NULL AND `Users`.`created_at` BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.PostgresEngineScope: {
				Sql: "SELECT * FROM \"users\" INNER JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"C\".\"id\" IS NULL AND \"Users\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.SqliteEngineScope: {
				Sql: "SELECT * FROM \"users\" INNER JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"C\".\"id\" IS NULL AND \"Users\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.OracleEngineScope: {
				Sql: "SELECT * FROM \"users\" INNER JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"C\".\"id\" IS NULL AND \"Users\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
		}

		runTests(t, query, tests)
	})

	t.Run("Left Join", func(t *testing.T) {
		contactQuery := querybuilder.NewQueryFromTable("contacts").
			Select("id", "user_id", "name", "email").
			WhereNull("user_id").As("C")

		query := querybuilder.NewQueryFromTable("users").
			LeftJoinWithQueryAndConditions(*contactQuery, func(query *querybuilder.Query) *querybuilder.Query {
				return query.WhereNull("C.id").WhereBetween(
					"Users.created_at", "2020-01-01", "2020-01-02",
				)
			})

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql: "SELECT * FROM [users] LEFT JOIN (SELECT [id], [user_id], [name], [email] FROM [contacts] WHERE [user_id] IS NULL) AS [C] ON [C].[id] IS NULL AND [Users].[created_at] BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.MySqlEngineScope: {
				Sql: "SELECT * FROM `users` LEFT JOIN (SELECT `id`, `user_id`, `name`, `email` FROM `contacts` WHERE `user_id` IS NULL) AS `C` ON `C`.`id` IS NULL AND `Users`.`created_at` BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.PostgresEngineScope: {
				Sql: "SELECT * FROM \"users\" LEFT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"C\".\"id\" IS NULL AND \"Users\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.SqliteEngineScope: {
				Sql: "SELECT * FROM \"users\" LEFT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"C\".\"id\" IS NULL AND \"Users\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.OracleEngineScope: {
				Sql: "SELECT * FROM \"users\" LEFT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"C\".\"id\" IS NULL AND \"Users\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
		}

		runTests(t, query, tests)
	})

	t.Run("Right Join", func(t *testing.T) {
		contactQuery := querybuilder.NewQueryFromTable("contacts").
			Select("id", "user_id", "name", "email").
			WhereNull("user_id").As("C")

		query := querybuilder.NewQueryFromTable("users").
			RightJoinWithQueryAndConditions(*contactQuery, func(query *querybuilder.Query) *querybuilder.Query {
				return query.WhereNull("C.id").WhereBetween(
					"Users.created_at", "2020-01-01", "2020-01-02",
				)
			})

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql: "SELECT * FROM [users] RIGHT JOIN (SELECT [id], [user_id], [name], [email] FROM [contacts] WHERE [user_id] IS NULL) AS [C] ON [C].[id] IS NULL AND [Users].[created_at] BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.MySqlEngineScope: {
				Sql: "SELECT * FROM `users` RIGHT JOIN (SELECT `id`, `user_id`, `name`, `email` FROM `contacts` WHERE `user_id` IS NULL) AS `C` ON `C`.`id` IS NULL AND `Users`.`created_at` BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.PostgresEngineScope: {
				Sql: "SELECT * FROM \"users\" RIGHT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"C\".\"id\" IS NULL AND \"Users\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.SqliteEngineScope: {
				Sql: "SELECT * FROM \"users\" RIGHT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"C\".\"id\" IS NULL AND \"Users\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.OracleEngineScope: {
				Sql: "SELECT * FROM \"users\" RIGHT JOIN (SELECT \"id\", \"user_id\", \"name\", \"email\" FROM \"contacts\" WHERE \"user_id\" IS NULL) AS \"C\" ON \"C\".\"id\" IS NULL AND \"Users\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
		}

		runTests(t, query, tests)
	})
}

func TestJoinWithComplexConditions(t *testing.T) {
	t.Run("Inner Join", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("users as U").
			InnerJoinWithConditions("contacts as C", func(query *querybuilder.Query) *querybuilder.Query {
				return query.WhereNull("U.user_id").WhereBetween(
					"C.created_at", "2020-01-01", "2020-01-02",
				)
			})

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql: "SELECT * FROM [users] AS [U] INNER JOIN [contacts] AS [C] ON [U].[user_id] IS NULL AND [C].[created_at] BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.MySqlEngineScope: {
				Sql: "SELECT * FROM `users` AS `U` INNER JOIN `contacts` AS `C` ON `U`.`user_id` IS NULL AND `C`.`created_at` BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.PostgresEngineScope: {
				Sql: "SELECT * FROM \"users\" AS \"U\" INNER JOIN \"contacts\" AS \"C\" ON \"U\".\"user_id\" IS NULL AND \"C\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.SqliteEngineScope: {
				Sql: "SELECT * FROM \"users\" AS \"U\" INNER JOIN \"contacts\" AS \"C\" ON \"U\".\"user_id\" IS NULL AND \"C\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.OracleEngineScope: {
				Sql: "SELECT * FROM \"users\" \"U\" INNER JOIN \"contacts\" \"C\" ON \"U\".\"user_id\" IS NULL AND \"C\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
		}

		runTests(t, query, tests)
	})

	t.Run("Left Join", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("users as U").
			LeftJoinWithConditions("contacts as C", func(query *querybuilder.Query) *querybuilder.Query {
				return query.WhereNull("U.user_id").WhereBetween(
					"C.created_at", "2020-01-01", "2020-01-02",
				)
			})

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql: "SELECT * FROM [users] AS [U] LEFT JOIN [contacts] AS [C] ON [U].[user_id] IS NULL AND [C].[created_at] BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.MySqlEngineScope: {
				Sql: "SELECT * FROM `users` AS `U` LEFT JOIN `contacts` AS `C` ON `U`.`user_id` IS NULL AND `C`.`created_at` BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.PostgresEngineScope: {
				Sql: "SELECT * FROM \"users\" AS \"U\" LEFT JOIN \"contacts\" AS \"C\" ON \"U\".\"user_id\" IS NULL AND \"C\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.SqliteEngineScope: {
				Sql: "SELECT * FROM \"users\" AS \"U\" LEFT JOIN \"contacts\" AS \"C\" ON \"U\".\"user_id\" IS NULL AND \"C\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.OracleEngineScope: {
				Sql: "SELECT * FROM \"users\" \"U\" LEFT JOIN \"contacts\" \"C\" ON \"U\".\"user_id\" IS NULL AND \"C\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
		}

		runTests(t, query, tests)
	})

	t.Run("Right Join", func(t *testing.T) {
		query := querybuilder.NewQueryFromTable("users as U").
			RightJoinWithConditions("contacts as C", func(query *querybuilder.Query) *querybuilder.Query {
				return query.WhereNull("U.user_id").WhereBetween(
					"C.created_at", "2020-01-01", "2020-01-02",
				)
			})

		tests := map[string]querybuilder.QueryResult{
			querybuilder.SqlServerEngineScope: {
				Sql: "SELECT * FROM [users] AS [U] RIGHT JOIN [contacts] AS [C] ON [U].[user_id] IS NULL AND [C].[created_at] BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.MySqlEngineScope: {
				Sql: "SELECT * FROM `users` AS `U` RIGHT JOIN `contacts` AS `C` ON `U`.`user_id` IS NULL AND `C`.`created_at` BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.PostgresEngineScope: {
				Sql: "SELECT * FROM \"users\" AS \"U\" RIGHT JOIN \"contacts\" AS \"C\" ON \"U\".\"user_id\" IS NULL AND \"C\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.SqliteEngineScope: {
				Sql: "SELECT * FROM \"users\" AS \"U\" RIGHT JOIN \"contacts\" AS \"C\" ON \"U\".\"user_id\" IS NULL AND \"C\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
			querybuilder.OracleEngineScope: {
				Sql: "SELECT * FROM \"users\" \"U\" RIGHT JOIN \"contacts\" \"C\" ON \"U\".\"user_id\" IS NULL AND \"C\".\"created_at\" BETWEEN $0 AND $1",
				Parameters: map[int]interface{}{
					0: "2020-01-01",
					1: "2020-01-02",
				},
			},
		}

		runTests(t, query, tests)
	})

}
