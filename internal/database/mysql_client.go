package database

import (
	"database/sql"
	"phones-review/internal/logs"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLClient struct {
	*sql.DB
}

func NewSQLClient(source string) *MySQLClient {
	db, err := sql.Open("mysql", source)

	if err != nil {
		logs.Log().Errorf("cannot create db tentant: %s", err.Error())
		panic("cannot create db")
	}

	err = db.Ping()

	if err != nil {
		logs.Log().Warn("cannot connect to mysql: %s", err.Error())
	}

	return &MySQLClient{db}
}

func (this *MySQLClient) ViewStats() sql.DBStats {
	return this.Stats()
}
