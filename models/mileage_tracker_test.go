package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func mockDistanceMatrixAPI(api_url string) Matrix_Response {
	var response Matrix_Response
	var test_res = []byte(Test_matrix_response)
	json.Unmarshal(test_res, &response)
	return response
}
func mockSnapToRoadsAPI(api_url string) Snapped_Points_Response {
	var response Snapped_Points_Response
	var test_res = []byte(Test_snap_response)
	json.Unmarshal(test_res, &response)
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
func TestFormatSnapAPIResponse(t *testing.T) {
	input := Test_mileage_points
	api_string := input.formatSnapAPICall()
	actual := mockSnapToRoadsAPI(api_string)
	fmt.Print(actual)
	var expected = Test_formatted_snap_res
	for i, point := range actual.Snapped_Points {
		var expected_point = expected.Snapped_Points[i]
		if point.Latitude != expected_point.Latitude {
			t.Errorf("Format Snap API Res Failed at a Latitude Location Compare => expected: %f, got: %f", expected_point.Latitude, point.Latitude)
		}
		if point.Longitude != expected_point.Longitude {
			t.Errorf("Format Snap API Res Failed at a Longitude Location Compare => expected: %f, got: %f", expected_point.Longitude, point.Longitude)
		}
	}
}
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

}

func TestSnapCalculation(t *testing.T) {

}

func TestResponseCompare(t *testing.T) {

}
