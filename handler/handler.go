package handler

import (
	"database/sql"

	"github.com/imrochamatheus/challenge-back-end-hit/config"
)

var (
	db *sql.DB
)

func InitializeHandlers() {
	db = config.GetDbInstance()
}