package user

import (
	"context"
	"errors"
	conn "financial-api/db"
	auth "financial-api/middleware"
	"math"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// possible expanison after org chart release
type Role string

const (
	EMPLOYEE  Role = "EMPLOYEE"
	MANAGER   Role = "MANAGER"
	CHIEF     Role = "CHIEF"
	EXECUTIVE Role = "EXECUTIVE"
	FINANCE   Role = "FINANCE"
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
	case "CHIEF":
		roleParse = CHIEF
	default:
		roleParse = EMPLOYEE
	}
	return roleParse
}

type User_Overview struct {
	ID                      string              `json:"id" bson:"_id"`
	Manager_ID              string              `json:"manager_id" bson:"manager_id"`
	Name                    string              `json:"name" bson:"name"`
	Last_Login              time.Time           `json:"last_login" bson:"last_login"`
	Is_Active               bool                `json:"is_active" bson:"is_active"`
	Role                    string              `json:"role" bson:"role"`
	Incomplete_Action_Count int                 `json:"incomplete_action_count" bson:"incomplete_action_count"`
	Mileage_Requests        User_Mileage        `json:"mileage_requests" bson:"mileage_requests"`
	Check_Requests          User_Check_Requests `json:"check_requests" bson:"check_requests"`
	Petty_Cash_Requests     User_Petty_Cash     `json:"petty_cash_requests" bson:"petty_cash_requests"`
}

type User_Detail struct {
	ID                      string              `json:"id" bson:"_id"`
	Manager_ID              string              `json:"manager_id" bson:"manager_id"`
	Name                    string              `json:"name" bson:"name"`
	Role                    string              `json:"role" bson:"role"`
	Vehicles                []Vehicle           `json:"vehicles" bson:"vehicles"`
	Last_Login              time.Time           `json:"last_login" bson:"last_login"`
	Incomplete_Actions      []Action            `json:"incomplete_actions" bson:"incomplete_actions"`
	Incomplete_Action_Count int                 `json:"incomplete_action_count" bson:"incomplete_action_count"`
	Mileage_Requests        User_Mileage        `json:"mileage_requests" bson:"mileage_requests"`
	Check_Requests          User_Check_Requests `json:"check_requests" bson:"check_requests"`
	Petty_Cash_Requests     User_Petty_Cash     `json:"petty_cash_requests" bson:"petty_cash_requests"`
}

type User_Agg_Mileage struct {
	Parking       float64           `json:"parking" bson:"parking"`
	Tolls         float64           `json:"tolls" bson:"tolls"`
	Mileage       int               `json:"mileage" bson:"mileage"`
	User          User              `json:"user" bson:"user"`
	Reimbursement float64           `json:"reimbursement" bson:"reimbursement"`
	Requests      []Mileage_Request `json:"requests" bson:"requests"`
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
	Current_Status string    `json:"current_status" bson:"current_status"`
	Is_Active      bool      `json:"is_active" bson:"is_active"`
}

type Vehicle struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type Mileage_Request struct {
	ID                string    `json:"id" bson:"_id"`
	Grant_ID          string    `json:"grant_id" bson:"grant_id"`
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
	Current_Status    string    `json:"current_status" bson:"current_status"`
	Is_Active         bool      `json:"is_active" bson:"is_active"`
}

type User_Action_Info struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Role string `json:"role" bson:"role"`
}

type User struct {
	ID                 string    `json:"id" bson:"_id"`
	Email              string    `json:"email" bson:"email"`
	Name               string    `json:"name" bson:"name"`
	Last_Login         time.Time `json:"last_login" bson:"last_login"`
	Vehicles           []Vehicle `json:"vehicles" bson:"vehicles"`
	InComplete_Actions []Action  `json:"incomplete_actions" bson:"incomplete_actions"`
	Manager_ID         string    `json:"manager_id" bson:"manager_id"`
	Manager_Email	   string	 `json:"manager_email" bson:"manager_email"`
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
	CHIEF_APPROVED        Status = "CHIEF_APPROVED"
	FINANCE_APPROVED      Status = "FINANACE_APPROVED"
	ORGANIZATION_APPROVED Status = "ORG_APPROVED"
	REJECTED              Status = "REJECTED"
	ARCHIVED              Status = "ARCHIVED"
)

