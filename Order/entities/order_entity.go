package entities

import "time"

const (
	OrderStatusPending = "pending"
	OrderStatusPaid    = "paid"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Stock struct {
	ID                string    `json:"id"`
	ProductID         string    `json:"product_id"`
	AvailableQuantity int       `json:"available_quantity"`
	ReservedQuantity  int       `json:"reserved_quantity"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	Status     string    `json:"status"`
	TotalPrice float32   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID           int     `json:"id"`
	OrderID      int     `json:"order_id"`
	ProductID    int     `json:"product_id"`
	Quantity     int     `json:"quantity"`
	PricePerUnit float32 `json:"price_per_unit"`
}

type OrderRequest struct {
	Order      Order
	OrderItems []*OrderItem
}

type OrderUsecase interface {
	CreateOrder(order *Order, orderItems []*OrderItem) error
	UpdateStatus() error
}

type OrderRepository interface {
	CreateOrder(o *Order, orderItems []*OrderItem) error
	UpdateStatus() error
}
