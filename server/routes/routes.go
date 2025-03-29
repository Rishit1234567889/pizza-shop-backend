package routes

import (
	"github.com/Rishit1234567889/pizza-shop/handler"
	"github.com/Rishit1234567889/pizza-shop/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, messagePublisher service.IMessagePublisher, websocketHandler handler.IWebSocketHandler) {

	router := r.Group("/")

	wsr := router.Group("/ws")
	{
		RegisterWebSocketRoutes(wsr, websocketHandler)
	}

	or := router.Group("/orders")
	{
		RegisterOrderRoutes(or, messagePublisher)
	}

}
