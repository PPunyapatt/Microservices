package repository

import (
	"fmt"
	"order/entities"

	"gorm.io/gorm"
)

type orderRepoDB struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepoDB {
	return &orderRepoDB{
		db: db,
	}
}

func (o *orderRepoDB) CreateOrder(order *entities.Order, orderItems []*entities.OrderItem) error {
	fmt.Println("order: ", order)
	if err := o.db.Create(&order).Error; err != nil {
		return err
	}

	fmt.Println("orderItems: ", orderItems)
	if err := o.db.Create(orderItems).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepoDB) UpdateStatus() error {
	return nil
}