type Action struct {
	ID           string `json:"id" bson:"_id"`
	User         User_Action_Info
	Request_Type string    `json:"request_type" bson:"request_type"`
	Request_ID   string    `json:"request_id" bson:"request_id"`
	Status       string    `json:"status" bson:"status"`
	Created_At   time.Time `json:"created_at" bson:"created_at"`
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
	Current_Status string     `json:"current_status" bson:"current_status"`
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
	Grant_IDS     []string   `json:"grant_ids" bson:"grant_ids"`
	Requests   []Mileage_Request   `json:"requests" bson:"requests"`
}
type User_Mileage struct {
	Vehicles      []Vehicle         `json:"vehicles" bson:"vehicles"`
	Mileage       int               `json:"mileage" bson:"mileage"`
	Tolls         float64           `json:"tolls" bson:"tolls"`
	Parking       float64           `json:"parking" bson:"parking"`
	Reimbursement float64           `json:"reimbursement" bson:"reimbursement"`
	Requests      []Mileage_Request `json:"requests" bson:"requests"`
	Last_Request  Mileage_Request   `json:"last_request" bson:"last_request"`
}
type User_Petty_Cash struct {
	User_ID      string               `json:"user_id" bson:"user_id"`
	User         User                 `json:"user" bson:"user"`
	Total_Amount float64              `json:"total_amount" bson:"total_amount"`
	Requests     []Petty_Cash_Request `json:"requests" bson:"requests"`
	Last_Request Petty_Cash_Request   `json:"last_request" bson:"last_request"`
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

type User_Check_Requests struct {
	ID           string          `json:"id" bson:"_id"`
	Name         string          `json:"name" bson:"name"`
	Start_Date   string          `json:"start_date" bson:"start_date"`
	End_Date     string          `json:"end_date" bson:"end_date"`
	Total_Amount float64         `json:"total_amount" bson:"total_amount"`
	Vendors      []Vendor        `json:"vendors" bson:"vendors"`
	Requests     []Check_Request `json:"requests" bson:"requests"`
}

// can optimize this function with a switch to search certain arrays based on the user's role
// double check two lower functions for manager id and manager role
// func setManagerID(email string, employee_role string) string {
// 	// var manager_id string
// 	// managers := []Manager{
// 	// 	{"101614399314441368253", []string{"emp1@norainc.org", "emp2@norainc.org", "emp3@norainc.org"}},
// 	// 	{"123", []string{"emp4@norainc.org", "emp5@norainc.org", "emp6@norainc.org"}},
// 	// }
// 	// chiefs := []Manager{
// 	// 	{"456", []string{"manager1@norainc.org"}},
// 	// 	{"116601745736489768774", []string{"sthompson@norainc.org"}},
// 	// }
// 	// var executive = "101112"
// 	// var finance = "131415"
// 	// switch employee_role {
// 	// case "EMPLOYEE":
// 	// 	for i := range managers {
// 	// 		var employeesArr = managers[i].Employees
// 	// 		for s := range employeesArr {
// 	// 			if employeesArr[s] == email {
// 	// 				manager_id = managers[i].ID
// 	// 			}

// 	// 		}
// 	// 		if manager_id == "" {
// 	// 			manager_id = executive
// 	// 		}
// 	// 	}
// 	// case "MANAGER":
// 	// 	for i := range chiefs {
// 	// 		var employeesArr = chiefs[i].Employees
// 	// 		for s := range employeesArr {
// 	// 			if employeesArr[s] == email {
// 	// 				manager_id = chiefs[i].ID
// 	// 			}
// 	// 		}
// 	// 		if manager_id == "" {
// 	// 			manager_id = executive
// 	// 		}
// 	// 	}
// 	// case "CHIEF":
// 	// 	manager_id = executive
// 	// case "EXECUTIVE":
// 	// 	manager_id = finance
// 	// }
// 	return manager_id
// }
// func setRole(email string) string {
// 	var employee_role string
// 	employees := []string{"emp1@norainc.org", "emp2@norainc.org"}
// 	for i := range employees {
// 		if employees[i] == email {
// 			employee_role = "EMPLOYEE"
// 		}
// 	}
// 	managers := []string{"manager1@norainc.org", "sthompson@norainc.org"}
// 	for i := range managers {
// 		if managers[i] == email {
// 			employee_role = "MANAGER"
// 		}

// 	}
// 	chiefs := []string{"coo@norainc.org", "cfo@norainc.org", "cmo@norainc.org"}
// 	for i := range chiefs {
// 		if chiefs[i] == email {
// 			employee_role = "CHIEF"
// 		}

// 	}
// 	finance_team := []string{"finance1@norainc.org", "finance2@norainc.org"}
// 	for i := range finance_team {
// 		if finance_team[i] == email {
// 			employee_role = "FINANCE"
// 		}

// 	}
// 	executives := []string{"ceo@norainc.org"}
// 	for i := range executives {
// 		if executives[i] == email {
// 			employee_role = "EXECUTIVE"
// 		}
// 	}
// 	return employee_role
// }
// func (u *User) Create(id string, name string, email string) (User, error) {
// 	collection := conn.Db.Collection("users")
// 	println("user info in create function: %s", id, name, email)
// 	u.ID = id
// 	u.Name = name
// 	u.Email = email
// 	u.Last_Login = time.Now()
// 	u.Is_Active = true
// 	u.Email = email
// 	u.Vehicles = []Vehicle{}
// 	u.InComplete_Actions = []Action{}
// 	role := setRole(email)
// 	manager_id := setManagerID(email, role)
// 	u.Manager_ID = manager_id
// 	u.Role = role
// 	_, err := collection.InsertOne(context.TODO(), u)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return *u, nil
// }
// func (u *User) Login(id string, name string, email string) (User, error) {
// 	var user User
// 	collection := conn.Db.Collection("users")
// 	filter := bson.D{{Key: "_id", Value: id}}
// 	update := bson.D{{Key: "$set", Value: bson.M{"last_login": time.Now()}}}
// 	upsert := true
// 	after := options.After
// 	opt := options.FindOneAndUpdateOptions{
// 		ReturnDocument: &after,
// 		Upsert:         &upsert,
// 	}
// turn this into async await
// 	err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&user)
// 	if err == mongo.ErrNoDocuments {
// 		println("user info in login err: %s", id, name, email)
// 		newUser, createErr := user.Create(id, name, email)
// 		if createErr != nil {
// 			panic(createErr)
// 		}
// 		return newUser, nil
// 	}
// 	return user, nil
// }
// automatically search through database for user info
func (u *User) Login(id string, name string, email string) (User, error) {
	var user User
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	update := bson.D{{Key: "$set", Value: bson.M{"last_login": time.Now()}}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	upsert := true
	after := options.After
	opt := FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert: &upsert,
	}
	if user.Google_ID == "" && user.Manager_ID != "" {
		var updatedUser User
		update := bson.D{{Key: "$set", Value: bson.M{"last_login": time.Now()}}, {Key: "$set", Value: bson.M{"google_id": id} }}
		err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&updatedUser)
		if err != nil {
			panic(err)
		}
		return updatedUser
	}
	if user.Manager_ID == "" && user.Google_ID == "" {
		var manager User
		var updatedUser User
		mgrFilter := bson.D{{Key: "email", Value: user.Manager_Email}}
		err := collection.FindOne(context.TODO(), mgrFilter).Decode(&manager)
		update := bson.D{{Key: "$set", Value: bson.M{"last_login": time.Now()}}, {Key: "$set", Value: bson.M{"manager_id": manager.ID}}}
		err := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&updatedUser)
		if err != nil {
			panic(err)
		}
		return updatedUser
	}
	var updatedUser User
	updateErr := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&user)
	if updateErr != nil {
		panic(updateErr)
	}
	return updatedUser, nil
	
}

