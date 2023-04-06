package methods

import (
	"strings"
)

type CurrentUser struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

func NewRequestUser(request_type string, request_category string) CurrentUser {
	if strings.ToLower(strings.TrimSpace(request_type)) == "mileage" {
		return CurrentUser{
			ID:   "109157735191825776845",
			Name: "Finance Requests",
		}
	} else {
		current_user := NewRequestHandler(request_category)
		return current_user
	}
}

func NewRequestHandler(category string) CurrentUser {
	switch category {
	case "ADMINISTRATIVE":
	case "IHBT":
		return CurrentUser{
			ID:   "111724747499299581748",
			Name: "Bianca Griffin",
		}
	case "IOP":
	case "PEERS":
	case "MENS_HOUSE":
		return CurrentUser{
			ID:   "100323196510465985735",
			Name: "Jeff Ward",
		}
	case "INTAKE":
		return CurrentUser{
			ID:   "114839465093091142050",
			Name: "Cynthia Woods",
		}
	case "ACT_TEAM":
		return CurrentUser{
			ID:   "109904620612313503312",
			Name: "Joyce Jordan",
		}
	case "FINANCE":
		return CurrentUser{
			ID:   "111876803051097580983",
			Name: "Lisa Fuentes",
		}
	case "LORAIN":
		return CurrentUser{
			ID:   "102294853567984211361",
			Name: "Ron Giusti",
		}
	case "NEXT_STEP":
		return CurrentUser{
			ID:   "104213502974305953852",
			Name: "Deborah Baker",
		}
	case "PERKINS":
		return CurrentUser{
			ID:   "101753058297288934715",
			Name: "Charlesetta Hurt",
		}
	case "PREVENTION":
		return CurrentUser{
			ID:   "106544031410769042454",
			Name: "Lauretta Amanor",
		}
	default:
		return CurrentUser{
			ID:   "109157735191825776845",
			Name: "Finance Requests",
		}
	}
	return CurrentUser{
		ID:   "109157735191825776845",
		Name: "Finance Requests",
	}
}
