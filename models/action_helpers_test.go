package models

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestReturnPrevAdminID(t *testing.T) {
	action_history := []Action{
		{
			ID:           "1",
			Request_ID:   "1",
			Request_Type: MILEAGE,
			User:         "1",
			Status:       CREATED,
			Created_At:   time.Now(),
		},
		{
			ID:           "2",
			Request_ID:   "1",
			Request_Type: MILEAGE,
			User:         "2",
			Status:       PENDING,
			Created_At:   time.Now(),
		},
		{
			ID:           "3",
			Request_ID:   "1",
			Request_Type: MILEAGE,
			User:         "3",
			Status:       REJECTED,
			Created_At:   time.Now(),
		},
		{
			ID:           "4",
			Request_ID:   "1",
			Request_Type: MILEAGE,
			User:         "4",
			Status:       REJECTED_EDIT,
			Created_At:   time.Now(),
		},
		{
			ID:           "5",
			Request_ID:   "1",
			Request_Type: MILEAGE,
			User:         "5",
			Status:       PENDING,
			Created_At:   time.Now(),
		},
	}
	requestor_id := "1"
	expected := "3"
	actual := ReturnPrevAdminID(action_history, requestor_id)
	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
func TestFormRequestType(t *testing.T) {
	mileage_actual := format_request_type(MILEAGE)
	mileage_expected := "mileage_requests"
	if mileage_actual != mileage_expected {
		t.Errorf("Mileage Test Case failed, expected: '%s', got: '%s'", mileage_expected, mileage_actual)
	}
	check_action := format_request_type(CHECK)
	check_expected := "check_requests"
	if check_action != check_expected {
		t.Errorf("Check Request Test Case failed, expected: '%s', got: '%s'", check_expected, check_action)
	}
	petty_cash_actual := format_request_type(PETTY_CASH)
	petty_cash_expected := "petty_cash_requests"
	if petty_cash_actual != petty_cash_expected {
		t.Errorf("Petty Cash Test Case failed, expected: '%s', got: '%s'", petty_cash_expected, petty_cash_actual)
	}
}
func TestExecEmailHandler(t *testing.T) {
	var category Category = ADMINISTRATIVE
	var current_status Status = FINANCE_APPROVED
	var exec_review bool = true
	var to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = IOP
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = INTAKE
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = PEERS
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = ACT_TEAM
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = IHBT
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = PERKINS
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = MENS_HOUSE
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = NEXT_STEP
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = LORAIN
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = PREVENTION
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	category = FINANCE
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	current_status = CREATED
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	current_status = PENDING
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	current_status = MANAGER_APPROVED
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	current_status = SUPERVISOR_APPROVED
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	current_status = FINANCE_APPROVED
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	current_status = ORGANIZATION_APPROVED
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
}
func TestAdminEmailHandler(t *testing.T) {
	// test for finance approved
	var category Category = ADMINISTRATIVE
	var current_status Status = FINANCE_APPROVED
	var exec_review bool = false
	var to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "abradley@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected abradley@norainc.org, got %s", to_email)
	}
	// test for rejected
	current_status = REJECTED
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "" {
		t.Errorf("UserEmailHandlerTest: expected empty string, got %s", to_email)
	}
	// test for organization approved
	current_status = ORGANIZATION_APPROVED
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "" {
		t.Errorf("UserEmailHandlerTest: expected empty string, got %s", to_email)
	}
	// test for default supervisor approved
	current_status = SUPERVISOR_APPROVED
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
	// test for default executive  and supervisor approved
	current_status = EXECUTIVE_APPROVED
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
	// test for lorain manager approved
	current_status = MANAGER_APPROVED
	category = LORAIN
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	// test for next step manager approved
	current_status = MANAGER_APPROVED
	category = NEXT_STEP
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "cwoods@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected cwoods@norainc.org, got %s", to_email)
	}
	// test for perkins manager approved
	current_status = MANAGER_APPROVED
	category = PERKINS
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	// test for prevention manager approved
	current_status = MANAGER_APPROVED
	category = PREVENTION
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "cwoods@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected cwoods@norainc.org, got %s", to_email)
	}
	// tests for default manager approved
	current_status = MANAGER_APPROVED
	category = IOP
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
	current_status = MANAGER_APPROVED
	category = INTAKE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
	current_status = MANAGER_APPROVED
	category = PEERS
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
	current_status = MANAGER_APPROVED
	category = ACT_TEAM
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
	current_status = MANAGER_APPROVED
	category = FINANCE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
	current_status = MANAGER_APPROVED
	category = IHBT
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
	current_status = MANAGER_APPROVED
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
}
func TestCreatedEmailHandler(t *testing.T) {
	// created test level cases
	current_status := CREATED
	category := ADMINISTRATIVE
	var exec_review bool = false
	to_email := UserEmailHandler(category, current_status, exec_review)
	if to_email != "bgriffin@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected bgriffin@norainc.org, got %s", to_email)
	}
	category = IOP
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	category = INTAKE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "cwoods@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected cwoods@norainc.org, got %s", to_email)
	}
	category = PEERS
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	category = ACT_TEAM
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jjordan@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jjordan@norainc.org, got %s", to_email)
	}
	category = IHBT
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "bgriffin@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected bgriffin@norainc.org, got %s", to_email)
	}
	category = FINANCE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "lfuentes@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected lfuentes@norainc.org, got %s", to_email)
	}
	category = LORAIN
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "rgiusti@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected rgiusti@norainc.org, got %s", to_email)
	}
	category = MENS_HOUSE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	category = NEXT_STEP
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "dbaker@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected dbaker@norainc.org, got %s", to_email)
	}
	category = PERKINS
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "churt@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected churt@norainc.org, got %s", to_email)
	}
	category = PREVENTION
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "lamanor@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected lamanor@norainc.org, got %s", to_email)
	}
	category = ""
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
}
func TestPendingEmailHandler(t *testing.T) {
	// pending test cases
	current_status := PENDING
	category := ADMINISTRATIVE
	var exec_review bool = false
	to_email := UserEmailHandler(category, current_status, exec_review)
	if to_email != "bgriffin@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected bgriffin@norainc.org, got %s", to_email)
	}
	category = IOP
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	category = INTAKE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "cwoods@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected cwoods@norainc.org, got %s", to_email)
	}
	category = PEERS
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	category = ACT_TEAM
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jjordan@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jjordan@norainc.org, got %s", to_email)
	}
	category = IHBT
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "bgriffin@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected bgriffin@norainc.org, got %s", to_email)
	}
	category = FINANCE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "lfuentes@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected lfuentes@norainc.org, got %s", to_email)
	}
	category = LORAIN
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "rgiusti@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected rgiusti@norainc.org, got %s", to_email)
	}
	category = MENS_HOUSE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	category = NEXT_STEP
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "dbaker@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected dbaker@norainc.org, got %s", to_email)
	}
	category = PERKINS
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "churt@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected churt@norainc.org, got %s", to_email)
	}
	category = PREVENTION
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "lamanor@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected lamanor@norainc.org, got %s", to_email)
	}
	category = ""
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
}
func TestRejectedEmailHandler(t *testing.T) {
	// rejected edit test cases
	current_status := REJECTED_EDIT
	category := ADMINISTRATIVE
	var exec_review bool = false
	to_email := UserEmailHandler(category, current_status, exec_review)
	if to_email != "bgriffin@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected bgriffin@norainc.org, got %s", to_email)
	}
	category = IOP
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	category = INTAKE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "cwoods@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected cwoods@norainc.org, got %s", to_email)
	}
	category = PEERS
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	category = ACT_TEAM
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jjordan@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jjordan@norainc.org, got %s", to_email)
	}
	category = IHBT
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "bgriffin@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected bgriffin@norainc.org, got %s", to_email)
	}
	category = FINANCE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "lfuentes@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected lfuentes@norainc.org, got %s", to_email)
	}
	category = LORAIN
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "rgiusti@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected rgiusti@norainc.org, got %s", to_email)
	}
	category = MENS_HOUSE
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "jward@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected jward@norainc.org, got %s", to_email)
	}
	category = NEXT_STEP
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "dbaker@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected dbaker@norainc.org, got %s", to_email)
	}
	category = PERKINS
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "churt@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected churt@norainc.org, got %s", to_email)
	}
	category = PREVENTION
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "lamanor@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected lamanor@norainc.org, got %s", to_email)
	}
	category = ""
	to_email = UserEmailHandler(category, current_status, exec_review)
	if to_email != "finance_requests@norainc.org" {
		t.Errorf("UserEmailHandlerTest: expected finance_requests@norainc.org, got %s", to_email)
	}
}
func TestCheckStatus(t *testing.T) {
	request_info := Request_Info{
		User_ID:        "1",
		Current_User:   "1",
		Current_Status: CREATED,
		Type:           MILEAGE,
		ID:             "001",
	}
	var new_status = PENDING
	var expected_response = false
	if expected_response != request_info.CheckStatus(new_status) {
		t.Errorf("Request Status Check: expected %t, got %t", expected_response, !expected_response)
	}
	request_info.Current_Status = PENDING
	new_status = MANAGER_APPROVED
	expected_response = false
	if expected_response != request_info.CheckStatus(new_status) {
		t.Errorf("Request Status Check: expected %t, got %t", expected_response, !expected_response)
	}
	request_info.Current_Status = REJECTED
	new_status = REJECTED_EDIT
	expected_response = false
	if expected_response != request_info.CheckStatus(new_status) {
		t.Errorf("Request Status Check: expected %t, got %t", expected_response, !expected_response)
	}
	request_info.Current_Status = REJECTED_EDIT
	new_status = REJECTED_EDIT
	expected_response = true
	if expected_response != request_info.CheckStatus(new_status) {
		t.Errorf("Request Status Check: expected %t, got %t", expected_response, !expected_response)
	}
}

