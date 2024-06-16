package utils

import (
	"fmt"
	"fullstack_toko/backend/model/domain"
	"fullstack_toko/backend/model/web"
	"net/http"

	"reflect"

	"github.com/go-playground/validator/v10"
)

func ConvertOrderItemsIntoWeb(ord *domain.Order_items) *web.Order_items {
	return &web.Order_items{
		Id:         ord.Id,
		Status:     ord.Status,
		Quantity:   ord.Quantity,
		TotalPrice: ord.TotalPrice,
		ProductID:  ord.ProductID,
		OrderID:    ord.OrderID,
		UserID:     ord.UserID,
	}
}

func ConvertOrderItemsIntoSlice(ord []domain.Order_items) []web.Order_items {
	or := make([]web.Order_items, 0, len(ord))
	for _, v := range ord {
		or = append(or, *ConvertOrderItemsIntoWeb(&v))
	}
	return or
}

func ConvertUserIntoWeb(u *domain.User) *web.User {
	return &web.User{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
}

func ConvertProfileIntoWeb(u *domain.UserProfile) *web.UserProfile {
	return &web.UserProfile{
		Id:        u.Id,
		Url_image: u.Url_image,
		Name:      u.Name,
		Deskripsi: u.Deskripsi,
		UserId:    u.UserId,
	}
}

func ConvertProductIntoWeb(p *domain.Product) *web.Product {
	return &web.Product{
		Id:        p.Id,
		Url_image: p.Url_image,
		Name:      p.Name,
		Deskripsi: p.Deskripsi,
		Price:     p.Price,
		Quantity:  p.Quantity,
		Userid:    p.Userid,
		Category:  p.Category,
	}
}

func ConvertOrdersIntoWeb(ord *domain.Orders) *web.Orders {
	return &web.Orders{
		Id:          ord.Id,
		LastUpdated: ord.LastUpdated,
		UserID:      ord.UserID,
	}
}

func ConvertOrdersIntoSlice(ord []domain.Orders) []web.Orders {
	slc := make([]web.Orders, 0, len(ord))

	for _, v := range ord {
		slc = append(slc, *ConvertOrdersIntoWeb(&v))
	}

	return slc
}

func ConvertProductIntoSlice(p []domain.Product) []web.Product {
	product := make([]web.Product, 0, len(p))
	for _, v := range p {
		product = append(product, *ConvertProductIntoWeb(&v))
	}
	return product
}

func ValidateStruct(validate *validator.Validate, stx any) map[string]any {
	errorList := make(map[string]any)
	if err := validate.Struct(stx); err != nil {
		errors := err.(validator.ValidationErrors)
		for _, er := range errors {
			va := reflect.TypeOf(stx).Elem()
			field, _ := va.FieldByName(er.StructField())
			fieldName := field.Tag.Get("json")
			var errMsg string
			switch er.Tag() {
			case "required":
				errMsg = fmt.Sprintf("form %v wajib di isi", fieldName)
			case "email":
				errMsg = fmt.Sprintf("form %v tidak sesuai format email", fieldName)
			case "eqfield":
				errMsg = fmt.Sprintf("form %v harus sama dengan form %v", fieldName, er.Param())
			case "oneof":
				errMsg = fmt.Sprintf("only accepted one of = %v", er.Param())
			case "min":
				errMsg = fmt.Sprintf("minimal di isi dengan %v karakter", er.Param())
			}
			errorList["error"+fieldName] = errMsg
		}
	}
	return errorList
}

func GetTokenAccessFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")

	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}
