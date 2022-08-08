package requests

import "time"

type Address struct {
	Website  string `json:"website" bson:"website"`
	Street   string `json:"street" bson:"street"`
	City     string `json:"city" bson:"city"`
	State    string `json:"state" bson:"state"`
	Zip_Code int64  `json:"zip" bson:"zip"`
}

type Vendor struct {
	Name    string  `json:"name" bson:"name"`
	Address Address `json:"address" bson:"address"`
}

type Purchase struct {
	Grant_ID        string  `json:"grant_id" bson:"grant_id"`
	Grant_Line_Item string  `json:"line_item" bson:"line_item"`
	Description     string  `json:"description" bson:"description"`
	Amount          float64 `json:"amount" bson:"amount"`
}

type Check_Request struct {
	ID             string     `json:"id" bson:"_id"`
	Date           time.Time  `json:"date" bson:"date"`
	Vendor         Vendor     `json:"vendor" bson:"vendor"`
	Description    string     `json:"description" bson:"description"`
	Purchases      []Purchase `json:"purchases" bson:"purchases"`
	Receipts       []string   `json:"receipts" bson:"receipts"`
	Order_Total    float64    `json:"order_total" bson:"order_total"`
	Credit_Card    string     `json:"credit_card" bson:"credit_card"`
	Created_At     time.Time  `json:"created_at" bson:"created_at"`
	User_ID        string     `json:"user_id" bson:"user_id"`
	Action_History []Action   `json:"action_history" bson:"action_history"`
	Current_Status Status     `json:"current_status" bson:"current_status"`
	Is_Active      bool       `json:"is_active" bson:"is_active"`
}
