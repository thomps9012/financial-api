package requests

import (
	"context"
	conn "financial-api/m/db"
	user "financial-api/m/models/user"
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
	ID         string    `json:"id" bson:"_id"`
	User_ID    string    `json:"user_id" bson:"user_id"`
	Status     Status    `json:"status" bson:"status"`
	Created_At time.Time `json:"created_at" bson:"created_at"`
}

type Request_Type string

const (
	MILEAGE    Request_Type = "mileage_requests"
	CHECK      Request_Type = "check_requests"
	PETTY_CASH Request_Type = "petty_cash_requests"
)

func Approve(request_id string, manager_id string, manager_role user.Role, request_type Request_Type) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(string(request_type))
	filter := bson.D{{Key: "_id", Value: request_id}}
	// possible expansion here
	var current_status Status
	switch manager_role {
	case user.MANAGER:
		current_status = MANAGER_APPROVED
	case user.FINANCE:
		current_status = FINANCE_APPROVED
	case user.EXECUTIVE:
		current_status = ORGANIZATION_APPROVED
	}
	current_action := &Action{
		ID:         uuid.NewString(),
		User_ID:    manager_id,
		Status:     current_status,
		Created_At: time.Now(),
	}
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_status": current_status}}}
	// updates the request
	err := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	// finds the manager of the person approving the request
	var req_user user.User
	manager_id, find_mgr_err := req_user.FindMgrID(manager_id)
	if find_mgr_err != nil {
		panic(err)
	}
	// adds the item_id to the manager of the person approving the request
	var manager user.User
	update_user, update_err := manager.AddNotification(request_id, manager_id)
	if update_err != nil {
		panic(err)
	}
	return update_user, nil
}

func Reject(request_id string, manager_id string, request_type Request_Type) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(string(request_type))
	var milage_req Mileage_Request
	// manager id is the id of the manager making the rejection
	current_action := &Action{
		ID:         uuid.NewString(),
		User_ID:    manager_id,
		Status:     REJECTED,
		Created_At: time.Now(),
	}
	filter := bson.D{{Key: "_id", Value: request_id}}
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_status": REJECTED}}}
	// updates the request
	err := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	milage_req.Current_Status = REJECTED
	// now adding a notification to the original user who made the request
	var user user.User
	update_user, update_err := user.AddNotification(milage_req.ID, milage_req.User_ID)
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
