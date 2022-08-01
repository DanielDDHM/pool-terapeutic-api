package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pool struct {
	Id          primitive.ObjectID `bson:"id" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name" validate:"required"`
	Address     interface{}        `bson:"address" json:"address" validate:"required"`
	Equipaments Equipaments        `bson:"equipaments" json:"equipaments" validate:"required"`
	Environment string             `bson:"environment" json:"environment"`
	Center      Centers            `bson:"center" json:"center" validate:"required"`
	Photo       string             `bson:"photo" json:"photo"`
	IsActive    bool               `bson:"isactive" json:"isactive"`
}
