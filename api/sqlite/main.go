package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Connection struct {
	db *sql.DB
}

// NewConnection creates a new database connection and runs migrations
func NewConnection() (*Connection, error) {
	dbPath := "../db/email_signups.db"
	migrationPath := "../db/migrations.sql"

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		fmt.Println("Database not found. Creating...")

		db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=rwc", dbPath))
		if err != nil {
			return nil, fmt.Errorf("failed to create database: %v", err)
		}

		err = runMigrations(db, migrationPath)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("failed to run migrations: %v", err)
		}

		fmt.Println("Migrations completed successfully")
		return &Connection{db}, nil
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=rw", dbPath))
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Database found. Skipping migrations")
	return &Connection{db}, nil
}

// runMigrations runs SQL migrations from the specified file
func runMigrations(db *sql.DB, migrationPath string) error {
	_, err := os.Stat(migrationPath)
	if os.IsNotExist(err) {
		log.Println("No migrations file found. Skipping migrations.")
		return nil
	}

	migrationBytes, err := os.ReadFile(migrationPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations file: %v", err)
	}

	migrationSQL := string(migrationBytes)

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	_, err = tx.Exec(migrationSQL)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to execute migrations: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}

func (c *Connection) Close() {
	c.db.Close()
}

func (c *Connection) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return c.db.Query(query, args...)
}

func (c *Connection) Exec(query string, args ...interface{}) (sql.Result, error) {
	return c.db.Exec(query, args...)
}
