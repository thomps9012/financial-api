package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database
func InitDB() {
	// change on production
	os.Setenv("ATLAS_URI", "mongodb+srv://spars01:H0YXCAGHoUihHcSZ@cluster0.wuezj.mongodb.net/TEST_finance_records?retryWrites=true&w=majority")
	os.Setenv("DB_NAME", "TEST_finance_records")
	ATLAS_URI := os.Getenv("ATLAS_URI")
	clientOptions := options.Client().ApplyURI(ATLAS_URI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("this is the client err", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("this is the ping err", err)
	}

	fmt.Println("Connected to MongoDB!")

	// change on production
	dbName := os.Getenv("DB_NAME")
	Db = client.Database(dbName)
}

func CloseDB() {
	err := Db.Client().Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed!")
}
