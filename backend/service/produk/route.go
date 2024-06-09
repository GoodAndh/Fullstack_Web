package produk

import (
	"errors"
	"fullstack_toko/backend/app"
	"fullstack_toko/backend/exception"
	"fullstack_toko/backend/model/web"
	"fullstack_toko/backend/utils"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
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

	router.POST("/api/v1/upload/:productID", app.JwtMiddleware(h.handleUploadImg, h.userService))

	router.GET("/api/v1/public/:img", app.HandleWithMiddleware(h.handleShowIMG))

}

func (h *Handler) handleAllProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodGet:
		pr, err := h.service.GetAllProduct(r.Context())
		if err != nil {
			log.Println("error:", err)
			exception.JsonBadRequest(w, err.Error())
			w.Write([]byte(err.Error()))
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

func (h *Handler) handleUploadImg(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodPost:
		// get userID from context
		userID, ok := app.GetUserIDfromContext(r.Context())
		if !ok {
			exception.JsonUnauthorized(w, "invalid token")
			return
		}

		ps, err := h.service.GetByUserID(r.Context(), userID)
		if err != nil {
			exception.JsonInternalError(w, err.Error())
			return
		}

		// get productID from params
		prID := params.ByName("productID")
		pIDS, err := strconv.Atoi(prID)
		if err != nil {
			exception.JsonBadRequest(w, "number only accepted")
			return
		}

		// validate productIDs
		if err := checkIfIDisExist(ps, pIDS); err != nil {
			exception.JsonBadRequest(w, err.Error())
			return
		}

		file, fileHeader, err := r.FormFile("imageReceiver")
		if err != nil {
			exception.JsonInternalError(w, err.Error())
			return
		}
		defer file.Close()

		// check if directory alr exist

		_, err = os.Stat("./public/product")
		if errors.Is(err, fs.ErrNotExist) {
			err := os.MkdirAll("./public/product", os.ModePerm)
			if err != nil {
				exception.JsonInternalError(w, err.Error())
				return
			}
		}

		fileDestination, err := os.Create("./public/product/" + fileHeader.Filename)
		if err != nil {
			exception.JsonInternalError(w, err.Error())
			return
		}
		_, err = io.Copy(fileDestination, file)
		if err != nil {
			exception.JsonInternalError(w, err.Error())
			return
		}

		// update product image

		if err := h.service.UpdateProduct(r.Context(), &web.ProductUpdatePayload{
			Id:        pIDS,
			Userid:    userID,
			Url_image: fileHeader.Filename,
		}); err != nil {
			exception.JsonInternalError(w, err.Error())
			return
		}

		exception.WriteJson(w, http.StatusOK, "status ok", "sukses upload img", map[string]any{
			"file_path": fileHeader.Filename,
		})
		return

	}
}

func (h *Handler) handleShowIMG(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodGet:
		fileName := params.ByName("img")
		_, err := os.Stat("./public/product/" + fileName)
		if errors.Is(err, fs.ErrNotExist) {
			exception.JsonBadRequest(w, "cannot find filename")
			return
		}
		file, _ := os.Open("./public/product/" + fileName)
		defer file.Close()

		img, err := io.ReadAll(file)
		if err != nil {
			exception.JsonInternalError(w, err.Error())
			return
		}
		w.Write(img)
	}
}
