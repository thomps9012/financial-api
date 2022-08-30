package requests

import (
	"context"
	"errors"
	conn "financial-api/db"
	grant "financial-api/models/grants"
	user "financial-api/models/user"
	"math"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Current_User   string    `json:"current_user" bson:"current_user"`
	Current_Status string    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type Petty_Cash_Overview struct {
	ID             string      `json:"id" bson:"_id"`
	User_ID        string      `json:"user_id" bson:"user_id"`
	User           user.User   `json:"user" bson:"user"`
	Grant_ID       string      `json:"grant_id" bson:"grant_id"`
	Amount         float64     `json:"amount" bson:"amount"`
	Grant          grant.Grant `json:"grant" bson:"grant"`
	Date           time.Time   `json:"date" bson:"date"`
	Created_At     time.Time   `json:"created_at" bson:"created_at"`
	Current_Status string      `json:"current_status" bson:"current_status"`
	Is_Active      bool        `json:"is_active" bson:"is_active"`
}

type User_Petty_Cash struct {
	User_ID      string               `json:"user_id" bson:"user_id"`
	User         user.User            `json:"user" bson:"user"`
	Total_Amount float64              `json:"total_amount" bson:"total_amount"`
	Requests     []Petty_Cash_Request `json:"requests" bson:"requests"`
	Last_Request Petty_Cash_Request   `json:"last_request" bson:"last_request"`
}

type Grant_Petty_Cash struct {
	Grant        grant.Grant          `json:"grant" bson:"grant"`
	Total_Amount float64              `json:"total_amount" bson:"total_amount"`
	Requests     []Petty_Cash_Request `json:"requests" bson:"requests"`
}

func (p *Petty_Cash_Request) FindByID(id string) (Petty_Cash_Request, error) {
	var petty_cash Petty_Cash_Request
	collection := conn.Db.Collection("petty_cash_requests")
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&petty_cash)
	if err != nil {
		panic(err)
	}
	return petty_cash, nil
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

func (p *Petty_Cash_Request) Create(requestor user.User) (Petty_Cash_Request, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	p.ID = uuid.NewString()
	p.Created_At = time.Now()
	p.Is_Active = true
	p.Current_Status = "PENDING"
	p.User_ID = requestor.ID
	p.Current_User = requestor.Manager_ID
	first_action := &Action{
		ID: uuid.NewString(),
		User: user.User_Action_Info{
			ID:   requestor.ID,
			Role: requestor.Role,
			Name: requestor.Name,
		},
		Request_Type: "petty_cash_requests",
		Request_ID:   p.ID,
		Status:       "PENDING",
		Created_At:   time.Now(),
	}
	p.Action_History = append(p.Action_History, *first_action)
	_, err := collection.InsertOne(context.TODO(), *p)
	if err != nil {
		panic(err)
	}
	var manager user.User
	update_user, update_err := manager.AddNotification(user.Action(*first_action), requestor.Manager_ID)
	if update_err != nil {
		panic(update_err)
	}
	if !update_user {
		return Petty_Cash_Request{}, update_err
	}
	return *p, nil
}

func (p *Petty_Cash_Request) Update(request Petty_Cash_Request, requestor user.User) (Petty_Cash_Request, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	if request.Current_Status == "REJECTED" {
		update_action := &Action{
			ID: uuid.NewString(),
			User: user.User_Action_Info{
				ID:   requestor.ID,
				Role: requestor.Role,
				Name: requestor.Name,
			},
			Request_Type: "petty_cash_requests",
			Request_ID:   request.ID,
			Status:       "REJECTED_EDIT",
			Created_At:   time.Now(),
		}
		request.Current_Status = "PENDING"
		request.Current_User = requestor.Manager_ID
		request.Action_History = append(request.Action_History, *update_action)
		var manager user.User
		update_user, update_err := manager.AddNotification(user.Action(*update_action), requestor.Manager_ID)
		if update_err != nil {
			panic(update_err)
		}
		if !update_user {
			return Petty_Cash_Request{}, errors.New("failed to update manager")
		}
	}
	var petty_cash_req Petty_Cash_Request
	filter := bson.D{{Key: "_id", Value: request.ID}}
	after := options.After
	opts := options.FindOneAndReplaceOptions{
		ReturnDocument: &after,
	}
	err := collection.FindOneAndReplace(context.TODO(), filter, request, &opts).Decode(&petty_cash_req)
	if err != nil {
		panic(err)
	}
	return petty_cash_req, nil
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
			Amount:         petty_cash_req.Amount,
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
func (p *Petty_Cash_Request) FindByUser(user_id string, start_date string, end_date string) (User_Petty_Cash, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	var filter bson.D
	layout := "2006-01-02T15:04:05.000Z"
	if start_date != "" && end_date != "" {
		start, err := time.Parse(layout, start_date)
		if err != nil {
			panic(err)
		}
		end, enderr := time.Parse(layout, end_date)
		if enderr != nil {
			panic(err)
		}
		filter = bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: bson.M{"$gte": start}}, {Key: "date", Value: bson.M{"$lte": end}}}
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
	last_request_date := time.Date(2000, time.April,
		34, 25, 72, 01, 0, time.UTC)
	var last_request Petty_Cash_Request
	var requests []Petty_Cash_Request
	for cursor.Next(context.TODO()) {
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requests = append(requests, petty_cash_req)
		if petty_cash_req.Date.After(last_request_date) {
			last_request = petty_cash_req
		}
		total_amount += math.Round(petty_cash_req.Amount*100) / 100
		total_amount = math.Round(total_amount*100) / 100
	}
	petty_cash_overview := &User_Petty_Cash{
		User_ID:      user_id,
		User:         user_info,
		Total_Amount: total_amount,
		Requests:     requests,
		Last_Request: last_request,
	}
	return *petty_cash_overview, nil
}
func (g *Grant_Petty_Cash) FindByGrant(grant_id string, start_date string, end_date string) (Grant_Petty_Cash, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	var filter bson.D
	layout := "2006-01-02T15:04:05.000Z"
	if start_date != "" && end_date != "" {
		start, err := time.Parse(layout, start_date)
		if err != nil {
			panic(err)
		}
		end, enderr := time.Parse(layout, end_date)
		if enderr != nil {
			panic(err)
		}
		filter = bson.D{{Key: "grant_id", Value: grant_id}, {Key: "date", Value: bson.M{"$gte": start}}, {Key: "date", Value: bson.M{"$lte": end}}}
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
	var requests []Petty_Cash_Request
	for cursor.Next(context.TODO()) {
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requests = append(requests, petty_cash_req)
		total_amount += petty_cash_req.Amount
	}
	petty_cash_overview := &Grant_Petty_Cash{
		Grant:        grant_info,
		Total_Amount: total_amount,
		Requests:     requests,
	}
	return *petty_cash_overview, nil
}
