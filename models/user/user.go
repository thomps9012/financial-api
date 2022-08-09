package user

import (
	"context"
	conn "financial-api/m/db"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type Vehicle struct {
	ID          string `json:"id" bson:"_id"`
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

type Manager struct {
	ID        string
	Employees []string
}

func findManagerID(email string) string {
	var manager_id string
	managers := []Manager{
		{"test1", []string{"id1", "id2", "id3", "id4", "id5"}},
		{"test2", []string{"id6", "id7", "id8", "id9", "id10"}},
		{"test3", []string{"id11", "id12", "id13", "id14", "id15"}},
	}
	for i := range managers {
		var employeesArr = managers[i].Employees
		for s := range employeesArr {
			if employeesArr[s] == email {
				manager_id = managers[i].ID
			}
		}
	}
	return manager_id
}

func (u *User) Create(email string) (string, error) {
	collection := conn.DB.Collection("users")
	filter := bson.D{{Key: "email", Value: email}}
	var user User
	findErr := collection.FindOne(context.TODO(), filter).Decode(&user)
	if findErr == nil {
		return "", fmt.Errorf("account already created")
	}
	u.ID = uuid.NewString()
	u.Last_Login = time.Now()
	u.Is_Active = true
	manager_id := findManagerID(email)
	if manager_id != "" {
		u.Manager_ID = manager_id
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
	filter := bson.D{{Key: "_id", Value: user_id}}
	vehicle := &Vehicle{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
	result, err := collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$push", Value: bson.M{"vehicles": vehicle}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return "", err
	}
	return vehicle.ID, nil
}

func (u *User) Deactivate(user_id string) (bool, error) {
	collection := conn.DB.Collection("users")
	filter := bson.D{{Key: "_id", Value: user_id}}
	result, err := collection.UpdateOne(context.TODO(), filter, bson.D{{Key: "$set", Value: bson.M{"is_active": false}}})
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount == 0 {
		return false, err
	}
	return true, nil
}
