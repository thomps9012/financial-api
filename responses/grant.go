package responses

import (
	"financial-api/models"
)

type GrantRes struct {
	Message string       `json:"message"`
	Data    models.Grant `json:"data"`
}
type GrantsRes struct {
	Message string         `json:"message"`
	Data    []models.Grant `json:"data"`
}

func AllGrants(data []models.Grant) GrantsRes {
	return GrantsRes{
		Message: "All Grants Data",
		Data:    data,
	}
}
func OneGrant(data models.Grant) GrantRes {
	return GrantRes{
		Message: "Grant Data Found for Grant: " + data.Name,
		Data:    data,
	}
}
func GrantCheckRequests(grant models.Grant, data []models.Check_Request_Overview) CheckOverviewsRes {
	return CheckOverviewsRes{
		Message: "Check Requests for Grant " + grant.Name + " Found",
		Data:    data,
	}
}
func GrantMileage(grant models.Grant, data []models.Mileage_Overview) MileageOverviewsRes {
	return MileageOverviewsRes{
		Message: "Mileage Requests for Grant " + grant.Name + " Found",
		Data:    data,
	}
}
func GrantPettyCash(grant models.Grant, data []models.Petty_Cash_Overview) PettyCashOverviewsRes {
	return PettyCashOverviewsRes{
		Message: "Petty Cash Requests for Grant " + grant.Name + " Found",
		Data:    data,
	}
}
