package app

import (
	"context"
	"fullstack_toko/backend/exception"
	"fullstack_toko/backend/model/web"
	"fullstack_toko/backend/utils"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type contextUserKey string

const UserKey contextUserKey = "userID"
const UserEmailKey contextUserKey = "email"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,refresh_token")
		next.ServeHTTP(w, r)
	})
}

func HandleWithMiddleware(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		handler(w, r, p)

	}
}

func JwtMiddleware(handler httprouter.Handle, userService web.UserService) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		//get token the from user request
		tokenString := utils.GetTokenAccessFromRequest(r)

		//validate the JWT
		token, err := exception.ValidateJwt(tokenString)
		if err != nil {
			exception.JsonUnauthorized(w, err.Error(),nil)
			return
		}

		if !token.Valid {
			exception.JsonUnauthorized(w, "invalid token",nil)
			return
		}

		//fetch the userid from db(id from token)
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)
		userID, _ := strconv.Atoi(str)
		u, err := userService.GetByID(r.Context(), userID)
		if err != nil {
			exception.JsonUnauthorized(w, err.Error(),nil)
			return
		}

		//set context userID to the userID
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.Id)

		r = r.WithContext(ctx)

		handler(w, r, p)
	}
}

// use this to get userID from the context
func GetUserIDfromContext(ctx context.Context) (int, bool) {

	userID := ctx.Value(UserKey).(int)
	return userID, userID > 0
}
