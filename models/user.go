package models

import (
	"context"
	"errors"
	conn "financial-api/db"
	auth "financial-api/middleware"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User_Overview struct {
	ID                      string              `json:"id" bson:"_id"`
	Name                    string              `json:"name" bson:"name"`
	Last_Login              time.Time           `json:"last_login" bson:"last_login"`
	Is_Active               bool                `json:"is_active" bson:"is_active"`
	Permissions             []auth.Permission   `json:"permissions" bson:"permissions"`
	Admin                   bool                `json:"admin" bson:"admin"`
	Incomplete_Action_Count int                 `json:"incomplete_action_count" bson:"incomplete_action_count"`
	Mileage_Requests        User_Mileage        `json:"mileage_requests" bson:"mileage_requests"`
	Check_Requests          User_Check_Requests `json:"check_requests" bson:"check_requests"`
	Petty_Cash_Requests     User_Petty_Cash     `json:"petty_cash_requests" bson:"petty_cash_requests"`
}

type User_Detail struct {
	ID                      string              `json:"id" bson:"_id"`
	Name                    string              `json:"name" bson:"name"`
	Permissions             []auth.Permission   `json:"permissions" bson:"permissions"`
	Admin                   bool                `json:"admin" bson:"admin"`
	Vehicles                []Vehicle           `json:"vehicles" bson:"vehicles"`
	Last_Login              time.Time           `json:"last_login" bson:"last_login"`
	Incomplete_Actions      []Action            `json:"incomplete_actions" bson:"incomplete_actions"`
	Incomplete_Action_Count int                 `json:"incomplete_action_count" bson:"incomplete_action_count"`
	Mileage_Requests        User_Mileage        `json:"mileage_requests" bson:"mileage_requests"`
	Check_Requests          User_Check_Requests `json:"check_requests" bson:"check_requests"`
	Petty_Cash_Requests     User_Petty_Cash     `json:"petty_cash_requests" bson:"petty_cash_requests"`
}

type Vehicle struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type User struct {
	ID                 string            `json:"id" bson:"_id"`
	Email              string            `json:"email" bson:"email"`
	Name               string            `json:"name" bson:"name"`
	Last_Login         time.Time         `json:"last_login" bson:"last_login"`
	Vehicles           []Vehicle         `json:"vehicles" bson:"vehicles"`
	InComplete_Actions []Action          `json:"incomplete_actions" bson:"incomplete_actions"`
	Is_Active          bool              `json:"is_active" bson:"is_active"`
	Admin              bool              `json:"admin" bson:"admin"`
	Permissions        []auth.Permission `json:"permissions" bson:"permissions"`
}

// update to create or login based on user existing in database
func (u *User) Login(id string) (string, error) {
	var user User
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		panic("there was an error logging you into your account")
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	var updatedUser User
	update := bson.D{{Key: "$set", Value: bson.M{"last_login": time.Now()}}}
	updateErr := collection.FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&updatedUser)
	if updateErr != nil {
		panic(updateErr)
	}
	token, tokenErr := auth.GenerateToken(updatedUser.ID, updatedUser.Name, updatedUser.Admin, updatedUser.Permissions)
	if tokenErr != nil {
		panic(tokenErr)
	}
	return token, nil
}

var admin_arr = [10]string{"churt", "rgiusti", "dbaker", "bgriffin", "lamanor", "lfuentes", "jward", "cwoods", "abradley", "finance_requests"}
var mgr_arr = [6]string{"churt", "rgiusti", "dbaker", "bgriffin", "lamanor", "lfuentes"}
var supervisor_arr = [2]string{"jward", "cwoods"}
var exec_arr = [1]string{"abradley"}
var fin_arr = [1]string{"finance_requests"}

func isAdmin(user_email string) bool {
	user_name := strings.Split(user_email, "@")[0]
	for _, name := range admin_arr {
		if name == user_name {
			return true
		}
	}
	return false
}

func setPermissions(user_email string) []auth.Permission {
	user_name := strings.Split(user_email, "@")[0]
	for _, name := range exec_arr {
		if name == user_name {
			return []auth.Permission{auth.EXECUTIVE}
		}
	}
	for _, name := range fin_arr {
		if name == user_name {
			return []auth.Permission{auth.FINANCE_TEAM}
		}
	}
	for _, name := range supervisor_arr {
		if name == user_name {
			return []auth.Permission{auth.SUPERVISOR, auth.MANAGER}
		}
	}
	for _, name := range mgr_arr {
		if name == user_name {
			return []auth.Permission{auth.MANAGER}
		}
	}
	return []auth.Permission{auth.EMPLOYEE}
}

func (u *User) Create() (string, error) {
	collection := conn.Db.Collection("users")
	u.Is_Active = true
	u.InComplete_Actions = make([]Action, 0)
	u.Last_Login = time.Now()
	u.Vehicles = make([]Vehicle, 0)
	u.Admin = isAdmin(u.Email)
	u.Permissions = setPermissions(u.Email)
	_, err := collection.InsertOne(context.TODO(), *u)
	if err != nil {
		panic(err)
	}
	token, tokenErr := auth.GenerateToken(u.ID, u.Name, u.Admin, u.Permissions)
	if tokenErr != nil {
		panic(tokenErr)
	}
	return token, nil
}

func (u *User) Exists(id string) (bool, error) {
	var user User
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *User) FindID(email string) (string, error) {
	var user User
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return "", err
	}
	return user.ID, nil
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

func (u *User) AddNotification(incomplete_action Action, user_id string) (bool, error) {
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: user_id}}
	updateManager := collection.FindOneAndUpdate(context.TODO(), filter, bson.D{{Key: "$push", Value: bson.M{"incomplete_actions": incomplete_action}}})
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
