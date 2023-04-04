package models

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TradeOrder struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Timestamp time.Time          `bson:"timestamp"`
	Symbol    string             `bson:"symbol"`
	OrderType string             `bson:"order_type"`
	Quantity  int                `bson:"quantity"`
	Price     decimal.Decimal    `bson:"price"`
	Status    string             `bson:"status"`
}
