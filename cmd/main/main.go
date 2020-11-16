package main

import (
	"phones-review/gadgets/smartphones/web"
	"phones-review/internal/database"
	"phones-review/internal/logs"

	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 1
)

func main() {
	_ = logs.InitLogger()

	client := database.NewSQLClient("root:burbujas27@tcp(localhost:3306)/phones_review")
	doMigrate(client, "phones_review")

	handler := web.NewCreateSmartphoneHandler(client)
	mux := Routes(handler)
	server := NewServer(mux)

	server.Run()
}

func doMigrate(client *database.MySQLClient, dbName string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migrations version info is %d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Error("no migration needed")
	}
}
