package models

import (
	"context"
	database "financial-api/db"
	"financial-api/methods"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Action struct {
	ID         string    `json:"id" bson:"_id"`
	User       string    `json:"user" bson:"user"`
	Status     string    `json:"status" bson:"status"`
	Created_At time.Time `json:"created_at" bson:"created_at"`
}

type ApproveAction struct {
	Action  Action              `json:"action"`
	NewUser methods.CurrentUser `json:"new_user"`
}

type RejectAction struct {
	Action               Action              `json:"action"`
	NewUser              methods.CurrentUser `json:"new_user"`
	LastUserBeforeReject methods.CurrentUser `json:"last_user"`
}

type IncompleteAction struct {
	ID          string `json:"id" bson:"_id"`
	ActionID    string `json:"action_id" bson:"action_id"`
	UserID      string `json:"user_id" bson:"user_id"`
	RequestID   string `json:"request_id" bson:"request_id"`
	RequestType string `json:"request_type" bson:"request_type"`
}

func NormalizeType(request_type string) string {
	return strings.ToLower(strings.TrimSpace(request_type))
}
func CreateIncompleteAction(request_type string, request_id string, action Action, user string) error {
	incomplete_action := new(IncompleteAction)
	incomplete_action.ID = uuid.NewString()
	incomplete_action.ActionID = action.ID
	incomplete_action.UserID = user
	incomplete_action.RequestID = request_id
	incomplete_action.RequestType = NormalizeType(request_type)
	collection, err := database.Use("incomplete_actions")
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.TODO(), incomplete_action)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	return nil
}
func ClearIncompleteAction(request_type string, request_id string, action_taker string) error {
	collection, err := database.Use("incomplete_actions")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "request_id", Value: request_id}, {Key: "request_type", Value: request_type}, {Key: "user_id", Value: action_taker}}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
func ClearRequestAssociatedActions(request_id string) error {
	collection, err := database.Use("incomplete_actions")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "request_id", Value: request_id}}
	_, err = collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
func FirstActions(user_id string) []Action {
	return []Action{
		{
			ID:         uuid.NewString(),
			User:       user_id,
			Status:     "CREATED",
			Created_At: time.Now(),
		},
	}
}
func ApproveStatusHandler(category Category, current_status string, exec_review bool) string {
	if exec_review {
		return "PENDING"
	} else if current_status == "FINANCE_APPROVED" {
		return "ORGANIZATION_APPROVED"
	} else if current_status == "SUPERVISOR_APPROVED" {
		return "FINANCE_APPROVED"
	} else if current_status == "MANAGER_APPROVED" {
		return "SUPERVISOR_APPROVED"
	} else {
		switch category {
		case ADMINISTRATIVE:
		case IHBT:
		case PEERS:
		case IOP:
		case MENS_HOUSE:
			return "SUPERVISOR_APPROVED"
		case INTAKE:
		case ACT_TEAM:
		case LORAIN:
		case FINANCE:
		case NEXT_STEP:
		case PERKINS:
		case PREVENTION:
			return "MANAGER_APPROVED"
		}
	}
	return "MANAGER_APPROVED"
}

func NewUserHandler(category Category, current_status string, exec_review bool) methods.CurrentUser {
	if exec_review {
		return methods.FINANCE_FULFILLMENT
	}
	if current_status == "ORGANIZATION_APPROVED" {
		return methods.END_USER
	}
	if current_status == "FINANCE_APPROVED" {
		return methods.FINANCE_FULFILLMENT
	}
	if current_status == "SUPERVISOR_APPROVED" {
		return methods.FINANCE_SUPERVISOR
	} else {
		switch category {
		case PERKINS:
			return methods.PERKINS_SUPERVISOR
		case LORAIN:
			return methods.LORAIN_SUPERVISOR
		case NEXT_STEP:
			return methods.NEXT_STEP_SUPERVISOR
		case PREVENTION:
			return methods.PREVENTION_SUPERVISOR
		default:
			return methods.FINANCE_SUPERVISOR
		}
	}
}
func ApproveMileageHandler(user_id string, current_status string) ApproveAction {
	if current_status == "FINANCE_APPROVED" {
		return ApproveAction{
			Action: Action{
				ID:         uuid.NewString(),
				User:       user_id,
				Status:     "ORGANIZATION_APPROVED",
				Created_At: time.Now(),
			},
			NewUser: methods.END_USER,
		}
	}
	return ApproveAction{
		Action: Action{
			ID:         uuid.NewString(),
			User:       user_id,
			Status:     "FINANCE_APPROVED",
			Created_At: time.Now(),
		},
		NewUser: methods.FINANCE_FULFILLMENT,
	}
}
func ApproveRequest(request_type string, user_id string, request_category Category, current_status string) ApproveAction {
	switch NormalizeType(request_type) {
	case "mileage":
		action := ApproveMileageHandler(user_id, current_status)
		return action
	default:
		new_user := NewUserHandler(request_category, current_status, false)
		new_status := ApproveStatusHandler(request_category, current_status, false)
		return ApproveAction{
			Action: Action{
				ID:         uuid.NewString(),
				User:       user_id,
				Status:     new_status,
				Created_At: time.Now(),
			},
			NewUser: new_user,
		}
	}
}
func RejectRequest(request_creator string, current_user string) RejectAction {
	return RejectAction{
		Action: Action{
			ID:         uuid.NewString(),
			User:       current_user,
			Status:     "REJECTED",
			Created_At: time.Now(),
		},
		NewUser: methods.CurrentUser{
			ID: request_creator,
		},
		LastUserBeforeReject: methods.CurrentUser{
			ID: current_user,
		},
	}
}
