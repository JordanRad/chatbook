package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// This method returns a pointer to an established connection to a given database
func ConnectToDatabase(username, password, host string, port int, dbName string) *sql.DB {
	dbConnectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, host, port, dbName)

	var db *sql.DB
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Maximum Idle Connections
	db.SetMaxIdleConns(16)
	// Maximum Open Connections
	db.SetMaxOpenConns(32)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(20 * time.Second)

	if err = db.Ping(); err != nil {
		log.Fatalf("Unable to reach database: %v", err)
	}

	log.Print("Connected to database \n")
	return db
}
