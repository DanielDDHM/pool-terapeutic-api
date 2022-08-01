package routes

import (
	"github.com/DanielDDHM/pool-api/controllers"
	"github.com/gin-gonic/gin"
)

func BodyParts(routes *gin.RouterGroup) {
	routes.GET("/", controllers.GetAllBodyPart())
	routes.GET("/:id", controllers.GetBodyPart())
	routes.POST("/", controllers.CreateBodyPart())
	routes.PUT("/:id", controllers.EditBodyPart())
	routes.PUT("/:id/info", controllers.EditBodyPartInfo())
	routes.DELETE("/:id", controllers.DeleteBodyPart())
}
