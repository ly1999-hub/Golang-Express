package models

import "go.mongodb.org/mongo-driver/bson/primitive"
type ReqLogin struct {
	Username string `json:"username",xml:"username",form :"username",formvalue:"username",queryparam:"username" `
	Password string `json:"password",xml:"password",form :"password",formvalue:"password",queryparam:"password"`
}

type ResLogin struct {
	Token string `json:"token"`
}

type X struct {
	Text string `json:"text"`
}
type Users struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Phone string             `bson:"phone"`
	Age   int8               `bson:"age"`
}
