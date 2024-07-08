package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func initDb() {
	var err error
	db, err = sql.Open("sqlite3", "./todos.sql")
	if err != nil {
		log.Fatal(err)
	}
	cmd := `CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		created_at DATETIME
	);`
	_, err = db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
}
