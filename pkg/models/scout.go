package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Scout is scout
type Scout struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Firstname string             `bson:"firstname,omitempty" json:"firstname,omitempty"`
	Rank      string             `bson:"rank,omitempty" json:"rank,omitempty"`
	Den       int32              `bson:"den,omitempty" json:"den,omitempty"`
	Address   string             `bson:"address,omitempty" json:"address,omitempty"`
	Parents   string             `bson:"parents,omitempty" json:"parents,omitempty"`
	Phone     string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
}
