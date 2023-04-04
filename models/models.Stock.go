package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Stock struct {
	Id       primitive.ObjectID `json:"_id,omitempty"`
	Symbol   string             `json:"full_name,omitempty" validate:"required"`
	Finance  string             `json:"userName,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
}
