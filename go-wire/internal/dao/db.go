package dao

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// var ProviderDB = wire.NewSet(NewDB)

func NewDB() (db *sql.DB, cf func(), err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/user?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4,utf8")
	cf = func() { fmt.Println("db close"); db.Close() }
	fmt.Println("mysql ping", db.Ping())
	return
}
