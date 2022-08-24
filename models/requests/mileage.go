package requests

import (
	"context"
	"errors"
	conn "financial-api/db"
	user "financial-api/models/user"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Mileage_Request struct {
	ID                string    `json:"id" bson:"_id"`
	Grant_ID          string    `json:"grant_id" bson:"grant_id"`
	User_ID           string    `json:"user_id" bson:"user_id"`
	Date              time.Time `json:"date" bson:"date"`
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
	Current_Status    string    `json:"current_status" bson:"current_status"`
	Is_Active         bool      `json:"is_active" bson:"is_active"`
}

type Mileage_Overview struct {
	ID             string    `json:"id" bson:"_id"`
	Grant_ID       string    `json:"grant_id" bson:"grant_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	User           user.User `json:"user" bson:"user"`
	Date           time.Time `json:"date" bson:"date"`
	Trip_Mileage   int       `json:"trip_mileage" bson:"trip_mileage"`
	Reimbursement  float64   `json:"reimbursement" bson:"reimbursement"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	Current_Status string    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type Monthly_Mileage_Overview struct {
	User_ID       string     `json:"user_id" bson:"user_id"`
	Grant_IDS     []string   `json:"grant_id" bson:"grant_id"`
	Name          string     `json:"name" bson:"name"`
	Month         time.Month `json:"month" bson:"month"`
	Year          int        `json:"year" bson:"year"`
	Mileage       int        `json:"mileage" bson:"mileage"`
	Tolls         float64    `json:"tolls" bson:"tolls"`
	Parking       float64    `json:"parking" bson:"parking"`
	Current_User  string     `json:"current_user" bson:"current_user"`
	Reimbursement float64    `json:"reimbursement" bson:"reimbursement"`
	Request_IDS   []string   `json:"request_ids" bson:"request_ids"`
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

func (m *Mileage_Request) Create(requestor user.User) (Mileage_Request, error) {
	fmt.Printf("%s\n", m.Date)
	collection := conn.Db.Collection("mileage_requests")
	var currentMileageRate = 62.5
	m.ID = uuid.NewString()
	m.Created_At = time.Now()
	m.Is_Active = true
	m.User_ID = requestor.ID
	m.Current_Status = "PENDING"
	m.Current_User = requestor.Manager_ID
	m.Trip_Mileage = m.End_Odometer - m.Start_Odometer
	first_action := &Action{
		ID:           uuid.NewString(),
		User:         requestor,
		Request_Type: "mileage_requests",
		Request_ID:   m.ID,
		Status:       "PENDING",
		Created_At:   time.Now(),
	}
	m.Action_History = append(m.Action_History, *first_action)
	m.Reimbursement = float64(m.Trip_Mileage)*currentMileageRate/100 + m.Tolls + m.Parking
	_, insert_err := collection.InsertOne(context.TODO(), *m)
	if insert_err != nil {
		panic(insert_err)
	}
	var manager user.User
	update_user, update_err := manager.AddNotification(user.Action(*first_action), requestor.Manager_ID)
	if update_err != nil {
		panic(update_err)
	}
	if !update_user {
		return *m, errors.New("failed to update manager")
	}
	return *m, nil

}

func (m *Mileage_Request) Update(request Mileage_Request, requestor user.User) (Mileage_Request, error) {
	if request.Current_Status == "REJECTED" {
		update_action := &Action{
			ID:           uuid.NewString(),
			User:         requestor,
			Request_Type: "mileage_requests",
			Request_ID:   request.ID,
			Status:       "PENDING",
			Created_At:   time.Now(),
		}
		request.Current_User = requestor.Manager_ID
		request.Action_History = append(request.Action_History, *update_action)
		var manager user.User
		update_user, update_err := manager.AddNotification(user.Action(*update_action), requestor.Manager_ID)
		if update_err != nil {
			panic(update_err)
		}
		if !update_user {
			return *m, errors.New("failed to update manager")
		}
	}
	var mileage_req Mileage_Request
	collection := conn.Db.Collection("mileage_requests")
	filter := bson.D{{Key: "_id", Value: request.ID}}
	err := collection.FindOneAndReplace(context.TODO(), filter, request).Decode(&mileage_req)
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
		var user user.User
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
