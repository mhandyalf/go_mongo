package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name" validate:"required"`
	Age      int                `json:"age" bson:"age" validate:"required"`
	Position string             `json:"position" bson:"position"`
}
