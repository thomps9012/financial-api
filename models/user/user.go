package user

import (
	"context"
	conn "financial-api/m/db"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// possible expanison after org chart release
type Role string

const (
	EMPLOYEE  Role = "EMPLOYEE"
	MANAGER   Role = "MANAGER"
	FINANCE   Role = "FINANCE"
	EXECUTIVE Role = "EXECUTIVE"
)

type Petty_Cash_Request struct {
	ID             string    `json:"id" bson:"_id"`
	User_ID        string    `json:"user_id" bson:"user_id"`
	Grant_ID       string    `json:"grant_id" bson:"grant_id"`
	Date           time.Time `json:"date" bson:"date"`
	Description    string    `json:"description" bson:"description"`
	Amount         float64   `json:"amount" bson:"amount"`
	Receipts       []string  `json:"receipts" bson:"receipts"`
	Created_At     time.Time `json:"created_at" bson:"created_at"`
	Action_History []Action  `json:"action_history" bson:"action_history"`
	Current_Status Status    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type Vehicle struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type Mileage_Request struct {
	ID                string    `json:"id" bson:"_id"`
	User_ID           string    `json:"user_id" bson:"user_id"`
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
	Action_History    []Action  `json:"action_history" bson:"action_history"`
	Current_Status    Status    `json:"current_status" bson:"current_status"`
	Is_Active         bool      `json:"is_active" bson:"is_active"`
}

type User struct {
	ID                 string    `json:"id" bson:"_id"`
	Email              string    `json:"email" bson:"email"`
	Name               string    `json:"name" bson:"name"`
	Last_Login         time.Time `json:"last_login" bson:"last_login"`
	Vehicles           []Vehicle `json:"vehicles" bson:"vehicles"`
	InComplete_Actions []string  `json:"incomplete_actions" bson:"incomplete_actions"`
	Manager_ID         string    `json:"manager_id" bson:"manager_id"`
	Is_Active          bool      `json:"is_active" bson:"is_active"`
	Role               Role      `json:"role" bson:"role"`
}

type Manager struct {
	ID        string
	Employees []string
}

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
	Grant_Line_Item string  `json:"line_item" bson:"line_item"`
	Description     string  `json:"description" bson:"description"`
	Amount          float64 `json:"amount" bson:"amount"`
}

type Check_Request struct {
	ID             string     `json:"id" bson:"_id"`
	Date           time.Time  `json:"date" bson:"date"`
	Vendor         Vendor     `json:"vendor" bson:"vendor"`
	Description    string     `json:"description" bson:"description"`
	Grant_ID       string     `json:"grant_id" bson:"grant_id"`
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

type User_Monthly_Mileage struct {
	ID            string     `json:"id" bson:"_id"`
	Name          string     `json:"name" bson:"name"`
	Vehicles      []Vehicle  `json:"vehicles" bson:"vehicles"`
	Month         time.Month `json:"month" bson:"month"`
	Year          int        `json:"year" bson:"year"`
	Mileage       int64      `json:"mileage" bson:"mileage"`
	Tolls         float64    `json:"tolls" bson:"tolls"`
	Parking       float64    `json:"parking" bson:"parking"`
	Reimbursement float64    `json:"reimbursement" bson:"reimbursement"`
	Request_IDS   []string   `json:"request_ids" bson:"request_ids"`
}
type User_Monthly_Petty_Cash struct {
	ID           string     `json:"id" bson:"_id"`
	Name         string     `json:"name" bson:"name"`
	Month        time.Month `json:"month" bson:"month"`
	Year         int        `json:"year" bson:"year"`
	Total_Amount float64    `json:"total_amount" bson:"total_amount"`
	Receipts     []string   `json:"receipts" bson:"receipts"`
}

type User_Agg_Check_Requests struct {
	ID           string     `json:"id" bson:"_id"`
	Name         string     `json:"name" bson:"name"`
	Start_Date   string     `json:"start_date" bson:"start_date"`
	End_Date     string     `json:"end_date" bson:"end_date"`
	Total_Amount float64    `json:"total_amount" bson:"total_amount"`
	Vendors      []Vendor   `json:"vendors" bson:"vendors"`
	Purchases    []Purchase `json:"purchases" bson:"purchases"`
	Receipts     []string   `json:"receipts" bson:"receipts"`
}

// can optimize this function with a switch to search certain arrays based on the user's role
func setManagerID(email string, employee_role Role) string {
	var manager_id string
	managers := []Manager{
		{"test1", []string{"id1", "id2", "id3", "id4", "id5"}},
		{"test2", []string{"id6", "id7", "id8", "id9", "id10"}},
		{"test3", []string{"id11", "id12", "id13", "id14", "id15"}},
	}
	executives := []Manager{
		{"test4", []string{"id16", "id22", "id32", "id41", "id35"}},
		{"test5", []string{"id66", "id57", "id84", "id92", "id101"}},
		{"test6", []string{"id161", "id312", "id123", "id149", "id915"}},
	}
	var finance = "finance_id"
	N_A := "N/A"
	switch employee_role {
	case EMPLOYEE:
		for i := range managers {
			var employeesArr = managers[i].Employees
			for s := range employeesArr {
				if employeesArr[s] == email {
					manager_id = managers[i].ID
				}
			}
		}
	case MANAGER:
		for i := range executives {
			var employeesArr = managers[i].Employees
			for s := range employeesArr {
				if employeesArr[s] == email {
					manager_id = managers[i].ID
				}
			}
		}
	case EXECUTIVE:
		manager_id = finance
	default:
		manager_id = N_A
	}
	return manager_id
}

func (u *User) Create(email string, role Role) (string, error) {
	collection := conn.DB.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	var user User
	findErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if findErr == nil {
		return "", fmt.Errorf("account already created")
	}
	u.ID = uuid.NewString()
	u.Role = role
	u.Last_Login = time.Now()
	u.Is_Active = true
	if role != EXECUTIVE {
		manager_id := setManagerID(email, role)
		u.Manager_ID = manager_id
	} else {
		u.Manager_ID = "N/A"
	}
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		panic(err)
	}
	return u.ID, nil
}

func (u *User) Login(email string) (bool, error) {
	collection := conn.DB.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	result, err := collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: bson.M{"last_login": time.Now()}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (u *User) AddVehicle(name string, description string, user_id string) (string, error) {
	collection := conn.DB.Collection("users")
	vehicle := &Vehicle{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$push", Value: bson.M{"vehicles": vehicle}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return "", err
	}
	return vehicle.ID, nil
}
func (u *User) RemoveVehicle(vehicle_id string, user_id string) (bool, error) {
	collection := conn.DB.Collection("users")
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$pull", Value: bson.M{"vehicles": bson.M{"_id": vehicle_id}}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (u *User) Deactivate(user_id string) (bool, error) {
	collection := conn.DB.Collection("users")
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$set", Value: bson.M{"is_active": false}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (u *User) FindByID(user_id string) (User, error) {
	collection := conn.DB.Collection("users")
	var user User
	filter := bson.D{{Key: "_id", Value: user_id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (u *User) FindMgrID(user_id string) (string, error) {
	collection := conn.DB.Collection("users")
	var user User
	filter := bson.D{{Key: "_id", Value: user_id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		panic(err)
	}
	return user.Manager_ID, nil
}

func (u *User) AddNotification(item_id string, user_id string) (bool, error) {
	collection := conn.DB.Collection("users")
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$push", Value: bson.M{"incomplete_actions": item_id}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (u *User) ClearNotifications(user_id string) (bool, error) {
	collection := conn.DB.Collection("users")
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$pullAll", Value: "incomplete_actions"}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (u *User) ClearNotification(item_id string, user_id string) (bool, error) {
	collection := conn.DB.Collection("users")
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$pull", Value: bson.M{"incomplete_actions": item_id}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (u *User) MonthlyMileage(user_id string, month int, year int) (User_Monthly_Mileage, error) {
	collection := conn.DB.Collection("mileage_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
	end_month := month + 1
	start_date := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	end_date := time.Date(year, time.Month(end_month), 0, 0, 0, 0, 0, time.UTC)
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "$gte", Value: bson.M{"date": start_date}}, {Key: "$lte", Value: bson.M{"date": end_date}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var mileage int64
	tolls := 0.0
	parking := 0.0
	reimbursement := 0.0
	var request_ids []string
	for cursor.Next(context.TODO()) {
		var mileage_req Mileage_Request
		decode_err := cursor.Decode(&mileage_req)
		if decode_err != nil {
			panic(decode_err)
		}
		request_ids = append(request_ids, mileage_req.ID)
		mileage += mileage_req.Trip_Mileage
		tolls += mileage_req.Tolls
		parking += mileage_req.Parking
		reimbursement += mileage_req.Reimbursement
	}
	return User_Monthly_Mileage{
		ID:            user_id,
		Name:          result.Name,
		Vehicles:      result.Vehicles,
		Month:         time.Month(month),
		Year:          year,
		Mileage:       mileage,
		Parking:       parking,
		Tolls:         tolls,
		Reimbursement: reimbursement,
		Request_IDS:   request_ids,
	}, nil
}

func (u *User) MonthlyPettyCash(user_id string, month int, year int) (User_Monthly_Petty_Cash, error) {
	collection := conn.DB.Collection("petty_cash_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
	end_month := month + 1
	start_date := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	end_date := time.Date(year, time.Month(end_month), 0, 0, 0, 0, 0, time.UTC)
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "$gte", Value: bson.M{"date": start_date}}, {Key: "$lte", Value: bson.M{"date": end_date}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	total_amount := 0.0
	var receipts []string
	for cursor.Next(context.TODO()) {
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		receipts = append(receipts, petty_cash_req.Receipts...)
		total_amount += petty_cash_req.Amount
	}
	return User_Monthly_Petty_Cash{
		ID:           user_id,
		Name:         result.Name,
		Month:        time.Month(month),
		Year:         year,
		Total_Amount: total_amount,
		Receipts:     receipts,
	}, nil
}

func (u *User) AggregateChecks(user_id string, start_date string, end_date string) (User_Agg_Check_Requests, error) {
	collection := conn.DB.Collection("check_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
	var filter bson.D
	if start_date != "" && end_date != "" {
		filter = bson.D{{Key: "user_id", Value: user_id}, {Key: "$gte", Value: bson.M{"date": start_date}}, {Key: "$lte", Value: bson.M{"date": end_date}}}
	} else {
		filter = bson.D{{Key: "user_id", Value: user_id}}
	}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	total_amount := 0.0
	var vendors []Vendor
	var receipts []string
	var purchases []Purchase
	for cursor.Next(context.TODO()) {
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		vendors = append(vendors, check_req.Vendor)
		purchases = append(purchases, check_req.Purchases...)
		receipts = append(receipts, check_req.Receipts...)
		total_amount += check_req.Order_Total
	}
	return User_Agg_Check_Requests{
		ID:           user_id,
		Name:         result.Name,
		Start_Date:   start_date,
		End_Date:     end_date,
		Total_Amount: total_amount,
		Vendors:      vendors,
		Purchases:    purchases,
		Receipts:     receipts,
	}, nil
}
