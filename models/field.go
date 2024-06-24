package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Field struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Price  float64            `json:"price" bson:"price"`
	Status string             `json:"status" bson:"status"`
}
