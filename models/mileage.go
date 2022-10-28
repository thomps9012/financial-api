package models

import (
	"context"
	"errors"
	conn "financial-api/db"
	"financial-api/middleware"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mileage_Request struct {
	ID                string    `json:"id" bson:"_id"`
	Grant_ID          string    `json:"grant_id" bson:"grant_id"`
	User_ID           string    `json:"user_id" bson:"user_id"`
	Date              time.Time `json:"date" bson:"date"`
	Category          Category  `json:"category" bson:"category"`
	Starting_Location string    `json:"starting_location" bson:"starting_location"`
	Destination       string    `json:"destination" bson:"destination"`
	Trip_Purpose      string    `json:"trip_purpose" bson:"trip_purpose"`
	Start_Odometer    int       `json:"start_odometer" bson:"start_odometer"`
	End_Odometer      int       `json:"end_odometer" bson:"end_odometer"`
	Tolls             float64   `json:"tolls" bson:"tolls"`
	Parking           float64   `json:"parking" bson:"parking"`
	Trip_Mileage      int       `json:"trip_mileage" bson:"trip_mileage"`
	Reimbursement     float64   `json:"reimbursement" bson:"reimbursement"`
	Created_At        time.Time `json:"created_at" bson:"created_at"`
	Action_History    []Action  `json:"action_history" bson:"action_history"`
	Current_User      string    `json:"current_user" bson:"current_user"`
	Current_Status    Status    `json:"current_status" bson:"current_status"`
	Is_Active         bool      `json:"is_active" bson:"is_active"`
}

type Mileage_Overview struct {
	ID             string    `json:"id" bson:"_id"`
	Grant_ID       string    `json:"grant_id" bson:"grant_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	User           User      `json:"user" bson:"user"`
	Date           time.Time `json:"date" bson:"date"`
	Trip_Mileage   int       `json:"trip_mileage" bson:"trip_mileage"`
	Reimbursement  float64   `json:"reimbursement" bson:"reimbursement"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	Current_Status Status    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}
type Grant_Mileage_Overview struct {
	Grant         Grant             `json:"grant" bson:"grant"`
	Mileage       int               `json:"mileage" bson:"mileage"`
	Tolls         float64           `json:"tolls" bson:"tolls"`
	Parking       float64           `json:"parking" bson:"parking"`
	Reimbursement float64           `json:"reimbursement" bson:"reimbursement"`
	Requests      []Mileage_Request `json:"requests" bson:"requests"`
}
type Monthly_Mileage_Overview struct {
	User_ID       string            `json:"user_id" bson:"user_id"`
	Grant_IDS     []string          `json:"grant_id" bson:"grant_id"`
	Name          string            `json:"name" bson:"name"`
	Month         time.Month        `json:"month" bson:"month"`
	Year          int               `json:"year" bson:"year"`
	Mileage       int               `json:"mileage" bson:"mileage"`
	Tolls         float64           `json:"tolls" bson:"tolls"`
	Parking       float64           `json:"parking" bson:"parking"`
	Current_User  string            `json:"current_user" bson:"current_user"`
	Reimbursement float64           `json:"reimbursement" bson:"reimbursement"`
	Requests      []Mileage_Request `json:"requests" bson:"requests"`
}

type User_Monthly_Mileage struct {
	ID            string            `json:"id" bson:"_id"`
	Name          string            `json:"name" bson:"name"`
	Vehicles      []Vehicle         `json:"vehicles" bson:"vehicles"`
	Month         time.Month        `json:"month" bson:"month"`
	Year          int               `json:"year" bson:"year"`
	Mileage       int               `json:"mileage" bson:"mileage"`
	Tolls         float64           `json:"tolls" bson:"tolls"`
	Parking       float64           `json:"parking" bson:"parking"`
	Reimbursement float64           `json:"reimbursement" bson:"reimbursement"`
	Grant_IDS     []string          `json:"grant_ids" bson:"grant_ids"`
	Requests      []Mileage_Request `json:"requests" bson:"requests"`
}

type User_Mileage struct {
	Vehicles      []Vehicle         `json:"vehicles" bson:"vehicles"`
	Mileage       int               `json:"mileage" bson:"mileage"`
	Tolls         float64           `json:"tolls" bson:"tolls"`
	Parking       float64           `json:"parking" bson:"parking"`
	Reimbursement float64           `json:"reimbursement" bson:"reimbursement"`
	Requests      []Mileage_Request `json:"requests" bson:"requests"`
	Last_Request  Mileage_Request   `json:"last_request" bson:"last_request"`
}

type User_Agg_Mileage struct {
	Parking       float64           `json:"parking" bson:"parking"`
	Tolls         float64           `json:"tolls" bson:"tolls"`
	Mileage       int               `json:"mileage" bson:"mileage"`
	User          User              `json:"user" bson:"user"`
	Reimbursement float64           `json:"reimbursement" bson:"reimbursement"`
	Requests      []Mileage_Request `json:"requests" bson:"requests"`
}

func (g *Grant_Mileage_Overview) FindByGrant(grant_id string, start_date string, end_date string) (Grant_Mileage_Overview, error) {
	collection := conn.Db.Collection("mileage_requests")
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
	var requests []Mileage_Request
	total_reimbursement := 0.00
	total_tolls := 0.00
	total_parking := 0.00
	total_mileage := 0

	for cursor.Next(context.TODO()) {
		var mileage_req Mileage_Request
		decode_err := cursor.Decode(&mileage_req)
		if decode_err != nil {
			panic(decode_err)
		}
		total_reimbursement += mileage_req.Reimbursement
		total_tolls += mileage_req.Tolls
		total_parking += mileage_req.Parking
		total_mileage += mileage_req.Trip_Mileage
		requests = append(requests, mileage_req)
	}
	return Grant_Mileage_Overview{
		Grant:         grant_info,
		Mileage:       total_mileage,
		Tolls:         total_tolls,
		Parking:       total_parking,
		Reimbursement: total_reimbursement,
		Requests:      requests,
	}, nil
}

func (m *Mileage_Request) Exists(user_id string, date time.Time, start int, end int) (bool, error) {
	var milage_req Mileage_Request
	collection := conn.Db.Collection("mileage_requests")
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: date}, {Key: "start_odometer", Value: start}, {Key: "end_odometer", Value: end}}
	fmt.Printf("%s\n", filter.Map())
	err := collection.FindOne(context.TODO(), filter).Decode(&milage_req)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (m *Mileage_Request) Create(requestor User) (Mileage_Request, error) {
	collection := conn.Db.Collection("mileage_requests")
	var currentMileageRate = 62.5
	m.ID = uuid.NewString()
	m.Created_At = time.Now()
	m.Is_Active = true
	m.User_ID = requestor.ID
	m.Current_Status = PENDING
	m.Trip_Mileage = m.End_Odometer - m.Start_Odometer
	m.Reimbursement = float64(m.Trip_Mileage)*currentMileageRate/100 + m.Tolls + m.Parking
	// set current user based off category
	current_user_email := UserEmailHandler(m.Category, PENDING, false)
	var user User
	current_user_id, err := user.FindID(current_user_email)
	if err != nil {
		panic(err)
	}
	m.Current_User = current_user_id
	// set first action and apply to mileage request
	first_action := &Action{
		ID:         uuid.NewString(),
		User:       m.User_ID,
		Request_ID: m.ID,
		Status:     PENDING,
		Created_At: time.Now(),
	}
	m.Action_History = append(m.Action_History, *first_action)
	// insert the document
	_, insert_err := collection.InsertOne(context.TODO(), *m)
	if insert_err != nil {
		panic(insert_err)
	}
	// notify the current user based off the above
	update_user, update_err := middleware.SendEmail([]string{current_user_email}, "Mileage", requestor.Name, time.Now())
	if update_err != nil {
		panic(update_err)
	}
	if !update_user {
		return *m, errors.New("failed to update appropiate admin staff")
	}
	return *m, nil
}

