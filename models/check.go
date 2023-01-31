package models

import (
	"context"
	conn "financial-api/db"
	"financial-api/middleware"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Category       Category   `json:"category" bson:"category"`
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
	Current_User   string     `json:"current_user" bson:"current_user"`
	Current_Status Status     `json:"current_status" bson:"current_status"`
	Is_Active      bool       `json:"is_active" bson:"is_active"`
}
type Check_Request_Overview struct {
	ID             string    `json:"id" bson:"_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	User           User      `json:"user" bson:"user"`
	Grant_ID       string    `json:"grant_id" bson:"grant_id"`
	Grant          Grant     `json:"grant" bson:"grant"`
	Date           time.Time `json:"date" bson:"date"`
	Vendor         Vendor    `json:"vendor" bson:"vendor"`
	Order_Total    float64   `json:"order_total" bson:"order_total"`
	Current_Status Status    `json:"current_status" bson:"current_status"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

func (c *Check_Request) DeleteAll() bool {
	collection := conn.Db.Collection("check_requests")
	record_count, _ := collection.CountDocuments(context.TODO(), bson.D{{}})
	cleared, _ := collection.DeleteMany(context.TODO(), bson.D{{}})
	return cleared.DeletedCount == record_count
}
func (c *Check_Request) Exists(user_id string, vendor_name string, order_total float64, date time.Time) bool {
	collection := conn.Db.Collection("check_requests")
	var check_req Check_Request
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: date}, {Key: "order_total", Value: order_total}, {Key: "vendor.name", Value: vendor_name}}
	err := collection.FindOne(context.TODO(), filter).Decode(&check_req)
	if err != nil {
		return false
	}
	return true
}

