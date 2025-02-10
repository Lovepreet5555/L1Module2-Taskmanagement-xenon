package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

// QueryDB executes a SQL query and returns the result.
func QueryDB(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	return rows, nil
}

// ExecDB executes a SQL statement and returns the result.
func ExecDB(query string, args ...interface{}) (sql.Result, error) {
	result, err := DB.Exec(query, args...)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return nil, err
	}
	return result, nil
}