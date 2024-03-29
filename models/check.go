package models

import (
	"context"
	database "financial-api/db"
	"financial-api/methods"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Vendor struct {
	Name         string `json:"name" bson:"name" validate:"required"`
	Website      string `json:"website" bson:"website"`
	AddressLine1 string `json:"address_line_one" bson:"address_line_one" validate:"required"`
	AddressLine2 string `json:"address_line_two" bson:"address_line_two"`
}
type Purchase struct {
	Grant_Line_Item string  `json:"grant_line_item" bson:"grant_line_item" validate:"required"`
	Description     string  `json:"description" bson:"description" validate:"required"`
	Amount          float64 `json:"amount" bson:"amount" validate:"required"`
}
type CheckRequestDetailResponse struct {
	ID             string       `json:"id" bson:"_id"`
	Grant          Grant        `json:"grant" bson:"grant"`
	RequestCreator UserNameInfo `json:"request_creator" bson:"request_creator"`
	Date           time.Time    `json:"date" bson:"date"`
	Category       Category     `json:"category" bson:"category"`
	Vendor         Vendor       `json:"vendor" bson:"vendor"`
	Description    string       `json:"description" bson:"description"`
	Purchases      []Purchase   `json:"purchases" bson:"purchases"`
	Receipts       []string     `json:"receipts" bson:"receipts"`
	Order_Total    float64      `json:"order_total" bson:"order_total"`
	Credit_Card    string       `json:"credit_card" bson:"credit_card"`
	Created_At     time.Time    `json:"created_at" bson:"created_at"`
	Action_History []Action     `json:"action_history" bson:"action_history"`
	Current_User   UserNameInfo `json:"current_user" bson:"current_user"`
	Current_Status string       `json:"current_status" bson:"current_status"`
	Is_Active      bool         `json:"is_active" bson:"is_active"`
}
type Check_Request struct {
	ID                      string     `json:"id" bson:"_id"`
	Grant_ID                string     `json:"grant_id" bson:"grant_id"`
	User_ID                 string     `json:"user_id" bson:"user_id"`
	Date                    time.Time  `json:"date" bson:"date"`
	Category                Category   `json:"category" bson:"category"`
	Vendor                  Vendor     `json:"vendor" bson:"vendor"`
	Description             string     `json:"description" bson:"description"`
	Purchases               []Purchase `json:"purchases" bson:"purchases"`
	Receipts                []string   `json:"receipts" bson:"receipts"`
	Order_Total             float64    `json:"order_total" bson:"order_total"`
	Credit_Card             string     `json:"credit_card" bson:"credit_card"`
	Created_At              time.Time  `json:"created_at" bson:"created_at"`
	Action_History          []Action   `json:"action_history" bson:"action_history"`
	Current_User            string     `json:"current_user" bson:"current_user"`
	Current_Status          string     `json:"current_status" bson:"current_status"`
	Last_User_Before_Reject string     `json:"last_user_before_reject" bson:"last_user_before_reject"`
	Is_Active               bool       `json:"is_active" bson:"is_active"`
}
type CheckRequestInput struct {
	Grant_ID    string     `json:"grant_id" bson:"grant_id" validate:"required"`
	Date        time.Time  `json:"date" bson:"date" validate:"required"`
	Category    Category   `json:"category" bson:"category" validate:"required"`
	Vendor      Vendor     `json:"vendor" bson:"vendor" validate:"required"`
	Description string     `json:"description" bson:"description" validate:"required"`
	Purchases   []Purchase `json:"purchases" bson:"purchases" validate:"required,dive,required"`
	Receipts    []string   `json:"receipts" bson:"receipts" validate:"required"`
	Credit_Card string     `json:"credit_card" bson:"credit_card" validate:"required"`
}
type EditCheckInput struct {
	ID          string     `json:"id" bson:"_id" validate:"required"`
	User_ID     string     `json:"user_id" bson:"user_id" validate:"required"`
	Grant_ID    string     `json:"grant_id" bson:"grant_id" validate:"required"`
	Date        time.Time  `json:"date" bson:"date" validate:"required"`
	Category    Category   `json:"category" bson:"category" validate:"required"`
	Vendor      Vendor     `json:"vendor" bson:"vendor" validate:"required"`
	Description string     `json:"description" bson:"description" validate:"required"`
	Purchases   []Purchase `json:"purchases" bson:"purchases" validate:"required,dive,required"`
	Receipts    []string   `json:"receipts" bson:"receipts" validate:"required"`
	Credit_Card string     `json:"credit_card" bson:"credit_card" validate:"required"`
}
type Check_Request_Overview struct {
	ID             string    `json:"id" bson:"_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	Date           time.Time `json:"date" bson:"date"`
	Order_Total    float64   `json:"order_total" bson:"order_total"`
	Current_Status string    `json:"current_status" bson:"current_status"`
	Current_User   string    `json:"current_user" bson:"current_user"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}
type FindCheckInput struct {
	CheckID string `json:"check_id" bson:"check_id" validate:"required"`
}

func GetUserCheckRequests(user_id string) ([]Check_Request_Overview, error) {
	collection, err := database.Use("check_requests")
	requests := make([]Check_Request_Overview, 0)
	if err != nil {
		return []Check_Request_Overview{}, err
	}
	filter := bson.D{{Key: "user_id", Value: user_id}}
	projection := bson.D{{Key: "_id", Value: 1}, {Key: "user_id", Value: 1}, {Key: "date", Value: 1}, {Key: "order_total", Value: 1}, {Key: "current_user", Value: 1}, {Key: "current_status", Value: 1}, {Key: "is_active", Value: 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Check_Request_Overview{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Check_Request_Overview{}, err
	}
	return requests, nil
}
func GetUserCheckRequestDetail(user_id string) ([]Check_Request, error) {
	collection, err := database.Use("check_requests")
	requests := make([]Check_Request, 0)
	if err != nil {
		return []Check_Request{}, err
	}
	filter := bson.D{{Key: "user_id", Value: user_id}}
	projection := bson.D{{Key: "action_history", Value: 0}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Check_Request{}, err
	}
	err = cursor.All(context.TODO(), &requests)
	if err != nil {
		return []Check_Request{}, err
	}
	return requests, nil
}
func (ci *CheckRequestInput) Exists(user_id string) (bool, *CustomError) {
	check_req_coll, err := database.Use("check_requests")
	if err != nil {
		return false, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	date_filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: ci.Date}}
	grant_filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "grant_id", Value: ci.Grant_ID}, {Key: "category", Value: ci.Category}, {Key: "description", Value: ci.Description}}
	purchase_filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "purchases", Value: ci.Purchases}}
	date_count, err := check_req_coll.CountDocuments(context.TODO(), date_filter)
	if err != nil {
		return false, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	grant_count, err := check_req_coll.CountDocuments(context.TODO(), grant_filter)
	if err != nil {
		return false, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	purchase_count, err := check_req_coll.CountDocuments(context.TODO(), purchase_filter)
	if err != nil {
		return false, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	return date_count+grant_count+purchase_count > 0, nil

}
func (ci *CheckRequestInput) CreateCheckRequest(user_id string) (Check_Request_Overview, *CustomError) {
	new_request := new(Check_Request)
	new_request.ID = uuid.NewString()
	new_request.Grant_ID = ci.Grant_ID
	new_request.User_ID = user_id
	new_request.Date = ci.Date
	new_request.Category = ci.Category
	new_request.Vendor = ci.Vendor
	new_request.Description = ci.Description
	new_request.Purchases = ci.Purchases
	new_request.Receipts = ci.Receipts
	new_request.Credit_Card = ci.Credit_Card
	new_request.Created_At = time.Now()
	new_request.Last_User_Before_Reject = bson.TypeNull.String()
	new_request.Is_Active = true
	first_action := FirstActions(user_id)
	new_request.Action_History = first_action
	current_user := methods.NewRequestUser("check_request", string(new_request.Category), user_id)
	new_request.Current_User = current_user.ID
	new_request.Current_Status = "PENDING"
	total := 0.0
	for _, purchase := range ci.Purchases {
		total += purchase.Amount
	}
	new_request.Order_Total = total
	check_req_coll, err := database.Use("check_requests")
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	_, err = check_req_coll.InsertOne(context.TODO(), new_request)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	err = CreateIncompleteAction("check", new_request.ID, first_action[0], current_user.ID)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}

	return Check_Request_Overview{
		ID:             new_request.ID,
		User_ID:        user_id,
		Date:           new_request.Date,
		Order_Total:    new_request.Order_Total,
		Current_Status: new_request.Current_Status,
		Current_User:   current_user.Name,
		Is_Active:      true,
	}, nil
}
func (ec *EditCheckInput) EditCheckRequest() (Check_Request_Overview, *CustomError) {
	collection, err := database.Use("check_requests")
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "_id", Value: ec.ID}}
	opts := options.FindOne().SetProjection(bson.D{{Key: "action_history", Value: 1}, {Key: "current_status", Value: 1}, {Key: "current_user", Value: 1}, {Key: "last_user_before_reject", Value: 1}})
	request := new(Check_Request)
	err = collection.FindOne(context.TODO(), filter, opts).Decode(&request)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	if request.Current_Status != "PENDING" && request.Current_Status != "REJECTED" && request.Current_Status != "REJECTED_EDIT_PENDING_REVIEW" {
		return Check_Request_Overview{}, &CustomError{
			Status:  423,
			Message: "Request is being processed by organization",
		}
	}
	err = ClearRequestAssociatedActions(ec.ID)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	switch request.Current_Status {
	case "REJECTED":
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       ec.User_ID,
			Status:     "REJECTED_EDIT",
			Created_At: time.Now(),
		}
		res, err := ec.SaveEdits(edit_action, "REJECTED_EDIT_PENDING_REVIEW", request.Last_User_Before_Reject)
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		current_user, err := FindUserName(res.Current_User)
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		return Check_Request_Overview{
			ID:             res.ID,
			User_ID:        res.User_ID,
			Date:           res.Date,
			Order_Total:    res.Order_Total,
			Current_Status: res.Current_Status,
			Current_User:   current_user,
			Is_Active:      res.Is_Active,
		}, nil
	case "REJECTED_EDIT_PENDING_REVIEW":
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       ec.User_ID,
			Status:     "REJECTED_EDIT",
			Created_At: time.Now(),
		}
		res, err := ec.SaveEdits(edit_action, "REJECTED_EDIT_PENDING_REVIEW", request.Current_User)
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		current_user, err := FindUserName(res.Current_User)
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		return Check_Request_Overview{
			ID:             res.ID,
			User_ID:        res.User_ID,
			Date:           res.Date,
			Order_Total:    res.Order_Total,
			Current_Status: res.Current_Status,
			Current_User:   current_user,
			Is_Active:      res.Is_Active,
		}, nil
	case "PENDING":
		edit_action := Action{
			ID:         uuid.NewString(),
			User:       ec.User_ID,
			Status:     "PENDING_EDIT",
			Created_At: time.Now(),
		}
		res, err := ec.SaveEdits(edit_action, "PENDING", request.Current_User)
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		current_user, err := FindUserName(res.Current_User)
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		if err != nil {
			return Check_Request_Overview{
					ID:      res.ID,
					User_ID: res.User_ID,
					Date:    res.Date,
				}, &CustomError{
					Status:  500,
					Message: err.Error(),
				}
		}
		return Check_Request_Overview{
			ID:             res.ID,
			User_ID:        res.User_ID,
			Date:           res.Date,
			Order_Total:    res.Order_Total,
			Current_Status: res.Current_Status,
			Current_User:   current_user,
			Is_Active:      res.Is_Active,
		}, nil
	}

	return Check_Request_Overview{}, &CustomError{
		Status:  423,
		Message: "Request is being processed by organization",
	}
}
func (ec *EditCheckInput) SaveEdits(action Action, new_status string, new_user string) (Check_Request, error) {
	collection, err := database.Use("check_requests")
	if err != nil {
		return Check_Request{}, err
	}
	new_total := 0.0
	for _, purchase := range ec.Purchases {
		new_total += purchase.Amount
	}
	err = CreateIncompleteAction("check", ec.ID, action, new_user)
	if err != nil {
		return Check_Request{}, err
	}
	var update bson.D
	if new_status == "REJECTED_EDIT_PENDING_REVIEW" {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "grant_id", Value: ec.Grant_ID}, {Key: "date", Value: ec.Date}, {Key: "category", Value: ec.Category}, {Key: "vendor", Value: ec.Vendor}, {Key: "description", Value: ec.Description}, {Key: "purchases", Value: ec.Purchases}, {Key: "receipts", Value: ec.Receipts}, {Key: "credit_card", Value: ec.Credit_Card}, {Key: "order_total", Value: new_total}, {Key: "current_status", Value: new_status}, {Key: "current_user", Value: new_user}, {Key: "last_user_before_reject", Value: "null"}}}, {Key: "$push", Value: bson.D{{Key: "action_history", Value: action}}}}
	} else {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "grant_id", Value: ec.Grant_ID}, {Key: "date", Value: ec.Date}, {Key: "category", Value: ec.Category}, {Key: "vendor", Value: ec.Vendor}, {Key: "description", Value: ec.Description}, {Key: "purchases", Value: ec.Purchases}, {Key: "receipts", Value: ec.Receipts}, {Key: "credit_card", Value: ec.Credit_Card}, {Key: "order_total", Value: new_total}, {Key: "current_status", Value: new_status}, {Key: "current_user", Value: new_user}}}, {Key: "$push", Value: bson.D{{Key: "action_history", Value: action}}}}
	}
	filter := bson.D{{Key: "_id", Value: ec.ID}}
	check_req := new(Check_Request)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&check_req)
	if err != nil {
		return Check_Request{}, err
	}

	return *check_req, nil
}
func DeleteCheckRequest(request_id string, user_id string, user_admin bool) (Check_Request_Overview, *CustomError) {
	collection, err := database.Use("check_requests")
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "_id", Value: request_id}}
	request_info := new(Check_Request_Overview)
	if !user_admin {
		err = collection.FindOne(context.TODO(), filter).Decode(&request_info)
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
		if request_info.User_ID != user_id {
			return Check_Request_Overview{}, &CustomError{
				Status:  403,
				Message: "Unauthorized Deletion",
			}
		}
	}
	err = collection.FindOneAndDelete(context.TODO(), filter).Decode(&request_info)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
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
func CheckRequestDetail(check_id string, user_id string, user_admin bool) (CheckRequestDetailResponse, *CustomError) {
	collection, err := database.Use("check_requests")
	if err != nil {
		return CheckRequestDetailResponse{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: check_id}}}}
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
	pipeline := mongo.Pipeline{filter, grant_stage, user_stage, current_user_stage, unwind_current_user, unwind_request_creator, unwind_grant}
	request_detail := make([]CheckRequestDetailResponse, 0)
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return CheckRequestDetailResponse{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	err = cursor.All(context.TODO(), &request_detail)
	if err != nil {
		return CheckRequestDetailResponse{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	if len(request_detail) == 0 {
		return CheckRequestDetailResponse{}, &CustomError{
			Status:  404,
			Message: "Request not found",
		}
	}
	if !user_admin {
		if request_detail[0].RequestCreator.ID != user_id {
			return CheckRequestDetailResponse{}, &CustomError{
				Status:  403,
				Message: "Unauthorized",
			}
		}
	}
	request_detail[0].Order_Total = ToFixed(request_detail[0].Order_Total, 2)
	return request_detail[0], nil
}
func (c *Check_Request) Approve(user_id string) (Check_Request_Overview, *CustomError) {
	collection, err := database.Use("check_requests")
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "_id", Value: c.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	if c.Current_User != user_id {
		return Check_Request_Overview{}, &CustomError{
			Status:  403,
			Message: "Unauthorized",
		}
	}
	new_action := ApproveRequest("check_request", c.Current_User, c.Category, c.Current_Status)
	c.Action_History = append(c.Action_History, new_action.Action)
	var update bson.D
	if new_action.Action.Status == "ORGANIZATION_APPROVED" {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "current_user", Value: new_action.NewUser.ID}, {Key: "is_active", Value: false}, {Key: "action_history", Value: c.Action_History}, {Key: "current_status", Value: new_action.Action.Status}}}}
	} else {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "current_user", Value: new_action.NewUser.ID}, {Key: "action_history", Value: c.Action_History}, {Key: "current_status", Value: new_action.Action.Status}}}}
		err = CreateIncompleteAction("check", c.ID, new_action.Action, new_action.NewUser.ID)
		if err != nil {
			return Check_Request_Overview{}, &CustomError{
				Status:  500,
				Message: err.Error(),
			}
		}
	}
	response := new(Check_Request_Overview)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&response)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	response.Current_User = new_action.NewUser.Name
	return *response, nil
}
func (c *Check_Request) Reject(user_id string) (Check_Request_Overview, *CustomError) {
	collection, err := database.Use("check_requests")
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	filter := bson.D{{Key: "_id", Value: c.ID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	if c.Current_User != user_id {
		return Check_Request_Overview{}, &CustomError{
			Status:  403,
			Message: "Unauthorized",
		}
	}
	reject_info := RejectRequest(c.User_ID, c.Current_User)
	c.Action_History = append(c.Action_History, reject_info.Action)
	err = CreateIncompleteAction("check", c.ID, reject_info.Action, c.User_ID)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "current_user", Value: reject_info.NewUser.ID}, {Key: "action_history", Value: c.Action_History}, {Key: "current_status", Value: "REJECTED"}, {Key: "last_user_before_reject", Value: reject_info.LastUserBeforeReject.ID}}}}
	response := new(Check_Request_Overview)
	upsert := true
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		Upsert:         &upsert,
		ReturnDocument: &after,
	}
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, &opts).Decode(&response)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	current_user_name, err := FindUserName(reject_info.NewUser.ID)
	if err != nil {
		return Check_Request_Overview{}, &CustomError{
			Status:  500,
			Message: err.Error(),
		}
	}
	response.Current_User = current_user_name

	return *response, nil
}
func MonthlyCheckRequests(month int, year int) ([]Check_Request_Overview, error) {
	collection, err := database.Use("check_requests")
	if err != nil {
		return []Check_Request_Overview{}, err
	}
	response := make([]Check_Request_Overview, 0)
	start_date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	end_date := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
	filter := bson.D{{Key: "date", Value: bson.D{{Key: "$lte", Value: end_date}, {Key: "$gte", Value: start_date}}}}
	projection := bson.D{{Key: "_id", Value: 1}, {Key: "user_id", Value: 1}, {Key: "date", Value: 1}, {Key: "order_total", Value: 1}, {Key: "current_user", Value: 1}, {Key: "current_status", Value: 1}, {Key: "is_active", Value: 1}}
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		return []Check_Request_Overview{}, err
	}
	err = cursor.All(context.TODO(), &response)
	if err != nil {
		return []Check_Request_Overview{}, err
	}
	return response, nil
}
