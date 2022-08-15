package user

import (
	"context"
	conn "financial-api/db"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// possible expanison after org chart release
type Role string

const (
	EMPLOYEE  Role = "EMPLOYEE"
	MANAGER   Role = "MANAGER"
	FINANCE   Role = "FINANCE"
	EXECUTIVE Role = "EXECUTIVE"
)

func (u User) ParseRole(role string) Role {
	var roleParse Role
	switch role {
	case "EMPLOYEE":
		roleParse = EMPLOYEE
	case "MANAGER":
		roleParse = MANAGER
	case "FINANCE":
		roleParse = FINANCE
	case "EXECUTIVE":
		roleParse = EXECUTIVE
	default:
		roleParse = EMPLOYEE
	}
	return roleParse
}

type User_Overview struct {
	ID                      string                  `json:"id" bson:"_id"`
	Manager_ID              string                  `json:"manager_id" bson:"manager_id"`
	Name                    string                  `json:"name" bson:"name"`
	Last_Login              time.Time               `json:"last_login" bson:"last_login"`
	Is_Active               bool                    `json:"is_active" bson:"is_active"`
	Role                    string                  `json:"role" bson:"role"`
	Incomplete_Action_Count int                     `json:"incomplete_action_count" bson:"incomplete_action_count"`
	Mileage_Requests        User_Agg_Mileage        `json:"mileage_requests" bson:"mileage_requests"`
	Check_Requests          User_Agg_Check_Requests `json:"check_requests" bson:"check_requests"`
	Petty_Cash_Requests     User_Agg_Petty_Cash     `json:"petty_cash_requests" bson:"petty_cash_requests"`
}

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
	Start_Odometer    int       `json:"start_odometer" bson:"start_odometer"`
	End_Odometer      int       `json:"end_odometer" bson:"end_odometer"`
	Tolls             float64   `json:"tolls" bson:"tolls"`
	Parking           float64   `json:"parking" bson:"parking"`
	Trip_Mileage      int       `json:"trip_mileage" bson:"trip_mileage"`
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
	Role               string    `json:"role" bson:"role"`
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
	Zip_Code int    `json:"zip" bson:"zip"`
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
	Mileage       int        `json:"mileage" bson:"mileage"`
	Tolls         float64    `json:"tolls" bson:"tolls"`
	Parking       float64    `json:"parking" bson:"parking"`
	Reimbursement float64    `json:"reimbursement" bson:"reimbursement"`
	Request_IDS   []string   `json:"request_ids" bson:"request_ids"`
}
type User_Agg_Mileage struct {
	Vehicles      []Vehicle `json:"vehicles" bson:"vehicles"`
	Mileage       int       `json:"mileage" bson:"mileage"`
	Tolls         float64   `json:"tolls" bson:"tolls"`
	Parking       float64   `json:"parking" bson:"parking"`
	Reimbursement float64   `json:"reimbursement" bson:"reimbursement"`
	Request_IDS   []string  `json:"request_ids" bson:"request_ids"`
}
type User_Agg_Petty_Cash struct {
	Total_Amount float64  `json:"total_amount" bson:"total_amount"`
	Receipts     []string `json:"receipts" bson:"receipts"`
	Request_IDS  []string `json:"request_ids" bson:"request_ids"`
}
type User_Monthly_Petty_Cash struct {
	ID           string     `json:"id" bson:"_id"`
	Name         string     `json:"name" bson:"name"`
	Month        time.Month `json:"month" bson:"month"`
	Year         int        `json:"year" bson:"year"`
	Total_Amount float64    `json:"total_amount" bson:"total_amount"`
	Request_IDS  []string   `json:"request_ids" bson:"request_ids"`
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
	Request_IDS  []string   `json:"request_ids" bson:"request_ids"`
}

