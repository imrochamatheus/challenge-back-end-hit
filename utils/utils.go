package utils

import (
	"database/sql"
	"os"
)

func ReadQueryFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}

func PrepareQuery(db *sql.DB, path string) (*sql.Stmt, error) {
	query, err := ReadQueryFile(path)
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	return stmt, nil
}
