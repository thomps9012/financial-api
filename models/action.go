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
func ApproveStatusHandler(current_user_id string, category Category, exec_review bool) string {
	if exec_review {
		return "PENDING"
	}
	if current_user_id == methods.FINANCE_FULFILLMENT.ID {
		return "ORGANIZATION_APPROVED"
	}
	if current_user_id == methods.FINANCE_SUPERVISOR.ID {
		return "FINANCE_APPROVED"
	}
	switch category {
	case IOP:
		return "MANAGER_APPROVED"
	case INTAKE:
		return "MANAGER_APPROVED"
	case PEERS:
		return "MANAGER_APPROVED"
	case ACT_TEAM:
		return "MANAGER_APPROVED"
	case IHBT:
		return "MANAGER_APPROVED"
	case PERKINS:
		if current_user_id == methods.PERKINS_SUPERVISOR.ID {
			return "SUPERVISOR_APPROVED"
		}
		return "MANAGER_APPROVED"
	case MENS_HOUSE:
		return "MANAGER_APPROVED"
	case NEXT_STEP:
		if current_user_id == methods.NEXT_STEP_SUPERVISOR.ID {
			return "SUPERVISOR_APPROVED"
		}
		return "MANAGER_APPROVED"
	case LORAIN:
		if current_user_id == methods.LORAIN_SUPERVISOR.ID {
			return "SUPERVISOR_APPROVED"
		}
		return "MANAGER_APPROVED"
	case PREVENTION:
		if current_user_id == methods.PREVENTION_SUPERVISOR.ID {
			return "SUPERVISOR_APPROVED"
		}
		return "MANAGER_APPROVED"
	case ADMINISTRATIVE:
		return "SUPERVISOR_APPROVED"
	case FINANCE:
		return "MANAGER_APPROVED"
	default:
		return "MANAGER_APPROVED"
	}
}
func NewUserHandler(current_user_id string, category Category, exec_review bool) methods.CurrentUser {
	if exec_review {
		return methods.FINANCE_FULFILLMENT
	}
	if current_user_id == methods.FINANCE_FULFILLMENT.ID {
		return methods.END_USER
	}
	if current_user_id == methods.FINANCE_SUPERVISOR.ID {
		return methods.FINANCE_FULFILLMENT
	}
	if current_user_id == methods.FINANCE_MANAGER.ID {
		return methods.FINANCE_SUPERVISOR
	}
	switch category {
	case IOP:
		if current_user_id == methods.IOP_MANAGER.ID {
			return methods.FINANCE_SUPERVISOR
		}
		return methods.IOP_MANAGER
	case INTAKE:
		if current_user_id == methods.INTAKE_MANAGER.ID {
			return methods.FINANCE_SUPERVISOR
		}
		return methods.INTAKE_MANAGER
	case PEERS:
		if current_user_id == methods.PEER_MANAGER.ID {
			return methods.FINANCE_SUPERVISOR
		}
		return methods.PEER_MANAGER
	case ACT_TEAM:
		if current_user_id == methods.ACT_MANAGER.ID {
			return methods.FINANCE_SUPERVISOR
		}
		return methods.ACT_MANAGER
	case IHBT:
		if current_user_id == methods.IHBT_MANAGER.ID {
			return methods.FINANCE_SUPERVISOR
		}
		return methods.IHBT_MANAGER
	case PERKINS:
		if current_user_id == methods.PERKINS_SUPERVISOR.ID {
			return methods.FINANCE_SUPERVISOR
		}
		if current_user_id == methods.PERKINS_MANAGER.ID {
			return methods.PERKINS_SUPERVISOR
		}
		return methods.PERKINS_MANAGER
	case MENS_HOUSE:
		if current_user_id == methods.MENS_HOUSE_MANAGER.ID {
			return methods.FINANCE_SUPERVISOR
		}
		return methods.MENS_HOUSE_MANAGER
	case NEXT_STEP:
		if current_user_id == methods.NEXT_STEP_SUPERVISOR.ID {
			return methods.FINANCE_SUPERVISOR
		}
		if current_user_id == methods.NEXT_STEP_MANAGER.ID {
			return methods.NEXT_STEP_SUPERVISOR
		}
		return methods.NEXT_STEP_MANAGER
	case LORAIN:
		if current_user_id == methods.LORAIN_SUPERVISOR.ID {
			return methods.FINANCE_SUPERVISOR
		}
		if current_user_id == methods.LORAIN_MANAGER.ID {
			return methods.LORAIN_SUPERVISOR
		}
		return methods.LORAIN_MANAGER
	case PREVENTION:
		if current_user_id == methods.PREVENTION_SUPERVISOR.ID {
			return methods.FINANCE_SUPERVISOR
		}
		if current_user_id == methods.PREVENTION_MANAGER.ID {
			return methods.PREVENTION_SUPERVISOR
		}
		return methods.PREVENTION_MANAGER
	case ADMINISTRATIVE:
		if current_user_id == methods.ADMINISTRATIVE_SUPERVISOR.ID {
			return methods.FINANCE_SUPERVISOR
		}
		return methods.ADMINISTRATIVE_SUPERVISOR
	case FINANCE:
		if current_user_id == methods.FINANCE_SUPERVISOR.ID {
			return methods.FINANCE_FULFILLMENT
		}
		if current_user_id == methods.FINANCE_MANAGER.ID {
			return methods.FINANCE_SUPERVISOR
		}
		return methods.FINANCE_MANAGER
	default:
		return methods.FINANCE_SUPERVISOR
	}
}
func ApproveMileageHandler(user_id string, current_status string) ApproveAction {
	if current_status == "FINANCE_APPROVED" || user_id == methods.FINANCE_FULFILLMENT.ID {
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
func ApproveRequest(request_type string, current_user string, request_category Category, current_status string) ApproveAction {
	switch NormalizeType(request_type) {
	case "mileage":
		action := ApproveMileageHandler(current_user, current_status)
		return action
	default:
		new_user := NewUserHandler(current_user, request_category, false)
		new_status := ApproveStatusHandler(current_user, request_category, false)
		return ApproveAction{
			Action: Action{
				ID:         uuid.NewString(),
				User:       current_user,
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
