package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DBProvider(cnf DBConfig) *sql.DB {
	if db != nil {
		return db
	}
	
	dsn := fmt.Sprintf("%s:%s@(%s:%v)/%s", cnf.Username, cnf.Password, cnf.Host, cnf.Port, cnf.Database)
	db, err := sql.Open(cnf.Provider, dsn)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}