package models

import (
	"context"
	conn "financial-api/db"
	"financial-api/middleware"
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
	Category       Category  `json:"category" bson:"category"`
	Date           time.Time `json:"date" bson:"date"`
	Description    string    `json:"description" bson:"description"`
	Amount         float64   `json:"amount" bson:"amount"`
	Receipts       []string  `json:"receipts" bson:"receipts"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	Action_History []Action  `json:"action_history" bson:"action_history"`
	Current_User   string    `json:"current_user" bson:"current_user"`
	Current_Status Status    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type Petty_Cash_Overview struct {
	ID             string    `json:"id" bson:"_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	User           User      `json:"user" bson:"user"`
	Grant_ID       string    `json:"grant_id" bson:"grant_id"`
	Amount         float64   `json:"amount" bson:"amount"`
	Grant          Grant     `json:"grant" bson:"grant"`
	Date           time.Time `json:"date" bson:"date"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	Current_Status Status    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type User_Petty_Cash struct {
	User_ID      string               `json:"user_id" bson:"user_id"`
	User         User                 `json:"user" bson:"user"`
	Total_Amount float64              `json:"total_amount" bson:"total_amount"`
	Requests     []Petty_Cash_Request `json:"requests" bson:"requests"`
	Last_Request Petty_Cash_Request   `json:"last_request" bson:"last_request"`
}

type User_Monthly_Petty_Cash struct {
	ID           string     `json:"id" bson:"_id"`
	Name         string     `json:"name" bson:"name"`
	Month        time.Month `json:"month" bson:"month"`
	Year         int        `json:"year" bson:"year"`
	Total_Amount float64    `json:"total_amount" bson:"total_amount"`
	Request_IDS  []string   `json:"request_ids" bson:"request_ids"`
	Receipts     []string   `json:"receipts" bson:"receipts"`
}

type Grant_Petty_Cash struct {
	Grant          Grant                `json:"grant" bson:"grant"`
	Total_Requests int                  `json:"total_requests" bson:"total_requests"`
	Total_Amount   float64              `json:"total_amount" bson:"total_amount"`
	Requests       []Petty_Cash_Request `json:"requests" bson:"requests"`
}

func (p *Petty_Cash_Request) DeleteAll() bool {
	collection := conn.Db.Collection("petty_cash_requests")
	record_count, _ := collection.CountDocuments(context.TODO(), bson.D{{}})
	cleared, _ := collection.DeleteMany(context.TODO(), bson.D{{}})
	return cleared.DeletedCount == record_count
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

func (p *Petty_Cash_Request) Create(requestor User) (Petty_Cash_Request, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	p.ID = uuid.NewString()
	p.Created_At = time.Now()
	p.Is_Active = true
	p.Current_Status = PENDING
	p.User_ID = requestor.ID
	first_action := &Action{
		ID:           uuid.NewString(),
		User:         requestor.ID,
		Request_ID:   p.ID,
		Request_Type: PETTY_CASH,
		Status:       CREATED,
		Created_At:   time.Now(),
	}
	// current_user_email := UserEmailHandler(c.Category, PENDING, false)
	current_user_email := UserEmailHandler(p.Category, MANAGER_APPROVED, false)
	var user User
	current_user_id, err := user.FindID(current_user_email)
	if err != nil {
		panic(err)
	}
	p.Current_User = current_user_id
	p.Action_History = append(p.Action_History, *first_action)
	user.AddNotification(*first_action, current_user_id)
	_, insert_err := collection.InsertOne(context.TODO(), *p)
	if insert_err != nil {
		panic(insert_err)
	}
	update_user, update_err := middleware.SendEmail([]string{current_user_email}, "Petty Cash", requestor.Name, time.Now())
	if update_err != nil {
		panic(update_err)
	}
	if !update_user {
		return Petty_Cash_Request{}, update_err
	}
	return *p, nil
}

func (p *Petty_Cash_Request) Update(request Petty_Cash_Request, requestor User) (Petty_Cash_Request, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	if request.Current_Status == REJECTED {
		update_action := &Action{
			ID:           uuid.NewString(),
			User:         requestor.ID,
			Request_ID:   request.ID,
			Request_Type: PETTY_CASH,
			Status:       REJECTED_EDIT,
			Created_At:   time.Now(),
		}
		prev_user_id := ReturnPrevAdminID(request.Action_History, requestor.ID)
		request.Current_User = prev_user_id
		var user User
		current_user, err := user.FindByID(prev_user_id)
		if err != nil {
			panic(err)
		}
		var current_user_email = current_user.Email
		request.Action_History = append(request.Action_History, *update_action)
		_, clear_notification_err := current_user.ClearNotification(request.ID, requestor.ID)
		if clear_notification_err != nil {
			panic(clear_notification_err)
		}
		_, notify_err := current_user.AddNotification(*update_action, prev_user_id)
		if notify_err != nil {
			panic(notify_err)
		}
		update_user, update_err := middleware.SendEmail([]string{current_user_email}, "Check Request", requestor.Name, time.Now())
		if update_err != nil {
			panic(update_err)
		}
		if !update_user {
			panic("failed to update appropiate admin staff")
		}
	} else {
		update_action := &Action{
			ID:         uuid.NewString(),
			User:       requestor.ID,
			Request_ID: request.ID,
			Status:     EDIT,
			Created_At: time.Now(),
		}
		request.Action_History = append(request.Action_History, *update_action)
	}
	var petty_cash_req Petty_Cash_Request
	request.Current_Status = PENDING
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

func (p *Petty_Cash_Overview) FindAll() ([]Petty_Cash_Overview, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	var overviews []Petty_Cash_Overview
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var user User
		var grant Grant
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
	var user User
	user_info, user_err := user.FindByID(user_id)
	if user_err != nil {
		panic(user_err)
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
		total_amount += math.Round(petty_cash_req.Amount*100) / 100
		total_amount = math.Round(total_amount*100) / 100
	}
	petty_cash_overview := &User_Petty_Cash{
		User:         user_info,
		Total_Amount: total_amount,
		Requests:     requests,
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
	var grant Grant
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
		Grant:          grant_info,
		Total_Requests: len(requests),
		Total_Amount:   total_amount,
		Requests:       requests,
	}
	return *petty_cash_overview, nil
}

func (u *User) FindPettyCash(user_id string) (User_Petty_Cash, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	filter := bson.D{{Key: "user_id", Value: user_id}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
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
	return User_Petty_Cash{
		Total_Amount: total_amount,
		Requests:     requests,
		Last_Request: last_request,
		User_ID:      user_id,
	}, nil
}

type UserAggPettyCash struct {
	User_ID        string             `json:"user_id" bson:"user_id"`
	User           User               `json:"user" bson:"user"`
	Total_Amount   float64            `json:"total_amount" bson:"total_amount"`
	Total_Requests int                `json:"total_requests" bson:"total_requests"`
	Last_Request   Petty_Cash_Request `json:"last_request" bson:"last_request"`
}

func (u *User) AggUserPettyCash(user_id string) (UserAggPettyCash, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	filter := bson.D{{Key: "user_id", Value: user_id}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
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
	return UserAggPettyCash{
		Total_Amount:   total_amount,
		Total_Requests: len(requests),
		Last_Request:   last_request,
		User_ID:        user_id,
	}, nil
}

func (u *User) MonthlyPettyCash(user_id string, month int, year int) (User_Monthly_Petty_Cash, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
	end_month := month + 1
	start_date := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	end_date := time.Date(year, time.Month(end_month), 0, 0, 0, 0, 0, time.UTC)
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: bson.M{"$gte": start_date}}, {Key: "date", Value: bson.M{"$lte": end_date}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	total_amount := 0.0
	var receipts []string
	var requestIDs []string
	for cursor.Next(context.TODO()) {
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requestIDs = append(requestIDs, petty_cash_req.ID)
		receipts = append(receipts, petty_cash_req.Receipts...)
		total_amount += petty_cash_req.Amount
	}
	return User_Monthly_Petty_Cash{
		ID:           user_id,
		Name:         result.Name,
		Month:        time.Month(month),
		Year:         year,
		Total_Amount: total_amount,
		Request_IDS:  requestIDs,
		Receipts:     receipts,
	}, nil
}
