package handler

import "time"

type ProductReq struct {
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	Category     string  `json:"category"`
	Description  string  `json:"description"`
	Rating       int     `json:"rating"`
	NumReviews   int     `json:"num_reviews"`
	Price        float64 `json:"price"`
	CountInStock int     `json:"count_in_stock"`
}

type ProductRes struct {
	ID           int64      `json:"id"`
	Name         string     `json:"name"`
	Image        string     `json:"image"`
	Category     string     `json:"category"`
	Description  string     `json:"description"`
	Rating       int        `json:"rating"`
	NumReviews   int        `json:"num_reviews"`
	Price        float64    `json:"price"`
	CountInStock int        `json:"count_in_stock"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

//* ORDERS
type OrderReq struct {
	Items         []OrderItem `json:"items"`
	PaymentMethod string  `json:"payment_method"`
	TaxPrice      float64 `json:"tax_price"`
	ShippingPrice float64 `json:"shipping_price"`
	TotalPrice    float64 `json:"total_price"`
}

type OrderItem struct {
	Name      string  `json:"name"`
	Quantity  int     `json:"quantity"`
	Image     string  `json:"image"`
	Price     float64 `json:"price"`
	ProductID int64   `json:"product_id"`
}

type OrderRes struct {
	ID            int64      `json:"id"`
	Items         []OrderItem `json:"items"`
	PaymentMethod string     `json:"payment_method"`
	TaxPrice      float64    `json:"tax_price"`
	ShippingPrice float64    `json:"shipping_price"`
	TotalPrice    float64    `json:"total_price"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}