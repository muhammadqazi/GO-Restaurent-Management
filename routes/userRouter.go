package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/muhammadqazi/restaurent-management/controllers"
)

func UserRoutes(incommingRoutes *gin.Engine) {

	incommingRoutes.GET("/users", controller.GetUsers())
	incommingRoutes.GET("/users/:user_id", controller.GetUser())
	incommingRoutes.POST("/users/signup", controller.SignUp())
	incommingRoutes.POST("/users/signin ", controller.SignIn())
}
