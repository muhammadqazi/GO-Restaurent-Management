package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/muhammadqazi/restaurent-management/controllers"
)

func TableRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.GET("/tables", controller.GetTables())
	incommingRoutes.GET("/tables/:table_id", controller.GetTable())
	incommingRoutes.POST("/tables", controller.CreateTable())
	incommingRoutes.PATCH("/tables/:table_id", controller.UpdateTable())
}
