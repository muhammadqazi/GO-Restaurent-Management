package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/muhammadqazi/restaurent-management/controllers"
)

func OrderRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.GET("/orders", controller.GetOrders())
	incommingRoutes.GET("/orders/:order_id", controller.GetOrder())
	incommingRoutes.POST("/orders", controller.CreateOrder())
	incommingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
}
