package models

import (
	"context"
	database "financial-api/db"
	"time"

	"github.com/google/uuid"
)

type Status string

const (
	CREATED               Status = "CREATED"
	PENDING               Status = "PENDING"
	MANAGER_APPROVED      Status = "MANAGER_APPROVED"
	MANAGER_REJECTED      Status = "MANAGER_REJECTED"
	SUPERVISOR_APPROVED   Status = "SUPERVISOR_APPROVED"
	SUPERVISOR_REJECTED   Status = "SUPERVISOR_REJECTED"
	FINANCE_APPROVED      Status = "FINANCE_APPROVED"
	FINANCE_REJECTED      Status = "FINANCE_REJECTED"
	ORGANIZATION_APPROVED Status = "ORGANIZATION_APPROVED"
	ORGANIZATION_REJECTED Status = "ORGANIZATION_REJECTED"
	REJECTED_EDIT         Status = "REJECTED_EDIT"
	EDIT                  Status = "EDIT"
	ARCHIVED              Status = "ARCHIVED"
)

type Category string

const (
	IOP            Category = "IOP"
	INTAKE         Category = "INTAKE"
	PEERS          Category = "PEERS"
	ACT_TEAM       Category = "ACT_TEAM"
	IHBT           Category = "IHBT"
	PERKINS        Category = "PERKINS"
	MENS_HOUSE     Category = "MENS_HOUSE"
	NEXT_STEP      Category = "NEXT_STEP"
	LORAIN         Category = "LORAIN"
	PREVENTION     Category = "PREVENTION"
	ADMINISTRATIVE Category = "ADMINISTRATIVE"
	FINANCE        Category = "FINANCE"
)

type Request_Type string

const (
	MILEAGE    Request_Type = "MILEAGE"
	CHECK      Request_Type = "CHECK"
	PETTY_CASH Request_Type = "PETTY_CASH"
)

type Request_Response struct {
	User_ID        string `json:"user_id" bson:"user_id"`
	Current_Status string
	Success        bool
}

type Request_Info struct {
	User_ID        string       `json:"user_id" bson:"user_id"`
	Current_User   string       `json:"current_user" bson:"current_user"`
	Current_Status Status       `json:"current_status" bson:"current_status"`
	Type           Request_Type `json:"type" bson:"type"`
	ID             string       `json:"id" bson:"_id"`
}

type Request_Info_With_Action_History struct {
	User_ID        string       `json:"user_id" bson:"user_id"`
	Current_User   string       `json:"current_user" bson:"current_user"`
	Current_Status Status       `json:"current_status" bson:"current_status"`
	Type           Request_Type `json:"type" bson:"type"`
	ID             string       `json:"id" bson:"_id"`
	Action_History []Action     `json:"action_history" bson:"action_history"`
}

type MonthlyRequestInput struct {
	Month time.Month `json:"month" bson:"month" validate:"required"`
	Year  int        `json:"year" bson:"year" validate:"required"`
}

type ApproveRejectRequest struct {
	RequestID string `json:"request_id" bson:"request_id" validate:"required"`
}

type ErrorLog struct {
	ID           string    `json:"id" bson:"_id"`
	UserID       string    `json:"user_id" bson:"user_id" validate:"required"`
	Error        string    `json:"error" bson:"error" validate:"required"`
	ErrorPath    string    `json:"error_path" bson:"error_path" validate:"required"`
	ErrorMessage string    `json:"error_message" bson:"error_message" validate:"required"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}

type ErrorLogOverview struct {
	ID           string    `json:"id" bson:"_id"`
	ErrorMessage string    `json:"error_message" bson:"error_message" validate:"required"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
}

