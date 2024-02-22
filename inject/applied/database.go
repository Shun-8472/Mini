package applied

import (
	"demo/internal/applied/database"
	"demo/internal/applied/database/mysql"
)

func InitDatabaseConnect() database.Database {
	return mysql.NewConnect()
}
