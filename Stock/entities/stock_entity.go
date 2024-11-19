package entities

import "time"

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type Stock struct {
	ID                int       `json:"id"`
	ProductID         int       `json:"product_id"`
	AvailableQuantity int       `json:"available_quantity"`
	ReservedQuantity  int       `json:"reserved_quantity"`
	UpdateAt          time.Time `json:"update_at"`
}

type StockRepository interface {
	CreateProductAndStock(product *Product, stock *Stock) error
	ReservedStock(orderItems []interface{}) error
	ReleaseStock(orderItems []interface{}) error
	CommitStock(orderItems []interface{}) error
}

type StockUsecases interface {
	CreateProductAndStock(product *Product, stock *Stock) error
	ReservedStock(orderItems []interface{}) error
	ReleaseStock(orderItems []interface{}) error
	CommitStock(orderItems []interface{}) error
}