func TestMileageUpdateRequest(t *testing.T) {
	var action_created = time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC)
	var action_created_2 = time.Date(2021, 1, 10, 1, 0, 0, 0, time.UTC)
	test_action := Action{
		ID:           "199",
		Request_ID:   "777",
		Request_Type: MILEAGE,
		User:         "2",
		Status:       REJECTED,
		Created_At:   action_created_2,
	}
	first_action := Action{
		ID:           "198",
		Request_ID:   "777",
		Request_Type: MILEAGE,
		User:         "1",
		Status:       CREATED,
		Created_At:   action_created,
	}
	user_id := "1"
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	// mt.Run("success full response", func(mt *mtest.T) {
	// 	var collection = mt.Coll
	// 	var filter = bson.D{{Key: "_id", Value: test_action.Request_ID}}
	// 	var update = bson.D{{Key: "$push", Value: bson.M{"action_history": test_action}}, {Key: "$set", Value: bson.M{"current_user": user_id}}, {Key: "$set", Value: bson.M{"current_status": test_action.Status}}}
	// 	var req_created = time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC)
	// 	successful_response := Mileage_Request{
	// 		ID:                "777",
	// 		Grant_ID:          "SOR_PEER",
	// 		User_ID:           "2",
	// 		Category:          IOP,
	// 		Starting_Location: "A",
	// 		Destination:       "B",
	// 		Trip_Purpose:      "Test",
	// 		Start_Odometer:    0,
	// 		End_Odometer:      1,
	// 		Tolls:             0.0,
	// 		Parking:           0.0,
	// 		Trip_Mileage:      1,
	// 		Reimbursement:     0.62,
	// 		Created_At:        req_created,
	// 		Action_History: []Action{
	// 			first_action,
	// 			test_action,
	// 		},
	// 		Current_User:   "1",
	// 		Current_Status: REJECTED,
	// 		Is_Active:      true,
	// 	}
	// 	mt.AddMockResponses(bson.D{
	// 		{Key: "ok", Value: 1},
	// 		{Key: "value", Value: bson.D{
	// 			{Key: "_id", Value: test_action.Request_ID},
	// 			{Key: "grant_id", Value: successful_response.Grant_ID},
	// 			{Key: "user_id", Value: successful_response.User_ID},
	// 			{Key: "category", Value: successful_response.Category},
	// 			{Key: "starting_location", Value: successful_response.Starting_Location},
	// 			{Key: "destination", Value: successful_response.Destination},
	// 			{Key: "trip_purpose", Value: successful_response.Trip_Purpose},
	// 			{Key: "start_odometer", Value: successful_response.Start_Odometer},
	// 			{Key: "end_odometer", Value: successful_response.End_Odometer},
	// 			{Key: "tolls", Value: successful_response.Tolls},
	// 			{Key: "parking", Value: successful_response.Parking},
	// 			{Key: "trip_mileage", Value: successful_response.Trip_Mileage},
	// 			{Key: "reimbursement", Value: successful_response.Reimbursement},
	// 			{Key: "created_at", Value: req_created},
	// 			{Key: "action_history", Value: []Action{first_action, test_action}},
	// 			{Key: "current_user", Value: user_id},
	// 			{Key: "current_status", Value: test_action.Status},
	// 			{Key: "is_active", Value: true},
	// 		}},
	// 	})
	// 	var mileage_res Mileage_Request
	// 	collection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&mileage_res)
	// 	assert.Equal(t, &successful_response, &mileage_res)

	// })
	mt.Run("success response info", func(mt *mtest.T) {
		var collection = mt.Coll
		var filter = bson.D{{Key: "_id", Value: test_action.Request_ID}}
		var update = bson.D{{Key: "$push", Value: bson.M{"action_history": test_action}}, {Key: "$set", Value: bson.M{"current_user": user_id}}, {Key: "$set", Value: bson.M{"current_status": test_action.Status}}}
		successful_response := Request_Info_With_Action_History{
			User_ID:        "2",
			Current_User:   "1",
			Current_Status: REJECTED,
			ID:             "777",
			Type:           MILEAGE,
			Action_History: []Action{
				first_action,
				test_action,
			},
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "user_id", Value: successful_response.User_ID},
				{Key: "current_user", Value: user_id},
				{Key: "current_status", Value: test_action.Status},
				{Key: "_id", Value: test_action.Request_ID},
				{Key: "type", Value: MILEAGE},
				{Key: "action_history", Value: []Action{first_action, test_action}},
			}},
		})
		var mileage_res Request_Info_With_Action_History
		collection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&mileage_res)
		assert.Equal(t, &successful_response, &mileage_res)

	})
	mt.Run("unsuccessful response", func(mt *mtest.T) {
		var collection = mt.Coll
		var filter = bson.D{{Key: "_id", Value: test_action.Request_ID}}
		var update = bson.D{{Key: "$push", Value: bson.M{"action_history": test_action}}, {Key: "$set", Value: bson.M{"current_user": user_id}}, {Key: "$set", Value: bson.M{"current_status": test_action.Status}}}
		unsuccessful_response := Request_Info_With_Action_History{
			User_ID:        "2",
			Current_User:   "1",
			Current_Status: REJECTED,
			ID:             "777",
			Type:           MILEAGE,
			Action_History: []Action{
				first_action,
				test_action,
			},
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "user_id", Value: user_id},
				{Key: "current_user", Value: user_id},
				{Key: "current_status", Value: test_action.Status},
				{Key: "_id", Value: test_action.Request_ID},
				{Key: "type", Value: MILEAGE},
				{Key: "action_history", Value: []Action{first_action, test_action}},
			}},
		})
		var mileage_res Request_Info_With_Action_History
		collection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&mileage_res)
		assert.NotEqual(t, &unsuccessful_response, &mileage_res)
	})
}

