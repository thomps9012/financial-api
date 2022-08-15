package requests

import (
	"context"
	conn "financial-api/m/db"
	grant "financial-api/m/models/grants"
	user "financial-api/m/models/user"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Petty_Cash_Request struct {
	ID             string    `json:"id" bson:"_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	Grant_ID       string    `json:"grant_id" bson:"grant_id"`
	Date           time.Time `json:"date" bson:"date"`
	Description    string    `json:"description" bson:"description"`
	Amount         float64   `json:"amount" bson:"amount"`
	Receipts       []string  `json:"receipts" bson:"receipts"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	Action_History []Action  `json:"action_history" bson:"action_history"`
	Current_Status string    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type Petty_Cash_Overview struct {
	ID             string      `json:"id" bson:"_id"`
	User_ID        string      `json:"user_id" bson:"user_id"`
	User           user.User   `json:"user" bson:"user"`
	Grant_ID       string      `json:"grant_id" bson:"grant_id"`
	Grant          grant.Grant `json:"grant" bson:"grant"`
	Date           time.Time   `json:"date" bson:"date"`
	Created_At     time.Time   `json:"created_at" bson:"created_at"`
	Current_Status string      `json:"current_status" bson:"current_status"`
	Is_Active      bool        `json:"is_active" bson:"is_active"`
}

type User_Petty_Cash struct {
	User_ID      string    `json:"user_id" bson:"user_id"`
	User         user.User `json:"user" bson:"user"`
	Last_Request time.Time `json:"last_request" bson:"last_request"`
	Total_Amount float64   `json:"total_amount" bson:"total_amount"`
}

type Grant_Petty_Cash struct {
	Grant_ID     string      `json:"grant_id" bson:"grant_id"`
	Grant        grant.Grant `json:"grant" bson:"grant"`
	Last_Request time.Time   `json:"last_request" bson:"last_request"`
	Total_Amount float64     `json:"total_amount" bson:"total_amount"`
}

func (p *Petty_Cash_Request) Exists(user_id string, amount float64, date time.Time) (bool, error) {
	var petty_cash_req Petty_Cash_Request
	collection := conn.Db.Collection("petty_cash_requests")
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: p.Date}, {Key: "amount", Value: p.Amount}}
	err := collection.FindOne(context.TODO(), filter).Decode(&petty_cash_req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *Petty_Cash_Request) Create(user_id string) (Petty_Cash_Request, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	p.ID = uuid.NewString()
	p.Created_At = time.Now()
	p.Is_Active = true
	p.Current_Status = "PENDING"
	p.User_ID = user_id
	first_action := &Action{
		ID:         uuid.NewString(),
		User_ID:    user_id,
		Status:     "PENDING",
		Created_At: time.Now(),
	}
	p.Action_History = append(p.Action_History, *first_action)
	_, err := collection.InsertOne(context.TODO(), *p)
	if err != nil {
		panic(err)
	}
	var req_user user.User
	manager_id, mgr_find_err := req_user.FindMgrID(user_id)
	if mgr_find_err != nil {
		panic(mgr_find_err)
	}
	var manager user.User
	update_user, update_err := manager.AddNotification(p.ID, manager_id)
	if update_err != nil {
		panic(update_err)
	}
	if !update_user {
		return Petty_Cash_Request{}, update_err
	}
	return *p, nil
}

func (p *Petty_Cash_Request) Update(request Petty_Cash_Request, user_id string) (bool, error) {
	collection := conn.Db.Collection("petty_cash_requests")
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
	if current_status != "PENDING" && current_status != "REJECTED" {
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
	collection := conn.Db.Collection("petty_cash_requests")
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
	if current_status != "PENDING" && current_status != "REJECTED" {
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

func (p *Petty_Cash_Overview) FindAll() ([]Petty_Cash_Overview, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	var overviews []Petty_Cash_Overview
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var user user.User
		var grant grant.Grant
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		grant_info, grant_err := grant.Find(petty_cash_req.Grant_ID)
		if grant_err != nil {
			panic(grant_err)
		}
		user_info, user_err := user.FindByID(petty_cash_req.User_ID)
		if user_err != nil {
			panic(user_err)
		}
		petty_cash_overview := &Petty_Cash_Overview{
			ID:             petty_cash_req.ID,
			User_ID:        petty_cash_req.User_ID,
			User:           user_info,
			Grant_ID:       petty_cash_req.Grant_ID,
			Grant:          grant_info,
			Date:           petty_cash_req.Date,
			Current_Status: petty_cash_req.Current_Status,
			Created_At:     petty_cash_req.Created_At,
			Is_Active:      petty_cash_req.Is_Active,
		}
		overviews = append(overviews, *petty_cash_overview)
	}
	return overviews, nil
}

// refactor inputs to start and end dates to allow for flexibility in data search
func (u *User_Petty_Cash) FindByUser(user_id string, start_date string, end_date string) (User_Petty_Cash, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	var filter bson.D
	if start_date != "" && end_date != "" {
		filter = bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: bson.M{"$gte": start_date}}, {Key: "date", Value: bson.M{"$lte": end_date}}}
	} else {
		filter = bson.D{{Key: "user_id", Value: user_id}}
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var user user.User
	user_info, user_err := user.FindByID(user_id)
	if user_err != nil {
		panic(user_err)
	}
	total_amount := 0.0
	last_request := time.Date(2020, time.April,
		34, 25, 72, 01, 0, time.UTC)
	for cursor.Next(context.TODO()) {
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		if petty_cash_req.Date.After(last_request) {
			last_request = petty_cash_req.Date
		}
		total_amount += petty_cash_req.Amount
	}
	petty_cash_overview := &User_Petty_Cash{
		User_ID:      user_id,
		User:         user_info,
		Last_Request: last_request,
		Total_Amount: total_amount,
	}
	return *petty_cash_overview, nil
}
func (g *Grant_Petty_Cash) FindByGrant(grant_id string, start_date string, end_date string) (Grant_Petty_Cash, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	var filter bson.D
	if start_date != "" && end_date != "" {
		filter = bson.D{{Key: "grant_id", Value: grant_id}, {Key: "$gte", Value: bson.M{"date": start_date}}, {Key: "$lte", Value: bson.M{"date": end_date}}}
	} else {
		filter = bson.D{{Key: "grant_id", Value: grant_id}}
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var grant grant.Grant
	grant_info, grant_err := grant.Find(grant_id)
	if grant_err != nil {
		panic(grant_err)
	}
	total_amount := 0.0
	last_request := time.Date(2020, time.April,
		34, 25, 72, 01, 0, time.UTC)
	for cursor.Next(context.TODO()) {
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		if petty_cash_req.Date.After(last_request) {
			last_request = petty_cash_req.Date
		}
		total_amount += petty_cash_req.Amount
	}
	petty_cash_overview := &Grant_Petty_Cash{
		Grant_ID:     grant_id,
		Grant:        grant_info,
		Last_Request: last_request,
		Total_Amount: total_amount,
	}
	return *petty_cash_overview, nil
}
