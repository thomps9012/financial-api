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
	collection, err := database.Use("petty_cash_requests")
	requests := make([]Petty_Cash_Overview, 0)
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	filter := bson.D{{Key: "user_id", Value: user_id}}
	projection := bson.D{{Key: "_id", Value: 1}, {Key: "user_id", Value: 1}, {Key: "date", Value: 1}, {Key: "amount", Value: 1}, {Key: "current_user", Value: 1}, {Key: "current_status", Value: 1}, {Key: "is_active", Value: 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Petty_Cash_Overview{}, err
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
func GetUserPettyCashDetail(user_id string) ([]Petty_Cash_Request, error) {
	collection, err := database.Use("petty_cash_requests")
	requests := make([]Petty_Cash_Request, 0)
	if err != nil {
		return []Petty_Cash_Request{}, err
	}
	filter := bson.D{{Key: "user_id", Value: user_id}}
	projection := bson.D{{Key: "action_history", Value: 0}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Petty_Cash_Request{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Petty_Cash_Request{}, err
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
	first_action := FirstActions(user_id)
	new_request.Action_History = first_action
	current_user := methods.NewRequestUser("petty_cash", string(pi.Category))
	new_request.Current_User = current_user.ID
	new_request.Current_Status = "PENDING"
	check_req_coll, err := database.Use("petty_cash_requests")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	_, err = check_req_coll.InsertOne(context.TODO(), new_request)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	err = CreateIncompleteAction("PETTY_CASH", new_request.ID, first_action[0], current_user.ID)
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
	collection, err := database.Use("petty_cash_requests")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	filter := bson.D{{Key: "_id", Value: ep.ID}}
	opts := options.FindOne().SetProjection(bson.D{{Key: "action_history", Value: 1}, {Key: "current_status", Value: 1}, {Key: "current_user", Value: 1}, {Key: "last_user_before_reject", Value: 1}})
	request := new(Petty_Cash_Request)
	err = collection.FindOne(context.TODO(), filter, opts).Decode(&request)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	if request.Current_Status == "REJECTED" || request.Current_Status == "REJECTED_EDIT_PENDING_REVIEW" {
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       ep.User_ID,
			Status:     "REJECTED_EDIT",
			Created_At: time.Now(),
		}
		res, err := ep.SaveEdits(edit_action, "REJECTED_EDIT_PENDING_REVIEW", request.Last_User_Before_Reject)
		if err != nil {
			return Petty_Cash_Overview{}, err
		}
		username, err := FindUserName(res.Current_User)
		if err != nil {
			return Petty_Cash_Overview{}, err
		}
		return Petty_Cash_Overview{
			ID:             res.ID,
			User_ID:        res.User_ID,
			Amount:         res.Amount,
			Date:           res.Date,
			Current_User:   username,
			Current_Status: res.Current_Status,
			Is_Active:      res.Is_Active,
		}, nil
	}
	if request.Current_Status == "PENDING" {
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       ep.User_ID,
			Status:     "PENDING_EDIT",
			Created_At: time.Now(),
		}
		res, err := ep.SaveEdits(edit_action, "PENDING", request.Current_User)
		if err != nil {
			return Petty_Cash_Overview{}, err
		}
		username, err := FindUserName(res.Current_User)
		if err != nil {
			return Petty_Cash_Overview{}, err
		}
		return Petty_Cash_Overview{
			ID:             res.ID,
			User_ID:        res.User_ID,
			Amount:         res.Amount,
			Date:           res.Date,
			Current_User:   username,
			Current_Status: res.Current_Status,
			Is_Active:      res.Is_Active,
		}, nil
	}
	return Petty_Cash_Overview{}, errors.New("this request is currently being processed by the finance team")
}
func (ep *EditPettyCash) SaveEdits(action Action, new_status string, new_user string) (Petty_Cash_Request, error) {
	collection, err := database.Use("petty_cash_requests")
	if err != nil {
		return Petty_Cash_Request{}, err
	}
	var update bson.D
	if new_status == "REJECTED_EDIT_PENDING_REVIEW" {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "grant_id", Value: ep.Grant_ID}, {Key: "date", Value: ep.Date}, {Key: "category", Value: ep.Category}, {Key: "amount", Value: ep.Amount}, {Key: "description", Value: ep.Description}, {Key: "receipts", Value: ep.Receipts}, {Key: "current_status", Value: new_status}, {Key: "current_user", Value: new_user}, {Key: "last_user_before_reject", Value: "null"}}}, {Key: "$push", Value: bson.D{{Key: "action_history", Value: action}}}}
	} else {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "grant_id", Value: ep.Grant_ID}, {Key: "date", Value: ep.Date}, {Key: "category", Value: ep.Category}, {Key: "amount", Value: ep.Amount}, {Key: "description", Value: ep.Description}, {Key: "receipts", Value: ep.Receipts}, {Key: "current_status", Value: new_status}, {Key: "current_user", Value: new_user}}}, {Key: "$push", Value: bson.D{{Key: "action_history", Value: action}}}}
	}
	filter := bson.D{{Key: "_id", Value: ep.ID}}
	request := new(Petty_Cash_Request)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&request)
	if err != nil {
		return Petty_Cash_Request{}, err
	}
	return *request, nil
}
func DeletePettyCash(request_id string) (Petty_Cash_Overview, error) {
	collection, err := database.Use("petty_cash_requests")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	filter := bson.D{{Key: "_id", Value: request_id}}
	request_info := new(Petty_Cash_Overview)
	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&request_info)
	if err != nil {
		return Petty_Cash_Overview{
			ID:             request_id,
			User_ID:        "",
			Amount:         0,
			Date:           time.Time{},
			Current_User:   "",
			Current_Status: "",
			Is_Active:      false,
		}, err
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
	collection, err := database.Use("petty_cash_requests")
	if err != nil {
		return Petty_Cash_Request{}, err
	}
	filter := bson.D{{Key: "_id", Value: petty_cash_id}}
	request_detail := new(Petty_Cash_Request)
	err = collection.FindOne(context.TODO(), filter).Decode(&request_detail)
	if err != nil {
		return Petty_Cash_Request{}, err
	}
	return *request_detail, nil
}
func (c *Petty_Cash_Request) Approve(user_id string) (Petty_Cash_Overview, error) {
	collection, err := database.Use("petty_cash_requests")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	filter := bson.D{{Key: "_id", Value: c.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	if user_id != c.Current_User {
		return Petty_Cash_Overview{}, errors.New("you are attempting to approve a request for which you are not authorized")
	}
	new_action, err := ApproveRequest("petty_cash", c.Current_User, c.Category, c.Current_Status)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	c.Action_History = append(c.Action_History, new_action.Action)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "current_user", Value: new_action.NewUser.ID}, {Key: "action_history", Value: c.Action_History}, {Key: "current_status", Value: new_action.Action.Status}}}}
	response := new(Petty_Cash_Overview)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&response)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	response.Current_User = new_action.NewUser.Name
	return *response, nil
}
func (c *Petty_Cash_Request) Reject(user_id string) (Petty_Cash_Overview, error) {
	collection, err := database.Use("petty_cash_requests")
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	filter := bson.D{{Key: "_id", Value: c.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		return Petty_Cash_Overview{}, err
	}
	if user_id != c.Current_User {
		return Petty_Cash_Overview{}, errors.New("you are attempting to reject a request for which you are not authorized")
	}
	reject_info := RejectRequest(c.User_ID, c.Current_User)
	c.Action_History = append(c.Action_History, reject_info.Action)
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "current_user", Value: reject_info.NewUser.ID}, {Key: "action_history", Value: c.Action_History}, {Key: "current_status", Value: "REJECTED"}, {Key: "last_user_before_reject", Value: reject_info.LastUserBeforeReject.ID}}}}
	response := new(Petty_Cash_Overview)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&response)
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
	collection, err := database.Use("petty_cash_requests")
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	response := make([]Petty_Cash_Overview, 0)
	start_date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	end_date := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	filter := bson.D{{Key: "date", Value: bson.D{{Key: "$lte", Value: end_date}, {Key: "$gte", Value: start_date}}}}
	projection := bson.D{{Key: "_id", Value: 1}, {Key: "user_id", Value: 1}, {Key: "date", Value: 1}, {Key: "amount", Value: 1}, {Key: "current_user", Value: 1}, {Key: "current_status", Value: 1}, {Key: "is_active", Value: 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
	err = cursor.All(context.TODO(), &response)
	if err != nil {
		return []Petty_Cash_Overview{}, err
	}
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
