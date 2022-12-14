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
func (g *Grant) BulkInsert() (bool, error) {
	grants := []interface{}{
		Grant{
			ID:   "H79TI082369",
			Name: "BCORR"},
		Grant{
			ID:   "H79SP082264",
			Name: "HIV Navigator"},
		Grant{
			ID:   "H79SP082475",
			Name: "SPF (HOPE 1)"},
		Grant{
			ID:   "SOR_PEER",
			Name: "SOR Peer"},
		Grant{
			ID:   "SOR_HOUSING",
			Name: "SOR Recovery Housing"},
		Grant{
			ID:   "SOR_TWR",
			Name: "SOR 2.0 - Together We Rise"},
		Grant{
			ID:   "TANF",
			Name: "TANF"},
		Grant{
			ID:   "2020-JY-FX-0014",
			Name: "JSBT (OJJDP) - Jumpstart For A Better Tomorrow"},
		Grant{
			ID:   "SOR_LORAIN",
			Name: "SOR Lorain 2.0"},
		Grant{
			ID:   "H79SP081048",
			Name: "STOP Grant"},
		Grant{
			ID:   "H79TI083370",
			Name: "BSW (Bridge to Success Workforce)"},
		Grant{
			ID:   "H79SM085150",
			Name: "CCBHC"},
		Grant{
			ID:   "H79TI083662",
			Name: "IOP New Syrenity Intensive outpatient Program"},
		Grant{
			ID:   "H79TI085495",
			Name: "RAP AID (Recover from Addition to Prevent Aids)"},
		Grant{
			ID:   "H79TI085410",
			Name: "N MAT (NORA Medication-Assisted Treatment Program)"},
	}
	collection := conn.Db.Collection("grants")
	result, err := collection.InsertMany(context.TODO(), grants)
	if err != nil {
		panic(err)
	}
	for _, id := range result.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}
	return true, nil
}
