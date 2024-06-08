package api

import (
	"database/sql"
	"fullstack_toko/backend/app"
	"fullstack_toko/backend/exception"
	"fullstack_toko/backend/service/cart"
	"fullstack_toko/backend/service/order"
	"fullstack_toko/backend/service/produk"
	"fullstack_toko/backend/service/token"
	"fullstack_toko/backend/service/user"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type ApiServer struct {
	Addr string
	Db   *sql.DB
}

func NewServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		Addr: addr,
		Db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := httprouter.New()
	router.NotFound =exception.NotFoundHandler()
	var validator *validator.Validate = validator.New()

	userRepo := user.NewRepository()
	userService := user.NewService(userRepo, s.Db)
	userHandler := user.NewHandler(userService, validator)
	userHandler.RegistierRoute(router)

	productRepo := produk.NewRepository()
	productService := produk.NewService(productRepo, s.Db)
	productHandler := produk.NewHandler(productService, validator, userService)
	productHandler.RegistierRoute(router)

	orderRepo := order.NewRepository()
	orderService := order.NewService(orderRepo, s.Db)
	cartHandler := cart.NewHandler(userService,productService,orderService,validator)
	cartHandler.RegistierRoute(router)

	token:=token.TokenHandler()
	token.RegisterRoute(router)

	handler:=cors.Default().Handler(router)
	return http.ListenAndServe(s.Addr, app.Middleware(handler))
}