func (m *Mileage_Request) Update(request Mileage_Request, requestor User) (Mileage_Request, error) {
	if request.Current_Status == REJECTED {
		update_action := &Action{
			ID:         uuid.NewString(),
			User:       requestor.ID,
			Request_ID: request.ID,
			Status:     REJECTED_EDIT,
			Created_At: time.Now(),
		}
		// find out status of previous action
		request.Current_Status = PENDING
		prev_user_id := request.Action_History[len(request.Action_History)-1].User
		prev_pre_rejection_action_status := request.Action_History[len(request.Action_History)-2].Status
		// make appropiate notification based off previous action
		var user User
		current_user, err := user.FindByID(prev_user_id)
		if err != nil {
			panic(err)
		}
		var current_user_email = current_user.Email
		// set the current status back to the previous action status
		request.Current_Status = prev_pre_rejection_action_status
		request.Action_History = append(request.Action_History, *update_action)
		// notify the current user based off the above
		update_user, update_err := middleware.SendEmail([]string{current_user_email}, "Mileage", requestor.Name, time.Now())
		if update_err != nil {
			panic(update_err)
		}
		if !update_user {
			panic("failed to update appropiate admin staff")
		}
	}
	// these are all other updates made
	var mileage_req Mileage_Request
	collection := conn.Db.Collection("mileage_requests")
	filter := bson.D{{Key: "_id", Value: request.ID}}
	after := options.After
	opts := options.FindOneAndReplaceOptions{
		ReturnDocument: &after,
	}
	err := collection.FindOneAndReplace(context.TODO(), filter, request, &opts).Decode(&mileage_req)
	if err != nil {
		panic(err)
	}
	return mileage_req, nil
}

