package models

import (
	"testing"
)

func mockDistanceMatrixAPI(api_url string) []byte {
	var test_res = []byte(Test_snap_response)
	return test_res
}
func mockSnapToRoadsAPI(api_url string) []byte {
	var test_res = []byte(Test_snap_response)
	return test_res
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

}
func TestFormatMatrixAPIResponse(t *testing.T) {

}

func TestDistanceBetweenTwoPoints(t *testing.T) {

}

func TestSnapCalculation(t *testing.T) {

}

func TestResponseCompare(t *testing.T) {

}
