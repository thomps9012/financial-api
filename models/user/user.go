package user

import "time"

type Vehicle struct {
	ID          int16  `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type User struct {
	ID            string    `json:"id" bson:"_id"`
	Email_Address string    `json:"email_address" bson:"email_address"`
	Name          string    `json:"name" bson:"name"`
	Last_Login    time.Time `json:"last_login" bson:"last_login"`
	Vehicles      []Vehicle `json:"vehicles" bson:"vehicles"`
	Manager_ID    string    `json:"manager_id" bson:"manager_id"`
	Is_Active     bool      `json:"is_active" bson:"is_active"`
}
