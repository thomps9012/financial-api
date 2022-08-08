package requests

import "time"

type Status string

const (
	PENDING               Status = "PENDING"
	MANAGER_APPROVED      Status = "MANAGER_APPROVED"
	FINANCE_APPROVED      Status = "FINANACE_APPROVED"
	ORGANIZATION_APPROVED Status = "ORG_APPROVED"
	REJECTED              Status = "REJECTED"
	ARCHIVED              Status = "ARCHIVED"
)

type Action struct {
	ID         string    `json:"id" bson:"_id"`
	User_ID    string    `json:"user_id" bson:"user_id"`
	Status     Status    `json:"status" bson:"status"`
	Created_At time.Time `json:"created_at" bson:"created_at"`
}
