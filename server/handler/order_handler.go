package handler

import (
	"github.com/Rishit1234567889/pizza-shop/constants"
	"github.com/Rishit1234567889/pizza-shop/service"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	messagePublisher service.IMessagePublisher
}

func (oh *OrderHandler) CreateOrder(ctx *gin.Context) {
	var payload map[string]interface{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(400, gin.H{
			"message":    "Bad Request",
			"statusCode": 400,
		})
	}
	payload["order_status"] = constants.ORDER_ORDERED
	oh.messagePublisher.PublishEvent(constants.KITCHEN_ORDER_QUEUE, payload)

	ctx.JSON(200, gin.H{
		"data":       payload,
		"statusCode": 200,
		"message":    "order accepted successfully",
	})
}

func GetOrderHandler(messagePublisher service.IMessagePublisher) *OrderHandler {
	return &OrderHandler{
		messagePublisher: messagePublisher,
	}
}
