package web

import "context"

type Product struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Deskripsi string  `json:"deskripsi"`
	Category  string  `json:"category"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Userid    int     `json:"userid"`
	Url_image string  `json:"url_image"`
}

type ProductCreatePayload struct {
	Name      string  `json:"name" validate:"required"`
	Deskripsi string  `json:"deskripsi" validate:"required"`
	Category  string  `json:"category" validate:"required,oneof=electric consumable etc"`
	Price     float64 `json:"price" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Userid    int     `json:"userid" validate:"required"`
	Url_image string  `json:"url_image" validate:"required"`
}

type ProductUpdatePayload struct {
	Id        int     `json:"id" validate:"required"`
	Name      string  `json:"name"`
	Deskripsi string  `json:"deskripsi"`
	Category  string  `json:"category"`
	Price     int `json:"price"`
	Quantity  int     `json:"quantity"`
	Userid    int     `json:"userid" validate:"required"`
	Url_image string  `json:"url_image"`
}

type ProductService interface {
	GetAllProduct(ctx context.Context) ([]Product, error)
	GetById(ctx context.Context, id int) (*Product, error)
	CreateProduct(ctx context.Context, pr *ProductCreatePayload) (int, error)
	CreateProductStat(ctx context.Context, productID int) error
	UpdateProduct(ctx context.Context,  ps *ProductUpdatePayload) error 
	GetByUserID(ctx context.Context, userID int) ([]Product, error)
}
