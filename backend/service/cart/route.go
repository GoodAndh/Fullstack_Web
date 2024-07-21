package cart

import (
	"fullstack_toko/backend/app"
	"fullstack_toko/backend/exception"
	"fullstack_toko/backend/model/web"
	"fullstack_toko/backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	userService    web.UserService
	productService web.ProductService
	orderService   web.OrderService
	validator      *validator.Validate
}

func NewHandler(userService web.UserService, productService web.ProductService, orderService web.OrderService, validator *validator.Validate) *Handler {
	return &Handler{
		userService:    userService,
		productService: productService,
		orderService:   orderService,
		validator:      validator,
	}
}

func (h *Handler) RegistierRoute(router *httprouter.Router) {
	router.POST("/api/v1/apr/:id/:params", app.JwtMiddleware(h.handleParams, h.userService)) //used for order product by id,return order id

	router.GET("/api/v1/order/", app.JwtMiddleware(h.handleGetOrder, h.userService))
}

func (h *Handler) handleParams(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodPost:
		id := params.ByName("id")
		productID, err := strconv.Atoi(id)
		if err != nil {
			exception.JsonBadRequest(w, "only number accepted",nil)
			return
		}

		prm := params.ByName("params")
		switch prm {
		case "order":
			//get userID from context
			userID, ok := app.GetUserIDfromContext(r.Context())
			if !ok {
				exception.JsonUnauthorized(w, "invalid token",nil)
				return
			}

			var cart web.CartCheckoutPayload
			if err := exception.ParseJson(r, &cart); err != nil {
				exception.JsonInternalError(w, "server undermaintenance",err.Error())
				return
			}

			for i := range cart.Items {
				cart.Items[i].ProductID = productID
			}

			if errorList := utils.ValidateStruct(h.validator, &cart); len(errorList) != 0 {
				exception.WriteJson(w, http.StatusBadRequest, "bad request", "please fill the form", errorList)
				return
			}

			//get the product by id
			pR, err := h.productService.GetById(r.Context(), productID)
			if err != nil {
				log.Println("error while search the product by id,message:", err)
				exception.JsonInternalError(w, "internal server error",err.Error())
				return
			}

			orderID, TotalPrice, err := h.createOrders(r.Context(), pR, cart.Items, userID)
			if err != nil {
				exception.JsonBadRequest(w, err.Error(),err.Error())
				return
			}

			exception.WriteJson(w, http.StatusOK, "status ok ", "succes create order", map[string]interface{}{
				"orderID":    orderID,
				"totalPrice": TotalPrice,
				"productID":  productID,
			})

		default:

			exception.WriteJson(w, http.StatusNotFound, "status not found ", "not found method", nil)

		}
	}
}

func (h *Handler) handleGetOrder(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodGet:

		userID, ok := app.GetUserIDfromContext(r.Context())
		if !ok {
			exception.JsonUnauthorized(w, "invalid token",nil)
			return
		}

		ordSlice, err := h.orderService.SearchOrders(r.Context(), userID)
		if err != nil {
			exception.JsonInternalError(w, "server undermaintenance",err.Error())
			return
		}

		orderID := make([]int, len(ordSlice))
		for i, v := range ordSlice {
			orderID[i] = v.Id
		}

		ordItemSlice, err := h.orderService.SearchOrderItems(r.Context(), userID, orderID)
		if err != nil {
			exception.JsonInternalError(w, "server undermaintenance",err.Error())
			return
		}
		

		dataList := map[string]any{
			"order":      ordSlice,
			"orderItems": ordItemSlice,
		}

		exception.WriteJson(w, http.StatusOK, "status ok", "order below", dataList)

	}
}
