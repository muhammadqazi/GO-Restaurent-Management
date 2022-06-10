package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/muhammadqazi/restaurent-management/controllers"
)

func FoodRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.GET("/foods", controller.GetFoods())
	incommingRoutes.GET("/foods/:food_id", controller.GetFood())
	incommingRoutes.POST("/foods", controller.CreateFood())
	incommingRoutes.PATCH("/foods/:food_id", controller.UpdateFood())
}
