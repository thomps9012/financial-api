package models

import (
	"testing"
	"time"
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

func TestUserEmailHandler(t *testing.T) {
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

	// created test level cases
	current_status = CREATED
	category = ADMINISTRATIVE
	to_email = UserEmailHandler(category, current_status, exec_review)
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
	// pending test cases
	current_status = PENDING
	category = ADMINISTRATIVE
	to_email = UserEmailHandler(category, current_status, exec_review)
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
	// rejected edit test cases
	current_status = REJECTED_EDIT
	category = ADMINISTRATIVE
	to_email = UserEmailHandler(category, current_status, exec_review)
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
