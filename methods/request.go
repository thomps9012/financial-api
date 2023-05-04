package methods

import (
	"strings"
)

type CurrentUser struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

// ****CURRENT IDS ARE FOR SEED PURPOSES ONLY*****

// iop manager
var IOP_MANAGER = CurrentUser{
	ID: "c160b410-e6a8-4cbb-92c2-068112187612",
	// ID:   "100323196510465985735",
	Name: "Intensive Outpatient Manager",
}

// intake manager
var INTAKE_MANAGER = CurrentUser{
	ID: "c160b410-e6a8-4cbb-92c2-068112187612",
	// ID:   "114839465093091142050",
	Name: "Intake Manager",
}

// peer manager
var PEER_MANAGER = CurrentUser{
	ID: "c160b410-e6a8-4cbb-92c2-068112187612",
	// ID:   "100323196510465985735",
	Name: "Peer Support Manager",
}

// act manager
var ACT_MANAGER = CurrentUser{
	ID: "c160b410-e6a8-4cbb-92c2-068112187612",
	// ID:   "109904620612313503312",
	Name: "ACT Manager",
}

// ihbt manager
var IHBT_MANAGER = CurrentUser{
	ID: "c160b410-e6a8-4cbb-92c2-068112187612",
	// ID:   "111724747499299581748",
	Name: "IHBT Manager",
}

// perkins manager
var PERKINS_MANAGER = CurrentUser{
	ID: "c160b410-e6a8-4cbb-92c2-068112187612",
	// ID:   "101753058297288934715",
	Name: "Perkins Manager",
}

// perkins supervisor
var PERKINS_SUPERVISOR = CurrentUser{
	ID: "c160b410-e6a8-4cbb-92c2-068112187612",
	// ID:   "100323196510465985735",
	Name: "Perkins Supervisor",
}

// next step manager
var NEXT_STEP_MANAGER = CurrentUser{
	ID: "c160b410-e6a8-4cbb-92c2-068112187612",
	// ID:   "104213502974305953852",
	Name: "NEXT Step Manager",
}

// next step supervisor
var NEXT_STEP_SUPERVISOR = CurrentUser{
	ID: "d160b410-e6a8-4cbb-92c2-068112187305",
	// ID:   "114839465093091142050",
	Name: "NEXT Step Supervisor",
}

// lorain manager
var LORAIN_MANAGER = CurrentUser{
	ID: "c160b410-e6a8-4cbb-92c2-068112187612",
	// ID:   "102294853567984211361",
	Name: "Lorain Manager",
}

// lorain supervisor
var LORAIN_SUPERVISOR = CurrentUser{
	ID: "d160b410-e6a8-4cbb-92c2-068112187305",
	// ID:   "100323196510465985735",
	Name: "Lorain Supervisor",
}

// prevention manager
var PREVENTION_MANAGER = CurrentUser{
	ID: "d160b410-e6a8-4cbb-92c2-068112187305",
	// ID:   "106544031410769042454",
	Name: "Prevention Manager",
}

// prevention supervisor
var PREVENTION_SUPERVISOR = CurrentUser{
	ID: "d160b410-e6a8-4cbb-92c2-068112187305",
	// ID:   "114839465093091142050",
	Name: "Prevention Supervisor",
}

// men's house manager
var MENS_HOUSE_MANAGER = CurrentUser{
	ID: "d160b410-e6a8-4cbb-92c2-068112187305",
	// ID:   "100323196510465985735",
	Name: "Men's House Manager",
}

// administrative supervisor
var ADMINISTRATIVE_SUPERVISOR = CurrentUser{
	ID: "d160b410-e6a8-4cbb-92c2-068112187305",
	// ID:   "111724747499299581748",
	Name: "Administrative Manager",
}

// finance manager
var FINANCE_MANAGER = CurrentUser{
	ID: "d160b410-e6a8-4cbb-92c2-068112187305",
	// ID:   "111876803051097580983",
	Name: "Finance Manager",
}

