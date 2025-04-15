package db

import (
	"database/sql"
	"os"
	"path/filepath"
)

func getDatabasePath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".todo-cli", "todo.db")
}

func ensureDatabasePathExists() error {
	dbPath := getDatabasePath()
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(dbPath), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetDatabase returns a pointer to the database connection.
func GetDatabase() (*sql.DB, error) {
	dbPath := getDatabasePath()
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InitDatabase initializes the database connection and creates the database file if it doesn't exist.
func InitDatabase() (*sql.DB, error) {
	if err := ensureDatabasePathExists(); err != nil {
		return nil, err
	}
	db, err := GetDatabase()
	if err != nil {
		return nil, err
	}
	return db, nil
}
