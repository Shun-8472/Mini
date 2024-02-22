package applied

import (
	"demo/internal/applied/database"
	"demo/internal/applied/database/mysql"
)

func InitDatabaseConnection() database.Database {
	return mysql.NewConnect()
}
