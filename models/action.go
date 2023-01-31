package models

import (
	"context"
	conn "financial-api/db"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Action struct {
	ID           string       `json:"id" bson:"_id"`
	Request_ID   string       `json:"request_id" bson:"request_id"`
	Request_Type Request_Type `json:"request_type" bson:"request_type"`
	User         string       `json:"user" bson:"user"`
	Status       Status       `json:"status" bson:"status"`
	Created_At   time.Time    `json:"created_at" bson:"created_at"`
}

// test coverage
func ReturnPrevAdminID(action_history []Action, requestor_id string) string {
	for i := len(action_history) - 1; i > 0; i-- {
		if action_history[i].Status == REJECTED {
			return action_history[i].User
		}
	}
	return requestor_id
}

// test coverage
func format_request_type(request_type Request_Type) string {
	var lowered = strings.ToLower(string(request_type))
	var collection_name = lowered + "_requests"
	return collection_name
}

// func (a *Action) Create() error {
// 	var collection_name = format_request_type(a.Request_Type)
// 	var collection = db.GetCollection(collection_name)
// 	var err = collection.Insert(a)
// 	return err
// }

// func (a *Action) Get() error {
// 	var collection_name = format_request_type(a.Request_Type)
// 	var collection = db.GetCollection(collection_name)
// 	var err = collection.Find(bson.M{"_id": a.ID}).One(a)
// 	return err
// }

func (a *Action) Approve(request_id string, request_type Request_Type, new_status Status, request_category Category, exec_review bool) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(format_request_type(request_type))
	filter := bson.D{{Key: "_id", Value: request_id}}
	var request Request_Info
	err := collection.FindOne(context.TODO(), filter).Decode(&request)
	if err != nil {
		panic(err)
	}
	if request.Current_Status == new_status {
		panic("current action has already been taken")
	}
	current_action := &Action{
		ID:           uuid.NewString(),
		Request_ID:   request_id,
		Request_Type: request_type,
		User:         request.Current_User,
		Status:       new_status,
		Created_At:   time.Now(),
	}
	current_user_email := UserEmailHandler(request_category, new_status, exec_review)
	var user User
	var user_id string
	fmt.Println("\n current user email: ", current_user_email)
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
	if request.User_ID != user_id {
		// notify original requestor if they are not the same as the new current user
		requestor_notified, err := user.AddNotification(*current_action, request.User_ID)
		if err != nil {
			panic(err)
		}
		return requestor_notified, nil
	}
	return current_user_notified, nil
}

func (a *Action) Reject(request_id string, request_type Request_Type) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(format_request_type(request_type))
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
		ID:           uuid.NewString(),
		Request_ID:   request_id,
		Request_Type: request_type,
		User:         request.Current_User,
		Status:       REJECTED,
		Created_At:   time.Now(),
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
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_user": request.User_ID}}, {Key: "$set", Value: bson.M{"current_status": REJECTED}}}
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

func (a *Action) Archive(request_id string, request_type Request_Type, user_id string, admin bool) (bool, error) {
	// request type will be collection name
	// i.e. mileage_requests
	collection := conn.Db.Collection(format_request_type(request_type))
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
		ID:           uuid.NewString(),
		Request_ID:   request_id,
		Request_Type: request_type,
		User:         request.Current_User,
		Status:       ARCHIVED,
		Created_At:   time.Now(),
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
	update := bson.D{{Key: "$push", Value: bson.M{"action_history": *current_action}}, {Key: "$set", Value: bson.M{"current_status": ARCHIVED, "is_active": false}}, {Key: "$set", Value: bson.M{"current_user": "null"}}}
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
