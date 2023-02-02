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

}
func TestFormMatrixAPICall(t *testing.T) {

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
