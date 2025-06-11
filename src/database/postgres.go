package database

import (
	"fmt"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDatabaseConnectionArguments is a struct representing the properties expected when making a connection to a
// Postgres database environment.
type PostgresDatabaseConnectionArguments struct {
	DatabaseConnectionArguments

	SSLMode  string
	Timezone string
}

// NewPostgresDatabaseConnection opens and initializes a Postgres connection with GORM using the supplied connection
// arguments, a boolean describing whether to turn on "debug mode" automatically, and a slice of GORM options. Returns
// the DB connection pointer as well as any error that may have occurred.
func NewPostgresDatabaseConnection(
	connectionArguments *PostgresDatabaseConnectionArguments, shouldUseDebugMode bool, gormOptions ...gorm.Option,
) (*DatabaseConnection, error) {
	if connectionArguments == nil {
		return nil, fmt.Errorf("connection arguments for postgres cannot be nil")
	}
	if connectionArguments.Host == "" {
		return nil, fmt.Errorf("database host in connection arguments cannot be blank")
	}
	if connectionArguments.DatabaseName == "" {
		return nil, fmt.Errorf("database name in connection arguments cannot be blank")
	}
	if connectionArguments.Username == "" {
		return nil, fmt.Errorf("username in connection arguments cannot be blank")
	}
	connectionParams := []string{
		fmt.Sprintf("host=%s", connectionArguments.Host),
		fmt.Sprintf("dbname=%s", connectionArguments.DatabaseName),
		fmt.Sprintf("user=%s", connectionArguments.Username),
	}
	if connectionArguments.Password != "" {
		connectionParams = append(connectionParams, fmt.Sprintf("password=%s", connectionArguments.Password))
	}
	if connectionArguments.Port != "" {
		connectionParams = append(connectionParams, fmt.Sprintf("port=%s", connectionArguments.Port))
	}
	if connectionArguments.SSLMode != "" {
		connectionParams = append(connectionParams, fmt.Sprintf("sslmode=%s", connectionArguments.SSLMode))
	}
	if connectionArguments.Timezone != "" {
		connectionParams = append(connectionParams, fmt.Sprintf("timezone=%s", connectionArguments.Timezone))
	}
	postgresDialector := postgres.New(postgres.Config{
		DSN: strings.Join(connectionParams, " "),
	})
	return NewDatabaseConnection(postgresDialector, shouldUseDebugMode, gormOptions...)
}
