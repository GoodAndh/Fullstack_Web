package produk

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
	service     web.ProductService
	validator   *validator.Validate
	userService web.UserService
}

func NewHandler(service web.ProductService, validator *validator.Validate, userService web.UserService) *Handler {
	return &Handler{
		service:     service,
		validator:   validator,
		userService: userService,
	}
}

func (h *Handler) RegistierRoute(router *httprouter.Router) {
	router.GET("/api/v1/apr", h.handleAllProduct) //apr=all product

	router.GET("/api/v1/apr/:id", app.HandleWithMiddleware(h.handleGetById)) //used for searching product by the id

	router.POST("/api/v1/register/product", app.JwtMiddleware(h.handleRegister, h.userService))

}

func (h *Handler) handleAllProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodGet:
		pr, err := h.service.GetAllProduct(r.Context())
		if err != nil {
			exception.JsonBadRequest(w, err.Error())
			return
		}

		exception.WriteJson(w, http.StatusOK, "status ok ", "succes", pr)
	}

}

func (h *Handler) handleGetById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodGet:
		id := params.ByName("id")
		idNumber, err := strconv.Atoi(id)
		if err != nil {
			exception.JsonBadRequest(w, "only number accepted")
			return
		}

		pr, err := h.service.GetById(r.Context(), idNumber)
		if err != nil {
			exception.JsonBadRequest(w, err.Error())
			return
		}
		exception.WriteJson(w, http.StatusOK, "status ok ", "succes", pr)

	}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodPost:

		//get userID from context
		userID, ok := app.GetUserIDfromContext(r.Context())
		if !ok {
			exception.JsonUnauthorized(w, "invalid token")
			return
		}

		var payload web.ProductCreatePayload

		if err := exception.ParseJson(r, &payload); err != nil {
			log.Println("error while parsing payload,message:", err)
			exception.JsonInternalError(w, "server under maintenance")
			return
		}

		payload.Userid = userID

		if errorList := utils.ValidateStruct(h.validator, &payload); len(errorList) > 0 {
			exception.JsonBadRequest(w, "please fill the form", errorList)
			return
		}

		prID, err := h.service.CreateProduct(r.Context(), &payload)
		if err != nil {
			log.Println("create product error ,message:", err)
			exception.JsonInternalError(w, "server under maintenance")
			return
		}

		if err := h.service.CreateProductStat(r.Context(), prID); err != nil {
			log.Println("create product_stat error ,message:", err)
			exception.JsonInternalError(w, "server under maintenance")
			return
		}

		exception.WriteJson(w, http.StatusOK, "status ok", "succes create product", map[string]any{
			"productID": prID,
		})

	}
}
