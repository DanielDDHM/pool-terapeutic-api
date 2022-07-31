package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Login",
	})
}

func Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Logout",
	})
}
