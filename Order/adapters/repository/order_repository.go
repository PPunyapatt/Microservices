package repository

import (
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

func (o *orderRepoDB) CreateOrder(order entities.Order) error {
	return nil
}

func (o *orderRepoDB) UpdateStatus() error {
	return nil
}
