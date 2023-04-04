package models

import (
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MarketData struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Symbol       string             `bson:"symbol"`
	Price        decimal.Decimal    `bson:"price"`
	Volume       int64              `bson:"volume"`
	Open         decimal.Decimal    `bson:"open"`
	Close        decimal.Decimal    `bson:"close"`
	High         decimal.Decimal    `bson:"high"`
	Low          decimal.Decimal    `bson:"low"`
	TimeInterval string             `bson:"time_interval"`
	LastUpdated  time.Time          `bson:"last_updated"`
}
