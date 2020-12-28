package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Scout is scout
type Scout struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	firstname string             `bson:"firstname,omitempty" json:"firstname,omitempty"`
	rank      string             `bson:"rank,omitempty" json:"rank,omitempty"`
	den       int32              `bson:"den,omitempty" json:"den,omitempty"`
	address   string             `bson:"address,omitempty" json:"address,omitempty"`
	parents   string             `bson:"parents,omitempty" json:"parents,omitempty"`
	phone     string             `bson:"phone,omitempty" json:"phone,omitempty"`
	email     string             `bson:"email,omitempty" json:"email,omitempty"`
}
