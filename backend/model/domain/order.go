package domain

import "time"

type Orders struct {
	Id          int
	LastUpdated time.Time
	UserID int
}

type Order_items struct {
	Id int
	Status string
	Quantity int
	TotalPrice float64
	ProductID int
	OrderID int
	UserID int
}