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

type Mileage_Request struct {
	ID                string    `json:"id" bson:"_id"`
	User_ID           string    `json:"user_id" bson:"user_id"`
	Date              time.Time `json:"date" bson:"date"`
	Starting_Location string    `json:"starting_location" bson:"starting_location"`
	Destination       string    `json:"destination" bson:"destination"`
	Trip_Purpose      string    `json:"trip_purpose" bson:"trip_purpose"`
	Start_Odometer    int64     `json:"start_odometer" bson:"start_odometer"`
	End_Odometer      int64     `json:"end_odometer" bson:"end_odometer"`
	Tolls             float64   `json:"tolls" bson:"tolls"`
	Parking           float64   `json:"parking" bson:"parking"`
	Trip_Mileage      int64     `json:"trip_mileage" bson:"trip_mileage"`
	Reimbursement     float64   `json:"reimbursement" bson:"reimbursement"`
	Created_At        time.Time `json:"created_at" bson:"created_at"`
	Action_History    []Action  `json:"action_history" bson:"action_history"`
	Current_Status    Status    `json:"current_status" bson:"current_status"`
	Is_Active         bool      `json:"is_active" bson:"is_active"`
}

func (m *Mileage_Request) Create(user_id string) (Mileage_Request, error) {
	collection := conn.DB.Collection("mileage_requests")
	var milage_req Mileage_Request
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: m.Date}, {Key: "starting_location", Value: m.Starting_Location}, {Key: "destintation", Value: m.Destination}}
	err := collection.FindOne(context.TODO(), filter).Decode(&milage_req)
	if err == nil {
		return *m, fmt.Errorf("mileage request already created")
	}
	var currentMileageRate = 62.5
	m.ID = uuid.NewString()
	m.Created_At = time.Now()
	m.Is_Active = true
	m.User_ID = user_id
	m.Current_Status = PENDING
	m.Trip_Mileage = m.End_Odometer - m.Start_Odometer
	first_action := &Action{
		ID:         uuid.NewString(),
		User_ID:    user_id,
		Status:     PENDING,
		Created_At: time.Now(),
	}
	m.Action_History = append(m.Action_History, *first_action)
	m.Reimbursement = float64(m.Trip_Mileage)*currentMileageRate + m.Tolls + m.Parking
	var req_user user.User
	manager_id, mgr_find_err := req_user.FindMgrID(m.User_ID)
	if mgr_find_err != nil {
		panic(err)
	}
	var manager user.User
	update_user, update_err := manager.AddNotification(m.ID, manager_id)
	if update_err != nil {
		panic(err)
	}
	if !update_user {
		return *m, err
	}
	return *m, nil
}

func (m *Mileage_Request) Update(request Mileage_Request, user_id string) (bool, error) {
	collection := conn.DB.Collection("mileage_requests")
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
	result, update_err := collection.UpdateByID(context.TODO(), request.ID, request)
	if update_err != nil {
		panic(update_err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (m *Mileage_Request) Delete(request Mileage_Request, user_id string) (bool, error) {
	collection := conn.DB.Collection("mileage_requests")
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
