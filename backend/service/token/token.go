package token

import (
	"fmt"
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
	router.GET("/api/v1/refresh-token", h.handleRefreshToken)
}

func (h *Handler) handleRefreshToken(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		exception.JsonUnauthorized(w, fmt.Sprintf("cannot find the cookie,message:%s", err.Error()))
		return
	}

	token, err := exception.ValidateJwt(cookie.Value)
	if err != nil {
		exception.JsonUnauthorized(w, err.Error())
		return
	}

	if !token.Valid {
		exception.JsonUnauthorized(w, "invalid token")
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	str := claims["userID"].(string)
	userID, _ := strconv.Atoi(str)

	newToken, err := exception.CreateJwt(exception.SecretKey, userID, time.Minute*15)
	if err != nil {
		exception.JsonInternalError(w, err.Error())
		return
	}
	exception.SetCookie(w, newToken, "Authorization", time.Minute*10)

	exception.WriteJson(w, http.StatusOK, "status ok", "success", map[string]string{"auth": newToken})

}
