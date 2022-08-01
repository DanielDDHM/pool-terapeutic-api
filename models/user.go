package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"id" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"required"`
	Password  string             `bson:"password" json:"password" validate:"required"`
	Phone     string             `bson:"phone" json:"phone"`
	Photo     string             `bson:"photo" json:"photo"`
	BirthDate string             `bson:"birthdate" json:"birthdate"`
	IsAdmin   bool               `bson:"isadmin" json:"isadmin"`
}
