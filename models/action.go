package models

import (
	"context"
	conn "financial-api/db"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Action struct {
	ID         string    `json:"id" bson:"_id"`
	Request_ID string    `json:"request_id" bson:"request_id"`
	User       string    `json:"user" bson:"user"`
	Status     Status    `json:"status" bson:"status"`
	Created_At time.Time `json:"created_at" bson:"created_at"`
}

// func (a *Action) FindOne(request_id string, request_type string) (Request_Response, error) {
// 	collection := conn.Db.Collection(request_type)
// 	filter := bson.D{{Key: "_id", Value: request_id}}
// 	var mileage Mileage_Request
// 	var check Check_Request
// 	var petty Petty_Cash_Request
// 	switch request_type {
// 	case "mileage_requests":
// 		findErr := collection.FindOne(context.TODO(), filter).Decode(&mileage)
// 		if findErr != nil {
// 			panic(findErr)
// 		}
// 		return Request_Response{
// 			User_ID:        mileage.User_ID,
// 			Current_Status: mileage.Current_Status,
// 			Success:        true,
// 		}, nil
// 	case "check_requests":
// 		findErr := collection.FindOne(context.TODO(), filter).Decode(&check)
// 		if findErr != nil {
// 			panic(findErr)
// 		}
// 		return Request_Response{
// 			User_ID:        check.User_ID,
// 			Current_Status: check.Current_Status,
// 			Success:        true,
// 		}, nil

// 	case "petty_cash_requests":
// 		findErr := collection.FindOne(context.TODO(), filter).Decode(&petty)
// 		if findErr != nil {
// 			panic(findErr)
// 		}
// 		return Request_Response{
// 			User_ID:        petty.User_ID,
// 			Current_Status: petty.Current_Status,
// 			Success:        true,
// 		}, nil
// 	}
// 	return Request_Response{
// 		User_ID:        "",
// 		Current_Status: "ARCHIVED",
// 		Success:        false,
// 	}, errors.New("no request found")
// }

func (a *Action) Approve(request_id string, request_type string, new_status Status, request_category Category, exec_review bool) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(request_type)
	filter := bson.D{{Key: "_id", Value: request_id}}
	// possible expansion here
	var request Request_Info
	err := collection.FindOne(context.TODO(), filter).Decode(&request)
	if err != nil {
		panic(err)
	}
	if request.Current_Status == new_status {
		panic("current action has already been taken")
	}
	current_action := &Action{
		ID:         uuid.NewString(),
		Request_ID: request_id,
		User:       request.Current_User,
		Status:     new_status,
		Created_At: time.Now(),
	}
	current_user_email := UserEmailHandler(request_category, new_status, exec_review)
	var user User
	var user_id string
	if current_user_email == "" {
		user_id = request.User_ID
	} else {
		user_id, err = user.FindID(current_user_email)
		if err != nil {
			panic(err)
		}
	}
	// clear notification of previous user
	prev_user_clear_notification, err := user.ClearNotification(request_id, request.Current_User)
	if err != nil {
		panic(err)
	}
	if !prev_user_clear_notification {
		panic("error clearing the previous reviewer's notifications")
	}
	// updates the request
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_user": user_id}}, {Key: "$set", Value: bson.M{"current_status": new_status}}}
	updateDoc := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if updateDoc == nil {
		panic("error updating the document")
	}
	// add notification to new current user
	current_user_notified, err := user.AddNotification(*current_action, user_id)
	if err != nil {
		panic(err)
	}
	if !current_user_notified {
		panic("error notifying new request reviewer")
	}
	// notify original requestor
	requestor_notified, err := user.AddNotification(*current_action, request.User_ID)
	if err != nil {
		panic(err)
	}
	return requestor_notified, nil
}

func (a *Action) Reject(request_id string, request_type string) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(string(request_type))
	filter := bson.D{{Key: "_id", Value: request_id}}
	// possible expansion here
	var request Request_Info
	err := collection.FindOne(context.TODO(), filter).Decode(&request)
	if err != nil {
		panic(err)
	}
	if request.Current_Status == REJECTED {
		panic("current action has already been taken")
	}
	current_action := &Action{
		ID:         uuid.NewString(),
		Request_ID: request_id,
		User:       request.Current_User,
		Status:     REJECTED,
		Created_At: time.Now(),
	}
	var user User
	// clear prev user notification
	prev_user_clear_notification, err := user.ClearNotification(request_id, request.Current_User)
	if err != nil {
		panic(err)
	}
	if !prev_user_clear_notification {
		panic("error clearing the previous reviewer's notifications")
	}
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_user": request.User_ID}}, {Key: "$set", Value: bson.M{"current_status": "REJECTED"}}}
	// updates the request
	updateDoc := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if updateDoc == nil {
		panic("error updating the document")
	}
	// send notification to requestor
	requestor_notified, err := user.AddNotification(*current_action, request.User_ID)
	if err != nil {
		panic(err)
	}
	return requestor_notified, nil
}

func (a *Action) Archive(request_id string, request_type string, user_id string, admin bool) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(string(request_type))
	filter := bson.D{{Key: "_id", Value: request_id}}
	var request Request_Info
	err := collection.FindOne(context.TODO(), filter).Decode(&request)
	if err != nil {
		panic(err)
	}
	if !admin && user_id != request.User_ID {
		panic("you are unauthorized to archive this request")
	}
	if request.Current_Status == ARCHIVED {
		panic("current action has already been taken")
	}
	current_action := &Action{
		ID:         uuid.NewString(),
		Request_ID: request_id,
		User:       request.Current_User,
		Status:     ARCHIVED,
		Created_At: time.Now(),
	}
	var user User
	// clear prev user notification
	prev_user_clear_notification, err := user.ClearNotification(request_id, request.Current_User)
	if err != nil {
		panic(err)
	}
	if !prev_user_clear_notification {
		panic("error clearing the previous reviewer's notifications")
	}
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_status": ARCHIVED, "is_active": false}}, {Key: "$set", Value: bson.M{"current_user": bson.TypeNull}}}
	// updates the request
	updateDoc := collection.FindOneAndUpdate(context.TODO(), filter, update)
	if updateDoc == nil {
		panic("error updating the document")
	}
	// updates the original requestor
	requestor_notified, err := user.AddNotification(*current_action, request.User_ID)
	if err != nil {
		panic(err)
	}
	return requestor_notified, nil
}
