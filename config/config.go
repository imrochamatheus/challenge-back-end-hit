package config

import (
	"database/sql"
	"fmt"
)

var (
	db *sql.DB
)

func Init() error {
	var err error

	db, err = initializeDbConnection()

	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func GetDbInstance() *sql.DB {
	return db
}
