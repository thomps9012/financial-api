package handlers

import (
	"context"
	database "financial-api/db"
	"financial-api/models"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

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
		Permissions: []string{"NEXT_STEP_MANAGER", "PERKINS_SUPERVISOR", "PERKINS_MANAGER", "IHBT_MANAGER", "ACT_MANAGER", "PEER_SUPPORT_MANAGER", "INTAKE_MANAGER", "IOP_MANAGER"},
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
		Permissions: []string{"FINANCE_MANAGER", "ADMIN_MANAGER", "MENS_HOUSE_MANAGER", "PREVENTION_SUPERVISOR", "PREVENTION_MANAGER", "LORAIN_SUPERVISOR", "LORAIN_MANAGER", "NEXT_STEP_SUPERVISOR"},
	},
	models.User{
		ID:         "0d1ee9e2-dbe3-4a2a-b9cf-1ff27ce3a500",
		Email:      "hford@example.com",
		Name:       "Harrison Ford",
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
		Permissions: []string{"FINANCE_SUPERVISOR"},
	},
	models.User{
		ID:         "2e780f36-7829-4707-9a17-34fce224c53e",
		Email:      "chewy@example.com",
		Name:       "Chewbacca",
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
		Permissions: []string{"FINANCE_FULFILLMENT"},
	},
}

var mileageRequests = []interface{}{
	models.Mileage_Request{
		ID:                "c2e85479-827c-4030-80fa-fe0b657b26fa",
		Grant_ID:          "H79SP082264",
		User_ID:           "c160b410-e6a8-4cbb-92c2-068112187612",
		Date:              time.Now(),
		Category:          models.IOP,
		Starting_Location: "Cleveland",
		Destination:       "Akron",
		Trip_Purpose:      "Meeting",
		Start_Odometer:    10000,
		End_Odometer:      10200,
		Tolls:             5.50,
		Parking:           10.00,
		Trip_Mileage:      200,
		Reimbursement:     150.00,
		Created_At:        time.Now(),
		Action_History: []models.Action{{
			ID:         uuid.NewString(),
			User:       "c160b410-e6a8-4cbb-92c2-068112187612",
			Status:     "CREATED",
			Created_At: time.Now(),
		}},
		Current_User:            "0d1ee9e2-dbe3-4a2a-b9cf-1ff27ce3a500",
		Current_Status:          "PENDING",
		Last_User_Before_Reject: "",
		Is_Active:               true,
	},
	models.Mileage_Request{
		ID:                "3015d932-1b43-467f-8cbc-8d687ed8ef81",
		Grant_ID:          "H79TI082369",
		User_ID:           "c160b410-e6a8-4cbb-92c2-068112187612",
		Date:              time.Now().Add(-48 * time.Hour),
		Category:          models.PERKINS,
		Starting_Location: "Cleveland",
		Destination:       "Columbus",
		Trip_Purpose:      "Training",
		Start_Odometer:    15000,
		End_Odometer:      15200,
		Tolls:             7.50,
		Parking:           15.00,
		Trip_Mileage:      200,
		Reimbursement:     150.00,
		Created_At:        time.Now().Add(-48 * time.Hour),
		Action_History: []models.Action{{
			ID:         uuid.NewString(),
			User:       "c160b410-e6a8-4cbb-92c2-068112187612",
			Status:     "CREATED",
			Created_At: time.Now(),
		}},
		Current_User:            "0d1ee9e2-dbe3-4a2a-b9cf-1ff27ce3a500",
		Current_Status:          "PENDING",
		Last_User_Before_Reject: "null",
		Is_Active:               true,
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
				ID:         uuid.NewString(),
				User:       "d160b410-e6a8-4cbb-92c2-068112187305",
				Status:     "CREATED",
				Created_At: time.Now(),
			},
		},
		Current_User:            "0d1ee9e2-dbe3-4a2a-b9cf-1ff27ce3a500",
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
				ID:         uuid.NewString(),
				User:       "d160b410-e6a8-4cbb-92c2-068112187503",
				Status:     "CREATED",
				Created_At: time.Now(),
			},
		},
		Current_User:            "0d1ee9e2-dbe3-4a2a-b9cf-1ff27ce3a500",
		Current_Status:          "PENDING",
		Last_User_Before_Reject: "null",
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
				ID:         uuid.NewString(),
				User:       "c160b410-e6a8-4cbb-92c2-068112187612",
				Status:     "CREATED",
				Created_At: time.Now().Add(55 * -time.Hour),
			},
		},
		Current_User:            "d160b410-e6a8-4cbb-92c2-068112187305",
		Current_Status:          "PENDING",
		Last_User_Before_Reject: "null",
		Is_Active:               true,
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
				ID:         uuid.NewString(),
				User:       "d160b410-e6a8-4cbb-92c2-068112187503",
				Status:     "CREATED",
				Created_At: time.Now().Add(5 * -time.Hour),
			},
		},
		Current_User:            "0d1ee9e2-dbe3-4a2a-b9cf-1ff27ce3a500",
		Current_Status:          "PENDING",
		Last_User_Before_Reject: "null",
		Is_Active:               true,
	},
}

// @id seed-data
// @summary initial seed data
// @description loads seed data for testing data and development purposes
// @tags setup, no-cache
// @produce json
// @success 201
// @router /seeds [post]
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

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"grants":     g_res,
		"users":      u_res,
		"mileage":    m_res,
		"check":      c_res,
		"petty_cash": p_res,
	})
}

// @id delete-seed-data
// @summary deletes seed data
// @description removes seed data used for testing and development purposes
// @tags setup, no-cache
// @produce json
// @success 200
// @router /seeds [delete]
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

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"grants":     g_res,
		"users":      u_res,
		"mileage":    m_res,
		"check":      c_res,
		"petty_cash": p_res,
	})
}
