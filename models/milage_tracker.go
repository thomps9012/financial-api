package models

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Mileage_Points struct {
	LocationPoints []Location `json:"location_points" bson:"location_points"`
	Starting_Point Location   `json:"starting_point" bson:"starting_point"`
	Destination    Location   `json:"destination" bson:"destination"`
}

type Location struct {
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}

type Snapped_Point struct {
	Location Location `json:"location"`
	PlaceID  string   `json:"placeId"`
}

type Snapped_Points_Response struct {
	Snapped_Points []Snapped_Point `json:"snappedPoints"`
}
type Matrix_Sub_Element struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}
type Matrix_Elements struct {
	Distance Matrix_Sub_Element `json:"distance"`
	Duration Matrix_Sub_Element `json:"duration"`
	Status   string             `json:"status"`
}
type Matrix_Row struct {
	Elements []Matrix_Elements `json:"elements"`
}
type Matrix_Response struct {
	Destination_Addresses []string     `json:"destination_addresses"`
	Origin_Addresses      []string     `json:"origin_addresses"`
	Rows                  []Matrix_Row `json:"rows"`
	Status                string       `json:"status"`
}

type Variance_Level string

const (
	HIGH   = "HIGH"
	MEDIUM = "MEDIUM"
	LOW    = "LOW"
)

type ResponseCompare struct {
	Matrix_Distance   float64        `json:"matrix_distance"`
	Traveled_Distance float64        `json:"traveled_distance"`
	Difference        float64        `json:"difference"`
	Variance          Variance_Level `json:"variance"`
}

var API_KEY = os.Getenv("MAPS_API_KEY")

const SNAP_API_BASE = "https://roads.googleapis.com/v1/snapToRoads?path="
const MATRIX_API_BASE = "https://maps.googleapis.com/maps/api/distancematrix/json?origins="

func (m *Mileage_Points) formatSnapJSONPoints() string {
	var location_points = m.LocationPoints
	var api_points string
	var point_string string
	for i, point := range location_points {
		if i == len(location_points)-1 {
			point_string = strconv.FormatFloat(point.Latitude, 'f', 5, 64) + "," + strconv.FormatFloat(point.Longitude, 'f', 5, 64)
		} else {
			point_string = strconv.FormatFloat(point.Latitude, 'f', 5, 64) + "," + strconv.FormatFloat(point.Longitude, 'f', 5, 64) + "|"
		}
		api_points += point_string
	}
	return api_points
}

func (m *Mileage_Points) formatSnapAPICall() string {
	var api_points = m.formatSnapJSONPoints()
	return SNAP_API_BASE + api_points + "&interpolate=true&key=" + API_KEY
}

func (m *Mileage_Points) formatMatrixStart() string {
	var starting_point = m.Starting_Point
	return strconv.FormatFloat(starting_point.Latitude, 'f', 5, 64) + "," + strconv.FormatFloat(starting_point.Longitude, 'f', 5, 64)
}

func (m *Mileage_Points) formatMatrixDestination() string {
	var destination = m.Destination
	return strconv.FormatFloat(destination.Latitude, 'f', 5, 64) + "," + strconv.FormatFloat(destination.Longitude, 'f', 5, 64)
}

func (m *Mileage_Points) formatMatrixAPICall() string {
	var start_string = m.formatMatrixStart()
	var destination_string = m.formatMatrixDestination()
	return MATRIX_API_BASE + start_string + "&destinations=" + destination_string + "&units=imperial&key=" + API_KEY
}

func (m *Mileage_Points) callSnapAPI() (Snapped_Points_Response, error) {
	client := &http.Client{}
	api_url := m.formatSnapAPICall()
	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var snapped_points Snapped_Points_Response
	json_err := json.Unmarshal(body, &snapped_points)
	if json_err != nil {
		panic(json_err)
	}
	return snapped_points, nil
}

func (m *Mileage_Points) CallMatrixAPI() (Matrix_Response, error) {
	api_url := m.formatMatrixAPICall()
	client := &http.Client{}
	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var matrix_res Matrix_Response
	json_err := json.Unmarshal(body, &matrix_res)
	if json_err != nil {
		panic(json_err)
	}
	return matrix_res, nil
}

func calculateDistanceBetweenPoints(point_one Location, point_two Location) float64 {
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

func (s *Snapped_Points_Response) calculateSnapAPIDistance() float64 {
	panic("unimplemented function")
	var running_total float64
	for i, point := range s.Snapped_Points {
		if i != len(s.Snapped_Points)-1 {
			next_point := s.Snapped_Points[i+1].Location
			distance := calculateDistanceBetweenPoints(point.Location, next_point)
			running_total += distance
		}
	}
	return running_total
}

func (m *Mileage_Points) CalculatePreSnapDistance() float64 {
	var running_total float64
	for i, point := range m.LocationPoints {
		if i != len(m.LocationPoints)-1 {
			next_point := m.LocationPoints[i+1]
			distance := calculateDistanceBetweenPoints(point, next_point)
			running_total += distance
		}
	}
	return running_total
}

func (mr *Matrix_Response) CompareToMatrix(traveled_distance float64) (ResponseCompare, error) {
	matrix_distance, err := strconv.ParseFloat(strings.Split(mr.Rows[0].Elements[0].Distance.Text, " ")[0], 64)
	if err != nil {
		panic(err)
	}
	variance := math.Round(traveled_distance - matrix_distance)
	var variance_lvl Variance_Level
	if variance > 10 {
		variance_lvl = HIGH
	} else if variance > 1 {
		variance_lvl = MEDIUM
	} else {
		variance_lvl = LOW
	}
	return ResponseCompare{
		Matrix_Distance:   matrix_distance,
		Traveled_Distance: traveled_distance,
		Variance:          variance_lvl,
		Difference:        variance,
	}, nil
}
