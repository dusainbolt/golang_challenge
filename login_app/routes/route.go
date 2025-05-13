package routes

import (
	"login-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	app.POST("/signup", controllers.SignUp)
	app.POST("/signin", controllers.SignIn)
	app.POST("/refresh", controllers.Refresh)
}