// finance supervisor
var FINANCE_SUPERVISOR = CurrentUser{
	ID: "0d1ee9e2-dbe3-4a2a-b9cf-1ff27ce3a500",
	// ID:   "516510984003210861035",
	Name: "Finance Supervisor",
}

// finance fulfillment
var FINANCE_FULFILLMENT = CurrentUser{
	ID: "2e780f36-7829-4707-9a17-34fce224c53e",
	// ID:   "109157735191825776845",
	Name: "Finance Requests",
}

var END_USER = CurrentUser{
	ID:   "null",
	Name: "null",
}

func NewRequestUser(request_type string, request_category string, user_id string) CurrentUser {
	if strings.ToLower(strings.TrimSpace(request_type)) == "mileage" {
		if user_id == "" || user_id == FINANCE_SUPERVISOR.ID {
			return FINANCE_FULFILLMENT
		}
		return FINANCE_SUPERVISOR
	} else {
		current_user := NewRequestHandler(request_category, user_id)
		return current_user
	}
}

func NewRequestHandler(category string, user_id string) CurrentUser {
	switch category {
	case "ADMINISTRATIVE":
		if user_id == ADMINISTRATIVE_SUPERVISOR.ID {
			return FINANCE_SUPERVISOR
		}
		return ADMINISTRATIVE_SUPERVISOR
	case "IHBT":
		if user_id == IHBT_MANAGER.ID {
			return FINANCE_SUPERVISOR
		}
		return IHBT_MANAGER
	case "IOP":
		if user_id == IOP_MANAGER.ID {
			return FINANCE_SUPERVISOR
		}
		return IOP_MANAGER
	case "PEERS":
		if user_id == PEER_MANAGER.ID {
			return FINANCE_SUPERVISOR
		}
		return PEER_MANAGER
	case "MENS_HOUSE":
		if user_id == MENS_HOUSE_MANAGER.ID {
			return FINANCE_SUPERVISOR
		}
		return MENS_HOUSE_MANAGER
	case "INTAKE":
		if user_id == INTAKE_MANAGER.ID {
			return FINANCE_SUPERVISOR
		}
		return IHBT_MANAGER
	case "ACT_TEAM":
		if user_id == ACT_MANAGER.ID {
			return FINANCE_SUPERVISOR
		}
		return ACT_MANAGER
	case "FINANCE":
		if user_id == FINANCE_SUPERVISOR.ID {
			return FINANCE_FULFILLMENT
		}
		if user_id == FINANCE_MANAGER.ID {
			return FINANCE_SUPERVISOR
		}
		return FINANCE_MANAGER
	case "LORAIN":
		if user_id == LORAIN_SUPERVISOR.ID {
			return FINANCE_SUPERVISOR
		}
		if user_id == LORAIN_MANAGER.ID {
			return LORAIN_SUPERVISOR
		}
		return LORAIN_MANAGER
	case "NEXT_STEP":
		if user_id == NEXT_STEP_SUPERVISOR.ID {
			return FINANCE_SUPERVISOR
		}
		if user_id == NEXT_STEP_MANAGER.ID {
			return NEXT_STEP_SUPERVISOR
		}
		return NEXT_STEP_MANAGER
	case "PERKINS":
		if user_id == PERKINS_SUPERVISOR.ID {
			return FINANCE_SUPERVISOR
		}
		if user_id == PERKINS_MANAGER.ID {
			return PERKINS_SUPERVISOR
		}
		return PERKINS_MANAGER
	case "PREVENTION":
		if user_id == PREVENTION_SUPERVISOR.ID {
			return FINANCE_SUPERVISOR
		}
		if user_id == PREVENTION_MANAGER.ID {
			return PREVENTION_SUPERVISOR
		}
		return PREVENTION_MANAGER
	default:
		if user_id == FINANCE_SUPERVISOR.ID {
			return FINANCE_FULFILLMENT
		}
		return FINANCE_SUPERVISOR
	}
}
