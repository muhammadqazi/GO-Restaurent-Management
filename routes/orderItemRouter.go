package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/muhammadqazi/restaurent-management/controllers"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.GET("/orderItems", controller.GetOrderItems())
	incommingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incommingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOrder())
	incommingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incommingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())
}
