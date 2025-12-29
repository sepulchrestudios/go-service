package database

import (
	"gorm.io/gorm"
)

// Contract represents the interface that all database connection implementations must satisfy.
type Contract interface {
	// GetGORMDB returns the GORM DB pointer for the connection.
	GetGORMDB() *gorm.DB

	// IsUsingDebugMode returns a boolean describing whether "debug mode" is turned on for the connection.
	IsUsingDebugMode() bool
}
