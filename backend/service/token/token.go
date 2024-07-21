package token

import (
	"fullstack_toko/backend/exception"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func TokenHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoute(router *httprouter.Router) {
	router.POST("/api/v1/refresh-token", h.handleRefreshToken)
}

func (h *Handler) handleRefreshToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	Value := r.Header.Get("refresh_token")

	token, err := exception.ValidateJwt(Value)
	if err != nil {
		exception.JsonForbidden(w, err.Error(),nil)
		return
	}

	if !token.Valid {
		exception.JsonForbidden(w, "invalid refresh token",nil)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	str := claims["userID"].(string)
	userID, _ := strconv.Atoi(str)

	newToken, err := exception.CreateJwtAccesToken(exception.SecretKey, userID, time.Minute*10)
	if err != nil {
		exception.JsonInternalError(w, err.Error(),nil)
		return
	}

	exception.WriteJson(w, http.StatusOK, "status ok", "success", map[string]string{"auth": newToken})

}
