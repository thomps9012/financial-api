package requests

import (
	"context"
	conn "financial-api/m/db"
	user "financial-api/m/models/user"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Petty_Cash_Request struct {
	ID             string    `json:"id" bson:"_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	Date           time.Time `json:"date" bson:"date"`
	Description    string    `json:"description" bson:"description"`
	Amount         float64   `json:"amount" bson:"amount"`
	Receipts       []string  `json:"receipts" bson:"receipts"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	Action_History []Action  `json:"action_history" bson:"action_history"`
	Current_Status Status    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type Petty_Cash_Overview struct{}

type User_Petty_Cash struct{}

type Grant_Petty_Cash struct{}

func (p *Petty_Cash_Request) Create(user_id string) (string, error) {
	var petty_cash_req Petty_Cash_Request
	collection := conn.DB.Collection("petty_cash_requests")
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: p.Date}, {Key: "amount", Value: p.Amount}}
	err := collection.FindOne(context.TODO(), filter).Decode(&petty_cash_req)
	if err == nil {
		return "", fmt.Errorf("duplicate petty cash request")
	}
	p.ID = uuid.NewString()
	p.Created_At = time.Now()
	p.Is_Active = true
	p.Current_Status = PENDING
	p.User_ID = user_id
	first_action := &Action{
		ID:         uuid.NewString(),
		User_ID:    user_id,
		Status:     PENDING,
		Created_At: time.Now(),
	}
	p.Action_History = append(p.Action_History, *first_action)
	var req_user user.User
	manager_id, mgr_find_err := req_user.FindMgrID(user_id)
	if mgr_find_err != nil {
		panic(err)
	}
	var manager user.User
	update_user, update_err := manager.AddNotification(p.ID, manager_id)
	if update_err != nil {
		panic(err)
	}
	if !update_user {
		return "", err
	}
	return p.ID, nil
}

func (p *Petty_Cash_Request) Update(request Petty_Cash_Request, user_id string) (bool, error) {
	collection := conn.DB.Collection("petty_cash_requests")
	var petty_cash_req Petty_Cash_Request
	filter := bson.D{{Key: "request_id", Value: request.ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&petty_cash_req)
	if err != nil {
		panic(err)
	}
	if petty_cash_req.User_ID != user_id {
		panic("you are not the user who created this request")
	}
	current_status := petty_cash_req.Current_Status
	if current_status != PENDING && current_status != REJECTED {
		panic("this request is already being processed")
	}
	result, update_err := collection.UpdateByID(context.TODO(), request.ID, request)
	if update_err != nil {
		panic(update_err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (p *Petty_Cash_Request) Delete(request Petty_Cash_Request, user_id string) (bool, error) {
	collection := conn.DB.Collection("petty_cash_requests")
	var petty_cash_req Petty_Cash_Request
	filter := bson.D{{Key: "request_id", Value: request.ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&petty_cash_req)
	if err != nil {
		panic(err)
	}
	if petty_cash_req.User_ID != user_id {
		panic("you are not the user who created this request")
	}
	current_status := petty_cash_req.Current_Status
	if current_status != PENDING && current_status != REJECTED {
		panic("this request is already being processed")
	}
	result, update_err := collection.DeleteOne(context.TODO(), request.ID)
	if update_err != nil {
		panic(update_err)
	}
	if result.DeletedCount == 0 {
		return false, err
	}
	return true, nil
}
