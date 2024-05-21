package web

import (
	"context"
	"time"
)

type CartCheckoutItem struct {
	ProductID int `json:"productID"`
	Quantity  int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items []CartCheckoutItem `json:"items" validate:"required"`
}

type OrderitemsCreatePayload struct {
	Status     string  `json:"status" validate:"required,oneof=pending cancelled succes"`
	Quantity   int     `json:"required" validate:"required,number"`
	TotalPrice float64 `json:"totalPrice" validate:"required"`
	ProductID  int     `json:"productID" validate:"required"`
	OrderID    int     `json:"orderID" validate:"required"`
	UserID     int     `json:"userID" validate:"required"`
}

type Orders struct {
	Id          int `json:"id"`
	LastUpdated time.Time `json:"lastUpdated"`
	UserID      int `json:"userID"`
}

type Order_items struct {
	Id         int `json:"id"`
	Status     string `json:"status"`
	Quantity   int `json:"quantity"`
	TotalPrice float64 `json:"totalPrice"`
	ProductID  int `json:"productID"`
	OrderID    int `json:"orderID"`
	UserID     int `json:"userID"`
}

type OrderService interface {
	CreateOrders(ctx context.Context, userID int) (int, error)
	CreateOrderitems(ctx context.Context, ord *OrderitemsCreatePayload) error
	SearchOrders(ctx context.Context,  userID int) ([]Orders, error)
	SearchOrderItems(ctx context.Context,  userID int, orderID []int) ([]Order_items, error)

}
