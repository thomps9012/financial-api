package models

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

type Mileage_Points struct {
	LocationPoints []Location `json:"location_points" bson:"location_points"`
	Starting_Point Location   `json:"starting_point" bson:"starting_point"`
	Destination    Location   `json:"destination" bson:"destination"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Snapped_Point struct {
	Location Location `json:"location"`
}

type Snapped_Points_Response struct {
	Snapped_Points []Snapped_Point `json:"snappedPoints"`
	Original_Index int             `json:"originalIndex"`
	Place_ID       string          `json:"placeId"`
}
type Matrix_Sub_Element struct {
	Text  string `json:"text"`
	Value int    `json:"value"`
}
type Matrix_Elements struct {
	Distance Matrix_Sub_Element `json:"distance"`
	Duration Matrix_Sub_Element `json:"duration"`
}
type Matrix_Row struct {
	Elements []Matrix_Elements `json:"elements"`
}
type Matrix_Response struct {
	Destination_Addresses string       `json:"destination_addresses"`
	Origin_Addresses      string       `json:"origin_addresses"`
	Rows                  []Matrix_Row `json:"rows"`
	Status                `json:"status"`
}

type Variance_Level string

const (
	HIGH   = "HIGH"
	MEDIUM = "MEDIUM"
	LOW    = "LOW"
)

type ResponseCompare struct {
	Matrix_Distance  float64        `json:"matrix_distance"`
	Snapped_Distance float64        `json:"snapped_distance"`
	Difference       float64        `json:"difference"`
	Variance         Variance_Level `json:"variance"`
}

const API_KEY = "AIzaSyAf7mF7egyl3Ip35hN1n9gXP854_u5-Zsk"
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

func callAPI(api_url string) ([]byte, error) {
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
	return body, nil
}

func formatSnapAPIResponse(api_response []byte) (Snapped_Points_Response, error) {
	panic("unimplemented function")
}

func formatMatrixAPIResponse(api_response []byte) (Matrix_Response, error) {
	panic("unimplemented function")
}

func calculateDistanceBetweenPoints(point_one Location, point_two Location) float64 {
	panic("unimplemented function")
}

func calculateSnapAPIDistance(snapped_points Snapped_Points_Response) (float64, error) {
	panic("unimplemented function")
}

func compareSnapToMatrix(snapped_distance float64, matrix_res Matrix_Response) (ResponseCompare, error) {
	panic("unimplemented function")
}
