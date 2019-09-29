package db

import (
	"database/sql"
	"util/coreconfig"
	logger "util/log"
)

func DBConnection() (*sql.DB, error) {
	mysqlHost := coreconfig.CC.Mysql.Host
	mysqlPort := coreconfig.CC.Mysql.Port
	mysqlDatabase := coreconfig.CC.Mysql.Database
	mysqlUser := coreconfig.CC.Mysql.User
	mysqlPassword := coreconfig.CC.Mysql.Password
	// "root:JDqXPqm@tcp(127.0.0.1:3306)/login?charset=utf8"
	connStr := mysqlUser + ":" + mysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDatabase + "?charset=utf8"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		logger.Error.Println("Open mysql error, ", err.Error())
		return nil, err
	}
	return db, nil
}