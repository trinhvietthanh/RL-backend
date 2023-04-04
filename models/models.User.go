package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	FullName     string             `bson:"full_name,omitempty" validate:"required"`
	UserName     string             `bson:"userName,omitempty" validate:"required"`
	Password     string             `bson:"password,omitempty" validate:"required"`
	TraderOrders []TradeOrder       `bson:"traderOrders,omitempty""`
}
