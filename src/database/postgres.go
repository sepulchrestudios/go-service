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

// MakePostgresConfigFromDSN takes a DSN string and returns a postgres.Config instance that contains it.
func MakePostgresConfigFromDSN(DSN string) postgres.Config {
	return postgres.Config{
		DSN: DSN,
	}
}

// MakePostgresDSNFromConnectionArguments takes a PostgresDatabaseConnectionArguments struct pointer and returns a
// string containing the DSN to be used during connections.
func MakePostgresDSNFromConnectionArguments(connectionArguments *PostgresDatabaseConnectionArguments) string {
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
	return strings.Join(connectionParams, " ")
}

// MakePostgresDialectorFromConfig takes a postgres.Config instance and returns a gorm.Dialector instance that can be
// used for opening new GORM database connections.
func MakePostgresDialectorFromConfig(config postgres.Config) gorm.Dialector {
	return postgres.New(config)
}

// NewPostgresDatabaseConnection opens and initializes a Postgres connection with GORM using the supplied connection
// arguments, a boolean describing whether to turn on "debug mode" automatically, and a slice of GORM options. Returns
// the DB connection pointer as well as any error that may have occurred.
func NewPostgresDatabaseConnection(
	connectionArguments *PostgresDatabaseConnectionArguments, shouldUseDebugMode bool, gormOptions ...gorm.Option,
) (*DatabaseConnection, error) {
	err := ValidatePostgresConnectionArguments(connectionArguments)
	if err != nil {
		return nil, err
	}
	postgresDSN := MakePostgresDSNFromConnectionArguments(connectionArguments)
	postgresConfig := MakePostgresConfigFromDSN(postgresDSN)
	postgresDialector := MakePostgresDialectorFromConfig(postgresConfig)
	return NewDatabaseConnection(postgresDialector, shouldUseDebugMode, gormOptions...)
}

// ValidatePostgresConnectionArguments takes a PostgresDatabaseConnectionArguments struct pointer and returns an error
// if any of the expected fields are missing. Returns nil if the validation checks pass.
func ValidatePostgresConnectionArguments(connectionArguments *PostgresDatabaseConnectionArguments) error {
	if connectionArguments == nil {
		return fmt.Errorf("connection arguments for postgres cannot be nil")
	}
	if connectionArguments.Host == "" {
		return fmt.Errorf("database host in connection arguments cannot be blank")
	}
	if connectionArguments.DatabaseName == "" {
		return fmt.Errorf("database name in connection arguments cannot be blank")
	}
	if connectionArguments.Username == "" {
		return fmt.Errorf("username in connection arguments cannot be blank")
	}
	return nil
}
