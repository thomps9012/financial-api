package requests

import "time"

type Petty_Cash_Request struct {
	ID             string    `json:"id" bson:"_id"`
	Date           time.Time `json:"date" bson:"date"`
	Reimbursement  float64   `json:"reimbursement" bson:"reimbursement"`
	Receipts       []string  `json:"receipts" bson:"receipts"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	Action_History []Action  `json:"action_history" bson:"action_history"`
	Current_Status Status    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}
