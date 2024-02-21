package mysql

import (
	"database/sql"

	"demo/config"
	"demo/internal/applied/database"
)

var (
	DatabaseConnection = new(sql.DB)
	err                error
)

type Database struct {
}

func NewConnect() database.Database {
	return &Database{}
}

func (d Database) ConnectDatabase() {
	connectionInfo := config.GetMySqlAddress()

	DatabaseConnection, err = sql.Open("mysql", connectionInfo)
	if err != nil {
		panic("failed to connect database")
	}
}