// can optimize this function with a switch to search certain arrays based on the user's role
func setManagerID(email string, employee_role string) string {
	var manager_id string
	managers := []Manager{
		{"b771af77-bffe-495a-8777-56ef9a4a1f46", []string{"emp1@norainc.org", "emp65@norainc.org"}},
		{"5960679a-d2f3-475b-b142-00650f8f0ebf", []string{"emp7@norainc.org", "emp87@norainc.org"}},
		{"46092af5-a989-4977-9da8-ca7c84132421", []string{"emp9@norainc.org", "emp10@norainc.org"}},
		{"5e6288d5-9219-4c75-87cf-cdc53fde3958", []string{"emp19@norainc.org", "emp13@norainc.org"}},
		{"29fc8292-8051-4f41-873c-d74bb7241e43", []string{"emp39@norainc.org", "emp52@norainc.org"}},
		{"12b243ea-5654-4b53-92ad-f6f056fd86fe", []string{"emp4S9@norainc.org", "emp99@norainc.org"}},
	}
	executives := []Manager{
		{"68125e1f-21c1-4f60-aab0-8efff5dc158e", []string{"manager1@norainc.org", "manager5@norainc.org"}},
		{"cde4638b-4c33-4015-85fc-b0dd106a4b6b", []string{"manager2@norainc.org", "manager6@norainc.org"}},
		{"acbe6899-200f-4185-8624-b31e32c42b44", []string{"manager3@norainc.org", "manager4@norainc.org"}},
	}
	var finance = "cbaf2ee1-79d7-40da-bb4c-a6017d4fb705"
	N_A := "N/A"
	switch employee_role {
	case "EMPLOYEE":
		for i := range managers {
			var employeesArr = managers[i].Employees
			for s := range employeesArr {
				if employeesArr[s] == email {
					manager_id = managers[i].ID
				}
			}
		}
	case "MANAGER":
		for i := range executives {
			var employeesArr = executives[i].Employees
			for s := range employeesArr {
				if employeesArr[s] == email {
					manager_id = executives[i].ID
				}
			}
		}
	case "EXECUTIVE":
		manager_id = finance
	default:
		manager_id = N_A
	}
	return manager_id
}

func (u *User) Create(email string, role string) (User, error) {
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	var user User
	findErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if findErr == nil {
		return *u, fmt.Errorf("account already created")
	}
	u.ID = uuid.NewString()
	u.Role = role
	u.Last_Login = time.Now()
	u.Is_Active = true
	u.Email = email
	u.Vehicles = []Vehicle{}
	u.InComplete_Actions = []string{}
	manager_id := setManagerID(email, role)
	u.Manager_ID = manager_id
	_, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		panic(err)
	}
	return *u, nil
}

