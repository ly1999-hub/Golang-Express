package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Phone string             `bson:"phone"`
	Age   int8               `bson:"age"`
}
