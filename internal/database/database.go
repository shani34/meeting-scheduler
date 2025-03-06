package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/shani34/meeting-scheduler/internal/config"
	_ "github.com/lib/pq"
)

// DB represents the database connection
type DB struct {
	*sql.DB
}

// NewDB creates a new database connection
func NewDB(cfg *config.Config) (*DB, error) {
	connStr := cfg.GetDBConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to database")
	return &DB{db}, nil
}

// Close closes the database connection
func (db *DB) Close() {
	if err := db.DB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
} 