func TestPettyCashUpdateRequest(t *testing.T) {
	var action_created = time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC)
	var action_created_2 = time.Date(2021, 1, 10, 1, 0, 0, 0, time.UTC)
	test_action := Action{
		ID:           "199",
		Request_ID:   "777",
		Request_Type: PETTY_CASH,
		User:         "2",
		Status:       REJECTED,
		Created_At:   action_created_2,
	}
	first_action := Action{
		ID:           "198",
		Request_ID:   "777",
		Request_Type: PETTY_CASH,
		User:         "1",
		Status:       CREATED,
		Created_At:   action_created,
	}
	user_id := "1"
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success response info", func(mt *mtest.T) {
		var collection = mt.Coll
		var filter = bson.D{{Key: "_id", Value: test_action.Request_ID}}
		var update = bson.D{{Key: "$push", Value: bson.M{"action_history": test_action}}, {Key: "$set", Value: bson.M{"current_user": user_id}}, {Key: "$set", Value: bson.M{"current_status": test_action.Status}}}
		successful_response := Request_Info_With_Action_History{
			User_ID:        "2",
			Current_User:   "1",
			Current_Status: REJECTED,
			ID:             "777",
			Type:           PETTY_CASH,
			Action_History: []Action{
				first_action,
				test_action,
			},
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "user_id", Value: successful_response.User_ID},
				{Key: "current_user", Value: user_id},
				{Key: "current_status", Value: test_action.Status},
				{Key: "_id", Value: test_action.Request_ID},
				{Key: "type", Value: PETTY_CASH},
				{Key: "action_history", Value: []Action{first_action, test_action}},
			}},
		})
		var petty_cash Request_Info_With_Action_History
		collection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&petty_cash)
		assert.Equal(t, &successful_response, &petty_cash)
	})
	mt.Run("unsuccessful response", func(mt *mtest.T) {
		var collection = mt.Coll
		var filter = bson.D{{Key: "_id", Value: test_action.Request_ID}}
		var update = bson.D{{Key: "$push", Value: bson.M{"action_history": test_action}}, {Key: "$set", Value: bson.M{"current_user": user_id}}, {Key: "$set", Value: bson.M{"current_status": test_action.Status}}}
		unsuccessful_response := Request_Info_With_Action_History{
			User_ID:        "2",
			Current_User:   "1",
			Current_Status: REJECTED,
			ID:             "777",
			Type:           PETTY_CASH,
			Action_History: []Action{
				first_action,
				test_action,
			},
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "user_id", Value: user_id},
				{Key: "current_user", Value: user_id},
				{Key: "current_status", Value: test_action.Status},
				{Key: "_id", Value: test_action.Request_ID},
				{Key: "type", Value: PETTY_CASH},
				{Key: "action_history", Value: []Action{first_action, test_action}},
			}},
		})
		var petty_cash Request_Info_With_Action_History
		collection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&petty_cash)
		assert.NotEqual(t, &unsuccessful_response, &petty_cash)
	})
}