func (m *Mileage_Request) Delete(request Mileage_Request, user_id string) (bool, error) {
	collection := conn.Db.Collection("mileage_requests")
	var milage_req Mileage_Request
	filter := bson.D{{Key: "request_id", Value: request.ID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&milage_req)
	if err != nil {
		panic(err)
	}
	if milage_req.User_ID != user_id {
		panic("you are not the user who created this request")
	}
	current_status := m.Current_Status
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

func (m *Mileage_Request) FindByID(mileage_id string) (Mileage_Request, error) {
	collection := conn.Db.Collection("mileage_requests")
	var milage_req Mileage_Request
	filter := bson.D{{Key: "_id", Value: mileage_id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&milage_req)
	if err != nil {
		panic(err)
	}
	return milage_req, nil
}

func (m *Mileage_Overview) FindAll() ([]Mileage_Overview, error) {
	collection := conn.Db.Collection("mileage_requests")
	var overviews []Mileage_Overview
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var mileage_req Mileage_Request
		decode_err := cursor.Decode(&mileage_req)
		if decode_err != nil {
			panic(decode_err)
		}
		var user User
		user_info, user_err := user.FindByID(mileage_req.User_ID)
		if user_err != nil {
			panic(user_err)
		}
		mileage_overview := &Mileage_Overview{
			ID:             mileage_req.ID,
			Grant_ID:       mileage_req.Grant_ID,
			User_ID:        mileage_req.User_ID,
			User:           user_info,
			Date:           mileage_req.Date,
			Trip_Mileage:   mileage_req.Trip_Mileage,
			Reimbursement:  mileage_req.Reimbursement,
			Created_At:     mileage_req.Created_At,
			Current_Status: mileage_req.Current_Status,
			Is_Active:      mileage_req.Is_Active,
		}
		overviews = append(overviews, *mileage_overview)
	}
	return overviews, nil
}

func (u *User) MonthlyMileage(user_id string, month int, year int) (User_Monthly_Mileage, error) {
	collection := conn.Db.Collection("mileage_requests")
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
	var mileage int
	tolls := 0.0
	parking := 0.0
	reimbursement := 0.0
	var requests []Mileage_Request
	var grant_ids []string
	for cursor.Next(context.TODO()) {
		var mileage_req Mileage_Request
		decode_err := cursor.Decode(&mileage_req)
		if decode_err != nil {
			panic(decode_err)
		}
		grant_ids = append(grant_ids, mileage_req.Grant_ID)
		requests = append(requests, mileage_req)
		mileage += mileage_req.Trip_Mileage
		tolls += mileage_req.Tolls
		parking += mileage_req.Parking
		reimbursement += mileage_req.Reimbursement
	}
	return User_Monthly_Mileage{
		ID:            user_id,
		Grant_IDS:     grant_ids,
		Name:          result.Name,
		Vehicles:      result.Vehicles,
		Month:         time.Month(month),
		Year:          year,
		Mileage:       mileage,
		Parking:       parking,
		Tolls:         tolls,
		Reimbursement: reimbursement,
		Requests:      requests,
	}, nil
}

func (u *User) AggregateMileage(user_id string, start_date string, end_date string) (User_Agg_Mileage, error) {
	collection := conn.Db.Collection("mileage_requests")
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
	var requests []Mileage_Request
	total_reimbursement := 0.00
	total_tolls := 0.00
	total_parking := 0.00
	total_mileage := 0

	for cursor.Next(context.TODO()) {
		var mileage_req Mileage_Request
		decode_err := cursor.Decode(&mileage_req)
		if decode_err != nil {
			panic(decode_err)
		}
		total_reimbursement += mileage_req.Reimbursement
		total_tolls += mileage_req.Tolls
		total_parking += mileage_req.Parking
		total_mileage += mileage_req.Trip_Mileage
		requests = append(requests, mileage_req)
	}
	return User_Agg_Mileage{
		User:          result,
		Mileage:       total_mileage,
		Tolls:         total_tolls,
		Parking:       total_parking,
		Reimbursement: total_reimbursement,
		Requests:      requests,
	}, nil
}

func (u *User) FindMileage(user_id string) (User_Mileage, error) {
	collection := conn.Db.Collection("mileage_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
	filter := bson.D{{Key: "user_id", Value: user_id}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var mileage int
	tolls := 0.0
	parking := 0.0
	reimbursement := 0.0
	var requests []Mileage_Request
	last_request_date := time.Date(2000, time.April,
		34, 25, 72, 01, 0, time.UTC)
	var last_request Mileage_Request
	for cursor.Next(context.TODO()) {
		var mileage_req Mileage_Request
		decode_err := cursor.Decode(&mileage_req)
		if decode_err != nil {
			panic(decode_err)
		}
		if mileage_req.Date.After(last_request_date) {
			last_request = mileage_req
		}
		requests = append(requests, mileage_req)
		mileage += mileage_req.Trip_Mileage
		tolls += mileage_req.Tolls
		parking += mileage_req.Parking
		reimbursement += mileage_req.Reimbursement
	}
	return User_Mileage{
		Vehicles:      result.Vehicles,
		Mileage:       mileage,
		Parking:       parking,
		Tolls:         tolls,
		Reimbursement: reimbursement,
		Requests:      requests,
		Last_Request:  last_request,
	}, nil
}
