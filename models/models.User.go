package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty"`
	FullName string             `json:"full_name,omitempty" validate:"required"`
	UserName string             `json:"userName,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
