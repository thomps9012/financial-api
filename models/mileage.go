package models

import (
	"context"
	"errors"
	database "financial-api/db"
	"financial-api/methods"
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
	ID                string         `json:"id" bson:"_id"`
	Grant             []Grant        `json:"grant" bson:"grant"`
	RequestCreator    []UserNameInfo `json:"request_creator" bson:"request_creator"`
	Date              time.Time      `json:"date" bson:"date"`
	Category          Category       `json:"category" bson:"category"`
	Starting_Location string         `json:"starting_location" bson:"starting_location"`
	Destination       string         `json:"destination" bson:"destination"`
	Trip_Purpose      string         `json:"trip_purpose" bson:"trip_purpose"`
	Start_Odometer    int            `json:"start_odometer" bson:"start_odometer"`
	End_Odometer      int            `json:"end_odometer" bson:"end_odometer"`
	Tolls             float64        `json:"tolls" bson:"tolls"`
	Parking           float64        `json:"parking" bson:"parking"`
	Trip_Mileage      int            `json:"trip_mileage" bson:"trip_mileage"`
	Reimbursement     float64        `json:"reimbursement" bson:"reimbursement"`
	Created_At        time.Time      `json:"created_at" bson:"created_at"`
	Action_History    []Action       `json:"action_history" bson:"action_history"`
	Current_User      []UserNameInfo `json:"current_user" bson:"current_user"`
	Current_Status    string         `json:"current_status" bson:"current_status"`
	Is_Active         bool           `json:"is_active" bson:"is_active"`
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
	filter := bson.D{{Key: "user_id", Value: user_id}}
	projection := bson.D{{Key: "_id", Value: 1}, {Key: "user_id", Value: 1}, {Key: "date", Value: 1}, {Key: "reimbursement", Value: 1}, {Key: "current_user", Value: 1}, {Key: "current_status", Value: 1}, {Key: "is_active", Value: 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Mileage_Overview{}, err
	}
	// for _, request := range requests {
	// 	current_user_id := request.Current_User
	// 	user_name, err := FindUserName(current_user_id)
	// 	if err != nil {
	// 		request.Current_User = "N/A"
	// 	} else {
	// 		request.Current_User = user_name
	// 	}
	// }

	return requests, nil
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
	// for _, request := range requests {
	// 	current_user_id := request.Current_User
	// 	user_name, err := FindUserName(current_user_id)
	// 	if err != nil {
	// 		request.Current_User = "N/A"
	// 	} else {
	// 		request.Current_User = user_name
	// 	}
	// }

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
func (em *EditMileageInput) EditMileage() (Mileage_Overview, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, err
	}
	filter := bson.D{{Key: "_id", Value: em.ID}}
	opts := options.FindOne().SetProjection(bson.D{{Key: "action_history", Value: 1}, {Key: "current_status", Value: 1}, {Key: "current_user", Value: 1}, {Key: "last_user_before_reject", Value: 1}})
	request := new(Mileage_Request)
	err = collection.FindOne(context.TODO(), filter, opts).Decode(&request)
	if err != nil {
		return Mileage_Overview{}, err
	}
	if request.Current_Status != "PENDING" && request.Current_Status != "REJECTED" && request.Current_Status != "REJECTED_EDIT_PENDING_REVIEW" {
		return Mileage_Overview{}, errors.New("this request is currently being processed by the organization and not editable")
	}
	err = ClearRequestAssociatedActions(em.ID)
	if err != nil {
		return Mileage_Overview{}, err
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
	case "REJECTED_EDIT_PENDING_REVIEW":
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       em.User_ID,
			Status:     "REJECTED_EDIT",
			Created_At: time.Now(),
		}
		res, err := em.SaveEdits(edit_action, "REJECTED_EDIT_PENDING_REVIEW", request.Current_User)
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
	case "PENDING":
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       em.User_ID,
			Status:     "PENDING_EDIT",
			Created_At: time.Now(),
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
func DeleteMileage(mileage_id string, user_id string, user_admin bool) (Mileage_Overview, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, err
	}
	filter := bson.D{{Key: "_id", Value: mileage_id}}
	request_info := new(Mileage_Overview)
	if !user_admin {
		err = collection.FindOne(context.TODO(), filter).Decode(&request_info)
		if err != nil {
			return Mileage_Overview{}, err
		}
		if request_info.User_ID != user_id {
			return Mileage_Overview{}, errors.New("you're attempting to delete a request for which you are unauthorized")
		}
	}
	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&request_info)
	current_user_id := request_info.Current_User
	user_name, err := FindUserName(current_user_id)
	if err != nil {
		request_info.Current_User = "N/A"
	} else {
		request_info.Current_User = user_name
	}

	return *request_info, nil
}
func MileageDetail(mileage_id string) (MileageDetailResponse, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return MileageDetailResponse{}, err
	}
	filter := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: mileage_id}}}}
	grant_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "grants"}, {Key: "localField", Value: "grant_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "grant"},
		{Key: "pipeline", Value: bson.A{
			bson.D{{Key: "$project", Value: bson.D{{Key: "name", Value: 1}}}}}}}}}
	user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "user_id"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "request_creator"}, {Key: "pipeline", Value: bson.A{
		bson.D{{Key: "$project", Value: bson.D{{Key: "name", Value: 1}}}}}}}}}
	current_user_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "users"}, {Key: "localField", Value: "current_user"}, {Key: "foreignField", Value: "_id"}, {Key: "as", Value: "current_user"}, {Key: "pipeline", Value: bson.A{
		bson.D{{Key: "$project", Value: bson.D{{Key: "name", Value: 1}, {Key: "is_active", Value: 1}}}}}}}}}
	pipeline := mongo.Pipeline{filter, grant_stage, user_stage, current_user_stage}
	request_detail := make([]MileageDetailResponse, 0)
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return MileageDetailResponse{}, err
	}
	err = cursor.All(context.TODO(), &request_detail)
	if err != nil {
		return MileageDetailResponse{}, err
	}
	if len(request_detail) == 0 {
		return MileageDetailResponse{}, errors.New("no mileage request with that id found")
	}
	request_detail[0].Reimbursement = ToFixed(request_detail[0].Reimbursement, 2)
	return request_detail[0], nil
}
func (m *Mileage_Request) Approve(user_id string) (Mileage_Overview, error) {
	collection, err := database.Use("mileage_requests")
	if err != nil {
		return Mileage_Overview{}, err
	}
	filter := bson.D{{Key: "_id", Value: m.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return Mileage_Overview{}, err
	}
	if m.Current_User != user_id {
		return Mileage_Overview{}, errors.New("you are attempting to approve a request for which you are unauthorized")
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
			return Mileage_Overview{}, err
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
	filter := bson.D{{Key: "_id", Value: m.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&m)
	if err != nil {
		return Mileage_Overview{}, err
	}
	if m.Current_User != user_id {
		return Mileage_Overview{}, errors.New("you are attempting to reject a request for which you are unauthorized")
	}
	reject_info := RejectRequest(m.User_ID, m.Current_User)
	m.Action_History = append(m.Action_History, reject_info.Action)
	err = CreateIncompleteAction("mileage", m.ID, reject_info.Action, m.User_ID)
	if err != nil {
		return Mileage_Overview{}, err
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
	// replace with aggregate function for greater flexibility
	// for _, request := range response {
	// 	current_user_id := request.Current_User
	// 	user_name, err := FindUserName(current_user_id)
	// 	if err != nil {
	// 		request.Current_User = "N/A"
	// 	} else {
	// 		request.Current_User = user_name
	// 	}
	// }

	return response, nil
}
