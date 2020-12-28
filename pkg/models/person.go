package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Person is person
type Person struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`

	Lastname  string `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Firstname string `bson:"firstname,omitempty" json:"firstname,omitempty"`
	Age       int    `bson:"age,omitempty" json:"age"`
}
