package controllers

import (
	"backend-rltrading/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddTradeOrder(userID primitive.ObjectID, order models.TradeOrder) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	col := client.Database("mydb").Collection("users")

	filter := bson.M{"_id": userID}
	update := bson.M{"$push": bson.M{"trade_orders": order}}

	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func GetTraderOrders(userID primitive.ObjectID) ([]models.TradeOrder, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

}