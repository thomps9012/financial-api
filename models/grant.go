package models

import (
	"context"
	database "financial-api/db"

	"go.mongodb.org/mongo-driver/bson"
)

type Grant struct {
	ID   string `json:"id" bson:"_id" validate:"required"`
	Name string `json:"name" bson:"name"`
}

func GetAllGrants() ([]Grant, error) {
	grants_coll, err := database.Use("grants")
	if err != nil {
		return []Grant{}, err
	}
	cursor, err := grants_coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []Grant{}, err
	}
	grants := make([]Grant, 0)
	err = cursor.All(context.TODO(), &grants)
	if err != nil {
		return []Grant{}, err
	}
	database.CloseDB()
	return grants, nil
}
func (g *Grant) GetOneGrant() (Grant, error) {
	grants_coll, err := database.Use("grants")
	if err != nil {
		return Grant{}, err
	}
	grant := new(Grant)
	filter := bson.D{{Key: "_id", Value: g.ID}}
	err = grants_coll.FindOne(context.TODO(), filter).Decode(&grant)
	if err != nil {
		return Grant{}, err
	}
	database.CloseDB()
	return *grant, nil
}
func (g *Grant) GetGrantMileage() ([]Mileage_Overview, error) {
	mileage_coll, err := database.Use("mileage_requests")
	if err != nil {
		return []Mileage_Overview{}, err
	}
	data := make([]Mileage_Overview, 0)
	filter := bson.D{{Key: "grant_id", Value: g.ID}}
	cursor, err := mileage_coll.Find(context.TODO(), filter)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	database.CloseDB()
	return data, nil
}
func (g *Grant) GetGrantPettyCash() ([]Petty_Cash_Overview, error) {
	petty_cash_coll, err := database.Use("petty_cash_requests")
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	data := make([]Petty_Cash_Overview, 0)
	filter := bson.D{{Key: "grant_id", Value: g.ID}}
	cursor, err := petty_cash_coll.Find(context.TODO(), filter)
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	database.CloseDB()
	return data, nil
}
func (g *Grant) GetGrantCheckRequest() ([]Check_Request_Overview, error) {
	check_req_coll, err := database.Use("check_requests")
	if err != nil {
		return []Check_Request_Overview{}, err
	}
	data := make([]Check_Request_Overview, 0)
	filter := bson.D{{Key: "grant_id", Value: g.ID}}
	cursor, err := check_req_coll.Find(context.TODO(), filter)
	if err != nil {
		return []Check_Request_Overview{}, err
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return []Check_Request_Overview{}, err
	}
	database.CloseDB()
	return data, nil
}
