package user

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
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	service   web.UserService
	validator *validator.Validate
}

func NewHandler(service web.UserService, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

func (h *Handler) RegistierRoute(router *httprouter.Router) {
	router.POST("/api/v1/login", h.handleLogin)
	router.GET("/api/v1/me/", app.JwtMiddleware(h.handleUsers, h.service))

	router.POST("/api/v1/register/users", h.handleRegister)
	router.POST("/api/v1/register/users/img", h.handleUploadImg)

	router.POST("/api/v1/update/:params", app.JwtMiddleware(h.handleUpdateUsers, h.service))
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	var payload web.UserLoginPayload

	if err := exception.ParseJson(r, &payload); err != nil {
		exception.WriteJson(w, http.StatusInternalServerError, "internal server error", "server under maintenance", nil)
		return
	}

	u, err := h.service.GetUsername(r.Context(), payload.Username)
	if err != nil {
		exception.WriteJson(w, http.StatusBadRequest, "bad request", "invalid username or password", nil)
		return
	}

	if ok := exception.ComparePassword([]byte(payload.Password), u.Password); !ok {
		exception.WriteJson(w, http.StatusBadRequest, "bad request", "invalid username or password", nil)
		return
	}

	token, err := exception.CreateJwtAccesToken(exception.SecretKey, u.Id, time.Minute*15)
	if err != nil {
		exception.WriteJson(w, http.StatusInternalServerError, "status internal server error", err.Error(), nil)
		return
	}
	Refreshtoken, err := exception.CreateJwtRefreshToken(exception.SecretKey, u.Id, u.Email, time.Hour*1)
	if err != nil {
		exception.WriteJson(w, http.StatusInternalServerError, "status internal server error", err.Error(), nil)
		return
	}

	exception.WriteJson(w, http.StatusOK, "status ok", "success", map[string]string{"access_token": token,
		"refresh_token": Refreshtoken,
	})

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodPost:

		var payload web.UserRegisterPayload

		if err := exception.ParseJson(r, &payload); err != nil {
			exception.JsonInternalError(w, err.Error(), nil)
			return
		}

		if errorList := utils.ValidateStruct(h.validator, &payload); len(errorList) > 0 {
			exception.WriteJson(w, http.StatusBadRequest, "bad request", "please fill the field", errorList)
			return
		}

		hashedPassword, err := exception.HashPassword(payload.Password)
		if err != nil {
			log.Println("Hashed password error,message:", err)
			exception.JsonInternalError(w, err.Error(), nil)
			return
		}

		payload.Password = hashedPassword

		_, err = h.service.GetUsername(r.Context(), payload.Username)
		if err == nil {
			exception.JsonBadRequest(w, "username already in used", map[string]any{
				"errorusername": "username already in used",
			})
			return
		}

		_, err = h.service.GetByEmail(r.Context(), payload.Email)
		if err == nil {
			exception.JsonBadRequest(w, "email already in used", map[string]any{
				"erroremail": "email already in used",
			})
			return
		}

		userId, err := h.service.CreateUsers(r.Context(), &payload)
		if err != nil {
			log.Println("create users error,message:", err)
			exception.JsonInternalError(w, err.Error(), nil)

			return
		}

		if err := h.service.CreateUsersProfile(r.Context(), userId); err != nil {
			exception.JsonInternalError(w, err.Error(), nil)

			return
		}

		exception.WriteJson(w, http.StatusOK, "status ok", "create users succes", nil)

	}
}

func (h *Handler) handleUpdateUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	prm := params.ByName("params")
	switch prm {
	case "profile":
		//get userID from context
		userID, ok := app.GetUserIDfromContext(r.Context())
		if !ok {
			exception.JsonUnauthorized(w, "invalid token", nil)
			return
		}

		var payload web.UserProfileUpdatePayload

		if err := exception.ParseJson(r, &payload); err != nil {
			exception.JsonInternalError(w, "server under maintenance", err.Error())
			return
		}

		payload.UserId = userID

		if errorList := utils.ValidateStruct(h.validator, &payload); len(errorList) > 0 {
			exception.WriteJson(w, http.StatusBadRequest, "bad request", "please fill the field", errorList)
			return
		}

		if err := h.service.UpdateUserProfile(r.Context(), &payload); err != nil {

			if err.Error() == "no rows affeected" {
				exception.JsonBadRequest(w, "no rows affected", err.Error())
				return
			}

			exception.JsonInternalError(w, "server under maintenance", err.Error())

			return
		}

		exception.WriteJson(w, http.StatusOK, "status ok", "succes", nil)

	}

}

func (h *Handler) handleUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodGet:

		userID, ok := app.GetUserIDfromContext(r.Context())
		if !ok {
			exception.JsonUnauthorized(w, "invalid token", nil)
			return
		}

		users, err := h.service.GetByID(r.Context(), userID)
		if err != nil {
			exception.JsonForbidden(w, err.Error(), nil)
			return
		}

		profile, err := h.service.GetUserProfile(r.Context(), userID)
		if err != nil {
			exception.JsonForbidden(w, err.Error(), nil)
			return
		}

		exception.WriteJson(w, http.StatusOK, "status ok", "succes", map[string]any{
			"users":   users,
			"profile": profile,
		})
		return
	}
}

func (h *Handler) handleUploadImg(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	switch r.Method {
	case http.MethodPost:
		file, fileHeader, err := r.FormFile("imageReceiver")
		if err != nil {
			exception.JsonInternalError(w, err.Error(), nil)
			return
		}
		defer file.Close()

		//check if dir alr exist
		_, err = os.Stat("./public/users")
		if errors.Is(err, fs.ErrNotExist) {
			err := os.MkdirAll("./public/users", os.ModePerm)
			if err != nil {
				exception.JsonInternalError(w, err.Error(), nil)
				return
			}
		}

		fileDestination, err := os.Create("./public/users/" + fileHeader.Filename)
		if err != nil {
			exception.JsonInternalError(w, err.Error(), nil)
			return
		}

		_, err = io.Copy(fileDestination, file)
		if err != nil {
			exception.JsonInternalError(w, err.Error(), nil)
			return
		}

		exception.WriteJson(w, 200, "status ok ", "success upload profile image", map[string]any{
			"file_path": fileHeader.Filename,
		})
		return
	}
}
