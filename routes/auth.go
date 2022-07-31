package routes

import (
	"github.com/DanielDDHM/pool-api/controllers"
	"github.com/gin-gonic/gin"
)

func Auth(routes *gin.RouterGroup) {
	routes.POST("/login", controllers.Login)
	routes.POST("/logout", controllers.Logout)
}
