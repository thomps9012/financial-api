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

type Vehicle struct {
	ID          string `json:"id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
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
