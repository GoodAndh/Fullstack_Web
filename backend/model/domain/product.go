package domain

import "time"

type Product struct {
	Id        int
	Name      string
	Deskripsi string
	Category  string
	Price     float64
	Quantity  int
	Userid    int
	Url_image string
}

type ProductStat struct {
	Id        int
	CreatedAT time.Time
	LastUpdate time.Time
	ProductID int
}