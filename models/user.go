package models

import (
	"context"
	database "financial-api/db"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserLogin struct {
	ID    string `json:"id" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required"`
}

type LoginRes struct {
	UserID      string   `json:"user_id"`
	Token       string   `json:"token"`
	Admin       bool     `json:"admin" bson:"admin"`
	Permissions []string `json:"permissions" bson:"permissions"`
}

type User struct {
	ID          string    `json:"id" bson:"_id"`
	Email       string    `json:"email" bson:"email"`
	Name        string    `json:"name" bson:"name"`
	Last_Login  time.Time `json:"last_login" bson:"last_login"`
	Vehicles    []Vehicle `json:"vehicles" bson:"vehicles"`
	Is_Active   bool      `json:"is_active" bson:"is_active"`
	Admin       bool      `json:"admin" bson:"admin"`
	Permissions []string  `json:"permissions" bson:"permissions"`
}

type UserNameInfo struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	Is_Active bool   `json:"is_active" bson:"is_active"`
}

type PublicInfo struct {
	ID                  string                   `json:"id" bson:"_id"`
	Email               string                   `json:"email" bson:"email"`
	Name                string                   `json:"name" bson:"name"`
	Last_Login          time.Time                `json:"last_login" bson:"last_login"`
	Vehicles            []Vehicle                `json:"vehicles" bson:"vehicles"`
	InComplete_Actions  []IncompleteAction       `json:"incomplete_actions" bson:"incomplete_actions"`
	Is_Active           bool                     `json:"is_active" bson:"is_active"`
	Admin               bool                     `json:"admin" bson:"admin"`
	Permissions         []string                 `json:"permissions" bson:"permissions"`
	Mileage_Requests    []Mileage_Overview       `json:"mileage_requests" bson:"mileage_requests"`
	Petty_Cash_Requests []Petty_Cash_Overview    `json:"petty_cash_requests" bson:"petty_cash_requests"`
	Check_Requests      []Check_Request_Overview `json:"check_requests" bson:"check_requests"`
}

type Vehicle struct {
	ID          string `json:"id" bson:"_id" validate:"required"`
	Name        string `json:"name" bson:"name" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required"`
}

type VehicleInput struct {
	Name        string `json:"name" bson:"name" validate:"required"`
	Description string `json:"description" bson:"description" validate:"required"`
}

var admin_arr = [10]string{"churt", "rgiusti", "dbaker", "bgriffin", "lamanor", "lfuentes", "jward", "cwoods", "abradley", "finance_requests"}
var mgr_arr = [6]string{"churt", "rgiusti", "dbaker", "bgriffin", "lamanor", "lfuentes"}
var supervisor_arr = [2]string{"jward", "cwoods"}
var exec_arr = [1]string{"abradley"}
var fin_arr = [1]string{"finance_requests"}

