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

type Address struct {
	Website  string `json:"website" bson:"website"`
	Street   string `json:"street" bson:"street"`
	City     string `json:"city" bson:"city"`
	State    string `json:"state" bson:"state"`
	Zip_Code int    `json:"zip" bson:"zip"`
}

type Vendor struct {
	Name    string  `json:"name" bson:"name"`
	Address Address `json:"address" bson:"address"`
}

type Purchase struct {
	Grant_Line_Item string  `json:"line_item" bson:"line_item"`
	Description     string  `json:"description" bson:"description"`
	Amount          float64 `json:"amount" bson:"amount"`
}

type Check_Request struct {
	ID             string     `json:"id" bson:"_id"`
	Date           time.Time  `json:"date" bson:"date"`
	Vendor         Vendor     `json:"vendor" bson:"vendor"`
	Description    string     `json:"description" bson:"description"`
	Grant_ID       string     `json:"grant_id" bson:"grant_id"`
	Purchases      []Purchase `json:"purchases" bson:"purchases"`
	Receipts       []string   `json:"receipts" bson:"receipts"`
	Order_Total    float64    `json:"order_total" bson:"order_total"`
	Credit_Card    string     `json:"credit_card" bson:"credit_card"`
	Created_At     time.Time  `json:"created_at" bson:"created_at"`
	User_ID        string     `json:"user_id" bson:"user_id"`
	Action_History []Action   `json:"action_history" bson:"action_history"`
	Current_Status Status     `json:"current_status" bson:"current_status"`
	Is_Active      bool       `json:"is_active" bson:"is_active"`
}

type Check_Request_Overview struct {
	ID             string      `json:"id" bson:"_id"`
	User_ID        string      `json:"user_id" bson:"user_id"`
	User           user.User   `json:"user" bson:"user"`
	Grant_ID       string      `json:"grant_id" bson:"grant_id"`
	Grant          grant.Grant `json:"grant" bson:"grant"`
	Date           time.Time   `json:"date" bson:"date"`
	Vendor         Vendor      `json:"vendor" bson:"vendor"`
	Order_Total    float64     `json:"order_total" bson:"order_total"`
	Current_Status Status      `json:"current_status" bson:"current_status"`
	Created_At     time.Time   `json:"created_at" bson:"created_at"`
	Is_Active      bool        `json:"is_active" bson:"is_active"`
}
type User_Check_Overview struct {
	User_ID      string    `json:"user_id" bson:"user_id"`
	User         user.User `json:"user" bson:"user"`
	Vendors      []Vendor  `json:"vendors" bson:"vendors"`
	Total_Amount float64   `json:"total_amount" bson:"total_amount"`
	Credit_Cards string    `json:"credit_cards" bson:"credit_cards"`
	Last_Request time.Time `json:"last_request" bson:"last_request"`
}

type Grant_Check_Overview struct {
	Grant_ID     string      `json:"grant_id" bson:"grant_id"`
	Grant        grant.Grant `json:"grant" bson:"grant"`
	Vendors      []Vendor    `json:"vendors" bson:"vendors"`
	Total_Amount float64     `json:"total_amount" bson:"total_amount"`
	Credit_Cards string      `json:"credit_cards" bson:"credit_cards"`
	Last_Request time.Time   `json:"last_request" bson:"last_request"`
}

func (c *Check_Request) Exists(user_id string, vendor_name string, order_total float64, date time.Time) (bool, error) {
	collection := conn.Db.Collection("check_requests")
	var check_req Check_Request
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: date}, {Key: "order_total", Value: order_total}, {Key: "vendor.name", Value: vendor_name}}
	err := collection.FindOne(context.TODO(), filter).Decode(&check_req)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Check_Request) Create(user_id string) (string, error) {
	collection := conn.Db.Collection("check_requests")
	c.ID = uuid.NewString()
	c.Created_At = time.Now()
	c.Is_Active = true
	c.User_ID = user_id
	c.Current_Status = PENDING
	first_action := &Action{
		ID:         uuid.NewString(),
		User_ID:    user_id,
		Status:     "PENDING",
		Created_At: time.Now(),
	}
	c.Action_History = append(c.Action_History, *first_action)
	_, insert_err := collection.InsertOne(context.TODO(), *c)
	if insert_err != nil {
		panic(insert_err)
	}
	var req_user user.User
	manager_id, mgr_find_err := req_user.FindMgrID(user_id)
	if mgr_find_err != nil {
		panic(mgr_find_err)
	}
	// add in extra validation based on org chart here
	var manager user.User
	update_user, update_err := manager.AddNotification(c.ID, manager_id)
	if update_err != nil {
		panic(update_err)
	}
	if !update_user {
		return "", update_err
	}
	return c.ID, nil
}

