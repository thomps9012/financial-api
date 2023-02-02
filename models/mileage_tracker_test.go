package models

import (
	"encoding/json"
	"math"
	"testing"
)

func mockDistanceMatrixAPI(api_url string) Matrix_Response {
	var response Matrix_Response
	var test_res = []byte(Test_matrix_response)
	err := json.Unmarshal(test_res, &response)
	if err != nil {
		panic(err)
	}
	return response
}
func mockSnapToRoadsAPI(api_url string) Snapped_Points_Response {
	var response Snapped_Points_Response
	var test_res = []byte(Test_snap_response)
	err := json.Unmarshal(test_res, &response)
	if err != nil {
		panic(err)
	}
	return response
}

func TestSnapJSONToEncode(t *testing.T) {
	var expected = Test_encoded_path
	input := Test_mileage_points
	actual := input.formatSnapJSONPoints()
	if actual != expected {
		t.Errorf("JSON to Encode Failed => expected: %s, got: %s", expected, actual)
	}
}
func TestFormSnapAPICall(t *testing.T) {
	var expected = Test_snap_api_call
	input := Test_mileage_points
	actual := input.formatSnapAPICall()
	if actual != expected {
		t.Errorf("Format Snap API Call Failed => expected: %s, got: %s", expected, actual)
	}
}

func TestMatrixStartForm(t *testing.T) {
	var expected = Test_origin
	input := Test_mileage_points
	actual := input.formatMatrixStart()
	if actual != expected {
		t.Errorf("Format Matrix Start Failed => expected: %s, got: %s", expected, actual)
	}
}

func TestMatrixDestinationForm(t *testing.T) {
	var expected = Test_destination
	input := Test_mileage_points
	actual := input.formatMatrixDestination()
	if actual != expected {
		t.Errorf("Format Matrix Destination Failed => expected: %s, got: %s", expected, actual)
	}
}

func TestFormMatrixAPICall(t *testing.T) {
	var expected = Test_matrix_api_call
	input := Test_mileage_points
	actual := input.formatMatrixAPICall()
	if actual != expected {
		t.Errorf("Format Matrix API Call Failed => expected: %s, got: %s", expected, actual)
	}
}

//	func TestFormatSnapAPIResponse(t *testing.T) {
//		input := Test_mileage_points
//		api_string := input.formatSnapAPICall()
//		actual := mockSnapToRoadsAPI(api_string)
//		fmt.Print(actual)
//		var expected = Test_formatted_snap_res
//		for i, point := range actual.Snapped_Points {
//			var expected_point = expected.Snapped_Points[i].Location
//			var location = point.Location
//			if location.Latitude != expected_point.Latitude {
//				t.Errorf("Format Snap API Res Failed at a Latitude Location Compare => expected: %f, got: %f", expected_point.Latitude, location.Latitude)
//			}
//			if location.Longitude != expected_point.Longitude {
//				t.Errorf("Format Snap API Res Failed at a Longitude Location Compare => expected: %f, got: %f", expected_point.Longitude, location.Longitude)
//			}
//		}
//	}
func TestFormatMatrixAPIResponse(t *testing.T) {
	input := Test_mileage_points
	api_string := input.formatMatrixAPICall()
	actual := mockDistanceMatrixAPI(api_string)
	var expected = Test_formatted_matrix_res
	if actual.Status != "OK" {
		t.Errorf("Format Matrix API Res Failed at Status => got: %s", actual.Status)
	}
	println(actual.Destination_Addresses)
	var actual_destination_address = actual.Destination_Addresses[0]
	var expected_destination_address = expected.Destination_Addresses[0]
	if actual_destination_address != expected_destination_address {
		t.Errorf("Format Matrix API Res Failed at Destination Addresses => expected: %s, got: %s", expected_destination_address, actual_destination_address)
	}
	var actual_origin_address = actual.Origin_Addresses[0]
	var expected_origin_address = expected.Origin_Addresses[0]
	if actual_origin_address != expected_origin_address {
		t.Errorf("Format Matrix API Res Failed at Origin Addresses => expected: %s, got: %s", expected_origin_address, actual_origin_address)
	}
	var distance_expected = expected.Rows[0].Elements[0].Distance
	var distance_actual = actual.Rows[0].Elements[0].Distance
	if distance_actual.Text != distance_expected.Text {
		t.Errorf("Format Matrix API Res Failed at Distance Text => expected: %s, got: %s", distance_expected.Text, distance_actual.Text)
	}
	if distance_actual.Value != distance_expected.Value {
		t.Errorf("Format Matrix API Res Failed at Distance Value => expected: %v, got: %v", distance_expected.Value, distance_actual.Value)
	}
	var duration_expected = expected.Rows[0].Elements[0].Duration
	var duration_actual = actual.Rows[0].Elements[0].Duration
	if duration_actual.Text != duration_expected.Text {
		t.Errorf("Format Matrix API Res Failed at Duration Text => expected: %s, got: %s", duration_expected.Text, duration_actual.Text)
	}
	if duration_actual.Value != duration_expected.Value {
		t.Errorf("Format Matrix API Res Failed at Duration Value => expected: %v, got: %v", duration_expected.Value, duration_actual.Value)
	}
}

