package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/DanielDDHM/pool-api/config"
	"github.com/DanielDDHM/pool-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var users *mongo.Collection = config.GetCollections(config.DB, "User")
var validate = validator.New()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"error": validationErr.Error()})
			return
		}

		newUser := models.User{
			Id:        primitive.NewObjectID(),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			Phone:     user.Phone,
			BirthDate: user.Phone,
			IsAdmin:   user.IsAdmin,
		}

		userCreated, err := users.InsertOne(ctx, newUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
			return
		}

		// var userT *models.User
		// result := users.FindOne(ctx, bson.M{"id": userCreated.InsertedID})

		c.JSON(http.StatusCreated, map[string]interface{}{"user": userCreated.InsertedID})
	}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
