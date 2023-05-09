package responses

import (
	"financial-api/models"
)

type MyInfoRes struct {
	Message string            `json:"message"`
	Data    models.PublicInfo `json:"data"`
}

type VehicleRes struct {
	Message string         `json:"message"`
	Data    models.Vehicle `json:"data"`
}

func MyInfo(my_info models.PublicInfo) MyInfoRes {
	return MyInfoRes{
		Message: "Found the following request and user info for " + my_info.Name,
		Data:    my_info,
	}
}
func MyMileage(requests []models.Mileage_Request) MileagesRes {
	return MileagesRes{
		Message: "Your mileage request information",
		Data:    requests,
	}
}
func MyCheckRequests(requests []models.Check_Request) CheckRequestsRes {
	return CheckRequestsRes{
		Message: "Your check request information",
		Data:    requests,
	}
}
func MyPettyCash(requests []models.Petty_Cash_Request) PettyCashRequestsRes {
	return PettyCashRequestsRes{
		Message: "Your petty cash requests",
		Data:    requests,
	}
}
func AddVehicle(vehicle models.Vehicle) VehicleRes {
	return VehicleRes{
		Message: "New vehicle added to your account",
		Data:    vehicle,
	}
}
func EditVehicle(vehicle models.Vehicle) VehicleRes {
	return VehicleRes{
		Message: "Vehicle " + vehicle.ID + " successfully saved",
		Data:    vehicle,
	}
}
func RemoveVehicle(vehicle models.Vehicle) VehicleRes {
	return VehicleRes{
		Message: "Vehicle " + vehicle.Name + " removed from account",
		Data:    vehicle,
	}
}
