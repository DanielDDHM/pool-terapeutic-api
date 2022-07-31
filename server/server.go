package server

import (
	"log"
	"net/http"

	"github.com/DanielDDHM/pool-api/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func App() Server {
	return Server{
		port:   "4000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "UP",
			})
		})

		routes.Auth(v1.Group("/auth"))
		routes.User(v1.Group("/user"))
	}
	log.Print("Server is on port:", s.port)
	log.Fatal(router.Run(":" + s.port))
}
