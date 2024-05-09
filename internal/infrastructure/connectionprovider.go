package infrastructure

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v5"
)

type DatabaseConnectionProvider struct {
}

func (provider DatabaseConnectionProvider) GetConnetion() *pgx.Conn {
	connectionString := getConnectionString()
	connection, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

func getConnectionString() string {
	userName := "postgres"
	password, _ := url.Parse("Secret@1234")
	server := "localhost:5432"
	database := "Account"
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s/%s", userName, password.String(), server, database)
	return connectionString
}