func (ul *UserLogin) Exists() (bool, error) {
	users, err := database.Use("users")
	if err != nil {
		return false, err
	}
	filter := bson.D{{Key: "_id", Value: ul.ID}}
	count, err := users.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	database.CloseDB()
	return count > 0, nil
}
func isAdmin(user_email string) bool {
	user_name := strings.Split(user_email, "@")[0]
	for _, name := range admin_arr {
		if name == user_name {
			return true
		}
	}
	return false
}
func setPermissions(user_email string) []string {
	user_name := strings.Split(user_email, "@")[0]
	for _, name := range exec_arr {
		if name == user_name {
			return []string{"EXECUTIVE"}
		}
	}
	for _, name := range fin_arr {
		if name == user_name {
			return []string{"FINANCE_TEAM"}
		}
	}
	for _, name := range supervisor_arr {
		if name == user_name {
			return []string{"SUPERVISOR", "MANAGER"}
		}
	}
	for _, name := range mgr_arr {
		if name == user_name {
			return []string{"MANAGER"}
		}
	}
	return []string{"EMPLOYEE"}
}
func (u *User) Create(user UserLogin) (LoginRes, error) {
	users, err := database.Use("users")
	if err != nil {
		return LoginRes{}, err
	}
	u.ID = user.ID
	u.Email = user.Email
	u.Name = user.Name
	u.Is_Active = true
	u.Last_Login = time.Now()
	u.Vehicles = make([]Vehicle, 0)
	u.Admin = isAdmin(u.Email)
	u.Permissions = setPermissions(u.Email)
	_, err = users.InsertOne(context.TODO(), *u)
	if err != nil {
		return LoginRes{}, err
	}
	token, err := GenerateToken(u.ID, u.Name)
	if err != nil {
		return LoginRes{}, err
	}
	database.CloseDB()
	return LoginRes{
		UserID:      u.ID,
		Token:       token,
		Admin:       u.Admin,
		Permissions: u.Permissions,
	}, nil
}
func (u *User) Login(user UserLogin) (LoginRes, error) {
	users, err := database.Use("users")
	if err != nil {
		return LoginRes{}, err
	}
	filter := bson.D{{Key: "_id", Value: user.ID}}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	update := bson.D{{Key: "$set", Value: bson.M{"last_login": time.Now()}}}
	err = users.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&u)
	if err != nil {
		return LoginRes{}, err
	}
	token, err := GenerateToken(u.ID, u.Name)
	if err != nil {
		return LoginRes{}, err
	}
	database.CloseDB()
	return LoginRes{
		UserID:      u.ID,
		Token:       token,
		Admin:       u.Admin,
		Permissions: u.Permissions,
	}, nil
}
func GetPublicInfo(user_id string) (PublicInfo, error) {
	users, err := database.Use("users")
	if err != nil {
		return PublicInfo{}, err
	}
	user_info := make([]PublicInfo, 0)
	filter := bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: user_id}}}}
	mileage_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "mileage_requests"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "user_id"}, {Key: "as", Value: "mileage_requests"}}}}
	check_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "check_requests"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "user_id"}, {Key: "as", Value: "check_requests"}}}}
	petty_cash_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "petty_cash_requests"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "user_id"}, {Key: "as", Value: "petty_cash_requests"}}}}
	incomplete_action_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "incomplete_actions"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "user_id"}, {Key: "as", Value: "incomplete_actions"}}}}
	pipeline := mongo.Pipeline{filter, mileage_stage, check_stage, petty_cash_stage, incomplete_action_stage}
	cursor, err := users.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return PublicInfo{}, err
	}
	err = cursor.All(context.TODO(), &user_info)
	if err != nil {
		return PublicInfo{}, err
	}
	database.CloseDB()
	return user_info[0], nil
}
func FindUserName(user_id string) (string, error) {
	var user User
	users, err := database.Use("users")
	if err != nil {
		return "", err
	}
	filter := bson.D{{Key: "_id", Value: user_id}}
	projection := bson.D{{Key: "name", Value: 1}}
	opts := options.FindOne().SetProjection(projection)
	err = users.FindOne(context.TODO(), filter, opts).Decode(&user)
	if err != nil {
		return "", err
	}
	database.CloseDB()
	return user.Name, nil
}
func (u *User) AddVehicle(name string, description string) (Vehicle, error) {
	users, err := database.Use("users")
	if err != nil {
		return Vehicle{}, err
	}
	new_vehicle := new(Vehicle)
	new_vehicle.ID = uuid.NewString()
	new_vehicle.Name = name
	new_vehicle.Description = description
	filter := bson.D{{Key: "_id", Value: u.ID}}
	projection := bson.D{{Key: "vehicles", Value: 1}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "vehicles", Value: new_vehicle}}}}
	opts := options.FindOneAndUpdate().SetProjection(projection).SetReturnDocument(options.After)
	err = users.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&u)
	if err != nil {
		return Vehicle{}, err
	}
	database.CloseDB()
	return *new_vehicle, nil
}
func (u *User) EditVehicle(new_vehicle Vehicle) (Vehicle, error) {
	users, err := database.Use("users")
	if err != nil {
		return Vehicle{}, err
	}
	filter := bson.D{{Key: "_id", Value: u.ID}, {Key: "vehicles._id", Value: new_vehicle.ID}}
	projection := bson.D{{Key: "vehicles", Value: 1}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "vehicles.$.name", Value: new_vehicle.Name}, {Key: "vehicles.$.description", Value: new_vehicle.Description}}}}
	opts := options.FindOneAndUpdate().SetProjection(projection).SetReturnDocument(options.After)
	err = users.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&u)
	if err != nil {
		return Vehicle{}, err
	}
	database.CloseDB()
	return new_vehicle, nil
}
func (u *User) RemoveVehicle(vehicle_id string) (Vehicle, error) {
	users, err := database.Use("users")
	if err != nil {
		return Vehicle{}, err
	}
	filter := bson.D{{Key: "_id", Value: u.ID}}
	projection := bson.D{{Key: "vehicles", Value: 1}}
	update := bson.D{{Key: "$pull", Value: bson.D{{Key: "vehicles", Value: bson.D{{Key: "_id", Value: vehicle_id}}}}}}
	opts := options.FindOneAndUpdate().SetProjection(projection).SetReturnDocument(options.Before)
	err = users.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&u)
	if err != nil {
		return Vehicle{}, err
	}
	old_vehicle := new(Vehicle)
	for _, vehicle := range u.Vehicles {
		if vehicle.ID == vehicle_id {
			old_vehicle.ID = vehicle_id
			old_vehicle.Name = vehicle.Name
			old_vehicle.Description = vehicle.Description
		}
	}
	database.CloseDB()
	return *old_vehicle, nil
}
func FindAllUsers() ([]UserNameInfo, error) {
	users, err := database.Use("users")
	if err != nil {
		return []UserNameInfo{}, err
	}
	data := make([]UserNameInfo, 0)
	cursor, err := users.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return []UserNameInfo{}, err
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		return []UserNameInfo{}, err
	}
	database.CloseDB()
	return data, nil
}
func GetAllUsersPublicInfo() ([]PublicInfo, error) {
	users, err := database.Use("users")
	if err != nil {
		return []PublicInfo{}, err
	}
	mileage_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "mileage_requests"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "user_id"}, {Key: "as", Value: "mileage_requests"}}}}
	check_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "check_requests"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "user_id"}, {Key: "as", Value: "check_requests"}}}}
	petty_cash_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "petty_cash_requests"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "user_id"}, {Key: "as", Value: "petty_cash_requests"}}}}
	incomplete_action_stage := bson.D{{Key: "$lookup", Value: bson.D{{Key: "from", Value: "incomplete_actions"}, {Key: "localField", Value: "_id"}, {Key: "foreignField", Value: "user_id"}, {Key: "as", Value: "incomplete_actions"}}}}
	pipeline := mongo.Pipeline{mileage_stage, check_stage, petty_cash_stage, incomplete_action_stage}
	cursor, err := users.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return []PublicInfo{}, err
	}
	user_info := make([]PublicInfo, 0)
	err = cursor.All(context.TODO(), &user_info)
	if err != nil {
		return []PublicInfo{}, err
	}
	database.CloseDB()
	return user_info, nil
}
func (u *User) Deactivate() (PublicInfo, error) {
	users, err := database.Use("users")
	if err != nil {
		return PublicInfo{}, err
	}
	user_info := new(PublicInfo)
	filter := bson.D{{Key: "_id", Value: u.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "is_active", Value: false}}}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = users.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&user_info)
	if err != nil {
		return PublicInfo{}, err
	}
	database.CloseDB()
	return *user_info, nil
}
func GetUserIncompleteActions(user_id string) ([]IncompleteAction, error) {
	incomplete_actions, err := database.Use("incomplete_actions")
	if err != nil {
		return []IncompleteAction{{}}, err
	}
	actions := make([]IncompleteAction, 0)
	filter := bson.D{{Key: "user_id", Value: user_id}}
	cursor, err := incomplete_actions.Find(context.TODO(), filter)
	if err != nil {
		return []IncompleteAction{{}}, err
	}
	err = cursor.All(context.TODO(), &actions)
	if err != nil {
		return []IncompleteAction{{}}, err
	}
	database.CloseDB()
	return actions, nil
}
