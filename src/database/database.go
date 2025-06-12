package database

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// ErrCannotOpenDatabaseConnection is a sentinel error describing a failure to open a database connection.
var ErrCannotOpenDatabaseConnection = errors.New("cannot open database connection")

// ErrNoDatabaseConnectionReturned is a sentinel error describing a nil database connection being returned from GORM
// without an actual GORM error occurring at the same time.
var ErrNoDatabaseConnectionReturned = errors.New("nil database connection returned from GORM")

// DatabaseConnection is a struct representing an active DB connection.
type DatabaseConnection struct {
	db        *gorm.DB
	debugMode bool
}

// DatabaseConnectionArguments is a struct representing the general properties expected when making a connection
// to a database environment.
type DatabaseConnectionArguments struct {
	DatabaseName string
	Host         string
	Password     string
	Port         string
	Username     string
}

// NewDatabaseConnection opens and initializes a database connection based upon the GORM dialector, a boolean
// describing whether to turn on "debug mode" automatically, and a slice of GORM options. Returns the DB connection
// pointer as well as any error that may have occurred.
func NewDatabaseConnection(
	gormConnection gorm.Dialector, shouldUseDebugMode bool, gormOptions ...gorm.Option,
) (*DatabaseConnection, error) {
	db, err := gorm.Open(gormConnection, gormOptions...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrCannotOpenDatabaseConnection, err)
	}
	if db == nil {
		return nil, ErrNoDatabaseConnectionReturned
	}
	if shouldUseDebugMode {
		db = db.Debug()
	}
	return &DatabaseConnection{
		db:        db,
		debugMode: shouldUseDebugMode,
	}, nil
}

// GetGORMDB returns the GORM DB pointer for this connection.
func (dc *DatabaseConnection) GetGORMDB() *gorm.DB {
	if dc == nil {
		return nil
	}
	return dc.db
}

// IsUsingDebugMode returns a boolean describing whether "debug mode" is turned on for this connection.
func (dc *DatabaseConnection) IsUsingDebugMode() bool {
	if dc == nil {
		return false
	}
	return dc.debugMode
}
