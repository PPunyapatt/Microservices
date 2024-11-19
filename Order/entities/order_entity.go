package entities

import "time"

const (
	OrderStatusPending = "pending"
	OrderStatusPaid    = "paid"
)

type Customer struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

type Product struct {
	ID        string    `gorm:"type:int;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Stock struct {
	ID                string    `gorm:"type:int;primaryKey" json:"id"`
	ProductID         string    `gorm:"type:int;not null" json:"product_id"`
	AvailableQuantity int       `gorm:"not null" json:"available_quantity"`
	ReservedQuantity  int       `gorm:"not null" json:"reserved_quantity"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Order struct {
	ID         string    `gorm:"type:int;primaryKey" json:"id"`
	CustomerID uint      `gorm:"not null" json:"customer_id"`
	Status     string    `gorm:"type:varchar(255);not null" json:"status"`
	TotalPrice float64   `gorm:"type:decimal(10,2);not null" json:"total_price"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type OrderItem struct {
	ID           string  `gorm:"type:int;primaryKey" json:"id"`
	OrderID      string  `gorm:"type:int;not null" json:"order_id"`
	ProductID    string  `gorm:"type:int;not null" json:"product_id"`
	Quantity     int     `gorm:"not null" json:"quantity"`
	PricePerUnit float64 `gorm:"type:decimal(10,2);not null" json:"price_per_unit"`
}

type OrderRequest struct {
	Order     Order
	OrderItem OrderItem
}

type OrderUsecase interface {
	CreateOrder(order Order, orderItem OrderItem) error
	UpdateStatus() error
}

type OrderRepository interface {
	CreateOrder(o Order) error
	UpdateStatus() error
}
