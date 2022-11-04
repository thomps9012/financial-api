package models

type Status string

const (
	CREATED               Status = "CREATED"
	PENDING               Status = "PENDING"
	MANAGER_APPROVED      Status = "MANAGER_APPROVED"
	SUPERVISOR_APPROVED   Status = "SUPERVISOR_APPROVED"
	FINANCE_APPROVED      Status = "FINANCE_APPROVED"
	EXECUTIVE_APPROVED    Status = "EXECUTIVE_APPROVED"
	ORGANIZATION_APPROVED Status = "ORGANIZATION_APPROVED"
	REJECTED              Status = "REJECTED"
	REJECTED_EDIT         Status = "REJECTED_EDIT"
	EDIT                  Status = "EDIT"
	ARCHIVED              Status = "ARCHIVED"
)

type Category string

const (
	FINANCE        Category = "FINANCE"
	HOUSING        Category = "HOUSING"
	ADMINISTRATIVE Category = "ADMINISTRATIVE"
	FESTIVUS       Category = "FESTIVUS"
	TRANSPORTATION Category = "TRANSPORTATION"
	DEPARTMENT_1   Category = "DEPARTMENT_1"
	DEPARTMENT_2   Category = "DEPARTMENT_2"
	DEPARTMENT_3   Category = "DEPARTMENT_3"
)

type Request_Type string

const (
	MILEAGE    Request_Type = "MILEAGE"
	CHECK      Request_Type = "CHECK"
	PETTY_CASH Request_Type = "PETTY_CASH"
)

type Request_Response struct {
	User_ID        string `json:"user_id" bson:"user_id"`
	Current_Status string
	Success        bool
}

type Request_Info struct {
	User_ID        string `json:"user_id" bson:"user_id"`
	Current_User   string `json:"current_user" bson:"current_user"`
	Current_Status Status `json:"current_status" bson:"current_status"`
}

func UserEmailHandler(category Category, current_status Status, exec_review bool) string {
	if exec_review || current_status == FINANCE_APPROVED {
		return "test_exec@finance.com"
	} else if current_status == SUPERVISOR_APPROVED || current_status == EXECUTIVE_APPROVED {
		return "test_finance@finance.com"
	} else if current_status == MANAGER_APPROVED {
		return "test_supervisor@finance.com"
	} else if current_status == PENDING {
		return "test_manager@finance.com"
	} else if current_status == REJECTED || current_status == ORGANIZATION_APPROVED {
		return ""
	} else {
		return ""
	}
}
