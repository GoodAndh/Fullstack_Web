package exception

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("very-super-not-secret-anymore")

func CreateJwt(secret []byte, userId int, timeIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": strconv.Itoa(userId),
		"exp":    time.Now().Add(timeIn).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJwt(t string) (*jwt.Token, error) {

	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return SecretKey, nil
	})
}

func SetCookie(w http.ResponseWriter, token, name string, duration time.Duration) {
	cookie := http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  time.Now().Add(duration),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)
}
