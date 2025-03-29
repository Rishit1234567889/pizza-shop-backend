package routes

import (
	"github.com/Rishit1234567889/pizza-shop/handler"
	"github.com/Rishit1234567889/pizza-shop/service"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(router *gin.RouterGroup, messagePublisher service.IMessagePublisher) {

	oh := handler.GetOrderHandler(messagePublisher)

	router.POST(
		"/create",
		oh.CreateOrder,
	)
}
