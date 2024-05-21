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

func TestCreateOrder(t *testing.T) {
	router := SetUpRouter()


	t.Run("order a product", func(t *testing.T) {

		// original route ("/api/v1/apr/:id/:params") ,change `id` to id product  and `params` into order

		payload := web.CartCheckoutPayload{
			Items: []web.CartCheckoutItem{
				{
					Quantity: 100,
				},
			},
		}
		payloadBytes, err := json.Marshal(&payload)
		if err != nil {
			t.Fatal(err)
		}

		payloadReader := bytes.NewReader(payloadBytes)

		request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/apr/1/order", payloadReader)

		// request.Header.Add("Authorization", "your token invalid lol")
		request.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3MTYyMzAzMjYsInVzZXJJRCI6IjEifQ.W3UQ2bSVkXcmNKicRJuufjoHofRjwA3WsF2mMBLOUPU")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)
		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)

		assert.Equal(t, 200, result)

	})

}

func TestGetOrder(t *testing.T) {
	router := SetUpRouter()

	t.Run("get all order", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/order/", nil)

		request.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkQXQiOjE3MTYyMzAzMjYsInVzZXJJRCI6IjEifQ.W3UQ2bSVkXcmNKicRJuufjoHofRjwA3WsF2mMBLOUPU")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, request)
		var webResponse web.WebResponse
		result := WebResponseUnmarshal(recorder, &webResponse)
		log.Println(result)
		assert.Equal(t, http.StatusOK, result.Code)
	})



}
