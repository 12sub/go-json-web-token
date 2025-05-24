package routes

import (
	controllers "github.com/12sub/go-jwt/controllers"
	"github.com/gin-gonic/gin"
)

// Routes to signup and login pages
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Signup())
	incomingRoutes.POST("users/login", controllers.Login())
}
