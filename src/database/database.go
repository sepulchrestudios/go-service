package database

import (
	"gorm.io/gorm"
)

// DatabaseConnectionInterface represents the interface that all database connection implementations must satisfy.
type DatabaseConnectionInterface interface {
	// GetGORMDB returns the GORM DB pointer for the connection.
	GetGORMDB() *gorm.DB

	// IsUsingDebugMode returns a boolean describing whether "debug mode" is turned on for the connection.
	IsUsingDebugMode() bool
}