func TestCheckUpdateRequest(t *testing.T) {
	var action_created = time.Date(2021, 1, 1, 1, 0, 0, 0, time.UTC)
	var action_created_2 = time.Date(2021, 1, 10, 1, 0, 0, 0, time.UTC)
	test_action := Action{
		ID:           "199",
		Request_ID:   "777",
		Request_Type: CHECK,
		User:         "2",
		Status:       REJECTED,
		Created_At:   action_created_2,
	}
	first_action := Action{
		ID:           "198",
		Request_ID:   "777",
		Request_Type: CHECK,
		User:         "1",
		Status:       CREATED,
		Created_At:   action_created,
	}
	user_id := "1"
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success response info", func(mt *mtest.T) {
		var collection = mt.Coll
		var filter = bson.D{{Key: "_id", Value: test_action.Request_ID}}
		var update = bson.D{{Key: "$push", Value: bson.M{"action_history": test_action}}, {Key: "$set", Value: bson.M{"current_user": user_id}}, {Key: "$set", Value: bson.M{"current_status": test_action.Status}}}
		successful_response := Request_Info_With_Action_History{
			User_ID:        "2",
			Current_User:   "1",
			Current_Status: REJECTED,
			ID:             "777",
			Type:           CHECK,
			Action_History: []Action{
				first_action,
				test_action,
			},
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "user_id", Value: successful_response.User_ID},
				{Key: "current_user", Value: user_id},
				{Key: "current_status", Value: test_action.Status},
				{Key: "_id", Value: test_action.Request_ID},
				{Key: "type", Value: CHECK},
				{Key: "action_history", Value: []Action{first_action, test_action}},
			}},
		})
		var check Request_Info_With_Action_History
		collection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&check)
		assert.Equal(t, &successful_response, &check)
	})
	mt.Run("unsuccessful response", func(mt *mtest.T) {
		var collection = mt.Coll
		var filter = bson.D{{Key: "_id", Value: test_action.Request_ID}}
		var update = bson.D{{Key: "$push", Value: bson.M{"action_history": test_action}}, {Key: "$set", Value: bson.M{"current_user": user_id}}, {Key: "$set", Value: bson.M{"current_status": test_action.Status}}}
		unsuccessful_response := Request_Info_With_Action_History{
			User_ID:        "2",
			Current_User:   "1",
			Current_Status: REJECTED,
			ID:             "777",
			Type:           CHECK,
			Action_History: []Action{
				first_action,
				test_action,
			},
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "user_id", Value: user_id},
				{Key: "current_user", Value: user_id},
				{Key: "current_status", Value: test_action.Status},
				{Key: "_id", Value: test_action.Request_ID},
				{Key: "type", Value: CHECK},
				{Key: "action_history", Value: []Action{first_action, test_action}},
			}},
		})
		var check_res Request_Info_With_Action_History
		collection.FindOneAndUpdate(context.Background(), filter, update, options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&check_res)
		assert.NotEqual(t, &unsuccessful_response, &check_res)
	})
}

