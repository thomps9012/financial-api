package methods

import (
	"strings"
)

type CurrentUser struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

// iop manager
var iop_manager = CurrentUser{
	ID:   "100323196510465985735",
	Name: "Intensive Outpatient Manager",
}

// intake manager
var intake_manager = CurrentUser{
	ID:   "114839465093091142050",
	Name: "Intake Manager",
}

// peer manager
var peer_manager = CurrentUser{
	ID:   "100323196510465985735",
	Name: "Peer Support Manager",
}

// act manager
var act_manager = CurrentUser{
	ID:   "109904620612313503312",
	Name: "ACT Manager",
}

// ihbt manager
var ihbt_manager = CurrentUser{
	ID:   "111724747499299581748",
	Name: "IHBT Manager",
}

// perkins manager
var perkins_manager = CurrentUser{
	ID:   "101753058297288934715",
	Name: "Perkins Manager",
}

// perkins supervisor
var perkins_supervisor = CurrentUser{
	ID:   "100323196510465985735",
	Name: "Perkins Supervisor",
}

// next step manager
var next_step_manager = CurrentUser{
	ID:   "104213502974305953852",
	Name: "NEXT Step Manager",
}

// next step supervisor
var next_step_supervisor = CurrentUser{
	ID:   "114839465093091142050",
	Name: "NEXT Step Supervisor",
}

// lorain manager
var lorain_manager = CurrentUser{
	ID:   "102294853567984211361",
	Name: "Lorain Manager",
}

// lorain supervisor
var lorain_supervisor = CurrentUser{
	ID:   "100323196510465985735",
	Name: "Lorain Supervisor",
}

// prevention manager
var prevention_manager = CurrentUser{
	ID:   "106544031410769042454",
	Name: "Prevention Manager",
}

// prevention supervisor
var prevention_supervisor = CurrentUser{
	ID:   "114839465093091142050",
	Name: "Prevention Supervisor",
}

// men's house manager
var mens_house_manager = CurrentUser{
	ID:   "100323196510465985735",
	Name: "Men's House Manager",
}

// administrative manager
var admin_manager = CurrentUser{
	ID:   "111724747499299581748",
	Name: "Administrative Manager",
}

// finance manager
var finance_manager = CurrentUser{
	ID:   "111876803051097580983",
	Name: "Finance Manager",
}

// finance supervisor
var finance_supervisor = CurrentUser{
	ID:   "516510984003210861035",
	Name: "Finance Supervisor",
}

// finance fulfillment
var finance_fulfillment = CurrentUser{
	ID:   "109157735191825776845",
	Name: "Finance Requests",
}

func NewRequestUser(request_type string, request_category string, user_id string) CurrentUser {
	if strings.ToLower(strings.TrimSpace(request_type)) == "mileage" {
		if user_id == "" {
			return finance_fulfillment
		}
		return finance_supervisor
	} else {
		current_user := NewRequestHandler(request_category, user_id)
		return current_user
	}
}

func NewRequestHandler(category string, user_id string) CurrentUser {
	switch category {
	case "ADMINISTRATIVE":
		if user_id == admin_manager.ID {
			return finance_supervisor
		}
		return admin_manager
	case "IHBT":
		if user_id == ihbt_manager.ID {
			return finance_supervisor
		}
		return ihbt_manager
	case "IOP":
		if user_id == iop_manager.ID {
			return finance_supervisor
		}
		return iop_manager
	case "PEERS":
		if user_id == peer_manager.ID {
			return finance_supervisor
		}
		return peer_manager
	case "MENS_HOUSE":
		if user_id == mens_house_manager.ID {
			return finance_supervisor
		}
		return mens_house_manager
	case "INTAKE":
		if user_id == intake_manager.ID {
			return finance_supervisor
		}
		return ihbt_manager
	case "ACT_TEAM":
		if user_id == act_manager.ID {
			return finance_supervisor
		}
		return act_manager
	case "FINANCE":
		if user_id == finance_supervisor.ID {
			return finance_fulfillment
		}
		if user_id == finance_manager.ID {
			return finance_supervisor
		}
		return finance_manager
	case "LORAIN":
		if user_id == lorain_supervisor.ID {
			return finance_supervisor
		}
		if user_id == lorain_manager.ID {
			return lorain_supervisor
		}
		return lorain_manager
	case "NEXT_STEP":
		if user_id == next_step_supervisor.ID {
			return finance_supervisor
		}
		if user_id == next_step_manager.ID {
			return next_step_supervisor
		}
		return next_step_manager
	case "PERKINS":
		if user_id == perkins_supervisor.ID {
			return finance_supervisor
		}
		if user_id == perkins_manager.ID {
			return perkins_supervisor
		}
		return perkins_manager
	case "PREVENTION":
		if user_id == prevention_supervisor.ID {
			return finance_supervisor
		}
		if user_id == prevention_manager.ID {
			return prevention_supervisor
		}
		return prevention_manager
	default:
		if user_id == finance_supervisor.ID {
			return finance_fulfillment
		}
		return finance_supervisor
	}
}
