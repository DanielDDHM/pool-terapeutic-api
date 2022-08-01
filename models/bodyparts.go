package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BodyParts struct {
	Id          primitive.ObjectID `bson:"id" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Environment string             `bson:"Environment" json:"Environment" validate:"required"`
	Info        interface{}        `bson:"info" json:"info" validate:"required"`
}
