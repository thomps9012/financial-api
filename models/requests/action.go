package requests

import (
	"context"
	conn "financial-api/db"
	user "financial-api/models/user"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Status string

const (
	PENDING               Status = "PENDING"
	MANAGER_APPROVED      Status = "MANAGER_APPROVED"
	FINANCE_APPROVED      Status = "FINANACE_APPROVED"
	ORGANIZATION_APPROVED Status = "ORG_APPROVED"
	REJECTED              Status = "REJECTED"
	ARCHIVED              Status = "ARCHIVED"
)

type Action struct {
	ID           string    `json:"id" bson:"_id"`
	User         user.User `json:"user" bson:"user`
	Request_Type string    `json:"request_type" bson:"request_type"`
	Request_ID   string    `json:"request_id" bson:"request_id"`
	Status       string    `json:"status" bson:"status"`
	Created_At   time.Time `json:"created_at" bson:"created_at"`
}

type Request_Type string

const (
	MILEAGE    Request_Type = "mileage_requests"
	CHECK      Request_Type = "check_requests"
	PETTY_CASH Request_Type = "petty_cash_requests"
)

func Approve(request_id string, request_user_id string, manager_id string, manager_role string, request_type Request_Type) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(string(request_type))
	filter := bson.D{{Key: "_id", Value: request_id}}
	// possible expansion here
	var current_status string
	switch manager_role {
	case "MANAGER":
		current_status = "MANAGER_APPROVED"
	case "FINANCE":
		current_status = "FINANCE_APPROVED"
	case "EXECUTIVE":
		current_status = "ORGANIZATION_APPROVED"
	}
	var manager user.User
	manager_info, err := manager.FindByID(manager_id)
	if err != nil {
		panic(err)
	}
	current_action := &Action{
		ID:           uuid.NewString(),
		User:         manager_info,
		Request_Type: string(request_type),
		Request_ID:   request_id,
		Status:       current_status,
		Created_At:   time.Now(),
	}
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_status": current_status}}}
	// updates the request
	updateErr := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if updateErr != nil {
		panic(updateErr)
	}
	// adds the item to the manager of the person approving the request
	update_user, update_err := manager.AddNotification(user.Action(*current_action), manager_id)
	if update_err != nil {
		panic(err)
	}
	// adds the item to the original request user
	var request_user user.User
	_, requestErr := request_user.AddNotification(user.Action(*current_action), request_user_id)
	if requestErr != nil {
		panic(err)
	}
	return update_user, nil
}

func Reject(request_id string, request_user_id string, manager_id string, request_type Request_Type) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(string(request_type))
	// manager id is the id of the manager making the rejection
	var manager user.User
	manager_info, err := manager.FindByID(manager_id)
	if err != nil {
		panic(err)
	}
	current_action := &Action{
		ID:         uuid.NewString(),
		User:         manager_info,
		Request_Type: string(request_type),
		Request_ID:   request_id,
		Status:     "REJECTED",
		Created_At: time.Now(),
	}
	filter := bson.D{{Key: "_id", Value: request_id}}
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_status": REJECTED}}}
	// updates the request
	updateErr := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if updateErr != nil {
		panic(updateErr)
	}
	// now adding a notification to the original user who made the request
	var request_user user.User
	update_user, update_err := request_user.AddNotification(user.Action(*current_action), request_user_id)
	if update_err != nil {
		panic(err)
	}
	return update_user, nil
}

func Archive(request_id string, request_type Request_Type) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(string(request_type))
	filter := bson.D{{Key: "_id", Value: request_id}}
	update := bson.D{{Key: "$set", Value: bson.M{"current_status": ARCHIVED, "is_active": false}}}
	// updates the request
	err := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	return true, nil
}
