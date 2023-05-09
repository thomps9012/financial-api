package responses

import (
	"financial-api/models"
)

type UsersInfoRes struct {
	Message string              `json:"message"`
	Data    []models.PublicInfo `json:"data"`
}
type UserInfoRes struct {
	Message string            `json:"message"`
	Data    models.PublicInfo `json:"data"`
}
type UsersNameRes struct {
	Message string                `json:"message"`
	Data    []models.UserNameInfo `json:"data"`
}
type UserNameRes struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func AllUsers(user_info []models.UserNameInfo) UsersNameRes {
	return UsersNameRes{
		Message: "All user data",
		Data:    user_info,
	}
}
func OneUser(user_info models.PublicInfo) UserInfoRes {
	return UserInfoRes{
		Message: "Overview information for " + user_info.ID,
		Data:    user_info,
	}
}
func UserName(user_info models.UserNameInfo) UserNameRes {
	return UserNameRes{
		Message: "Name information for " + user_info.ID,
		Data:    user_info.Name,
	}
}
func DeactivateUser(user_info models.PublicInfo) NilRes {
	return NilRes{
		Message: "User " + user_info.Name + "'s account has been deactivated",
		Data:    "null",
	}
}
func UserMileage(user_id string, user_info []models.Mileage_Overview) MileageOverviewsRes {
	return MileageOverviewsRes{
		Message: "Mileage Requests for User " + user_id,
		Data:    user_info,
	}
}
func UserPettyCash(user_id string, user_info []models.Petty_Cash_Overview) PettyCashOverviewsRes {
	return PettyCashOverviewsRes{
		Message: "Petty Cash Requests for User " + user_id,
		Data:    user_info,
	}
}
func UserCheckRequests(user_id string, user_info []models.Check_Request_Overview) CheckOverviewsRes {
	return CheckOverviewsRes{
		Message: "Check Requests for User " + user_id,
		Data:    user_info,
	}
}
