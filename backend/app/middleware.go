package app

import (
	"context"
	"fullstack_toko/backend/exception"
	"fullstack_toko/backend/model/web"
	"fullstack_toko/backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type contextKey string

const UserKey contextKey = "userID"

func Middleware(next *httprouter.Router) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		tokenString := utils.GetTokenFromRequest(r)
		//validate the JWT
		token, err := exception.ValidateJwt(tokenString)
		if err != nil {
			log.Println("error validate jwt middleware ,message:", err)
			exception.WriteJson(w, http.StatusForbidden, "status forbidden", "invalid token", nil)
			return
		}

		if !token.Valid {
			exception.WriteJson(w, http.StatusForbidden, "status forbidden", "invalid token", nil)
			return
		}

		//fetch the userid from db(id from token)
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)
		userID, _ := strconv.Atoi(str)
		u, err := userService.GetByID(r.Context(), userID)
		if err != nil {
			exception.WriteJson(w, http.StatusForbidden, "status forbidden", "invalid token", nil)
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
	
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1, false
	}

	return userID, true
}
