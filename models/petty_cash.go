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

type Petty_Cash_Request struct {
	ID                      string    `json:"id" bson:"_id"`
	User_ID                 string    `json:"user_id" bson:"user_id"`
	Grant_ID                string    `json:"grant_id" bson:"grant_id"`
	Category                Category  `json:"category" bson:"category"`
	Date                    time.Time `json:"date" bson:"date"`
	Description             string    `json:"description" bson:"description"`
	Amount                  float64   `json:"amount" bson:"amount"`
	Receipts                []string  `json:"receipts" bson:"receipts"`
	Created_At              time.Time `json:"created_at" bson:"created_at"`
	Action_History          []Action  `json:"action_history" bson:"action_history"`
	Current_User            string    `json:"current_user" bson:"current_user"`
	Current_Status          string    `json:"current_status" bson:"current_status"`
	Last_User_Before_Reject string    `json:"last_user_before_reject" bson:"last_user_before_reject"`
	Is_Active               bool      `json:"is_active" bson:"is_active"`
}

type Petty_Cash_Overview struct {
	ID             string    `json:"id" bson:"_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	Amount         float64   `json:"amount" bson:"amount"`
	Grant          Grant     `json:"grant" bson:"grant"`
	Date           time.Time `json:"date" bson:"date"`
	Current_User   string    `json:"current_user" bson:"current_user"`
	Current_Status string    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type PettyCashInput struct {
	Grant_ID    string    `json:"grant_id" bson:"grant_id" validate:"required"`
	Category    Category  `json:"category" bson:"category" validate:"required"`
	Date        time.Time `json:"date" bson:"date" validate:"required"`
	Description string    `json:"description" bson:"description" validate:"required"`
	Amount      float64   `json:"amount" bson:"amount" validate:"required"`
	Receipts    []string  `json:"receipts" bson:"receipts" validate:"required"`
}
type EditPettyCash struct {
	ID          string    `json:"id" bson:"_id" validate:"required"`
	User_ID     string    `json:"user_id" bson:"user_id" validate:"required"`
	Grant_ID    string    `json:"grant_id" bson:"grant_id" validate:"required"`
	Category    Category  `json:"category" bson:"category" validate:"required"`
	Date        time.Time `json:"date" bson:"date" validate:"required"`
	Description string    `json:"description" bson:"description" validate:"required"`
	Amount      float64   `json:"amount" bson:"amount" validate:"required"`
	Receipts    []string  `json:"receipts" bson:"receipts" validate:"required"`
}
type FindPettyCashInput struct {
	PettyCashID string `json:"petty_cash_id" bson:"petty_cash_id" validate:"required"`
}

func GetUserPettyCash(user_id string) ([]Petty_Cash_Overview, error) {
	collection, err := database.Use("petty_cash")
	requests := make([]Petty_Cash_Overview, 0)
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	filter := bson.D{{"user_id", user_id}}
	projection := bson.D{{"_id", 1}, {"user_id", 1}, {"date", 1}, {"amount", 1}, {"current_user", 1}, {"current_status", 1}, {"is_active", 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Petty_Cash_Overview{}, err
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
func GetUserPettyCashDetail(user_id string) ([]Petty_Cash_Request, error) {
	collection, err := database.Use("petty_cash")
	requests := make([]Petty_Cash_Request, 0)
	if err != nil {
		return []Petty_Cash_Request{}, err
	}
	filter := bson.D{{"user_id", user_id}}
	projection := bson.D{{"action_history", 0}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Petty_Cash_Request{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Petty_Cash_Request{}, err
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
func (pi *PettyCashInput) CreatePettyCash(user_id string) (Petty_Cash_Overview, error) {
	new_request := new(Petty_Cash_Request)
	new_request.ID = uuid.NewString()
	new_request.Grant_ID = pi.Grant_ID
	new_request.User_ID = user_id
	new_request.Date = pi.Date
	new_request.Category = pi.Category
	new_request.Description = pi.Description
	new_request.Receipts = pi.Receipts
	new_request.Amount = pi.Amount
	new_request.Created_At = time.Now()
	new_request.Last_User_Before_Reject = bson.TypeNull.String()
	new_request.Is_Active = true
	first_action, _ := FirstActions("petty_cash", new_request.ID, user_id)
	new_request.Action_History = first_action
	current_user := methods.NewRequestUser("petty_cash", "nil")
	new_request.Current_User = current_user.ID
	new_request.Current_Status = "PENDING"
	check_req_coll, err := database.Use("petty_cash")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	_, err = check_req_coll.InsertOne(context.TODO(), new_request)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	return Petty_Cash_Overview{
		ID:             new_request.ID,
		User_ID:        user_id,
		Date:           new_request.Date,
		Amount:         new_request.Amount,
		Current_Status: new_request.Current_Status,
		Current_User:   current_user.Name,
		Is_Active:      true,
	}, nil
}
func (ep *EditPettyCash) EditPettyCash() (Petty_Cash_Overview, error) {
	collection, err := database.Use("petty_cash")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	filter := bson.D{{"_id", ep.ID}}
	opts := options.FindOne().SetProjection(bson.D{{"action_history", 1}, {"current_status", 1}, {"current_user", 1}})
	request := new(Petty_Cash_Request)
	err = collection.FindOne(context.TODO(), filter, opts).Decode(&request)
	if request.Current_Status == "REJECTED" {
		edit_action := Action{
			ID:           uuid.NewString(),
			Request_ID:   ep.ID,
			Request_Type: MILEAGE,
			User:         ep.User_ID,
			Status:       "REJECTED_EDIT",
			Created_At:   time.Now(),
		}
		err := ep.SaveEdits(edit_action, "REJECTED_EDIT", request.Last_User_Before_Reject)
		if err != nil {
			return Petty_Cash_Overview{}, err
		}
	}
	if request.Current_Status == "PENDING" {
		edit_action := Action{
			ID:           uuid.NewString(),
			Request_ID:   ep.ID,
			Request_Type: MILEAGE,
			User:         ep.User_ID,
			Status:       "PENDING_EDIT",
			Created_At:   time.Now(),
		}
		err := ep.SaveEdits(edit_action, "PENDING", request.Current_User)
		if err != nil {
			return Petty_Cash_Overview{}, err
		}
	}
	return Petty_Cash_Overview{}, errors.New("this request is currently being processed by the finance team")
}
func (ep *EditPettyCash) SaveEdits(action Action, new_status string, new_user string) error {
	collection, err := database.Use("petty_cash")
	if err != nil {
		return err
	}
	update := bson.D{{"$set", bson.D{{"grant_id", ep.Grant_ID}, {"date", ep.Date}, {"category", ep.Category}, {"amount", ep.Amount}, {"description", ep.Description}, {"receipts", ep.Receipts}, {"current_status", new_status}, {"current_user", new_user}}}, {"$push", bson.D{{"action_history", action}}}}
	filter := bson.D{{"_id", ep.ID}}
	mileage_req := new(Petty_Cash_Request)
	err = collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&mileage_req)
	if err != nil {
		return err
	}
	return nil
}
func DeletePettyCash(request_id string) (Petty_Cash_Overview, error) {
	collection, err := database.Use("petty_cash")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	filter := bson.D{{"_id", request_id}}
	request_info := new(Petty_Cash_Overview)
	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&request_info)
	if err != nil {
		return Petty_Cash_Overview{}, err
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
func PettyCashDetails(petty_cash_id string) (Petty_Cash_Request, error) {
	collection, err := database.Use("petty_cash")
	if err != nil {
		return Petty_Cash_Request{}, err
	}
	filter := bson.D{{"_id", petty_cash_id}}
	request_detail := new(Petty_Cash_Request)
	err = collection.FindOne(context.TODO(), filter).Decode(&request_detail)
	if err != nil {
		return Petty_Cash_Request{}, err
	}
	return *request_detail, nil
}
func (c *Petty_Cash_Request) Approve() (Petty_Cash_Overview, error) {
	collection, err := database.Use("petty_cash")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	filter := bson.D{{"_id", c.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	new_action, err := ApproveRequest("petty_cash", c.ID, c.Current_User, c.Category, c.Current_Status)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	c.Action_History = append(c.Action_History, new_action.Action)
	update := bson.D{{"$set", bson.D{{"current_user", new_action.NewUser.ID}, {"action_history", c.Action_History}, {"current_status", new_action.Action.Status}}}}
	response := new(Petty_Cash_Overview)
	err = collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&response)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	response.Current_User = new_action.NewUser.Name
	return *response, nil
}
func (c *Petty_Cash_Request) Reject() (Petty_Cash_Overview, error) {
	collection, err := database.Use("petty_cash")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	filter := bson.D{{"_id", c.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	reject_info := RejectRequest("petty_cash", c.ID, c.User_ID, c.Current_User)
	c.Action_History = append(c.Action_History, reject_info.Action)
	update := bson.D{{"$set", bson.D{{"current_user", reject_info.NewUser.ID}, {"action_history", c.Action_History}, {"current_status", "REJECTED"}, {"last_user_before_reject", reject_info.LastUserBeforeReject.ID}}}}
	response := new(Petty_Cash_Overview)
	err = collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&response)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	current_user_name, err := FindUserName(reject_info.NewUser.ID)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	response.Current_User = current_user_name
	return *response, nil
}
func MonthlyPettyCash(month int, year int) ([]Petty_Cash_Overview, error) {
	collection, err := database.Use("petty_cash")
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	response := make([]Petty_Cash_Overview, 0)
	start_date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	end_date := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	filter := bson.D{{"date", bson.D{{"$lte", end_date}, {"$gte", start_date}}}}
	projection := bson.D{{"_id", 1}, {"user_id", 1}, {"date", 1}, {"amount", 1}, {"current_user", 1}, {"current_status", 1}, {"is_active", 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	err = cursor.All(context.TODO(), &response)
	if err != nil {
		return []Petty_Cash_Overview{}, err
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
