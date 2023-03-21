package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type DB struct {
}

var db *mongo.Client

func Init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	info := os.Getenv("MONGOURI")

	var connectErr error
	db, connectErr = ConnectDB(info)
	if connectErr != nil {
		log.Fatal(connectErr)
	}
}

func ConnectDB(datasource string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(datasource))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// ping to database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connect to MongoDB")
	return client, nil
}
