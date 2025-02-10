package database

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() *sql.DB {
	var err error
	DB, err = sql.Open("sqlite", "./tasks.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		due_date DATE NOT NULL,
		status TEXT CHECK(status IN ('pending', 'in_progress', 'completed'))
	);`

	if _, err = DB.Exec(createTableSQL); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return DB
}