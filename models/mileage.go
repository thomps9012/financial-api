package models

import (
	"context"
	database "financial-api/db"
	"financial-api/methods"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
type MileageDetailResponse struct {
	ID                string       `json:"id" bson:"_id"`
	Grant             Grant        `json:"grant" bson:"grant"`
	RequestCreator    UserNameInfo `json:"request_creator" bson:"request_creator"`
	Date              time.Time    `json:"date" bson:"date"`
	Category          Category     `json:"category" bson:"category"`
	Starting_Location string       `json:"starting_location" bson:"starting_location"`
	Destination       string       `json:"destination" bson:"destination"`
	Trip_Purpose      string       `json:"trip_purpose" bson:"trip_purpose"`
	Start_Odometer    int          `json:"start_odometer" bson:"start_odometer"`
	End_Odometer      int          `json:"end_odometer" bson:"end_odometer"`
	Tolls             float64      `json:"tolls" bson:"tolls"`
	Parking           float64      `json:"parking" bson:"parking"`
	Trip_Mileage      int          `json:"trip_mileage" bson:"trip_mileage"`
	Reimbursement     float64      `json:"reimbursement" bson:"reimbursement"`
	Created_At        time.Time    `json:"created_at" bson:"created_at"`
	Action_History    []Action     `json:"action_history" bson:"action_history"`
	Current_User      UserNameInfo `json:"current_user" bson:"current_user"`
	Current_Status    string       `json:"current_status" bson:"current_status"`
	Is_Active         bool         `json:"is_active" bson:"is_active"`
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
type UserMileage struct {
	User               UserNameInfo       `json:"user" bson:"user"`
	TotalReimbursement float64            `json:"total_reimbursement" bson:"total_reimbursement"`
	StartDate          time.Time          `json:"start_date" bson:"end_date"`
	EndDate            time.Time          `json:"end_date" bson:"end_date"`
	Requests           []Mileage_Overview `json:"request" bson:"requests"`
}

type FindMileageInput struct {
	MileageID string `json:"mileage_id" bson:"mileage_id" validate:"required"`
}

var current_mileage_rate = .655

func endMonthDay(monthString string) string {
	switch monthString {
	case "01":
	case "03":
	case "05":
	case "07":
	case "08":
	case "10":
	case "12":
		return "-31"
	default:
		return "-30"
	}
	return "-30"
}
func GetUserMileage(user_id string, start_month string, end_month string) (*UserMileage, error) {
	collection, err := database.Use("mileage_requests")
	res := new(UserMileage)
	if err != nil {
		return res, err
	}
	user_info := new(UserNameInfo)
	var start_date time.Time
	var end_date time.Time
	layout := "2006-01-02"
	if start_month != "" && len(start_month) > 0 {
		start_date, _ = time.Parse(layout, start_month+"-01")
	}

	if end_month != "" && len(end_month) > 0 {
		end_day := endMonthDay(strings.Split(end_month, "-")[1])
		end_date, _ = time.Parse(layout, end_month+end_day)
	}
	requests := make([]Mileage_Overview, 0)
	user_name, err := FindUserName(user_id)
	if err != nil {
		return res, err
	}
	user_info.ID = user_id
	user_info.Name = user_name
	var filter bson.D
	var zero_time time.Time
	if start_date == zero_time && end_date == zero_time {
		filter = bson.D{{Key: "user_id", Value: user_id}}
	}
	if start_date == zero_time && end_date != zero_time {
		filter = bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: bson.D{{Key: "$lte", Value: end_date}}}}
	}
	if end_date == zero_time && start_date != zero_time {
		filter = bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: bson.D{{Key: "$gte", Value: start_date}}}}
	}
	if end_date != zero_time && start_date != zero_time {
		filter = bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: bson.D{{Key: "$gte", Value: start_date}, {Key: "$lte", Value: end_date}}}}
	}
	projection := bson.D{{Key: "_id", Value: 1}, {Key: "user_id", Value: 1}, {Key: "date", Value: 1}, {Key: "reimbursement", Value: 1}, {Key: "current_user", Value: 1}, {Key: "current_status", Value: 1}, {Key: "is_active", Value: 1}}
	opts := options.Find().SetProjection(projection).SetSort(bson.D{{Key: "date", Value: 1}})
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return res, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return res, err
	}
	total := 0.00
	for _, request := range requests {
		total = total + request.Reimbursement
	}
	res.TotalReimbursement = total
	res.Requests = requests
	if start_date == zero_time && len(requests) > 0 {
		res.StartDate = requests[0].Date
	} else {
		res.StartDate = start_date
	}
	if end_date == zero_time && len(requests) > 0 {
		res.EndDate = requests[len(requests)-1].Date
	} else {
		res.EndDate = end_date
	}
	res.User = *user_info
	return res, nil
}
func GetUserMileageDetail(user_id string) ([]Mileage_Request, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return []Mileage_Request{}, err
	}
	requests := make([]Mileage_Request, 0)
	filter := bson.D{{Key: "user_id", Value: user_id}}
	projection := bson.D{{Key: "action_history", Value: 0}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Mileage_Request{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Mileage_Request{}, err
	}
	return requests, nil
}
func (mi *MileageInput) DuplicateRequest(user_id string) (bool, error) {
	mileage_coll, err := database.Use("mileage_requests")
	if err != nil {
		return false, err
	}
	date_filter := bson.D{{Key: "date", Value: mi.Date}, {Key: "user_id", Value: user_id}}
	odometer_filter := bson.D{{Key: "start_odometer", Value: mi.Start_Odometer}, {Key: "end_odometer", Value: mi.End_Odometer}, {Key: "user_id", Value: user_id}}
	grant_category_filter := bson.D{{Key: "start_odometer", Value: mi.Start_Odometer}, {Key: "grant_id", Value: mi.Grant_ID}, {Key: "category", Value: mi.Category}, {Key: "user_id", Value: user_id}}
	same_date, err := mileage_coll.CountDocuments(context.TODO(), date_filter)
	if err != nil {
		return false, err
	}
	same_category, err := mileage_coll.CountDocuments(context.TODO(), grant_category_filter)
	if err != nil {
		return false, err
	}
	same_odometer, err := mileage_coll.CountDocuments(context.TODO(), odometer_filter)
	if err != nil {
		return false, err
	}
	return same_date+same_category+same_odometer > 0, nil
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
	first_action := FirstActions(user_id)
	new_request.Action_History = first_action
	current_user := methods.NewRequestUser("mileage", "nil", user_id)
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
	err = CreateIncompleteAction("mileage", new_request.ID, first_action[0], current_user.ID)
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
func (em *EditMileageInput) EditMileage() (Mileage_Overview, *CustomError) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "_id", Value: em.ID}}
	opts := options.FindOne().SetProjection(bson.D{{Key: "action_history", Value: 1}, {Key: "current_status", Value: 1}, {Key: "current_user", Value: 1}, {Key: "last_user_before_reject", Value: 1}})
	request := new(Mileage_Request)
	err = collection.FindOne(context.TODO(), filter, opts).Decode(&request)
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	if request.Current_Status != "PENDING" && request.Current_Status != "REJECTED" && request.Current_Status != "REJECTED_EDIT_PENDING_REVIEW" {
		return Mileage_Overview{}, &CustomError{
			Status:  423,
			Message: "Request is being processed by organization",
		}
	}
	err = ClearRequestAssociatedActions(em.ID)
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	switch request.Current_Status {
	case "REJECTED":
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       em.User_ID,
			Status:     "REJECTED_EDIT",
			Created_At: time.Now(),
		}
		res, err := em.SaveEdits(edit_action, "REJECTED_EDIT_PENDING_REVIEW", request.Last_User_Before_Reject)
		if err != nil {
			return Mileage_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		current_user, err := FindUserName(res.Current_User)
		if err != nil {
			return Mileage_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
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
	case "REJECTED_EDIT_PENDING_REVIEW":
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       em.User_ID,
			Status:     "REJECTED_EDIT",
			Created_At: time.Now(),
		}
		res, err := em.SaveEdits(edit_action, "REJECTED_EDIT_PENDING_REVIEW", request.Current_User)
		if err != nil {
			return Mileage_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		current_user, err := FindUserName(res.Current_User)
		if err != nil {
			return Mileage_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
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
	case "PENDING":
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       em.User_ID,
			Status:     "PENDING_EDIT",
			Created_At: time.Now(),
		}
		res, err := em.SaveEdits(edit_action, "PENDING", request.Current_User)
		if err != nil {
			return Mileage_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		current_user, err := FindUserName(res.Current_User)
		if err != nil {
			return Mileage_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
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

	return Mileage_Overview{}, &CustomError{
		Status:  423,
		Message: "Request is being processed by organization",
	}
}
func (em *EditMileageInput) SaveEdits(action Action, new_status string, new_user string) (Mileage_Request, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Request{}, err
	}
	new_mileage := em.End_Odometer - *em.Start_Odometer
	new_reimbursement := *em.Tolls + *em.Parking + float64(new_mileage)*current_mileage_rate
	err = CreateIncompleteAction("mileage", em.ID, action, new_user)
	if err != nil {
		return Mileage_Request{}, err
	}
	var update bson.D
	if new_status == "REJECTED_EDIT_PENDING_REVIEW" {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "reimbursement", Value: new_reimbursement}, {Key: "grant_id", Value: em.Grant_ID}, {Key: "date", Value: em.Date}, {Key: "category", Value: em.Category}, {Key: "starting_location", Value: em.Starting_Location}, {Key: "destination", Value: em.Destination}, {Key: "trip_purpose", Value: em.Trip_Purpose}, {Key: "start_odometer", Value: em.Start_Odometer}, {Key: "end_odometer", Value: em.End_Odometer}, {Key: "tolls", Value: em.Tolls}, {Key: "parking", Value: em.Parking}, {Key: "current_status", Value: new_status}, {Key: "current_user", Value: new_user}, {Key: "last_user_before_reject", Value: "null"}}}, {Key: "$push", Value: bson.D{{Key: "action_history", Value: action}}}}
	} else {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "reimbursement", Value: new_reimbursement}, {Key: "grant_id", Value: em.Grant_ID}, {Key: "date", Value: em.Date}, {Key: "category", Value: em.Category}, {Key: "starting_location", Value: em.Starting_Location}, {Key: "destination", Value: em.Destination}, {Key: "trip_purpose", Value: em.Trip_Purpose}, {Key: "start_odometer", Value: em.Start_Odometer}, {Key: "end_odometer", Value: em.End_Odometer}, {Key: "tolls", Value: em.Tolls}, {Key: "parking", Value: em.Parking}, {Key: "current_status", Value: new_status}, {Key: "current_user", Value: new_user}}}, {Key: "$push", Value: bson.D{{Key: "action_history", Value: action}}}}
	}
	filter := bson.D{{Key: "_id", Value: em.ID}}
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
func DeleteMileage(mileage_id string, user_id string, user_admin bool) (Mileage_Overview, *CustomError) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "_id", Value: mileage_id}}
	request_info := new(Mileage_Overview)
	if !user_admin {
		err = collection.FindOne(context.TODO(), filter).Decode(&request_info)
		if err != nil {
			return Mileage_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		if request_info.User_ID != user_id {
			return Mileage_Overview{}, &CustomError{
				Status:  403,
				Message: "Unauthorized",
			}
		}
	}
	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&request_info)
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
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
func MileageDetail(mileage_id string, user_id string, user_admin bool) (MileageDetailResponse, *CustomError) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return MileageDetailResponse{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: mileage_id}}}}
	grant_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "grants"}, {Key: "localField", Value: "grant_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "grant"},
		{Key: "pipeline", Value: bson.A{
			bson.D{{Key: "$project", Value: bson.D{{Key: "name", Value: 1}}}}}}}}}
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "request_creator"}, {Key: "pipeline", Value: bson.A{
		bson.D{{Key: "$project", Value: bson.D{{Key: "name", Value: 1}}}}}}}}}
	current_user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "current_user"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "current_user"}, {Key: "pipeline", Value: bson.A{
		bson.D{{Key: "$project", Value: bson.D{{Key: "name", Value: 1}, {Key: "is_active", Value: 1}}}}}}}}}
	unwind_request_creator := bson.D{{Key: "$unwind", Value: "$request_creator"}}
	unwind_current_user := bson.D{{Key: "$unwind", Value: "$current_user"}}
	unwind_grant := bson.D{{Key: "$unwind", Value: "$grant"}}
	pipeline := mongo.Pipeline{filter, grant_stage, user_stage, current_user_stage, unwind_request_creator, unwind_current_user, unwind_grant}
	request_detail := make([]MileageDetailResponse, 0)
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return MileageDetailResponse{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	err = cursor.All(context.TODO(), &request_detail)
	if err != nil {
		return MileageDetailResponse{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	if len(request_detail) == 0 {
		return MileageDetailResponse{}, &CustomError{
			Status:  404,
			Message: "No mileage request with that id found",
		}
	}
	if !user_admin {
		if request_detail[0].RequestCreator.ID != user_id {
			return MileageDetailResponse{}, &CustomError{
				Status:  403,
				Message: "You are not authorized to view this request",
			}
		}
	}
	request_detail[0].Reimbursement = ToFixed(request_detail[0].Reimbursement, 2)
	return request_detail[0], nil
}
func (m *Mileage_Request) Approve(user_id string) (Mileage_Overview, *CustomError) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "_id", Value: m.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	if m.Current_User != user_id {
		return Mileage_Overview{}, &CustomError{
			Status:  403,
			Message: "Unauthorized",
		}
	}
	new_action := ApproveRequest("mileage", m.Current_User, m.Category, m.Current_Status)
	m.Action_History = append(m.Action_History, new_action.Action)
	var update bson.D
	if new_action.Action.Status == "ORGANIZATION_APPROVED" {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "current_user", Value: new_action.NewUser.ID}, {Key: "is_active", Value: false}, {Key: "action_history", Value: m.Action_History}, {Key: "current_status", Value: new_action.Action.Status}}}}
	} else {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "current_user", Value: new_action.NewUser.ID}, {Key: "action_history", Value: m.Action_History}, {Key: "current_status", Value: new_action.Action.Status}}}}
		err = CreateIncompleteAction("mileage", m.ID, new_action.Action, new_action.NewUser.ID)
		if err != nil {
			return Mileage_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
	}
	response := new(Mileage_Overview)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&response)
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}

	response.Current_User = new_action.NewUser.Name
	return *response, nil
}
func (m *Mileage_Request) Reject(user_id string) (Mileage_Overview, *CustomError) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "_id", Value: m.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	if m.Current_User != user_id {
		return Mileage_Overview{}, &CustomError{
			Status:  403,
			Message: "Unauthorized",
		}
	}
	reject_info := RejectRequest(m.User_ID, m.Current_User)
	m.Action_History = append(m.Action_History, reject_info.Action)
	err = CreateIncompleteAction("mileage", m.ID, reject_info.Action, m.User_ID)
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "current_user", Value: reject_info.NewUser.ID}, {Key: "action_history", Value: m.Action_History}, {Key: "current_status", Value: "REJECTED"}, {Key: "last_user_before_reject", Value: reject_info.LastUserBeforeReject.ID}}}}
	response := new(Mileage_Overview)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&response)
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	current_user_name, err := FindUserName(reject_info.NewUser.ID)
	if err != nil {
		return Mileage_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
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
	start_date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	end_date := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	filter := bson.D{{Key: "date", Value: bson.D{{Key: "$lte", Value: end_date}, {Key: "$gte", Value: start_date}}}}
	projection := bson.D{{Key: "_id", Value: 1}, {Key: "user_id", Value: 1}, {Key: "date", Value: 1}, {Key: "reimbursement", Value: 1}, {Key: "current_user", Value: 1}, {Key: "current_status", Value: 1}, {Key: "is_active", Value: 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	err = cursor.All(context.TODO(), &response)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	return response, nil
}
