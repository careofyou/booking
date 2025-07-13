package main

import (
	"fmt"
	"log"

	"github.com/careofyou/booking.git/internal/db"
)

func main() {
	conn := config{
		maxOpenConns: 30,
		maxIdleConns: 30,
		maxIdleTime:  "15m",
	}

	// Database connection
	db, err := db.New(
		conn.maxOpenConns,
		conn.maxIdleConns,
		conn.maxIdleTime,
	)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Database connection pool established")

	app := &application{
		config: conn,
	}

	mux := app.mount()
	if err := app.run(mux); err != nil {
		fmt.Println("Connection error")
	}
	log.Println(app.run(mux))
}
