package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	FieldID   primitive.ObjectID `json:"field_id,omitempty" bson:"field_id,omitempty"`
	StartTime string             `json:"start_time" bson:"start_time"`
	EndTime   string             `json:"end_time" bson:"end_time"`
}
