package handler

import (
	"fmt"
	"net/http"
	"order/entities"

	"github.com/labstack/echo"
)

type OrderHandler struct {
	svc entities.OrderUsecase
}

func NewOrderHandler(svc entities.OrderUsecase) *OrderHandler {
	return &OrderHandler{svc}
}

func (o *OrderHandler) CreateOrder(c echo.Context) error {
	var orderRequest *entities.OrderRequest
	fmt.Println("start")
	if err := c.Bind(&orderRequest); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// fmt.Println("orderRequest: ", *orderRequest)

	err := o.svc.CreateOrder(&orderRequest.Order, orderRequest.OrderItems)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return nil
}
