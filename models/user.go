package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"id" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name,omitempty" validate:"required"`
	Email     string             `bson:"email" json:"email,omitempty" validate:"required"`
	Password  string             `bson:"password" json:"password,omitempty" validate:"required"`
	Phone     string             `bson:"phone" json:"phone,omitempty"`
	BirthDate string             `bson:"birthdate" json:"birthdate,omitempty"`
	IsAdmin   bool               `bson:"isadmin" json:"isadmin,omitempty"`
}
