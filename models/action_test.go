package models

import (
	"testing"
)

func TestCreateAction(t *testing.T) {
	var new_status = MANAGER_APPROVED
	var request_info = Request_Info{
		ID:           "111",
		Type:         MILEAGE,
		Current_User: "1",
	}
	var action Action
	new_action := action.Create(new_status, request_info)
	if new_status != new_action.Status {
		t.Errorf("Test Create Action Failed: expected %s, got: %s", new_status, new_action.Status)
	}
}

func TestGetAction(t *testing.T) {
	type test_case struct {
		request_id   string
		request_type Request_Type
		description  string
		expected     Request_Info
	}
	test_cases := []test_case{
		{
			request_id:   "5d308c52-21e2-43c6-b2eb-97d82e926294",
			request_type: MILEAGE,
			description:  "mileage test request",
			expected: Request_Info{
				ID:             "5d308c52-21e2-43c6-b2eb-97d82e926294",
				Type:           MILEAGE,
				Current_User:   "a651df6a5s1-65as1df6sda1f165ds-65asdf1",
				Current_Status: PENDING,
				User_ID:        "798b8670-3a6a-9f62-3267f732a5a5",
			},
		},
	}
	var action Action
	for _, test := range test_cases {
		t.Run(test.description, func(t *testing.T) {
			actual, _ := action.Get(test.request_id, test.request_type)
			if actual != test.expected {
				t.Errorf(test.description+" Failed: expected %s, got: %s", test.expected, actual)
			}
		})
	}
}