func TestGeneralUpdateRequest(t *testing.T) {
	test_action := Action{
		ID:           "199",
		Request_ID:   "777",
		Request_Type: MILEAGE,
		User:         "2",
		Status:       REJECTED,
		Created_At:   time.Now(),
	}
	switch test_action.Request_Type {
	case MILEAGE:
		TestMileageUpdateRequest(t)
	case PETTY_CASH:
		TestPettyCashUpdateRequest(t)
	case CHECK:
		TestCheckUpdateRequest(t)
	}
	test_action.Request_Type = CHECK
	switch test_action.Request_Type {
	case MILEAGE:
		TestMileageUpdateRequest(t)
	case PETTY_CASH:
		TestPettyCashUpdateRequest(t)
	case CHECK:
		TestCheckUpdateRequest(t)
	}
	test_action.Request_Type = PETTY_CASH
	switch test_action.Request_Type {
	case MILEAGE:
		TestMileageUpdateRequest(t)
	case PETTY_CASH:
		TestPettyCashUpdateRequest(t)
	case CHECK:
		TestCheckUpdateRequest(t)
	}
}

// func TestDetermineUserID(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock).CollectionName("users"))
// 	defer mt.Close()
// 	mt.Coll = mt.DB.Collection("users")
// 	mt.Run("find user by id", func(t *mtest.T) {
// 		var user User
// 		var email = "emp1@norainc.org"
// 		var confirm_id = "117035974203946272200"
// 		id, err := user.FindID(email)
// 		if err != nil {
// 			t.Errorf("Unexpected Error: %s", err)
// 		}
// 		if id != confirm_id {
// 			t.Errorf("Determine User ID Failed: expected %s, got %s", confirm_id, id)
// 		}
// 	})
// }
