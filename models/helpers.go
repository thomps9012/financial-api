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
	IOP            Category = "IOP"
	INTAKE         Category = "INTAKE"
	PEERS          Category = "PEERS"
	ACT_TEAM       Category = "ACT_TEAM"
	IHBT           Category = "IHBT"
	PERKINS        Category = "PERKINS"
	MENS_HOUSE     Category = "MENS_HOUSE"
	NEXT_STEP      Category = "NEXT_STEP"
	LORAIN         Category = "LORAIN"
	PREVENTION     Category = "PREVENTION"
	ADMINISTRATIVE Category = "ADMINISTRATIVE"
	FINANCE        Category = "FINANCE"
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
	// possible more build out of test scenarios here
	var to_email = ""
	if exec_review {
		to_email = "abradley@norainc.org"
	} else if current_status == REJECTED || current_status == ORGANIZATION_APPROVED || current_status == FINANCE_APPROVED {
		to_email = ""
	} else if current_status == SUPERVISOR_APPROVED {
		to_email = "finance_requests@norainc.org"
	} else if current_status == MANAGER_APPROVED {
		switch category {
		case LORAIN:
			to_email = "jward@norainc.org"
		case NEXT_STEP:
			to_email = "cwoods@norainc.org"
		case PERKINS:
			to_email = "jward@norainc.org"
		case PREVENTION:
			to_email = "cwoods@norainc.org"
		default:
			to_email = "finance_requests@norainc.org"
		}
	} else {
		switch category {
		case ADMINISTRATIVE:
			to_email = "bgriffin@norainc.org"
		case IOP:
			to_email = "jward@norainc.org"
		case INTAKE:
			to_email = "cwoods@norainc.org"
		case PEERS:
			to_email = "jward@norainc.org"
		case ACT_TEAM:
			to_email = "jjordan@norainc.org"
		case IHBT:
			to_email = "bgriffin@norainc.org"
		case FINANCE:
			to_email = "lfuentes@norainc.org"
		case LORAIN:
			to_email = "rgiusti@norainc.org"
		case MENS_HOUSE:
			to_email = "jward@norainc.org"
		case NEXT_STEP:
			to_email = "dbaker@norainc.org"
		case PERKINS:
			to_email = "churt@norainc.org"
		case PREVENTION:
			to_email = "lamanor@norainc.org"
		default:
			to_email = "finance_requests@norainc.org"
		}
	}
	return to_email
}
