package models

import (
	"context"
	"errors"
	conn "financial-api/db"
	auth "financial-api/middleware"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User_Overview struct {
	ID                      string            `json:"id" bson:"_id"`
	Name                    string            `json:"name" bson:"name"`
	Last_Login              time.Time         `json:"last_login" bson:"last_login"`
	Is_Active               bool              `json:"is_active" bson:"is_active"`
	Permissions             []auth.Permission `json:"permissions" bson:"permissions"`
	Admin                   bool              `json:"admin" bson:"admin"`
	Incomplete_Action_Count int               `json:"incomplete_action_count" bson:"incomplete_action_count"`
	Mileage_Requests        User_Mileage      `json:"mileage_requests" bson:"mileage_requests"`
	Check_Requests          UserAggChecks     `json:"check_requests" bson:"check_requests"`
	Petty_Cash_Requests     UserAggPettyCash  `json:"petty_cash_requests" bson:"petty_cash_requests"`
}

type User_Detail struct {
	ID                      string            `json:"id" bson:"_id"`
	Name                    string            `json:"name" bson:"name"`
	Permissions             []auth.Permission `json:"permissions" bson:"permissions"`
	Admin                   bool              `json:"admin" bson:"admin"`
	Vehicles                []Vehicle         `json:"vehicles" bson:"vehicles"`
	Last_Login              time.Time         `json:"last_login" bson:"last_login"`
	Incomplete_Actions      []Action          `json:"incomplete_actions" bson:"incomplete_actions"`
	Incomplete_Action_Count int               `json:"incomplete_action_count" bson:"incomplete_action_count"`
	Mileage_Requests        User_Mileage      `json:"mileage_requests" bson:"mileage_requests"`
	Check_Requests          UserAggChecks     `json:"check_requests" bson:"check_requests"`
	Petty_Cash_Requests     UserAggPettyCash  `json:"petty_cash_requests" bson:"petty_cash_requests"`
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

func (u *User) BulkInsert() bool {
	users := []interface{}{
		User{
			ID:                 uuid.NewString(),
			Email:              "test_exec@finance.com",
			Name:               "Test Executive",
			Last_Login:         time.Now(),
			Vehicles:           []Vehicle{},
			InComplete_Actions: []Action{},
			Is_Active:          true,
			Admin:              true,
			Permissions: []auth.Permission{
				auth.EXECUTIVE, auth.MANAGER, auth.SUPERVISOR,
			},
		},
		User{
			ID:                 uuid.NewString(),
			Email:              "test_supervisor@finance.com",
			Name:               "Test Supervisor",
			Last_Login:         time.Now(),
			Vehicles:           []Vehicle{},
			InComplete_Actions: []Action{},
			Is_Active:          true,
			Admin:              true,
			Permissions: []auth.Permission{
				auth.MANAGER, auth.SUPERVISOR,
			},
		},
		User{
			ID:                 uuid.NewString(),
			Email:              "test_manager@finance.com",
			Name:               "Test Manager",
			Last_Login:         time.Now(),
			Vehicles:           []Vehicle{},
			InComplete_Actions: []Action{},
			Is_Active:          true,
			Admin:              true,
			Permissions: []auth.Permission{
				auth.MANAGER,
			},
		},
		User{
			ID:                 uuid.NewString(),
			Email:              "test_financec@finance.com",
			Name:               "Test Finance",
			Last_Login:         time.Now(),
			Vehicles:           []Vehicle{},
			InComplete_Actions: []Action{},
			Is_Active:          true,
			Admin:              true,
			Permissions: []auth.Permission{
				auth.FINANCE_TEAM, auth.MANAGER, auth.SUPERVISOR,
			},
		},
	}
	collection := conn.Db.Collection("users")
	result, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		panic(err)
	}
	for _, id := range result.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}
	return true
}

func (u *User) DeleteAll() bool {
	collection := conn.Db.Collection("users")
	record_count, _ := collection.CountDocuments(context.TODO(), bson.D{{}})
	cleared, _ := collection.DeleteMany(context.TODO(), bson.D{{}})
	return cleared.DeletedCount == record_count
}

func (u *User) Login(email string) (string, error) {
	var user User
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
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

func (u *User) SetPermissions(permissions string) []auth.Permission {
	switch permissions {
	case "EXECUTIVE":
		return []auth.Permission{auth.EXECUTIVE, auth.SUPERVISOR, auth.MANAGER}
	case "FINANCE_TEAM":
		return []auth.Permission{auth.FINANCE_TEAM, auth.SUPERVISOR, auth.MANAGER}
	case "SUPERVISOR":
		return []auth.Permission{auth.SUPERVISOR, auth.MANAGER}
	case "MANAGER":
		return []auth.Permission{auth.MANAGER}
	default:
		return []auth.Permission{auth.EMPLOYEE}
	}
}

func (u *User) Create() (string, error) {
	collection := conn.Db.Collection("users")
	u.Is_Active = true
	u.InComplete_Actions = make([]Action, 0)
	u.Last_Login = time.Now()
	u.Vehicles = make([]Vehicle, 0)
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

func (u *User) ClearNotificationByID(notification_id string, user_id string) (bool, error) {
	collection := conn.Db.Collection("users")
	filter := bson.D{{Key: "_id", Value: user_id}}
	updateManager := collection.FindOneAndUpdate(context.TODO(), filter, bson.D{{Key: "$pull", Value: bson.M{"incomplete_actions": bson.D{{Key: "_id", Value: notification_id}}}}})
	if updateManager == nil {
		return false, errors.New("error clearing notification second lvl")
	}
	return true, nil
}
