package models

import (
	"context"
	conn "financial-api/db"
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

func (a *Action) Get(request_id string, request_type Request_Type) (Request_Info, error) {
	collection := conn.Db.Collection(format_request_type(request_type))
	filter := bson.D{{Key: "_id", Value: request_id}}
	var request Request_Info
	err := collection.FindOne(context.TODO(), filter).Decode(&request)
	if err != nil {
		panic(err)
	}
	request.ID = request_id
	request.Type = request_type
	return request, nil
}

// test coverage
func (a *Action) Create(new_status Status, request_info Request_Info) Action {
	return Action{
		ID:           uuid.NewString(),
		Request_ID:   request_info.ID,
		Request_Type: request_info.Type,
		User:         request_info.Current_User,
		Status:       new_status,
		Created_At:   time.Now(),
	}
}
