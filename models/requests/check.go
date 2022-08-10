package requests

import (
	"context"
	conn "financial-api/m/db"
	grant "financial-api/m/models/grants"
	user "financial-api/m/models/user"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Address struct {
	Website  string `json:"website" bson:"website"`
	Street   string `json:"street" bson:"street"`
	City     string `json:"city" bson:"city"`
	State    string `json:"state" bson:"state"`
	Zip_Code int64  `json:"zip" bson:"zip"`
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

// possibly add created at date in here
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
	ID             string    `json:"id" bson:"_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	User           user.User `json:"user" bson:"user"`
	Date           time.Time `json:"date" bson:"date"`
	Vendor         Vendor    `json:"vendor" bson:"vendor"`
	Order_Total    float64   `json:"order_total" bson:"order_total"`
	Credit_Card    string    `json:"credit_card" bson:"credit_card"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	Current_Status Status    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type Grant_Check_Overview struct {
	ID             string      `json:"id" bson:"_id"`
	Grant_ID       string      `json:"grant_id" bson:"grant_id"`
	Grant          grant.Grant `json:"grant" bson:"grant"`
	Date           time.Time   `json:"date" bson:"date"`
	Vendor         Vendor      `json:"vendor" bson:"vendor"`
	Order_Total    float64     `json:"order_total" bson:"order_total"`
	Current_Status Status      `json:"current_status" bson:"current_status"`
	Created_At     time.Time   `json:"created_at" bson:"created_at"`
	Is_Active      bool        `json:"is_active" bson:"is_active"`
}

func (c *Check_Request) Create(user_id string) (string, error) {
	collection := conn.DB.Collection("check_requests")
	var check_req Check_Request
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: c.Date}, {Key: "order_total", Value: c.Order_Total}, {Key: "vendor.name", Value: c.Vendor.Name}}
	err := collection.FindOne(context.TODO(), filter).Decode(&check_req)
	if err == nil {
		return "", fmt.Errorf("duplicate check request")
	}
	c.ID = uuid.NewString()
	c.Created_At = time.Now()
	c.Is_Active = true
	c.User_ID = user_id
	c.Current_Status = PENDING
	first_action := &Action{
		ID:         uuid.NewString(),
		User_ID:    user_id,
		Status:     PENDING,
		Created_At: time.Now(),
	}
	c.Action_History = append(c.Action_History, *first_action)
	var req_user user.User
	manager_id, mgr_find_err := req_user.FindMgrID(user_id)
	if mgr_find_err != nil {
		panic(err)
	}
	// add in extra validation based on org chart here
	var manager user.User
	update_user, update_err := manager.AddNotification(c.ID, manager_id)
	if update_err != nil {
		panic(err)
	}
	if !update_user {
		return "", err
	}
	return c.ID, nil
}

func (c *Check_Request) Update(request Check_Request, user_id string) (bool, error) {
	collection := conn.DB.Collection("check_request")
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
	collection := conn.DB.Collection("check_request")
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
	collection := conn.DB.Collection("check_request")
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

func (u *User_Check_Overview) FindByUser(user_id string) ([]User_Check_Overview, error) {
	var overviews []User_Check_Overview
	collection := conn.DB.Collection("check_request")
	filter := bson.D{{Key: "user_id", Value: user_id}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var user user.User
		var check_req User_Check_Overview
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		user_info, user_err := user.FindByID(check_req.User_ID)
		if user_err != nil {
			panic(user_err)
		}
		check_overview := &User_Check_Overview{
			ID:             check_req.ID,
			User_ID:        check_req.User_ID,
			User:           user_info,
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

func (g *Grant_Check_Overview) FindByGrant(grant_id string) ([]Grant_Check_Overview, error) {
	var overviews []Grant_Check_Overview
	collection := conn.DB.Collection("check_request")
	filter := bson.D{{Key: "grant_id", Value: grant_id}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var grant grant.Grant
		var check_req Grant_Check_Overview
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		grant_info, grant_err := grant.Find(check_req.Grant_ID)
		if grant_err != nil {
			panic(grant_err)
		}
		check_overview := &Grant_Check_Overview{
			ID:             check_req.ID,
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