func (el *ErrorLog) Save() (ErrorLogOverview, error) {
	el.ID = uuid.NewString()
	el.CreatedAt = time.Now()
	collection, err := database.Use("error_logs")
	if err != nil {
		return ErrorLogOverview{}, err
	}
	_, err = collection.InsertOne(context.TODO(), el)
	if err != nil {
		return ErrorLogOverview{}, err
	}
	return ErrorLogOverview{
		ID:           el.ID,
		ErrorMessage: el.ErrorMessage,
		CreatedAt:    el.CreatedAt,
	}, nil
}

// // test coverage
// func (r *Request_Info) CheckStatus(new_status Status) bool {
// 	return r.Current_Status != new_status
// }

// // loose test coverage
// func UpdateRequest(new_action Action, user_id string) (bool, error) {
// 	var mileage Mileage_Request
// 	var check Check_Request
// 	var petty Petty_Cash_Request
// 	switch new_action.Request_Type {
// 	case MILEAGE:
// 		_, err := mileage.UpdateActionHistory(new_action, user_id)
// 		if err != nil {
// 			panic(err)
// 		} else {
// 			return true, nil
// 		}
// 	case CHECK:
// 		_, err := check.UpdateActionHistory(new_action, user_id)
// 		if err != nil {
// 			panic(err)
// 		} else {
// 			return true, nil
// 		}
// 	case PETTY_CASH:
// 		_, err := petty.UpdateActionHistory(new_action, user_id)
// 		if err != nil {
// 			panic(err)
// 		} else {
// 			return true, nil
// 		}
// 	}
// 	return false, errors.New("request not updated")
// }

// func DetermineUserID(current_user_email string, request_info Request_Info) (string, error) {
// 	var user User
// 	if current_user_email == "" {
// 		user_id := request_info.User_ID
// 		return user_id, nil
// 	} else {
// 		user_id, err := user.FindID(current_user_email)
// 		if err != nil {
// 			panic(err)
// 		}
// 		return user_id, nil
// 	}
// }

// // test coverage
// func UserEmailHandler(category Category, current_status Status, exec_review bool) string {
// 	// possible more build out of test scenarios here
// 	var to_email = ""
// 	if exec_review || current_status == FINANCE_APPROVED {
// 		to_email = "abradley@norainc.org"
// 	} else if current_status == REJECTED || current_status == ORGANIZATION_APPROVED {
// 		to_email = ""
// 	} else if current_status == SUPERVISOR_APPROVED || current_status == EXECUTIVE_APPROVED {
// 		to_email = "finance_requests@norainc.org"
// 	} else if current_status == MANAGER_APPROVED {
// 		switch category {
// 		case LORAIN:
// 			to_email = "jward@norainc.org"
// 		case NEXT_STEP:
// 			to_email = "cwoods@norainc.org"
// 		case PERKINS:
// 			to_email = "jward@norainc.org"
// 		case PREVENTION:
// 			to_email = "cwoods@norainc.org"
// 		default:
// 			to_email = "finance_requests@norainc.org"
// 		}
// 	} else {
// 		switch category {
// 		case ADMINISTRATIVE:
// 			to_email = "bgriffin@norainc.org"
// 		case IOP:
// 			to_email = "jward@norainc.org"
// 		case INTAKE:
// 			to_email = "cwoods@norainc.org"
// 		case PEERS:
// 			to_email = "jward@norainc.org"
// 		case ACT_TEAM:
// 			to_email = "jjordan@norainc.org"
// 		case IHBT:
// 			to_email = "bgriffin@norainc.org"
// 		case FINANCE:
// 			to_email = "lfuentes@norainc.org"
// 		case LORAIN:
// 			to_email = "rgiusti@norainc.org"
// 		case MENS_HOUSE:
// 			to_email = "jward@norainc.org"
// 		case NEXT_STEP:
// 			to_email = "dbaker@norainc.org"
// 		case PERKINS:
// 			to_email = "churt@norainc.org"
// 		case PREVENTION:
// 			to_email = "lamanor@norainc.org"
// 		default:
// 			to_email = "finance_requests@norainc.org"
// 		}
// 	}
// 	return to_email
// }