func TestDistanceBetweenTwoPoints(t *testing.T) {
	expected := 129.047273
	point_a := Location{Latitude: 38.8976, Longitude: -77.0366}
	point_b := Location{Latitude: 39.9496, Longitude: -75.0366}
	actual := calculateDistanceBetweenPoints(point_a, point_b)
	if math.Round(actual) != math.Round(expected) {
		t.Errorf("Distance Between Two Points Failed => expected: %f, got: %f", expected, actual)
	}
}

// func TestSnapCalculation(t *testing.T) {
// 	expected := 0.518615
// 	actual := Test_formatted_snap_res.calculateSnapAPIDistance()
// 	if math.Round(actual) != math.Round(expected) {
// 		t.Errorf("Calculating Snapped API Distance Failed => expected: %f, got: %f", expected, actual)
// 	}
// }

func TestPointCalculation(t *testing.T) {
	dist_1 := 0.162064
	dist_2 := 0.04789
	dist_3 := 0.04389
	dist_4 := 0.03615
	dist_5 := 0.066283
	dist_6 := 0.04447
	dist_7 := 0.120761
	expected := dist_1 + dist_2 + dist_3 + dist_4 + dist_5 + dist_6 + dist_7
	input := Test_mileage_points
	actual := input.CalculatePreSnapDistance()
	if math.Round(expected) != math.Round(actual) {
		t.Errorf("Calculating Pre Snap Distance Failed => expected: %f, got: %f", expected, actual)
	}
}

func TestResponseCompare(t *testing.T) {
	input := Test_mileage_points
	var calculated_dist = input.CalculatePreSnapDistance()
	api_string := input.formatMatrixAPICall()
	matrix_res := mockDistanceMatrixAPI(api_string)
	expected_low := Test_low_variance
	expected_med := Test_med_variance
	expected_high := Test_high_variance
	var res_compare, err = matrix_res.CompareToMatrix(calculated_dist)
	if err != nil {
		t.Error(err)
	}
	if res_compare != expected_low {
		t.Errorf("Failed Test Response Compare => expected: %v, got: %v", expected_low, res_compare)
	}
	if res_compare == expected_med {
		t.Errorf("Failed Test Response Compare => expected: %v, got: %v", expected_low, res_compare)
	}
	if res_compare == expected_high {
		t.Errorf("Failed Test Response Compare => expected: %v, got: %v", expected_low, res_compare)
	}
	res_compare, err = matrix_res.CompareToMatrix(7.07)
	if err != nil {
		t.Error(err)
	}
	if res_compare == expected_low {
		t.Errorf("Failed Test Response Compare => expected: %v, got: %v", expected_low, res_compare)
	}
	if res_compare != expected_med {
		t.Errorf("Failed Test Response Compare => expected: %v, got: %v", expected_low, res_compare)
	}
	if res_compare == expected_high {
		t.Errorf("Failed Test Response Compare => expected: %v, got: %v", expected_low, res_compare)
	}
	res_compare, err = matrix_res.CompareToMatrix(20.02)
	if err != nil {
		t.Error(err)
	}
	if res_compare == expected_low {
		t.Errorf("Failed Test Response Compare => expected: %v, got: %v", expected_low, res_compare)
	}
	if res_compare == expected_med {
		t.Errorf("Failed Test Response Compare => expected: %v, got: %v", expected_low, res_compare)
	}
	if res_compare != expected_high {
		t.Errorf("Failed Test Response Compare => expected: %v, got: %v", expected_low, res_compare)
	}
}
