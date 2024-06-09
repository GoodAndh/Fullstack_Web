package testservice

import (
	"bytes"
	"encoding/json"
	"fullstack_toko/backend/model/web"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

// to get a new token its to run user test login ,get the token from its test message
func TestGetProduct(t *testing.T) {
	router := SetUpRouter()

	t.Run("get all product", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/apr", nil)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)

		assert.Equal(t, http.StatusOK, result.Code)

	})

	t.Run("get product by id success", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/apr/1", nil)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)

		assert.Equal(t, http.StatusOK, result.Code)

	})

	t.Run("get product by id fail", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/apr/100", nil)

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)

		assert.Equal(t, http.StatusBadRequest, result.Code)

	})

}

// to get token test the login user test
func TestCreateProduct(t *testing.T) {
	router := SetUpRouter()

	t.Run("create product success", func(t *testing.T) {
		payload := web.ProductCreatePayload{
			Name:      "rawon basi",
			Deskripsi: " nasi doang yg gagal dibuat(gk bisa masak nasi)",
			Category:  "consumable",
			Price:     10000,
			Quantity:  100,
			Url_image: "lol",
		}

		payloadBytes, err := json.Marshal(&payload)
		if err != nil {
			t.Fatal(err)
		}

		payloadReader := bytes.NewReader(payloadBytes)
		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/register/product", payloadReader)
		request.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3MTY5MTA4MjMsInVzZXJJRCI6IjEifQ.8_DU7oFZZioyMuhzU6aSEDMFEzvgXr1DmGyhCHApSkk")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)

		log.Println("result:", result)

		assert.Equal(t, http.StatusOK, result.Code)

	})

	t.Run("create product failed", func(t *testing.T) {
		payload := web.ProductCreatePayload{}

		payloadBytes, err := json.Marshal(&payload)
		if err != nil {
			t.Fatal(err)
		}

		payloadReader := bytes.NewReader(payloadBytes)
		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/register/product", payloadReader)
		request.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3MTYyMzAzMjYsInVzZXJJRCI6IjEifQ.W3UQ2bSVkXcmNKicRJuufjoHofRjwA3WsF2mMBLOUPU")
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)

		log.Println("result:", result)

		assert.Equal(t, http.StatusBadRequest, result.Code)

	})
}