func (u *User) Login(email string) (User, error) {
	var user User
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.M{"last_login": time.Now()}}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&user)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (u *User) AddVehicle(user_id string, name string, description string) (string, error) {
	collection := conn.Db.Collection("users")
	vehicle := &Vehicle{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$push", Value: bson.M{"vehicles": vehicle}}})
	if err != nil {
		panic(err)
	}
	println(result.ModifiedCount)
	if result.ModifiedCount == 0 {
		return "", err
	}
	return vehicle.ID, nil
}
func (u *User) RemoveVehicle(user_id string, vehicle_id string) (bool, error) {
	collection := conn.Db.Collection("users")
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$pull", Value: bson.M{"vehicles": bson.M{"_id": vehicle_id}}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

func (u *User) Deactivate(user_id string) (User, error) {
	var user User
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: user_id}}
	update := bson.D{{Key: "$set", Value: bson.M{"is_active": false}}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&user)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (u *User) FindByID(user_id string) (User, error) {
	collection := conn.Db.Collection("users")
	var user User
	filter := bson.D{{Key: "_id", Value: user_id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	user.Role = string(user.Role)
	if err != nil {
		panic(err)
	}
	return user, nil
}
func (u *User) Findall() ([]User, error) {
	collection := conn.Db.Collection("users")
	var userArr []User
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		userArr = append(userArr, user)
	}
	return userArr, nil
}

func (u *User) FindMgrID(user_id string) (string, error) {
	collection := conn.Db.Collection("users")
	var user User
	filter := bson.D{{Key: "_id", Value: user_id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		panic(err)
	}
	return user.Manager_ID, nil
}

func (u *User) AddNotification(item_id string, user_id string) (bool, error) {
	collection := conn.Db.Collection("users")
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
	collection := conn.Db.Collection("users")
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
	collection := conn.Db.Collection("users")
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$pull", Value: bson.M{"incomplete_actions": item_id}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}

// one of these should be deprecated
func (u *User) MonthlyMileage(user_id string, month int, year int) (User_Monthly_Mileage, error) {
	collection := conn.Db.Collection("mileage_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
	end_month := month + 1
	start_date := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	end_date := time.Date(year, time.Month(end_month), 0, 0, 0, 0, 0, time.UTC)
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: bson.M{"$gte": start_date}}, {Key: "date", Value: bson.M{"$lte": end_date}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var mileage int
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
	collection := conn.Db.Collection("petty_cash_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
	end_month := month + 1
	start_date := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	end_date := time.Date(year, time.Month(end_month), 0, 0, 0, 0, 0, time.UTC)
	filter := bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: bson.M{"$gte": start_date}}, {Key: "date", Value: bson.M{"$lte": end_date}}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	total_amount := 0.0
	var receipts []string
	var requestIDs []string
	for cursor.Next(context.TODO()) {
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requestIDs = append(requestIDs, petty_cash_req.ID)
		receipts = append(receipts, petty_cash_req.Receipts...)
		total_amount += petty_cash_req.Amount
	}
	return User_Monthly_Petty_Cash{
		ID:           user_id,
		Name:         result.Name,
		Month:        time.Month(month),
		Year:         year,
		Total_Amount: total_amount,
		Request_IDS:  requestIDs,
		Receipts:     receipts,
	}, nil
}

func (u *User) AggregateChecks(user_id string, start_date string, end_date string) (User_Agg_Check_Requests, error) {
	collection := conn.Db.Collection("check_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
	var filter bson.D
	layout := "2006-01-02T15:04:05.000Z"
	if start_date != "" && end_date != "" {
		start, err := time.Parse(layout, start_date)
		if err != nil {
			panic(err)
		}
		end, enderr := time.Parse(layout, end_date)
		if enderr != nil {
			panic(err)
		}
		filter = bson.D{{Key: "user_id", Value: user_id}, {Key: "date", Value: bson.M{"$gte": start}}, {Key: "date", Value: bson.M{"$lte": end}}}
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
	var requestIDs []string
	var vendorExists = make(map[Vendor]bool)
	for cursor.Next(context.TODO()) {
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requestIDs = append(requestIDs, check_req.ID)
		if !vendorExists[check_req.Vendor] {
			vendors = append(vendors, check_req.Vendor)
			vendorExists[check_req.Vendor] = true
		}
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
		Request_IDS:  requestIDs,
		Receipts:     receipts,
	}, nil
}

func (u *User) FindMileage(user_id string) (User_Agg_Mileage, error) {
	collection := conn.Db.Collection("mileage_requests")
	var user User
	result, err := user.FindByID(user_id)
	if err != nil {
		panic(err)
	}
	filter := bson.D{{Key: "user_id", Value: user_id}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var mileage int
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
	return User_Agg_Mileage{
		Vehicles:      result.Vehicles,
		Mileage:       mileage,
		Parking:       parking,
		Tolls:         tolls,
		Reimbursement: reimbursement,
		Request_IDS:   request_ids,
	}, nil
}

func (u *User) FindPettyCash(user_id string) (User_Agg_Petty_Cash, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	filter := bson.D{{Key: "user_id", Value: user_id}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	total_amount := 0.0
	var receipts []string
	var requestIDs []string
	for cursor.Next(context.TODO()) {
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requestIDs = append(requestIDs, petty_cash_req.ID)
		receipts = append(receipts, petty_cash_req.Receipts...)
		total_amount += petty_cash_req.Amount
	}
	return User_Agg_Petty_Cash{
		Total_Amount: total_amount,
		Request_IDS:  requestIDs,
		Receipts:     receipts,
	}, nil
}
