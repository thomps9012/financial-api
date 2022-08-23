package requests

import (
	"context"
	"errors"
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

type Request_Response struct {
	User_ID string `json:"user_id" bson:"user_id"`
	Success bool
}

func (a *Action) FindOne(request_id string, request_type string) (Request_Response, error) {
	collection := conn.Db.Collection(request_type)
	filter := bson.D{{Key: "_id", Value: request_id}}
	var mileage Mileage_Request
	var check Check_Request
	var petty Petty_Cash_Request
	switch request_type {
	case "mileage_requests":
		findErr := collection.FindOne(context.TODO(), filter).Decode(&mileage)
		if findErr != nil {
			panic(findErr)
		}
		return Request_Response{
			User_ID: mileage.User_ID,
			Success: true,
		}, nil
	case "check_requests":
		findErr := collection.FindOne(context.TODO(), filter).Decode(&check)
		if findErr != nil {
			panic(findErr)
		}
		return Request_Response{
			User_ID: check.User_ID,
			Success: true,
		}, nil

	case "petty_cash_requests":
		findErr := collection.FindOne(context.TODO(), filter).Decode(&petty)
		if findErr != nil {
			panic(findErr)
		}
		return Request_Response{
			User_ID: petty.User_ID,
			Success: true,
		}, nil
	}
	return Request_Response{
		User_ID: "",
		Success: false,
	}, errors.New("no request found")
}

func (a *Action) Approve(request_id string, request_user_id string, manager user.User, request_type string) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(request_type)
	filter := bson.D{{Key: "_id", Value: request_id}}
	// possible expansion here
	var current_status string
	switch manager.Role {
	case "MANAGER":
		current_status = "MANAGER_APPROVED"
	case "CHIEF":
		current_status = "CHIEF_APPROVED"
	case "FINANCE":
		current_status = "FINANCE_APPROVED"
	case "EXECUTIVE":
		current_status = "ORG_APPROVED"
	}
	current_action := &Action{
		ID:           uuid.NewString(),
		User:         manager,
		Request_Type: string(request_type),
		Request_ID:   request_id,
		Status:       current_status,
		Created_At:   time.Now(),
	}
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_user": manager.Manager_ID}}, {Key: "$set", Value: bson.M{"current_status": current_status}}}
	// updates the request
	updateDoc := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if updateDoc == nil {
		panic("error updating the document")
	}
	// adds the item to the manager of the person approving the request
	mgrNotified, update_err := manager.AddNotification(user.Action(*current_action), manager.Manager_ID)
	if update_err != nil {
		panic(update_err)
	}
	if !mgrNotified {
		panic("error notifying manager")
	}
	// clears the manager's notification queue
	updateCurrentMgr, clearQueueErr := manager.ClearNotification(request_id, manager.ID)
	if clearQueueErr != nil {
		panic(clearQueueErr)
	}
	if !updateCurrentMgr {
		panic("error clearing manager's notification")
	}
	// adds the item to the original request user
	var request_user user.User
	update_requestor, requestErr := request_user.AddNotification(user.Action(*current_action), request_user_id)
	if requestErr != nil {
		panic(requestErr)
	}
	return update_requestor, nil
}

func (a *Action) Reject(request_id string, request_user_id string, manager_id string, request_type string) (bool, error) {
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
		ID:           uuid.NewString(),
		User:         manager_info,
		Request_Type: string(request_type),
		Request_ID:   request_id,
		Status:       "REJECTED",
		Created_At:   time.Now(),
	}
	filter := bson.D{{Key: "_id", Value: request_id}}
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_user": request_user_id}}, {Key: "$set", Value: bson.M{"current_status": REJECTED}}}
	// updates the request
	updateDoc := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if updateDoc == nil {
		panic("error updating the document")
	}
	updateMgr, updateMgrErr := manager.ClearNotification(request_id, manager_id)
	if updateMgrErr != nil {
		panic(updateMgrErr)
	}
	if updateMgr != true {
		panic("error clearing manager notification")
	}
	// now adding a notification to the original user who made the request
	var request_user user.User
	update_user, update_err := request_user.AddNotification(user.Action(*current_action), request_user_id)
	if update_err != nil {
		panic(update_err)
	}
	return update_user, nil
}

func (a *Action) Archive(request_id string, request_type string, manager_id string) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(string(request_type))
	filter := bson.D{{Key: "_id", Value: request_id}}
	update := bson.D{{Key: "$set", Value: bson.M{"current_status": ARCHIVED, "is_active": false}}, {Key: "$set", Value: bson.M{"current_user": ""}}}
	// updates the request
	updateDoc := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if updateDoc == nil {
		panic("error updating the document")
	}
	var manager user.User
	updateMgr, updateMgrErr := manager.ClearNotification(request_id, manager_id)
	if updateMgrErr != nil {
		panic(updateMgrErr)
	}
	return updateMgr, nil
}
