package models

import (
	"context"
	conn "financial-api/db"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Grant struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

func (g *Grant) Find(grant_id string) (Grant, error) {
	collection := conn.Db.Collection("grants")
	var grant Grant
	filter := bson.D{{Key: "_id", Value: grant_id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&grant)
	if err != nil {
		panic(err)
	}
	return grant, nil
}
func (g *Grant) FindAll() ([]Grant, error) {
	collection := conn.Db.Collection("grants")
	var grantArr []Grant
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var grant Grant
		cursor.Decode(&grant)
		grantArr = append(grantArr, grant)
	}
	return grantArr, nil
}
func (g *Grant) BulkInsert() bool {
	grants := []interface{}{
		Grant{
			ID:   "H79TI082369",
			Name: "ROADS ACROSS AMERICA"},
		Grant{
			ID:   "H79SP082264",
			Name: "HOUSING"},
		Grant{
			ID:   "H79SP082475",
			Name: "PUBLIC TRANSIT"},
		Grant{
			ID:   "C718972H789087",
			Name: "PAPER AND PRINTING"},
		Grant{
			ID:   "Z651681AS6D5F15",
			Name: "SUPPLIES"},
		Grant{
			ID:   "G654A531F3A51",
			Name: "NONDESCRIPT GRANT"},
		Grant{
			ID:   "Y681651Y651YA3",
			Name: "NONDESCRIPT GRANT 2"},
	}
	collection := conn.Db.Collection("grants")
	result, err := collection.InsertMany(context.TODO(), grants)
	if err != nil {
		panic(err)
	}
	for _, id := range result.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}
	return true
}
