package methods

import (
	"strings"
)

type CurrentUser struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

var finance_user = CurrentUser{
	ID:   "109157735191825776845",
	Name: "Finance Requests",
}

func NewRequestUser(request_type string, request_category string, user_id string) CurrentUser {
	if strings.ToLower(strings.TrimSpace(request_type)) == "mileage" {
		if user_id == "" {
			return CurrentUser{
				ID:   "...id",
				Name: "Stephanie Bryant",
			}
		}
		return finance_user
	} else {
		current_user := NewRequestHandler(request_category, user_id)
		return current_user
	}
}

func NewRequestHandler(category string, user_id string) CurrentUser {
	switch category {
	case "ADMINISTRATIVE":
	case "IHBT":
		if user_id == "111724747499299581748" {
			return finance_user
		}
		return CurrentUser{
			ID:   "111724747499299581748",
			Name: "Bianca Griffin",
		}
	case "IOP":
	case "PEERS":
	case "MENS_HOUSE":
		if user_id == "100323196510465985735" {
			return finance_user
		}
		return CurrentUser{
			ID:   "100323196510465985735",
			Name: "Jeff Ward",
		}
	case "INTAKE":
		if user_id == "114839465093091142050" {
			return finance_user
		}
		return CurrentUser{
			ID:   "114839465093091142050",
			Name: "Cynthia Woods",
		}
	case "ACT_TEAM":
		if user_id == "109904620612313503312" {
			return finance_user
		}
		return CurrentUser{
			ID:   "109904620612313503312",
			Name: "Joyce Jordan",
		}
	case "FINANCE":
		if user_id == "111876803051097580983" {
			return finance_user
		}
		return CurrentUser{
			ID:   "111876803051097580983",
			Name: "Lisa Fuentes",
		}
	case "LORAIN":
		if user_id == "100323196510465985735" {
			return finance_user
		}
		if user_id == "102294853567984211361" {
			return CurrentUser{
				ID:   "100323196510465985735",
				Name: "Jeff Ward",
			}
		}
		return CurrentUser{
			ID:   "102294853567984211361",
			Name: "Ron Giusti",
		}
	case "NEXT_STEP":
		if user_id == "114839465093091142050" {
			return finance_user
		}
		if user_id == "104213502974305953852" {
			return CurrentUser{
				ID:   "114839465093091142050",
				Name: "Cynthia Woods",
			}
		}
		return CurrentUser{
			ID:   "104213502974305953852",
			Name: "Deborah Baker",
		}
	case "PERKINS":
		if user_id == "100323196510465985735" {
			return finance_user
		}
		if user_id == "101753058297288934715" {
			return CurrentUser{
				ID:   "100323196510465985735",
				Name: "Jeff Ward",
			}
		}
		return CurrentUser{
			ID:   "101753058297288934715",
			Name: "Charlesetta Hurt",
		}
	case "PREVENTION":
		if user_id == "114839465093091142050" {
			return finance_user
		}
		if user_id == "106544031410769042454" {
			return CurrentUser{
				ID:   "114839465093091142050",
				Name: "Cynthia Woods",
			}
		}
		return CurrentUser{
			ID:   "106544031410769042454",
			Name: "Lauretta Amanor",
		}
	default:
		if user_id == "109157735191825776845" {
			return CurrentUser{
				ID:   "...id",
				Name: "Stephanie Bryant",
			}
		}
		return finance_user
	}
	return finance_user
}
