package main

import (
	"fullstack_toko/backend/app"
	"fullstack_toko/backend/cmd/api"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := app.NewMysql(mysql.Config{
		User:                 "root",
		Passwd:               "r23password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "newshoponline",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(":3000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
