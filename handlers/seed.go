package handlers

import (
	"context"
	database "financial-api/db"
	"financial-api/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Seed data for the Action model
var actions = []models.Action{
	{
		ID:         "1",
		User:       "d160b410-e6a8-4cbb-92c2-068112187503",
		Status:     "APPROVED",
		Created_At: time.Now(),
	},
	{
		ID:         "2",
		User:       "c160b410-e6a8-4cbb-92c2-068112187612",
		Status:     "PENDING",
		Created_At: time.Now().Add(-24 * time.Hour),
	},
	{
		ID:         "3",
		User:       "d160b410-e6a8-4cbb-92c2-068112187305",
		Status:     "REJECTED",
		Created_At: time.Now().Add(-24 * time.Hour),
	},
}

// Seed data for the Mileage_Request model
var mileageRequests = []interface{}{
	models.Mileage_Request{
		ID:                      "c2e85479-827c-4030-80fa-fe0b657b26fa",
		Grant_ID:                "H79SP082264",
		User_ID:                 "c160b410-e6a8-4cbb-92c2-068112187612",
		Date:                    time.Now(),
		Category:                models.IOP,
		Starting_Location:       "Cleveland",
		Destination:             "Akron",
		Trip_Purpose:            "Meeting",
		Start_Odometer:          10000,
		End_Odometer:            10200,
		Tolls:                   5.50,
		Parking:                 10.00,
		Trip_Mileage:            200,
		Reimbursement:           150.00,
		Created_At:              time.Now(),
		Action_History:          []models.Action{actions[0]},
		Current_User:            "d160b410-e6a8-4cbb-92c2-068112187503",
		Current_Status:          "APPROVED",
		Last_User_Before_Reject: "",
		Is_Active:               true,
	},
	models.Mileage_Request{
		ID:                      "3015d932-1b43-467f-8cbc-8d687ed8ef81",
		Grant_ID:                "H79TI082369",
		User_ID:                 "c160b410-e6a8-4cbb-92c2-068112187612",
		Date:                    time.Now().Add(-48 * time.Hour),
		Category:                models.PERKINS,
		Starting_Location:       "Cleveland",
		Destination:             "Columbus",
		Trip_Purpose:            "Training",
		Start_Odometer:          15000,
		End_Odometer:            15200,
		Tolls:                   7.50,
		Parking:                 15.00,
		Trip_Mileage:            200,
		Reimbursement:           150.00,
		Created_At:              time.Now().Add(-48 * time.Hour),
		Action_History:          []models.Action{actions[1], actions[2]},
		Current_User:            "d160b410-e6a8-4cbb-92c2-068112187503",
		Current_Status:          "REJECTED",
		Last_User_Before_Reject: "d160b410-e6a8-4cbb-92c2-068112187305",
		Is_Active:               false,
	},
}

var grantSeeds = []interface{}{
	models.Grant{
		ID:   "H79TI082369",
		Name: "BCORR"},
	models.Grant{
		ID:   "H79SP082264",
		Name: "HIV Navigator"},
	models.Grant{
		ID:   "H79SP082475",
		Name: "SPF (HOPE 1)"},
	models.Grant{
		ID:   "SOR_PEER",
		Name: "SOR Peer"},
	models.Grant{
		ID:   "SOR_HOUSING",
		Name: "SOR Recovery Housing"},
	models.Grant{
		ID:   "SOR_TWR",
		Name: "SOR 2.0 - Together We Rise"},
	models.Grant{
		ID:   "TANF",
		Name: "TANF"},
	models.Grant{
		ID:   "2020-JY-FX-0014",
		Name: "JSBT (OJJDP) - Jumpstart For A Better Tomorrow"},
	models.Grant{
		ID:   "SOR_LORAIN",
		Name: "SOR Lorain 2.0"},
	models.Grant{
		ID:   "H79SP081048",
		Name: "STOP Grant"},
	models.Grant{
		ID:   "H79TI083370",
		Name: "BSW (Bridge to Success Workforce)"},
	models.Grant{
		ID:   "H79SM085150",
		Name: "CCBHC"},
	models.Grant{
		ID:   "H79TI083662",
		Name: "IOP New Syrenity Intensive outpatient Program"},
	models.Grant{
		ID:   "H79TI085495",
		Name: "RAP AID (Recover from Addition to Prevent Aids)"},
	models.Grant{
		ID:   "H79TI085410",
		Name: "N MAT (NORA Medication-Assisted Treatment Program)"},
}

var userSeeds = []interface{}{
	models.User{
		ID:         "d160b410-e6a8-4cbb-92c2-068112187503",
		Email:      "bob@example.com",
		Name:       "Bob Johnson",
		Last_Login: time.Now(),
		Vehicles: []models.Vehicle{
			{
				ID:          "1",
				Name:        "Honda Civic",
				Description: "2018 Model",
			},
			{
				ID:          "2",
				Name:        "Toyota Corolla",
				Description: "2017 Model",
			},
		},
		Is_Active:   true,
		Admin:       false,
		Permissions: []string{"EMPLOYEE"},
	},
	models.User{
		ID:         "c160b410-e6a8-4cbb-92c2-068112187612",
		Email:      "jane@example.com",
		Name:       "Jane Smith",
		Last_Login: time.Now(),
		Vehicles: []models.Vehicle{
			{
				ID:          "1",
				Name:        "Honda Civic",
				Description: "2018 Model",
			},
			{
				ID:          "2",
				Name:        "Toyota Corolla",
				Description: "2017 Model",
			},
		},
		Is_Active:   true,
		Admin:       true,
		Permissions: []string{"FINANCE_TEAM", "SUPERVISOR"},
	},
	models.User{
		ID:         "d160b410-e6a8-4cbb-92c2-068112187305",
		Email:      "john@example.com",
		Name:       "John Smith",
		Last_Login: time.Now(),
		Vehicles: []models.Vehicle{
			{
				ID:          "1",
				Name:        "Honda Civic",
				Description: "2018 Model",
			},
			{
				ID:          "2",
				Name:        "Toyota Corolla",
				Description: "2017 Model",
			},
		},
		Is_Active:   true,
		Admin:       true,
		Permissions: []string{"EXECUTIVE", "MANAGER"},
	},
}

var checkRequestSeeds = []interface{}{
	models.Check_Request{
		ID:       "cd304148-d143-4acc-a666-3854fd109e0f",
		Grant_ID: "2020-JY-FX-0014",
		User_ID:  "d160b410-e6a8-4cbb-92c2-068112187305",
		Date:     time.Now().Add(time.Hour * -72),
		Category: "ADMINISTRATIVE",
		Vendor: models.Vendor{
			Name:         "Office Depot",
			Website:      "www.test.com",
			AddressLine1: "123 st, Anytown, OH, 55555",
			AddressLine2: "",
		},
		Description: "Office Supplies for December",
		Purchases: []models.Purchase{
			{
				Description:     "Pens",
				Grant_Line_Item: "office supplies",
				Amount:          1.99,
			},
			{
				Description:     "Stapler",
				Grant_Line_Item: "office supplies",
				Amount:          9.99,
			},
		},
		Receipts: []string{
			"https://example.com/receipt1",
			"https://example.com/receipt2",
		},
		Order_Total: 13.97,
		Credit_Card: "**** **** **** 1234",
		Created_At:  time.Now(),
		Action_History: []models.Action{
			{
				ID:         "14d0dfae-8d43-48f3-8858-73c0098e8a14",
				User:       "d160b410-e6a8-4cbb-92c2-068112187305",
				Status:     "CREATED",
				Created_At: time.Now(),
			},
		},
		Current_User:            "c160b410-e6a8-4cbb-92c2-068112187612",
		Current_Status:          "PENDING",
		Last_User_Before_Reject: "null",
		Is_Active:               true,
	},
	models.Check_Request{
		ID:       "cd304148-d143-4acc-a666-3854fd109e0W",
		Grant_ID: "2020-JY-FX-0014",
		User_ID:  "d160b410-e6a8-4cbb-92c2-068112187503",
		Date:     time.Now().Add(time.Hour * -72),
		Category: "ADMINISTRATIVE",
		Vendor: models.Vendor{
			Name:         "Office Depot",
			Website:      "www.test.com",
			AddressLine1: "123 st, Anytown, OH, 55555",
			AddressLine2: "",
		},
		Description: "Office Supplies for December",
		Purchases: []models.Purchase{
			{
				Description:     "Pens",
				Grant_Line_Item: "office supplies",
				Amount:          1.99,
			},
			{
				Description:     "Stapler",
				Grant_Line_Item: "office supplies",
				Amount:          9.99,
			},
		},
		Receipts: []string{
			"https://example.com/receipt1",
			"https://example.com/receipt2",
		},
		Order_Total: 13.97,
		Credit_Card: "**** **** **** 1234",
		Created_At:  time.Now(),
		Action_History: []models.Action{
			{
				ID:         "14d0dfae-8d43-48f3-8858-73c0098e8a14",
				User:       "d160b410-e6a8-4cbb-92c2-068112187503",
				Status:     "CREATED",
				Created_At: time.Now(),
			},
			{
				ID:         "b2f12f44-66a0-45db-9646-708dfca7c9d7",
				User:       "d160b410-e6a8-4cbb-92c2-068112187305",
				Status:     "REJECTED",
				Created_At: time.Now(),
			},
		},
		Current_User:            "d160b410-e6a8-4cbb-92c2-068112187503",
		Current_Status:          "REJECTED",
		Last_User_Before_Reject: "d160b410-e6a8-4cbb-92c2-068112187305",
		Is_Active:               true,
	},
}

var pettyCashSeeds = []interface{}{
	models.Petty_Cash_Request{
		ID:          "81ebdc42-cd41-469f-a449-6ba30947f972",
		User_ID:     "c160b410-e6a8-4cbb-92c2-068112187612",
		Grant_ID:    "SOR_TWR",
		Category:    "ADMINISTRATIVE",
		Date:        time.Now().Add(99 * -time.Hour),
		Description: "Office supplies",
		Amount:      50.0,
		Receipts:    []string{"https://example.com/receipt1.jpg", "https://example.com/receipt2.jpg"},
		Created_At:  time.Now().Add(55 * -time.Hour),
		Action_History: []models.Action{
			{
				ID:         "90de4b14-503f-461b-9971-ccd778df0566",
				User:       "c160b410-e6a8-4cbb-92c2-068112187612",
				Status:     "CREATED",
				Created_At: time.Now().Add(55 * -time.Hour),
			},
			{
				ID:         "ddc15aee-5e79-4a47-86a4-1073ee93ea0b",
				User:       "d160b410-e6a8-4cbb-92c2-068112187305",
				Status:     "ORGANIZATION_APPROVED",
				Created_At: time.Now().Add(55 * -time.Hour),
			},
		},
		Current_User:            "c160b410-e6a8-4cbb-92c2-068112187612",
		Current_Status:          "ORGANIZATION_APPROVED",
		Last_User_Before_Reject: "null",
		Is_Active:               false,
	},
	models.Petty_Cash_Request{
		ID:          "81ebdc42-cd41-469f-a449-6ba30947f973",
		User_ID:     "d160b410-e6a8-4cbb-92c2-068112187503",
		Grant_ID:    "SOR_LORAIN",
		Category:    "ADMINISTRATIVE",
		Date:        time.Now().Add(17 * -time.Hour),
		Description: "Printer ink",
		Amount:      35.0,
		Receipts:    []string{"https://example.com/receipt1.jpg", "https://example.com/receipt2.jpg"},
		Created_At:  time.Now().Add(5 * -time.Hour),
		Action_History: []models.Action{
			{
				ID:         "90de4b14-503f-461b-9971-ccd778df0566",
				User:       "d160b410-e6a8-4cbb-92c2-068112187503",
				Status:     "CREATED",
				Created_At: time.Now().Add(5 * -time.Hour),
			},
		},
		Current_User:            "d160b410-e6a8-4cbb-92c2-068112187305",
		Current_Status:          "PENDING",
		Last_User_Before_Reject: "null",
		Is_Active:               false,
	},
}

func SeedData(c *fiber.Ctx) error {
	grants, err := database.Use("grants")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	g_res, err := grants.InsertMany(context.TODO(), grantSeeds)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	users, err := database.Use("users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	u_res, err := users.InsertMany(context.TODO(), userSeeds)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	mileage, err := database.Use("mileage_requests")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	m_res, err := mileage.InsertMany(context.TODO(), mileageRequests)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	check_requests, err := database.Use("check_requests")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	c_res, err := check_requests.InsertMany(context.TODO(), checkRequestSeeds)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	petty_cash, err := database.Use("petty_cash_requests")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	p_res, err := petty_cash.InsertMany(context.TODO(), pettyCashSeeds)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"grants":     g_res,
		"users":      u_res,
		"mileage":    m_res,
		"check":      c_res,
		"petty_cash": p_res,
	})
}

func DeleteSeeds(c *fiber.Ctx) error {
	grants, err := database.Use("grants")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	g_res, err := grants.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	users, err := database.Use("users")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	u_res, err := users.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	mileage, err := database.Use("mileage_requests")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	m_res, err := mileage.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	check_requests, err := database.Use("check_requests")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	c_res, err := check_requests.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	petty_cash, err := database.Use("petty_cash_requests")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	p_res, err := petty_cash.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}
	return c.Status(http.StatusAccepted).JSON(fiber.Map{
		"grants":     g_res,
		"users":      u_res,
		"mileage":    m_res,
		"check":      c_res,
		"petty_cash": p_res,
	})
}
