package products_db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/chertokdmitry/bot-pandora-11/src/utils/logger"
	"os"
)

func GetDB() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser, _ := os.LookupEnv("MYSQL_USERNAME")
	dbPass, _ := os.LookupEnv("MYSQL_PASS")
	dbName, _ := os.LookupEnv("MYSQL_DB")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		logger.Error("error when trying to connect to DB", err)
	}
	return db
}
