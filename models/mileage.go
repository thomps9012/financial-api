package models

import (
	"context"
	"errors"
	database "financial-api/db"
	"financial-api/methods"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mileage_Request struct {
	ID                      string    `json:"id" bson:"_id"`
	Grant_ID                string    `json:"grant_id" bson:"grant_id"`
	User_ID                 string    `json:"user_id" bson:"user_id"`
	Date                    time.Time `json:"date" bson:"date"`
	Category                Category  `json:"category" bson:"category"`
	Starting_Location       string    `json:"starting_location" bson:"starting_location"`
	Destination             string    `json:"destination" bson:"destination"`
	Trip_Purpose            string    `json:"trip_purpose" bson:"trip_purpose"`
	Start_Odometer          int       `json:"start_odometer" bson:"start_odometer"`
	End_Odometer            int       `json:"end_odometer" bson:"end_odometer"`
	Tolls                   float64   `json:"tolls" bson:"tolls"`
	Parking                 float64   `json:"parking" bson:"parking"`
	Trip_Mileage            int       `json:"trip_mileage" bson:"trip_mileage"`
	Reimbursement           float64   `json:"reimbursement" bson:"reimbursement"`
	Created_At              time.Time `json:"created_at" bson:"created_at"`
	Action_History          []Action  `json:"action_history" bson:"action_history"`
	Current_User            string    `json:"current_user" bson:"current_user"`
	Current_Status          string    `json:"current_status" bson:"current_status"`
	Last_User_Before_Reject string    `json:"last_user_before_reject" bson:"last_user_before_reject"`
	Is_Active               bool      `json:"is_active" bson:"is_active"`
}
type MileageInput struct {
	Grant_ID          string    `json:"grant_id" bson:"grant_id" validate:"required"`
	Date              time.Time `json:"date" bson:"date" validate:"required"`
	Category          Category  `json:"category" bson:"category" validate:"required"`
	Starting_Location string    `json:"starting_location" bson:"starting_location" validate:"required"`
	Destination       string    `json:"destination" bson:"destination" validate:"required"`
	Trip_Purpose      string    `json:"trip_purpose" bson:"trip_purpose" validate:"required"`
	Start_Odometer    *int      `json:"start_odometer" bson:"start_odometer" validate:"required"`
	End_Odometer      int       `json:"end_odometer" bson:"end_odometer" validate:"required"`
	Tolls             *float64  `json:"tolls" bson:"tolls" validate:"required"`
	Parking           *float64  `json:"parking" bson:"parking" validate:"required"`
}
type EditMileageInput struct {
	ID                string    `json:"id" bson:"_id" validate:"required"`
	User_ID           string    `json:"user_id" bson:"user_id" validate:"required"`
	Grant_ID          string    `json:"grant_id" bson:"grant_id" validate:"required"`
	Date              time.Time `json:"date" bson:"date" validate:"required"`
	Category          Category  `json:"category" bson:"category" validate:"required"`
	Starting_Location string    `json:"starting_location" bson:"starting_location" validate:"required"`
	Destination       string    `json:"destination" bson:"destination" validate:"required"`
	Trip_Purpose      string    `json:"trip_purpose" bson:"trip_purpose" validate:"required"`
	Start_Odometer    *int      `json:"start_odometer" bson:"start_odometer" validate:"required"`
	End_Odometer      int       `json:"end_odometer" bson:"end_odometer" validate:"required"`
	Tolls             *float64  `json:"tolls" bson:"tolls" validate:"required"`
	Parking           *float64  `json:"parking" bson:"parking" validate:"required"`
}
type Mileage_Overview struct {
	ID             string    `json:"id" bson:"_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	Date           time.Time `json:"date" bson:"date"`
	Reimbursement  float64   `json:"reimbursement" bson:"reimbursement"`
	Current_Status string    `json:"current_status" bson:"current_status"`
	Current_User   string    `json:"current_user" bson:"current_user"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}
type FindMileageInput struct {
	MileageID string `json:"mileage_id" bson:"mileage_id" validate:"required"`
}

var current_mileage_rate = .655

func GetUserMileage(user_id string) ([]Mileage_Overview, error) {
	collection, err := database.Use("mileage_requests")
	requests := make([]Mileage_Overview, 0)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	filter := bson.D{{"user_id", user_id}}
	projection := bson.D{{"_id", 1}, {"user_id", 1}, {"date", 1}, {"reimbursement", 1}, {"current_user", 1}, {"current_status", 1}, {"is_active", 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	for _, request := range requests {
		current_user_id := request.Current_User
		user_name, err := FindUserName(current_user_id)
		if err != nil {
			request.Current_User = "N/A"
		} else {
			request.Current_User = user_name
		}
	}
	return requests, nil
}
func GetUserMileageDetail(user_id string) ([]Mileage_Request, error) {
	collection, err := database.Use("mileage_requests")
	requests := make([]Mileage_Request, 0)
	filter := bson.D{{"user_id", user_id}}
	projection := bson.D{{"action_history", 0}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Mileage_Request{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Mileage_Request{}, err
	}
	for _, request := range requests {
		current_user_id := request.Current_User
		user_name, err := FindUserName(current_user_id)
		if err != nil {
			request.Current_User = "N/A"
		} else {
			request.Current_User = user_name
		}
	}
	return requests, nil
}
func (mi *MileageInput) CreateMileage(user_id string) (Mileage_Overview, error) {
	new_request := new(Mileage_Request)
	new_request.ID = uuid.NewString()
	new_request.Grant_ID = mi.Grant_ID
	new_request.User_ID = user_id
	new_request.Date = mi.Date
	new_request.Category = mi.Category
	new_request.Starting_Location = mi.Starting_Location
	new_request.Destination = mi.Destination
	new_request.Trip_Purpose = mi.Trip_Purpose
	new_request.Start_Odometer = *mi.Start_Odometer
	new_request.End_Odometer = mi.End_Odometer
	new_request.Trip_Mileage = mi.End_Odometer - *mi.Start_Odometer
	new_request.Parking = *mi.Parking
	new_request.Tolls = *mi.Tolls
	new_request.Created_At = time.Now()
	new_request.Last_User_Before_Reject = bson.TypeNull.String()
	new_request.Is_Active = true
	first_action, _ := FirstActions("mileage", new_request.ID, user_id)
	new_request.Action_History = first_action
	current_user := methods.NewRequestUser("mileage", "nil")
	new_request.Current_User = current_user.ID
	new_request.Current_Status = "PENDING"
	trip_sum := *mi.Tolls + *mi.Parking + float64(new_request.Trip_Mileage)*current_mileage_rate
	new_request.Reimbursement = trip_sum
	mileage_coll, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, err
	}
	_, err = mileage_coll.InsertOne(context.TODO(), new_request)
	if err != nil {
		return Mileage_Overview{}, err
	}
	return Mileage_Overview{
		ID:             new_request.ID,
		User_ID:        user_id,
		Date:           new_request.Date,
		Reimbursement:  new_request.Reimbursement,
		Current_Status: new_request.Current_Status,
		Current_User:   current_user.Name,
		Is_Active:      true,
	}, nil
}
func (em *EditMileageInput) EditMileage() (Mileage_Overview, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, err
	}
	filter := bson.D{{"_id", em.ID}}
	opts := options.FindOne().SetProjection(bson.D{{"action_history", 1}, {"current_status", 1}, {"current_user", 1}, {"last_user_before_reject", 1}})
	request := new(Mileage_Request)
	err = collection.FindOne(context.TODO(), filter, opts).Decode(&request)
	if err != nil {
		return Mileage_Overview{}, err
	}
	if request.Current_Status == "REJECTED" || request.Current_Status == "REJECTED_EDIT_PENDING_REVIEW" {
		edit_action := Action{
			ID:           uuid.NewString(),
			Request_ID:   em.ID,
			Request_Type: MILEAGE,
			User:         em.User_ID,
			Status:       "REJECTED_EDIT",
			Created_At:   time.Now(),
		}
		res, err := em.SaveEdits(edit_action, "REJECTED_EDIT_PENDING_REVIEW", request.Last_User_Before_Reject)
		if err != nil {
			return Mileage_Overview{}, err
		}
		current_user, err := FindUserName(res.Current_User)
		if err != nil {
			return Mileage_Overview{}, err
		}
		return Mileage_Overview{
			ID:             res.ID,
			User_ID:        res.User_ID,
			Date:           res.Date,
			Reimbursement:  res.Reimbursement,
			Current_Status: res.Current_Status,
			Current_User:   current_user,
			Is_Active:      res.Is_Active,
		}, nil
	}
	if request.Current_Status == "PENDING" {
		edit_action := Action{
			ID:           uuid.NewString(),
			Request_ID:   em.ID,
			Request_Type: MILEAGE,
			User:         em.User_ID,
			Status:       "PENDING_EDIT",
			Created_At:   time.Now(),
		}
		res, err := em.SaveEdits(edit_action, "PENDING", request.Current_User)
		if err != nil {
			return Mileage_Overview{}, err
		}
		current_user, err := FindUserName(res.Current_User)
		if err != nil {
			return Mileage_Overview{}, err
		}
		return Mileage_Overview{
			ID:             res.ID,
			User_ID:        res.User_ID,
			Date:           res.Date,
			Reimbursement:  res.Reimbursement,
			Current_Status: res.Current_Status,
			Current_User:   current_user,
			Is_Active:      res.Is_Active,
		}, nil
	}
	return Mileage_Overview{}, errors.New("this request is currently being processed by the organization and not editable")
}
func (em *EditMileageInput) SaveEdits(action Action, new_status string, new_user string) (Mileage_Request, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Request{}, err
	}
	new_mileage := em.End_Odometer - *em.Start_Odometer
	new_reimbursement := *em.Tolls + *em.Parking + float64(new_mileage)*current_mileage_rate
	var update bson.D
	if new_status == "REJECTED_EDIT_PENDING_REVIEW" {
		update = bson.D{{"$set", bson.D{{"reimbursement", new_reimbursement}, {"grant_id", em.Grant_ID}, {"date", em.Date}, {"category", em.Category}, {"starting_location", em.Starting_Location}, {"destination", em.Destination}, {"trip_purpose", em.Trip_Purpose}, {"start_odometer", em.Start_Odometer}, {"end_odometer", em.End_Odometer}, {"tolls", em.Tolls}, {"parking", em.Parking}, {"current_status", new_status}, {"current_user", new_user}, {"last_user_before_reject", "null"}}}, {"$push", bson.D{{"action_history", action}}}}
	} else {
		update = bson.D{{"$set", bson.D{{"reimbursement", new_reimbursement}, {"grant_id", em.Grant_ID}, {"date", em.Date}, {"category", em.Category}, {"starting_location", em.Starting_Location}, {"destination", em.Destination}, {"trip_purpose", em.Trip_Purpose}, {"start_odometer", em.Start_Odometer}, {"end_odometer", em.End_Odometer}, {"tolls", em.Tolls}, {"parking", em.Parking}, {"current_status", new_status}, {"current_user", new_user}}}, {"$push", bson.D{{"action_history", action}}}}
	}
	filter := bson.D{{"_id", em.ID}}
	mileage_req := new(Mileage_Request)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&mileage_req)
	if err != nil {
		return Mileage_Request{}, err
	}
	return *mileage_req, nil
}
func DeleteMileage(mileage_id string) (Mileage_Overview, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, err
	}
	filter := bson.D{{"_id", mileage_id}}
	request_info := new(Mileage_Overview)
	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&request_info)
	if err != nil {
		return Mileage_Overview{}, err
	}
	current_user_id := request_info.Current_User
	user_name, err := FindUserName(current_user_id)
	if err != nil {
		request_info.Current_User = "N/A"
	} else {
		request_info.Current_User = user_name
	}
	return *request_info, nil
}
func MileageDetail(mileage_id string) (Mileage_Request, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Request{}, err
	}
	filter := bson.D{{"_id", mileage_id}}
	request_detail := new(Mileage_Request)
	err = collection.FindOne(context.TODO(), filter).Decode(&request_detail)
	if err != nil {
		return Mileage_Request{}, err
	}
	return *request_detail, nil
}
func (m *Mileage_Request) Approve(user_id string) (Mileage_Overview, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, err
	}
	filter := bson.D{{"_id", m.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return Mileage_Overview{}, err
	}
	if m.Current_User != user_id {
		return Mileage_Overview{}, errors.New("you are attempting to approve a request for which you are unauthorized")
	}
	new_action, err := ApproveRequest("mileage", m.ID, m.Current_User, m.Category, m.Current_Status)
	if err != nil {
		return Mileage_Overview{}, err
	}
	m.Action_History = append(m.Action_History, new_action.Action)
	update := bson.D{{"$set", bson.D{{"current_user", new_action.NewUser.ID}, {"action_history", m.Action_History}, {"current_status", new_action.Action.Status}}}}
	response := new(Mileage_Overview)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&response)
	if err != nil {
		return Mileage_Overview{}, err
	}
	response.Current_User = new_action.NewUser.Name
	return *response, nil
}
func (m *Mileage_Request) Reject(user_id string) (Mileage_Overview, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, err
	}
	filter := bson.D{{"_id", m.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return Mileage_Overview{}, err
	}
	if m.Current_User != user_id {
		return Mileage_Overview{}, errors.New("you are attempting to reject a request for which you are unauthorized")
	}
	reject_info := RejectRequest("mileage", m.ID, m.User_ID, m.Current_User)
	m.Action_History = append(m.Action_History, reject_info.Action)
	update := bson.D{{"$set", bson.D{{"current_user", reject_info.NewUser.ID}, {"action_history", m.Action_History}, {"current_status", "REJECTED"}, {"last_user_before_reject", reject_info.LastUserBeforeReject.ID}}}}
	response := new(Mileage_Overview)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&response)
	if err != nil {
		return Mileage_Overview{}, err
	}
	current_user_name, err := FindUserName(reject_info.NewUser.ID)
	if err != nil {
		return Mileage_Overview{}, err
	}
	response.Current_User = current_user_name
	return *response, nil
}
func MonthlyMileage(month int, year int) ([]Mileage_Overview, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return []Mileage_Overview{}, err
	}
	response := make([]Mileage_Overview, 0)
	start_date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end_date := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.UTC)
	filter := bson.D{{"date", bson.D{{"$lte", end_date}, {"$gte", start_date}}}}
	projection := bson.D{{"_id", 1}, {"user_id", 1}, {"date", 1}, {"reimbursement", 1}, {"current_user", 1}, {"current_status", 1}, {"is_active", 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	err = cursor.All(context.TODO(), &response)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	for _, request := range response {
		current_user_id := request.Current_User
		user_name, err := FindUserName(current_user_id)
		if err != nil {
			request.Current_User = "N/A"
		} else {
			request.Current_User = user_name
		}
	}
	return response, nil
}