func (c *Check_Request) Create(requestor User) (string, error) {
	collection := conn.Db.Collection("check_requests")
	c.ID = uuid.NewString()
	c.Created_At = time.Now()
	c.Is_Active = true
	c.User_ID = requestor.ID
	c.Current_Status = PENDING
	// build in logic for setting current user id
	current_user_email := UserEmailHandler(c.Category, PENDING, false)
	fmt.Printf("current user email: %s", current_user_email)
	var user User
	// this breaks if current user email is not in databse
	// on production all managers will need to create signed accounts
	current_user_id, err := user.FindID(current_user_email)
	if err != nil {
		panic(err)
	}
	c.Current_User = current_user_id
	first_action := &Action{
		ID:           uuid.NewString(),
		User:         requestor.ID,
		Request_ID:   c.ID,
		Request_Type: CHECK,
		Status:       CREATED,
		Created_At:   time.Now(),
	}
	c.Action_History = append(c.Action_History, *first_action)
	user.AddNotification(*first_action, current_user_id)
	_, insert_err := collection.InsertOne(context.TODO(), *c)
	if insert_err != nil {
		panic(insert_err)
	}
	// add in extra validation based on org chart here
	update_user, update_err := middleware.SendEmail([]string{current_user_email}, "Check Request", requestor.Name, time.Now())
	if update_err != nil {
		panic(update_err)
	}
	if !update_user {
		return "", update_err
	}
	return c.ID, nil
}
func (c *Check_Request) Update(request Check_Request, requestor User) (Check_Request, error) {
	collection := conn.Db.Collection("check_requests")
	if request.Current_Status == REJECTED {
		update_action := &Action{
			ID:           uuid.NewString(),
			User:         requestor.ID,
			Request_ID:   request.ID,
			Request_Type: CHECK,
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
	// add in edit tracker
	var check_request Check_Request
	request.Current_Status = PENDING
	filter := bson.D{{Key: "_id", Value: request.ID}}
	after := options.After
	opts := options.FindOneAndReplaceOptions{
		ReturnDocument: &after,
	}
	err := collection.FindOneAndReplace(context.TODO(), filter, request, &opts).Decode(&check_request)
	if err != nil {
		panic(err)
	}
	return check_request, nil
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
		println("hit cursor")
		var user User
		var grant Grant
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			println("decode error")
			panic(decode_err)
		}
		println("grant ID", check_req.Grant_ID)
		grant_info, grant_err := grant.Find(check_req.Grant_ID)
		if grant_err != nil {
			println("grant error")
			panic(grant_err)
		}
		println("user ID", check_req.User_ID)
		user_info, user_err := user.FindByID(check_req.User_ID)
		if user_err != nil {
			println("user error")
			panic(user_err)
		}
		println(check_req.ID)
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
func (c *Check_Request) FindByID(check_id string) (Check_Request, error) {
	collection := conn.Db.Collection("check_requests")
	var check_req Check_Request
	filter := bson.D{{Key: "_id", Value: check_id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&check_req)
	if err != nil {
		panic(err)
	}
	return check_req, nil
}

type Grant_Check_Overview struct {
	Grant          Grant           `json:"grant" bson:"grant"`
	Vendors        []Vendor        `json:"vendors" bson:"vendors"`
	Total_Amount   float64         `json:"total_amount" bson:"total_amount"`
	Total_Requests int             `json:"total_requests" bson:"total_requests"`
	Credit_Cards   []string        `json:"credit_cards" bson:"credit_cards"`
	Requests       []Check_Request `json:"request_ids" bson:"request_ids"`
}

func (g *Grant_Check_Overview) FindByGrant(grant_id string, start_date string, end_date string) (Grant_Check_Overview, error) {
	collection := conn.Db.Collection("check_requests")
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
	var vendors []Vendor
	var credit_cards []string
	var exists = make(map[string]bool)
	var vendorExists = make(map[Vendor]bool)
	var requests []Check_Request
	for cursor.Next(context.TODO()) {
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requests = append(requests, check_req)
		if !vendorExists[check_req.Vendor] {
			vendors = append(vendors, check_req.Vendor)
			vendorExists[check_req.Vendor] = true
		}
		if !exists[check_req.Credit_Card] {
			credit_cards = append(credit_cards, check_req.Credit_Card)
			exists[check_req.Credit_Card] = true
		}
		total_amount += math.Round(check_req.Order_Total*100) / 100
		total_amount = math.Round(total_amount*100) / 100
	}
	check_overview := &Grant_Check_Overview{
		Grant:          grant_info,
		Vendors:        vendors,
		Credit_Cards:   credit_cards,
		Requests:       requests,
		Total_Requests: len(requests),
		Total_Amount:   total_amount,
	}
	return *check_overview, nil
}

type User_Check_Requests struct {
	User         User            `json:"user" bson:"user"`
	Start_Date   string          `json:"start_date" bson:"start_date"`
	End_Date     string          `json:"end_date" bson:"end_date"`
	Total_Amount float64         `json:"total_amount" bson:"total_amount"`
	Vendors      []Vendor        `json:"vendors" bson:"vendors"`
	Requests     []Check_Request `json:"requests" bson:"requests"`
}
type UserAggChecks struct {
	ID             string        `json:"id" bson:"_id"`
	Name           string        `json:"name" bson:"name"`
	Total_Amount   float64       `json:"total_amount" bson:"total_amount"`
	Total_Requests int           `json:"total_requests" bson:"total_requests"`
	Last_Request   Check_Request `json:"last_request" bson:"last_request"`
}

func (u *User) FindCheckReqs(user_id string, start_date string, end_date string) (User_Check_Requests, error) {
	collection := conn.Db.Collection("check_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
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
	total_amount := 0.0
	var vendors []Vendor
	var requests []Check_Request
	var vendorExists = make(map[Vendor]bool)
	for cursor.Next(context.TODO()) {
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requests = append(requests, check_req)
		if !vendorExists[check_req.Vendor] {
			vendors = append(vendors, check_req.Vendor)
			vendorExists[check_req.Vendor] = true
		}
		total_amount += check_req.Order_Total
	}
	return User_Check_Requests{
		User:         result,
		Total_Amount: total_amount,
		Start_Date:   start_date,
		End_Date:     end_date,
		Vendors:      vendors,
		Requests:     requests,
	}, nil
}
func (u *User) AggregateChecks(user_id string, start_date string, end_date string) (UserAggChecks, error) {
	collection := conn.Db.Collection("check_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
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
	total_amount := 0.0
	var requests []Check_Request
	for cursor.Next(context.TODO()) {
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requests = append(requests, check_req)
		total_amount += math.Round(check_req.Order_Total*100) / 100
		total_amount = math.Round(total_amount*100) / 100
	}
	var last_request Check_Request
	if len(requests) > 1 {
		last_request = requests[len(requests)-1]
	} else if len(requests) == 1 {
		last_request = requests[0]
	} else {
		last_request = Check_Request{}
	}
	return UserAggChecks{
		ID:             user_id,
		Name:           result.Name,
		Total_Amount:   total_amount,
		Total_Requests: len(requests),
		Last_Request:   last_request,
	}, nil
}
