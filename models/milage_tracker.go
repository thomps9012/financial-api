package models

import (
	"context"
	"financial-api/config"
	database "financial-api/db"
	"math"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"googlemaps.github.io/maps"
)

type TEST_LocationPoint struct {
	Latitude  float64 `json:"latitude" bson:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" bson:"longitude" validate:"required"`
}

type TEST_MatrixSubElement struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}
type TEST_MatrixElement struct {
	Distance TEST_MatrixSubElement `json:"distance"`
	Duration TEST_MatrixSubElement `json:"duration"`
	Status   string                `json:"status"`
}
type TEST_MatrixRow struct {
	Elements []TEST_MatrixElement `json:"elements"`
}

type TEST_MatrixResponse struct {
	Destination_Addresses []string         `json:"destination_addresses"`
	Origin_Addresses      []string         `json:"origin_addresses"`
	Rows                  []TEST_MatrixRow `json:"rows"`
	Status                string           `json:"status"`
}

type TEST_MileageTrackingInput struct {
	Grant_ID          string               `json:"grant_id" bson:"grant_id" validate:"required"`
	Category          Category             `json:"category" bson:"category" validate:"required"`
	Starting_Location string               `json:"starting_location" bson:"starting_location" validate:"required"`
	Destination       string               `json:"destination" bson:"destination" validate:"required"`
	Trip_Purpose      string               `json:"trip_purpose" bson:"trip_purpose" validate:"required"`
	Tracked_Points    []TEST_LocationPoint `json:"tracked_points" bson:"tracked_points" validate:"required,dive,required"`
	Tolls             *float64             `json:"tolls" bson:"tolls" validate:"required"`
	Parking           *float64             `json:"parking" bson:"parking" validate:"required"`
	Date              time.Time            `json:"date" bson:"date" validate:"required"`
}

type TEST_MileageTrackingRequest struct {
	ID                      string               `json:"id" bson:"_id"`
	Grant_ID                string               `json:"grant_id" bson:"grant_id"`
	User_ID                 string               `json:"user_id" bson:"user_id"`
	Date                    time.Time            `json:"date" bson:"date"`
	Category                Category             `json:"category" bson:"category"`
	Starting_Location       string               `json:"starting_location" bson:"starting_location"`
	Destination             string               `json:"destination" bson:"destination"`
	Trip_Purpose            string               `json:"trip_purpose" bson:"trip_purpose"`
	Tolls                   float64              `json:"tolls" bson:"tolls"`
	Parking                 float64              `json:"parking" bson:"parking"`
	Tracked_Points          []TEST_LocationPoint `json:"tracked_points" bson:"tracked_points"`
	Calculated_Distance     float64              `json:"calculated_distance" bson:"calculated_distance"`
	Trip_Mileage            float64              `json:"trip_mileage" bson:"trip_mileage"`
	Mileage_Variance        float64              `json:"mileage_variance" bson:"mileage_variance"`
	Variance_Level          TEST_VarianceLevel   `json:"variance_level" bson:"variance_level"`
	Reimbursement           float64              `json:"reimbursement" bson:"reimbursement"`
	Created_At              time.Time            `json:"created_at" bson:"created_at"`
	Action_History          []Action             `json:"action_history" bson:"action_history"`
	Current_User            string               `json:"current_user" bson:"current_user"`
	Current_Status          string               `json:"current_status" bson:"current_status"`
	Last_User_Before_Reject string               `json:"last_user_before_reject" bson:"last_user_before_reject"`
	Is_Active               bool                 `json:"is_active" bson:"is_active"`
}

type TEST_MileageTrackingRes struct {
	ID               string             `json:"id" bson:"_id"`
	User_ID          string             `json:"user_id" bson:"user_id"`
	Date             time.Time          `json:"date" bson:"date"`
	Reimbursement    float64            `json:"reimbursement" bson:"reimbursement"`
	Mileage          float64            `json:"mileage" bson:"mileage"`
	Trip_Mileage     float64            `json:"trip_mileage" bson:"trip_mileage"`
	Mileage_Variance float64            `json:"mileage_variance" bson:"mileage_variance"`
	Variance         TEST_VarianceLevel `json:"variance" bson:"variance"`
	Current_Status   string             `json:"current_status" bson:"current_status"`
	Current_User     string             `json:"current_user" bson:"current_user"`
	Is_Active        bool               `json:"is_active" bson:"is_active"`
}

type TEST_VarianceLevel string

const (
	HIGH   = "HIGH"
	MEDIUM = "MEDIUM"
	LOW    = "LOW"
)

type TEST_MileageDistance struct {
	MatrixDistance float64            `json:"matrix_distance"`
	Mileage        float64            `json:"mileage"`
	Difference     float64            `json:"difference"`
	Variance       TEST_VarianceLevel `json:"variance"`
}

var API_KEY = config.ENV("MAPS_API_KEY")

func (mti *TEST_MileageTrackingInput) DuplicateRequest(user_id string) (bool, error) {
	mileage_coll, err := database.Use("test_tracked_mileage")
	if err != nil {
		return false, err
	}
	odometer_filter := bson.D{{Key: "starting_location", Value: mti.Starting_Location}, {Key: "destination", Value: mti.Destination}, {Key: "user_id", Value: user_id}, {Key: "date", Value: mti.Date}}
	same_odometer, err := mileage_coll.CountDocuments(context.TODO(), odometer_filter)
	if err != nil {
		return false, err
	}
	return same_odometer > 0, nil
}

