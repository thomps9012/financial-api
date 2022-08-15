package seeds

import (
	"context"
	. "financial-api/Db"
	r "financial-api/models/requests"
	u "financial-api/models/user"
	"time"

	"github.com/google/uuid"
)

func UserBulkInsert() ([]string, error) {
	users := []interface{}{
		u.User{
			ID:         uuid.NewString(),
			Email:      "emp1@norainc.org",
			Name:       "Test One",
			Role:       "MANAGER",
			Last_Login: time.Now(),
			Vehicles:   []u.Vehicle{},
			Is_Active:  true,
		},
		u.User{
			ID:         uuid.NewString(),
			Email:      "emp2@norainc.org",
			Name:       "Test Two",
			Role:       "MANAGER",
			Last_Login: time.Now(),
			Vehicles:   []u.Vehicle{},
			Is_Active:  true,
		},
		u.User{
			ID:         uuid.NewString(),
			Email:      "emp3@norainc.org",
			Name:       "Test Three",
			Role:       "MANAGER",
			Last_Login: time.Now(),
			Vehicles:   []u.Vehicle{},
			Is_Active:  true,
		},
		u.User{
			ID:         uuid.NewString(),
			Email:      "exec1@norainc.org",
			Name:       "Executive One",
			Role:       "EXECUTIVE",
			Last_Login: time.Now(),
			Vehicles:   []u.Vehicle{},
			Is_Active:  true,
		},
		u.User{
			ID:         uuid.NewString(),
			Email:      "exec2@norainc.org",
			Name:       "Executive Two",
			Role:       "EXECUTIVE",
			Last_Login: time.Now(),
			Vehicles:   []u.Vehicle{},
			Is_Active:  true,
		},
		u.User{
			ID:         uuid.NewString(),
			Email:      "finance1@norainc.org",
			Name:       "Finance One",
			Role:       "FINANCE",
			Last_Login: time.Now(),
			Vehicles:   []u.Vehicle{},
			Is_Active:  true,
		},
		u.User{
			ID:         uuid.NewString(),
			Email:      "finance2@norainc.org",
			Name:       "Finance Two",
			Role:       "FINANCE",
			Last_Login: time.Now(),
			Vehicles:   []u.Vehicle{},
			Is_Active:  true,
		},
	}
	collection := Db.Collection("users")
	result, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		panic(err)
	}
	var userID_array []string
	for _, id := range result.InsertedIDs {
		println(id)
		userID_array = append(userID_array, id.(string))
	}
	return userID_array, nil
}
func CheckBulkInsert(userID_array []string) ([]string, error) {
	collection := Db.Collection("check_requests")
	var InsertedIDs []string
	for _, user_id := range userID_array {
		check_req := r.Check_Request{
			ID:          uuid.NewString(),
			Date:        time.Now(),
			Description: "an amazing check request",
			User_ID:     user_id,
			Vendor: r.Vendor{
				Name: "an awesome vendor",
				Address: r.Address{
					Website:  "www.vendor.com",
					Street:   "123 Street",
					City:     "City",
					State:    "State",
					Zip_Code: 55555,
				},
			},
			Grant_ID: "2020-JY-FX-0014",
			Purchases: []r.Purchase{
				{
					Grant_Line_Item: "line item 1",
					Description:     "an awesome purchase",
					Amount:          14.12,
				}},
			Receipts:       []string{},
			Order_Total:    14.12,
			Credit_Card:    "N/A",
			Created_At:     time.Now(),
			Action_History: []r.Action{},
			Current_Status: "PENDING",
			Is_Active:      true,
		}
		result, err := collection.InsertOne(context.TODO(), check_req)
		if err != nil {
			panic(err)
		}
		InsertedIDs = append(InsertedIDs, result.InsertedID.(string))
	}
	return InsertedIDs, nil
}
func PettyBulkInsert(userID_array []string) ([]string, error) {
	collection := Db.Collection("petty_cash_requests")
	var InsertedIDs []string
	for _, user_id := range userID_array {
		cash_req := r.Petty_Cash_Request{
			ID:             uuid.NewString(),
			Date:           time.Now(),
			Description:    "an amazing petty cash request",
			User_ID:        user_id,
			Amount:         9,
			Grant_ID:       "2020-JY-FX-0014",
			Receipts:       []string{},
			Created_At:     time.Now(),
			Action_History: []r.Action{},
			Current_Status: "PENDING",
			Is_Active:      true,
		}
		result, err := collection.InsertOne(context.TODO(), cash_req)
		if err != nil {
			panic(err)
		}
		InsertedIDs = append(InsertedIDs, result.InsertedID.(string))
	}
	return InsertedIDs, nil
}
func MileageBulkInsert(userID_array []string) ([]string, error) {
	collection := Db.Collection("mileage_requests")
	var InsertedIDs []string
	for _, user_id := range userID_array {
		mileage_req := r.Mileage_Request{
			ID:                uuid.NewString(),
			Date:              time.Now(),
			User_ID:           user_id,
			Starting_Location: "the start of trip",
			Destination:       "the destination",
			Trip_Purpose:      "to drive from point a to b",
			Start_Odometer:    79,
			End_Odometer:      99,
			Tolls:             0.0,
			Parking:           0.0,
			Trip_Mileage:      20,
			Reimbursement:     12.5,
			Created_At:        time.Now(),
			Action_History:    []r.Action{},
			Current_Status:    r.PENDING,
			Is_Active:         true,
		}
		result, err := collection.InsertOne(context.TODO(), mileage_req)
		if err != nil {
			panic(err)
		}
		InsertedIDs = append(InsertedIDs, result.InsertedID.(string))
	}
	return InsertedIDs, nil
}
