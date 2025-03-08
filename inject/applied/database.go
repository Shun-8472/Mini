package applied

import (
	"mini/internal/applied/database"
	"mini/internal/applied/database/mysql"
)

func InitDatabaseConnection() database.Database {
	return mysql.NewConnect()
}
