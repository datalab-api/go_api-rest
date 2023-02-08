package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Name   string             `json:"name,omitempty" validate:"required"`
	Adress string             `json:"adress,omitempty" validate:"required"`
	Phone  string             `json:"phone,omiteempty" validate:"required"`
}