func (c *Check_Request) Update(request Check_Request, user_id string) (bool, error) {
	collection := conn.Db.Collection("check_requests")
	var check_request Check_Request
	filter := bson.D{{Key: "request_id", Value: request.ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&check_request)
	if err != nil {
		panic(err)
	}
	if check_request.User_ID != user_id {
		panic("you are not the user who created this request")
	}
	current_status := check_request.Current_Status
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

func (c *Check_Request) Delete(request Check_Request, user_id string) (bool, error) {
	collection := conn.Db.Collection("check_requests")
	var check_request Check_Request
	filter := bson.D{{Key: "request_id", Value: request.ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&check_request)
	if err != nil {
		panic(err)
	}
	if check_request.User_ID != user_id {
		panic("you are not the user who created this request")
	}
	current_status := check_request.Current_Status
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

func (c *Check_Request_Overview) FindAll() ([]Check_Request_Overview, error) {
	collection := conn.Db.Collection("check_requests")
	var overviews []Check_Request_Overview
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var user user.User
		var grant grant.Grant
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		grant_info, grant_err := grant.Find(check_req.Grant_ID)
		if grant_err != nil {
			panic(grant_err)
		}
		user_info, user_err := user.FindByID(check_req.User_ID)
		if user_err != nil {
			panic(user_err)
		}
		check_overview := &Check_Request_Overview{
			ID:             check_req.ID,
			User_ID:        check_req.User_ID,
			User:           user_info,
			Grant_ID:       check_req.Grant_ID,
			Grant:          grant_info,
			Date:           check_req.Date,
			Vendor:         check_req.Vendor,
			Order_Total:    check_req.Order_Total,
			Current_Status: check_req.Current_Status,
			Created_At:     check_req.Created_At,
			Is_Active:      check_req.Is_Active,
		}
		overviews = append(overviews, *check_overview)
	}
	return overviews, nil
}

func (u *User_Check_Overview) FindByUser(user_id string, start_date string, end_date string) (User_Check_Overview, error) {
	collection := conn.Db.Collection("check_requests")
	var filter bson.D
	if start_date != "" && end_date != "" {
		filter = bson.D{{Key: "user_id", Value: user_id}, {Key: "$gte", Value: bson.M{"date": start_date}}, {Key: "$lte", Value: bson.M{"date": end_date}}}
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
	last_request := time.Date(2020, time.April,
		34, 25, 72, 01, 0, time.UTC)
	total_amount := 0.0
	var vendors []Vendor
	for cursor.Next(context.TODO()) {
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		vendors = append(vendors, check_req.Vendor)
		if check_req.Date.After(last_request) {
			last_request = check_req.Date
		}
		total_amount += check_req.Order_Total
	}
	check_overview := &User_Check_Overview{
		User_ID:      user_id,
		User:         user_info,
		Vendors:      vendors,
		Last_Request: last_request,
		Total_Amount: total_amount,
	}
	return *check_overview, nil
}

func (g *Grant_Check_Overview) FindByGrant(grant_id string, start_date string, end_date string) (Grant_Check_Overview, error) {
	collection := conn.Db.Collection("check_requests")
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
	last_request := time.Date(2020, time.April,
		34, 25, 72, 01, 0, time.UTC)
	total_amount := 0.0
	var vendors []Vendor
	for cursor.Next(context.TODO()) {
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		vendors = append(vendors, check_req.Vendor)
		if check_req.Date.After(last_request) {
			last_request = check_req.Date
		}
		total_amount += check_req.Order_Total
	}
	check_overview := &Grant_Check_Overview{
		Grant_ID:     grant_id,
		Grant:        grant_info,
		Vendors:      vendors,
		Last_Request: last_request,
		Total_Amount: total_amount,
	}
	return *check_overview, nil
}
