package testservice

import (
	"bytes"
	"encoding/json"
	"fullstack_toko/backend/model/web"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

// available token :
// "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3MTYyMzAzMjYsInVzZXJJRCI6IjEifQ.W3UQ2bSVkXcmNKicRJuufjoHofRjwA3WsF2mMBLOUPU"

// if err == malformed token its mean your token is invalid ,to get a new token is by login

func TestRegisterUser(t *testing.T) {
	roter := SetUpRouter()

	t.Run("register user failed", func(t *testing.T) {
		payload := web.UserRegisterPayload{
			Name:     "Zx",
			Username: "qwe",
			Password: "das",
			VPasword: "ad",
			Email:    "s",
		}

		payloadBytes, err := json.Marshal(&payload)
		if err != nil {
			t.Fatal(err)
		}
		payloadReader := bytes.NewReader(payloadBytes)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/register/users", payloadReader)
		request.Header.Add("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		roter.ServeHTTP(recorder, request)

		response := recorder.Result()

		// to get response body --->

		result, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}

		var webResponse web.WebResponse
		err = json.Unmarshal(result, &webResponse)
		if err != nil {
			t.Fatal(err)
		}

		log.Println("result:", webResponse)

		// --->

		assert.Equal(t, webResponse.Code, http.StatusBadRequest)
	})

	t.Run("register user success", func(t *testing.T) {

		// if the the output was bad request / 400 ,read the error message its neither username already in used or panic
		payload := web.UserRegisterPayload{
			Name:     "sikocaBerhasil",
			Username: "berhasilmas1", //berhasilmas
			Password: "berhasilmas1", //berhasilmas
			VPasword: "berhasilmas1",
			Email:    "berhasilmas1@gmail.com", //berhasilmas@gmail.com
		}

		payloadBytes, err := json.Marshal(&payload)
		if err != nil {
			t.Fatal(err)
		}
		payloadReader := bytes.NewReader(payloadBytes)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/register/users", payloadReader)
		request.Header.Add("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		roter.ServeHTTP(recorder, request)

		response := recorder.Result()

		// to get response body --->

		result, err := io.ReadAll(response.Body)
		if err != nil {
			t.Fatal(err)
		}

		var webResponse web.WebResponse
		err = json.Unmarshal(result, &webResponse)
		if err != nil {
			t.Fatal(err)
		}

		log.Println("result:", webResponse)

		// --->

		assert.Equal(t, response.Status, http.StatusOK)

	})

	

}

func TestLoginUser(t *testing.T) {
	router := SetUpRouter()

	t.Run("login user failed", func(t *testing.T) {
		requestBody := strings.NewReader(`{"username":"ini salah"}`)
		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", requestBody)


		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := recorder.Result()

		assert.Equal(t, response.StatusCode, 400)

	})

	t.Run("login user succed", func(t *testing.T) {
		payload := web.UserLoginPayload{
			Username: "berhasilmas1",
			Password: "berhasilmas1",
		}

		payloadBytes, err := json.Marshal(&payload)
		if err != nil {
			t.Fatal(err)
		}

		payloadReader := bytes.NewReader(payloadBytes)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/login", payloadReader)
		request.Header.Add("Content-Type", "application/json")

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)

		log.Println("your token :",result.Data)
		assert.Equal(t, 200, result.Code)

	})
}

func TestUserProfile(t *testing.T) {
	router := SetUpRouter()

	t.Run("update user profile fail", func(t *testing.T) {
		payload := web.UserProfileUpdatePayload{
			Url_image: "",
			Name:      "",
			Deskripsi: "",
		}
		payloadBytes, err := json.Marshal(&payload)
		if err != nil {
			t.Fatal(err)
		}

		payloadReader := bytes.NewReader(payloadBytes)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/update/profile", payloadReader)
		request.Header.Add("Content-Type", "application/json")
		// request.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3MTYyMzAzMjYsInVzZXJJRCI6IjEifQ.W3UQ2bSVkXcmNKicRJuufjoHofRjwA3WsF2mMBLOUPU")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)
		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)
		log.Println("result:", result)
		assert.Equal(t, result.Code, http.StatusForbidden)

	})

	t.Run("user profile update succcess", func(t *testing.T) {
		// if the test return error message "no rows affected" or status internal error its mean success but no rows was affected,try change the payload below

		payload := web.UserProfileUpdatePayload{
			Url_image: "ssdqe",
			Name:      "",
			Deskripsi: "",
		}

		payloadBytes, err := json.Marshal(&payload)
		if err != nil {
			t.Fatal(err)
		}

		payloadReader := bytes.NewReader(payloadBytes)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/update/profile", payloadReader)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3MTYyMzAzMjYsInVzZXJJRCI6IjEifQ.W3UQ2bSVkXcmNKicRJuufjoHofRjwA3WsF2mMBLOUPU")
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)
		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)
		log.Println("result:", result)
		assert.Equal(t, result.Code, http.StatusOK)

	})
}
