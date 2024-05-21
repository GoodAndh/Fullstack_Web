package testservice

import (
	"database/sql"
	"fullstack_toko/backend/app"
	"fullstack_toko/backend/service/cart"
	"fullstack_toko/backend/service/order"
	"fullstack_toko/backend/service/produk"
	"fullstack_toko/backend/service/user"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func SetUpRouter() http.Handler {
	router := httprouter.New()
	Db := DbStore()
	var validator *validator.Validate = validator.New()

	userRepo := user.NewRepository()
	userService := user.NewService(userRepo, Db)
	userHandler := user.NewHandler(userService, validator)
	userHandler.RegistierRoute(router)

	productRepo := produk.NewRepository()
	productService := produk.NewService(productRepo, Db)
	productHandler := produk.NewHandler(productService, validator, userService)
	productHandler.RegistierRoute(router)

	orderRepo := order.NewRepository()
	orderService := order.NewService(orderRepo, Db)
	cartHandler := cart.NewHandler(userService, productService, orderService, validator)
	cartHandler.RegistierRoute(router)

	return app.Middleware(router)
}

func NewMysql(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}

func DbStore() *sql.DB {
	db, err := NewMysql(mysql.Config{
		User:                 "root",
		Passwd:               "r23password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "newshoponline_test",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
