package responses

import (
	"financial-api/models"
)

type UsersInfoRes struct {
	Status  string              `json:"status"`
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    []models.PublicInfo `json:"data"`
}
type UserInfoRes struct {
	Status  string            `json:"status"`
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    models.PublicInfo `json:"data"`
}
type UsersNameRes struct {
	Status  string                `json:"status"`
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    []models.UserNameInfo `json:"data"`
}

func AllUsers(user_info []models.UserNameInfo) UsersNameRes {
	return UsersNameRes{
		Status:  "OK",
		Code:    200,
		Message: "All user data",
		Data:    user_info,
	}
}
func OneUser(user_info models.PublicInfo) UserInfoRes {
	return UserInfoRes{
		Status:  "OK",
		Code:    200,
		Message: "Overview information for " + user_info.ID,
		Data:    user_info,
	}
}
func DeactivateUser(user_info models.PublicInfo) NilRes {
	return NilRes{
		Status:  "OK",
		Code:    200,
		Message: "User " + user_info.Name + "'s account has been deactivated",
		Data:    "null",
	}
}
func UserMileage(user_id string, user_info []models.Mileage_Overview) MileageOverviewsRes {
	return MileageOverviewsRes{
		Status:  "OK",
		Code:    200,
		Message: "Mileage Requests for User " + user_id,
		Data:    user_info,
	}
}
func UserPettyCash(user_id string, user_info []models.Petty_Cash_Overview) PettyCashOverviewsRes {
	return PettyCashOverviewsRes{
		Status:  "OK",
		Code:    200,
		Message: "Petty Cash Requests for User " + user_id,
		Data:    user_info,
	}
}
func UserCheckRequests(user_id string, user_info []models.Check_Request_Overview) CheckOverviewsRes {
	return CheckOverviewsRes{
		Status:  "OK",
		Code:    200,
		Message: "Check Requests for User " + user_id,
		Data:    user_info,
	}
}
