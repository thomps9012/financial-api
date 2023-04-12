package responses

import (
	"financial-api/models"
)

type MyInfoRes struct {
	Status  string            `json:"status"`
	Code    int               `json:"code"`
	Message string            `json:"string"`
	Data    models.PublicInfo `json:"data"`
}

type VehicleRes struct {
	Status  string         `json:"status"`
	Code    int            `json:"code"`
	Message string         `json:"string"`
	Data    models.Vehicle `json:"data"`
}

func MyInfo(my_info models.PublicInfo) MyInfoRes {
	return MyInfoRes{
		Status:  "OK",
		Code:    200,
		Message: "Found the following request and user info for " + my_info.Name,
		Data:    my_info,
	}
}
func MyMileage(requests []models.Mileage_Request) MileagesRes {
	return MileagesRes{
		Status:  "OK",
		Code:    200,
		Message: "Your mileage request information",
		Data:    requests,
	}
}
func MyCheckRequests(requests []models.Check_Request) CheckRequestsRes {
	return CheckRequestsRes{
		Status:  "OK",
		Code:    200,
		Message: "Your check request information",
		Data:    requests,
	}
}
func MyPettyCash(requests []models.Petty_Cash_Request) PettyCashRequestsRes {
	return PettyCashRequestsRes{
		Status:  "OK",
		Code:    200,
		Message: "You petty cash requests",
		Data:    requests,
	}
}
func AddVehicle(vehicle models.Vehicle) VehicleRes {
	return VehicleRes{
		Status:  "CREATED",
		Code:    201,
		Message: "New vehicle added to your account",
		Data:    vehicle,
	}
}
func EditVehicle(vehicle models.Vehicle) VehicleRes {
	return VehicleRes{
		Status:  "OK",
		Code:    200,
		Message: "Vehicle " + vehicle.ID + " successfully saved",
		Data:    vehicle,
	}
}
func RemoveVehicle(vehicle models.Vehicle) VehicleRes {
	return VehicleRes{
		Status:  "OK",
		Code:    200,
		Message: "Vehicle " + vehicle.Name + " removed from account",
		Data:    vehicle,
	}
}
