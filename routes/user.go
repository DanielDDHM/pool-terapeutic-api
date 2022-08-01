package routes

import (
	"github.com/DanielDDHM/pool-api/controllers"
	"github.com/gin-gonic/gin"
)

func User(routes *gin.RouterGroup) {
	routes.POST("/", controllers.CreateUser())
	routes.GET("/:id", controllers.GetUser())
	routes.GET("/all", controllers.GetAllUser())
	routes.PUT("/:id", controllers.EditUser())
	routes.DELETE("/:id", controllers.DeleteUser())
}
