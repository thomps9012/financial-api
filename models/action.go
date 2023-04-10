package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Action struct {
	ID         string    `json:"id" bson:"_id"`
	User       string    `json:"user" bson:"user"`
	Status     string    `json:"status" bson:"status"`
	Created_At time.Time `json:"created_at" bson:"created_at"`
}

type ApproveAction struct {
	Action  Action    `json:"action"`
	NewUser UserLogin `json:"new_user"`
}

type RejectAction struct {
	Action               Action    `json:"action"`
	NewUser              UserLogin `json:"new_user"`
	LastUserBeforeReject UserLogin `json:"last_user"`
}

func NormalizeType(request_type string) string {
	return strings.ToLower(strings.TrimSpace(request_type))
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
func NewUserHandler(category Category, current_status string, exec_review bool) UserLogin {
	if exec_review || current_status == "FINANCE_APPROVED" {
		return UserLogin{
			ID:   "117117754499201658837",
			Name: "Anita Bradley",
		}
	}
	if current_status == "SUPERVISOR_APPROVED" {
		return UserLogin{
			ID:   "109157735191825776845",
			Name: "Finance Requests",
		}
	} else {
		switch category {
		case PERKINS:
		case LORAIN:
			return UserLogin{
				ID:   "109157735191825776845",
				Name: "Jeff Ward",
			}
		case NEXT_STEP:
		case PREVENTION:
			return UserLogin{
				ID:   "109157735191825776845",
				Name: "Cynthia Woods",
			}
		default:
			return UserLogin{
				ID:   "109157735191825776845",
				Name: "Finance Requests",
			}
		}
	}
	return UserLogin{
		ID:   "109157735191825776845",
		Name: "Finance Requests",
	}
}
func ApproveRequest(request_type string, user_id string, request_category Category, current_status string) (ApproveAction, error) {
	switch NormalizeType(request_type) {
	case "mileage":
		return ApproveAction{
			Action: Action{
				ID:         uuid.NewString(),
				User:       user_id,
				Status:     "FINANCE_APPROVED",
				Created_At: time.Now(),
			},
			NewUser: UserLogin{
				ID:   "117117754499201658837",
				Name: "Anita Bradley",
			},
		}, nil
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
		}, nil
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
		NewUser: UserLogin{
			ID: request_creator,
		},
		LastUserBeforeReject: UserLogin{
			ID: current_user,
		},
	}
}
