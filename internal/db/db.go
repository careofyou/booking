package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func New(maxOpenConns, maxIdleConns int, maxIdleTime string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		return nil, err
	}
	if maxOpenConns > 0 {
		db.SetMaxOpenConns(maxOpenConns)
	}
	db.SetMaxIdleConns(maxIdleConns)
	duration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	createTable(db)
	return db, nil
}

func createTable(db *sql.DB) {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
		)
	`

	_, err := db.Exec(createUsersTable)

	if err != nil {
		panic("Couldnt create users table")
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL, 
			datetime DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = db.Exec(createEventsTable)
	if err != nil {
		panic("Couldnt create events table")
	}

	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registration (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = db.Exec(createRegistrationTable)
	if err != nil {
		panic("Couldnt create registration table")
	}
}
