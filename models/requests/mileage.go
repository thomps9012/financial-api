package requests

import "time"

type Mileage_Request struct {
	ID                string    `json:"id" bson:"_id"`
	Date              time.Time `json:"date" bson:"date"`
	Starting_Location string    `json:"starting_location" bson:"starting_location"`
	Destination       string    `json:"destination" bson:"destination"`
	Trip_Purpose      string    `json:"trip_purpose" bson:"trip_purpose"`
	Start_Odometer    int64     `json:"start_odometer" bson:"start_odometer"`
	End_Odometer      int64     `json:"end_odometer" bson:"end_odometer"`
	Tolls             float64   `json:"tolls" bson:"tolls"`
	Parking           float64   `json:"parking" bson:"parking"`
	Trip_Mileage      int64     `json:"trip_mileage" bson:"trip_mileage"`
	Reimbursement     float64   `json:"reimbursement" bson:"reimbursement"`
	Created_At        time.Time `json:"created_at" bson:"created_at"`
	User_ID           string    `json:"user_id" bson:"user_id"`
	Action_History    []Action  `json:"action_history" bson:"action_history"`
	Current_Status    Status    `json:"current_status" bson:"current_status"`
	Is_Active         bool      `json:"is_active" bson:"is_active"`
}
