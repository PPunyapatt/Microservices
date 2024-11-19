package routes

import (
	"order/adapters/handler"

	"github.com/labstack/echo"
)

func Setup(orderHandler *handler.OrderHandler) {
	e := echo.New()
	e.POST("/api/v1/order", orderHandler.CreateOrder)
	e.Logger.Fatal(e.Start(":8777"))
}
