package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	MYSQL_DRIVER_NAME = "mysql"
)

// Database ...
type DB struct {
	SQL *sql.DB
}

//Database Connection ...
var dbConn = &DB{}

func ConnectSQL(host, port, username, pass, dbname string) (*DB, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username,
		pass,
		host,
		port,
		dbname,
	)
	openConnection, err := sql.Open(MYSQL_DRIVER_NAME, dataSourceName)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = openConnection
	return dbConn, err
}
