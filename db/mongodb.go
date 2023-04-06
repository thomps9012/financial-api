package database

import (
	"context"
	"financial-api/config"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var Collection *mongo.Collection

func Use(collection_name string) (*mongo.Collection, error) {
	ATLAS_URI := config.ENV("ATLAS_URI")
	clientOptions := options.Client().ApplyURI(ATLAS_URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	dbName := config.ENV("DB_NAME")
	DB = client.Database(dbName)
	Collection = DB.Collection(collection_name)
	return Collection, nil
}

func CloseDB() {
	err := DB.Client().Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed!")
}
