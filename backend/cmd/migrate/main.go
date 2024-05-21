package main

import (
	"fullstack_toko/backend/app"
	"log"
	"os"

	cfg "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// migrate database "mysql://root@tcp(localhost:3306)/database_baru1" -path cmd/migrate up
// migrate create -ext sql -dir [cmd/migrate] [namafile]
// powershell if migrate unkown command [$env:PATH += ";" + (go env GOPATH) + "\bin"]
//  migrate -path cmd/migrate -database "mysql://root:password@tcp(localhost:3306)/[databaseName]" version

func main() {
	db, err := app.NewMysql(cfg.Config{
		User:                 "root",
		Passwd:               "r23password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "newshoponline_test", //newshoponline_test ,newshoponline
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://", "mysql", driver)
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		err := m.Up()
		if err != nil {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		err := m.Down()
		if err != nil {
			log.Fatal(err)
		}
	}

}
