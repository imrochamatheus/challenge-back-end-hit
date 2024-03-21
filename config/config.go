package config

import (
	"database/sql"
	"log"
)

var (
	db *sql.DB
)

func Init() error {
	var err error

	db, err = initializeDbConnection()

	if err != nil {
		log.Printf("Error initializing database connection: %v", err)
	}

	return err
}

func GetDbInstance() *sql.DB {
	return db
}
