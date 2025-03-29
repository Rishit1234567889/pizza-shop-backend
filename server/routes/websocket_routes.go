package routes

import (
	"github.com/Rishit1234567889/pizza-shop/handler"
	"github.com/gin-gonic/gin"
)

func RegisterWebSocketRoutes(router *gin.RouterGroup, websocketHandler handler.IWebSocketHandler) {
	router.GET(
		"/",
		websocketHandler.HandleConnection,
	)
}