func (u *User) Create() (User, error){
	collection := conn.Db.Collection("users")
	newUser, err := collection.InsertOne(context.TODO(), *u)
	if err != nil {
		panic(err)
	}
	return newUser, nil
}

func (u *User) Exists(email string) (bool, error) {
	var user User
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return false, err
	}
	return true, nil
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

func (u *User) FindContextID(ctx context.Context) (User, error) {
	user_id := auth.ForID(ctx)
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

func (u *User) CheckAdmin(ctx context.Context) bool {
	user_role := auth.ForRole(ctx)
	return user_role != "EMPLOYEE"
}

func (u *User) LoggedIn(ctx context.Context) bool {
	user_id := auth.ForID(ctx)
	return user_id != ""
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

func (u *User) AddNotification(item Action, user_id string) (bool, error) {
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: user_id}}
	updateManager := collection.FindOneAndUpdate(context.TODO(), filter, bson.D{{Key: "$push", Value: bson.M{"incomplete_actions": item}}})
	if updateManager == nil {
		return false, errors.New("error notifying manager second lvl")
	}
	return true, nil
}

func (u *User) ClearNotifications(user_id string) (bool, error) {
	collection := conn.Db.Collection("users")
	result, err := collection.UpdateByID(context.TODO(), user_id, bson.D{{Key: "$set", Value: bson.M{"incomplete_actions": []Action{}}}})
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
	filter := bson.D{{Key: "_id", Value: user_id}}
	updateManager := collection.FindOneAndUpdate(context.TODO(), filter, bson.D{{Key: "$pull", Value: bson.M{"incomplete_actions": bson.D{{Key: "request_id", Value: item_id}}}}})
	if updateManager == nil {
		return false, errors.New("error clearing notification second lvl")
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
	var requests []Mileage_Request
	var grant_ids []string
	for cursor.Next(context.TODO()) {
		var mileage_req Mileage_Request
		decode_err := cursor.Decode(&mileage_req)
		if decode_err != nil {
			panic(decode_err)
		}
		grant_ids = append(grant_ids, mileage_req.Grant_ID)
		requests = append(requests, mileage_req)
		mileage += mileage_req.Trip_Mileage
		tolls += mileage_req.Tolls
		parking += mileage_req.Parking
		reimbursement += mileage_req.Reimbursement
	}
	return User_Monthly_Mileage{
		ID:            user_id,
		Grant_IDS:     grant_ids,
		Name:          result.Name,
		Vehicles:      result.Vehicles,
		Month:         time.Month(month),
		Year:          year,
		Mileage:       mileage,
		Parking:       parking,
		Tolls:         tolls,
		Reimbursement: reimbursement,
		Requests:   requests,
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

func (u *User) AggregateMileage(user_id string, start_date string, end_date string) (User_Agg_Mileage, error) {
	collection := conn.Db.Collection("mileage_requests")
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
	var requests []Mileage_Request
	total_reimbursement := 0.00
	total_tolls := 0.00
	total_parking := 0.00
	total_mileage := 0

	for cursor.Next(context.TODO()) {
		var mileage_req Mileage_Request
		decode_err := cursor.Decode(&mileage_req)
		if decode_err != nil {
			panic(decode_err)
		}
		total_reimbursement += mileage_req.Reimbursement
		total_tolls += mileage_req.Tolls
		total_parking += mileage_req.Parking
		total_mileage += mileage_req.Trip_Mileage
		requests = append(requests, mileage_req)
	}
	return User_Agg_Mileage{
		User:          result,
		Mileage:       total_mileage,
		Tolls:         total_tolls,
		Parking:       total_parking,
		Reimbursement: total_reimbursement,
		Requests:      requests,
	}, nil
}

func (u *User) AggregateChecks(user_id string, start_date string, end_date string) (User_Check_Requests, error) {
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
	var requests []Check_Request
	var vendorExists = make(map[Vendor]bool)
	for cursor.Next(context.TODO()) {
		var check_req Check_Request
		decode_err := cursor.Decode(&check_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requests = append(requests, check_req)
		if !vendorExists[check_req.Vendor] {
			vendors = append(vendors, check_req.Vendor)
			vendorExists[check_req.Vendor] = true
		}
		purchases = append(purchases, check_req.Purchases...)
		receipts = append(receipts, check_req.Receipts...)
		total_amount += check_req.Order_Total
	}
	return User_Check_Requests{
		ID:           user_id,
		Name:         result.Name,
		Start_Date:   start_date,
		End_Date:     end_date,
		Total_Amount: total_amount,
		Vendors:      vendors,
		Requests:     requests,
	}, nil
}

func (u *User) FindMileage(user_id string) (User_Mileage, error) {
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
	var requests []Mileage_Request
	last_request_date := time.Date(2000, time.April,
		34, 25, 72, 01, 0, time.UTC)
	var last_request Mileage_Request
	for cursor.Next(context.TODO()) {
		var mileage_req Mileage_Request
		decode_err := cursor.Decode(&mileage_req)
		if decode_err != nil {
			panic(decode_err)
		}
		if mileage_req.Date.After(last_request_date) {
			last_request = mileage_req
		}
		requests = append(requests, mileage_req)
		mileage += mileage_req.Trip_Mileage
		tolls += mileage_req.Tolls
		parking += mileage_req.Parking
		reimbursement += mileage_req.Reimbursement
	}
	return User_Mileage{
		Vehicles:      result.Vehicles,
		Mileage:       mileage,
		Parking:       parking,
		Tolls:         tolls,
		Reimbursement: reimbursement,
		Requests:      requests,
		Last_Request:  last_request,
	}, nil
}

func (u *User) FindPettyCash(user_id string) (User_Petty_Cash, error) {
	collection := conn.Db.Collection("petty_cash_requests")
	filter := bson.D{{Key: "user_id", Value: user_id}}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	total_amount := 0.0
	last_request_date := time.Date(2000, time.April,
		34, 25, 72, 01, 0, time.UTC)
	var last_request Petty_Cash_Request
	var requests []Petty_Cash_Request
	for cursor.Next(context.TODO()) {
		var petty_cash_req Petty_Cash_Request
		decode_err := cursor.Decode(&petty_cash_req)
		if decode_err != nil {
			panic(decode_err)
		}
		requests = append(requests, petty_cash_req)
		if petty_cash_req.Date.After(last_request_date) {
			last_request = petty_cash_req
		}
		total_amount += math.Round(petty_cash_req.Amount*100) / 100
		total_amount = math.Round(total_amount*100) / 100
	}
	return User_Petty_Cash{
		Total_Amount: total_amount,
		Requests:     requests,
		Last_Request: last_request,
		User_ID:      user_id,
	}, nil
}
