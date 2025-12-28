package database

import "errors"

// ErrCannotOpenDatabaseConnection is a sentinel error describing a failure to open a database connection.
var ErrCannotOpenDatabaseConnection = errors.New("cannot open database connection")

// ErrNoDatabaseConnectionReturned is a sentinel error describing a nil database connection being returned from GORM
// without an actual GORM error occurring at the same time.
var ErrNoDatabaseConnectionReturned = errors.New("nil database connection returned from GORM")

// ErrPostgresNoConnectionArguments is a sentinel error representing a nil connection arguments pointer when attempting
// to make a Postgres DB connection.
var ErrPostgresNoConnectionArguments = errors.New("connection arguments for postgres cannot be nil")

// ErrPostgresNoConnectionDatabaseHost is a sentinel error representing a blank database host string when attempting to
// make a Postgres DB connection.
var ErrPostgresNoConnectionDatabaseHost = errors.New("database host in connection arguments cannot be blank")

// ErrPostgresNoConnectionDatabaseName is a sentinel error representing a blank database name string when attempting to
// make a Postgres DB connection.
var ErrPostgresNoConnectionDatabaseName = errors.New("database name in connection arguments cannot be blank")

// ErrPostgresNoConnectionUsername is a sentinel error representing a blank username string when attempting to make a
// Postgres DB connection.
var ErrPostgresNoConnectionUsername = errors.New("username in connection arguments cannot be blank")