func (mti *TEST_MileageTrackingInput) CallMatrixAPI() (*maps.DistanceMatrixResponse, error) {
	c, err := maps.NewClient(maps.WithAPIKey(API_KEY))
	if err != nil {
		return nil, err
	}
	matrix_request := maps.DistanceMatrixRequest{
		Origins:      []string{mti.Starting_Location},
		Destinations: []string{mti.Destination},
		Mode:         maps.TravelModeDriving,
		Units:        maps.UnitsImperial,
	}
	matrix_response, err := c.DistanceMatrix(context.TODO(), &matrix_request)
	if err != nil {
		return nil, err
	}
	return matrix_response, nil
}

func DistanceBetweenLocationPoints(point_one TEST_LocationPoint, point_two TEST_LocationPoint) float64 {
	lat_one := (point_one.Latitude * math.Pi) / 180
	long_one := (point_one.Longitude * math.Pi) / 180
	lat_two := (point_two.Latitude * math.Pi) / 180
	long_two := (point_two.Longitude * math.Pi) / 180
	diff_long := long_two - long_one
	diff_lat := lat_two - lat_one
	haver_1 := math.Pow(math.Sin(diff_lat/2), 2) + math.Cos(lat_one)*math.Cos(lat_two)*math.Pow(math.Sin(diff_long/2), 2)
	haver_2 := 2 * math.Asin(math.Sqrt(haver_1))
	radius := 3956.00
	return haver_2 * radius
}

func (mti *TEST_MileageTrackingInput) CalculateTrackedDistance() float64 {
	var running_total float64
	for i, point := range mti.Tracked_Points {
		if i != len(mti.Tracked_Points)-1 {
			next_point := mti.Tracked_Points[i+1]
			distance := DistanceBetweenLocationPoints(point, next_point)
			running_total += distance
		}
	}
	return running_total
}

func CompareToMatrix(traveled_distance float64, matrix_res maps.DistanceMatrixResponse) (TEST_MileageDistance, error) {
	matrix_mileage := float64(matrix_res.Rows[0].Elements[0].Distance.Meters) / 1609.344
	variance := math.Round(traveled_distance - matrix_mileage)
	var variance_lvl TEST_VarianceLevel
	if variance > 10 || variance < 0 {
		variance_lvl = HIGH
	} else if variance > 1 {
		variance_lvl = MEDIUM
	} else {
		variance_lvl = LOW
	}
	return TEST_MileageDistance{
		MatrixDistance: matrix_mileage,
		Mileage:        traveled_distance,
		Variance:       variance_lvl,
		Difference:     variance,
	}, nil
}

func (mti *TEST_MileageTrackingInput) CreateRequest(user_id string) (*TEST_MileageTrackingRes, error) {
	matrix_res, err := mti.CallMatrixAPI()
	if err != nil {
		return nil, err
	}
	calculated_distance := mti.CalculateTrackedDistance()
	trip_variance, err := CompareToMatrix(calculated_distance, *matrix_res)
	if err != nil {
		return nil, err
	}
	first_action := FirstActions(user_id)

	new_request := new(TEST_MileageTrackingRequest)
	new_request.ID = uuid.NewString()
	new_request.Grant_ID = mti.Grant_ID
	new_request.User_ID = user_id
	new_request.Date = mti.Date
	new_request.Category = mti.Category
	new_request.Starting_Location = mti.Starting_Location
	new_request.Destination = mti.Destination
	new_request.Trip_Purpose = mti.Trip_Purpose
	new_request.Tolls = *mti.Tolls
	new_request.Parking = *mti.Parking
	new_request.Tracked_Points = mti.Tracked_Points
	new_request.Calculated_Distance = calculated_distance
	new_request.Trip_Mileage = trip_variance.MatrixDistance
	new_request.Mileage_Variance = trip_variance.Difference
	new_request.Variance_Level = trip_variance.Variance
	new_request.Reimbursement = *mti.Tolls + *mti.Parking + trip_variance.MatrixDistance*0.655
	new_request.Created_At = time.Now()
	new_request.Action_History = first_action
	new_request.Current_User = bson.TypeNull.String()
	new_request.Current_Status = "TESTING"
	new_request.Last_User_Before_Reject = bson.TypeNull.String()
	new_request.Is_Active = false

	collection, err := database.Use("test_tracked_mileage")
	if err != nil {
		return nil, err
	}
	_, err = collection.InsertOne(context.TODO(), new_request)
	if err != nil {
		return nil, err
	}
	return &TEST_MileageTrackingRes{
		ID:               new_request.ID,
		User_ID:          new_request.User_ID,
		Date:             new_request.Date,
		Reimbursement:    new_request.Reimbursement,
		Mileage:          new_request.Calculated_Distance,
		Trip_Mileage:     trip_variance.MatrixDistance,
		Mileage_Variance: trip_variance.Difference,
		Variance:         trip_variance.Variance,
		Current_Status:   new_request.Current_Status,
		Current_User:     new_request.Current_User,
		Is_Active:        new_request.Is_Active,
	}, nil
}
