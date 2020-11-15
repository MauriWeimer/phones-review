package database

import (
	"database/sql"
	"fmt"
)

type MySQLClient struct {
}

func NewSQLClient(source string) *sql.DB {
	db, err := sql.Open("mysql", source)

	if err != nil {
		_ = fmt.Errorf("cannot create db tentant: %s", err.Error())
		panic("cannot create db")
	}

	return db
}
