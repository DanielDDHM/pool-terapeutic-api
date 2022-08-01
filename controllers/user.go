package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/DanielDDHM/pool-api/config"
	"github.com/DanielDDHM/pool-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var users *mongo.Collection = config.GetCollections(config.DB, "User")
var validate = validator.New()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		id := c.Param("id")

		fmt.Println(id)

		var userT models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		fmt.Println(objId)

		err := users.FindOne(ctx, bson.M{"id": objId}).Decode(&userT)

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"data": userT})
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
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
			return
		}

		var createdUser models.User
		users.FindOne(ctx, bson.M{"_id": userCreated.InsertedID}).Decode(&createdUser)

		c.JSON(http.StatusCreated, map[string]interface{}{"user": createdUser})
	}
}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("id")
		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"data": err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"data": validationErr.Error()})
			return
		}

		update := bson.M{
			"name":      user.Name,
			"email":     user.Email,
			"password":  user.Password,
			"photo":     user.Photo,
			"phone":     user.Phone,
			"isadmin":   user.IsAdmin,
			"birthdate": user.BirthDate,
		}

		result, err := users.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
			return
		}

		//get updated user details
		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := users.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, map[string]interface{}{"data": updatedUser})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		result, err := users.DeleteOne(ctx, bson.M{"id": objId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound, map[string]interface{}{"data": "User with specified ID not found!"})
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{"data": "User successfully deleted!"})
	}
}

func GetAllUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var usersFind []models.User
		defer cancel()

		results, err := users.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{"Error": err.Error()})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)

		for results.Next(ctx) {
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, map[string]interface{}{"data": err.Error()})
			}

			usersFind = append(usersFind, singleUser)
		}

		c.JSON(http.StatusOK, map[string]interface{}{"data": usersFind})
	}
}
