package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func readQueryFile(path string) (string, error) {
	data, err := os.ReadFile(path)

	return string(data), err
}

func initializeDbConnection() (*sql.DB, error) {
	dbPath := "./db/main.db"
	err := os.MkdirAll("./db", os.ModePerm)

	if err != nil{
		fmt.Printf("error create database folder: %s", err)
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil{
		fmt.Printf("error initialize database connection: %s", err)
		return nil, err
	}

	createTable(db)

	return db, nil
}

func createTable(db *sql.DB) error {
	query, err := readQueryFile("./queries/create_planets_table.sql")

	if err != nil{
		fmt.Printf("unable to read read file with query to create planets table: %s", err)

		return err
	}

	_, err = db.Exec(query)

	if err != nil{
		fmt.Printf("error create planets table: %s", err)

		return err
	}

	return nil
}