package root

import (
	"errors"
	auth "financial-api/middleware"
	r "financial-api/models/requests"
	u "financial-api/models/user"
	"fmt"
	"math"
	"regexp"
	"time"

	"github.com/graphql-go/graphql"
)

var RootMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutations",
	Fields: graphql.Fields{
		// user mutations
		"sign_in": &graphql.Field{
			Type:        graphql.String,
			Description: "Either creates a new user or logs a user in based on their account history",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email, okEmail := p.Args["email"].(string)
				if !okEmail {
					panic(okEmail)
				}
				emailCheck, _ := regexp.MatchString("[a-z0-9!#$%&'*+/=?^_{|}~-]*@norainc.org", email)
				if !emailCheck {
					panic("must have a Northern Ohio Recovery Association Email to register")
				}
				name, okName := p.Args["name"].(string)
				if !okName {
					panic(okName)
				}
				id, okID := p.Args["id"].(string)
				if !okID {
					panic(okID)
				}
				var user u.User
				exists, _ := user.Exists(id)
				if exists {
					result, err := user.Login(id)
					if err != nil {
						panic(err)
					}
					return result, nil
				} else {
					result, err := user.Create(id, name, email)
					if err != nil {
						panic(err)
					}
					return result, nil
				}
			},
		},
		"deactivate_user": &graphql.Field{
			Type:        UserType,
			Description: "Deactivates a user by setting the status to inactive",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_id, idOK := p.Args["user_id"].(string)
				if !idOK {
					panic("you must enter a valid user id")
				}
				var user u.User
				userRole := auth.ForRole(p.Context)
				userID := auth.ForID(p.Context)
				if userRole == "EMPLOYEE" {
					if userID != user_id {
						panic("You are unauthorized to deactivate this user")
					}
				}
				result, err := user.Deactivate(user_id)
				if err != nil {
					panic(err)
				}
				return result, nil
			},
		},
		"add_vehicle": &graphql.Field{
			Type:        VehicleType,
			Description: "Allow a user to add a vehicle to their account",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_info := p.Context.Value("user")
				if user_info == nil {
					return nil, errors.New("Not logged in")
				}
				user_id, idOK := p.Args["user_id"].(string)
				if !idOK {
					panic("you must enter a valid user id")
				}
				name, nameOK := p.Args["name"].(string)
				if !nameOK {
					panic("you must enter a valid vehicle name")
				}
				description, descriptionOK := p.Args["description"].(string)
				if !descriptionOK {
					panic("you must enter a valid vehicle description")
				}
				var user u.User
				result, err := user.AddVehicle(user_id, name, description)
				println(result)
				if err != nil {
					panic(err)
				}
				return &u.Vehicle{
					ID:          result,
					Name:        name,
					Description: description,
				}, nil
			},
		},
		"remove_vehicle": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Allow a user to remove a vehicle from their account",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"vehicle_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_info := p.Context.Value("user")
				if user_info == nil {
					return nil, errors.New("Not logged in")
				}
				user_id, idOK := p.Args["user_id"].(string)
				if !idOK {
					panic("you must enter a valid user id")
				}
				vehicle_id, vehicle_idOK := p.Args["vehicle_id"].(string)
				if !vehicle_idOK {
					panic("you must enter a valid vehicle id")
				}
				var user u.User
				result, err := user.RemoveVehicle(user_id, vehicle_id)
				if err != nil {
					panic(err)
				}
				return result, nil
			},
		},
		// mileage mutations
		"create_mileage": &graphql.Field{
			Type:        MileageType,
			Description: "Creates a new mileage request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(MileageInputType),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_info := p.Context.Value("user")
				if user_info == nil {
					return nil, errors.New("Not logged in")
				}
				userID, isOK := p.Args["user_id"].(string)
				if !isOK {
					panic("must enter a valid user id")
				}
				mileageArgs := p.Args["request"].(map[string]interface{})
				date, dateisOK := mileageArgs["date"].(time.Time)
				if !dateisOK {
					panic("must enter a valid date")
				}
				start, startisOK := mileageArgs["starting_location"].(string)
				if !startisOK {
					panic("must enter a valid starting location")
				}
				destination, destinationisOK := mileageArgs["destination"].(string)
				if !destinationisOK {
					panic("must enter a valid destination")
				}
				purpose, purposeisOK := mileageArgs["trip_purpose"].(string)
				if !purposeisOK {
					panic("must enter a valid trip purpose")
				}
				start_odo, start_odoisOK := mileageArgs["start_odometer"].(int)
				if !start_odoisOK {
					panic("must enter a valid starting odometer")
				}
				end_odo, end_odoisOK := mileageArgs["end_odometer"].(int)
				if !end_odoisOK {
					panic("must enter a valid end odometer")
				}
				tolls, tollsisOK := mileageArgs["tolls"].(float64)
				if !tollsisOK {
					panic("must enter a valid tolls amount")
				}
				parking, parkingisOK := mileageArgs["parking"].(float64)
				if !parkingisOK {
					panic("must enter a valid parking amount")
				}
				mileage_req := &r.Mileage_Request{
					Date:              date,
					Starting_Location: start,
					Destination:       destination,
					Trip_Purpose:      purpose,
					Start_Odometer:    start_odo,
					End_Odometer:      end_odo,
					Tolls:             tolls,
					Parking:           parking,
				}
				exists, _ := mileage_req.Exists(userID, date, start_odo, end_odo)
				if exists {
					return nil, errors.New("mileage request already created")
				}
				mileage_req.Create(userID)
				return mileage_req, nil
			},
		},
		// petty cash mutations
		"create_petty_cash": &graphql.Field{
			Type:        PettyCashType,
			Description: "Creates a new petty cash request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(PettyCashInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_info := p.Context.Value("user")
				if user_info == nil {
					return nil, errors.New("Not logged in")
				}
				user_id, isOk := p.Args["user_id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				println(user_id)
				grant_id, grantOK := p.Args["grant_id"].(string)
				if !grantOK {
					panic("must enter a valid grant")
				}
				requestArgs := p.Args["request"].(map[string]interface{})
				amount, okAmount := requestArgs["amount"].(float64)
				if !okAmount {
					panic("must enter a valid amount")
				}
				date, okdate := requestArgs["date"].(time.Time)
				if !okdate {
					panic("must enter a valid date")
				}
				receiptArgs, receiptsOK := requestArgs["receipts"].([]interface{})
				if !receiptsOK {
					panic("must enter a valid receipt item")
				}
				var receipts []string
				for item := range receiptArgs {
					receipts = append(receipts, receiptArgs[item].(string))
				}
				description, descriptionOK := requestArgs["description"].(string)
				if !descriptionOK {
					panic("must enter a valid description")
				}
				petty_cash_req := &r.Petty_Cash_Request{
					Date:        date,
					Grant_ID:    grant_id,
					Amount:      amount,
					Description: description,
					Receipts:    receipts,
				}
				exists, _ := petty_cash_req.Exists(user_id, amount, date)
				if exists {
					return nil, errors.New("duplicate petty cash request")
				}
				petty_cash_req.Create(user_id)
				return petty_cash_req, nil
			},
		},
		// check request mutations
		"create_check_request": &graphql.Field{
			Type:        CheckRequestType,
			Description: "Creates a new check request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"vendor": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(VendorInput),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CheckRequestInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_info := p.Context.Value("user")
				if user_info == nil {
					return nil, errors.New("Not logged in")
				}
				vendor_input := p.Args["vendor"].(map[string]interface{})
				vendor_address_input := vendor_input["address"].(map[string]interface{})
				vendor_address := &r.Address{
					Website:  vendor_address_input["website"].(string),
					Street:   vendor_address_input["street"].(string),
					City:     vendor_address_input["city"].(string),
					State:    vendor_address_input["state"].(string),
					Zip_Code: vendor_address_input["zip"].(int),
				}
				vendor := &r.Vendor{
					Name:    vendor_input["name"].(string),
					Address: *vendor_address,
				}
				checkReqArgs := p.Args["request"].(map[string]interface{})
				purchases_input := checkReqArgs["purchases"].([]interface{})
				var purchases []r.Purchase
				var order_total = 0.0
				for _, purchase_item_obj := range purchases_input {
					fmt.Printf("%s\n", purchase_item_obj)
					purchase_item := purchase_item_obj.(map[string]interface{})
					amount := purchase_item["amount"].(float64)
					description := purchase_item["description"].(string)
					grant_line_item := purchase_item["grant_line_item"].(string)
					purchase := &r.Purchase{
						Grant_Line_Item: grant_line_item,
						Description:     description,
						Amount:          (math.Round(amount*100) / 100),
					}
					order_total += (math.Round(amount*100) / 100)
					order_total = math.Round(order_total*100) / 100
					purchases = append(purchases, *purchase)
				}
				receiptArgs := checkReqArgs["receipts"].([]interface{})
				var receipts []string
				for item := range receiptArgs {
					receipts = append(receipts, receiptArgs[item].(string))
				}
				check_request := &r.Check_Request{
					Date:        checkReqArgs["date"].(time.Time),
					Vendor:      *vendor,
					Description: checkReqArgs["description"].(string),
					Grant_ID:    checkReqArgs["grant_id"].(string),
					Purchases:   purchases,
					Order_Total: order_total,
					Receipts:    receipts,
					Credit_Card: checkReqArgs["credit_card"].(string),
				}
				user_id, isOk := p.Args["user_id"].(string)
				if !isOk {
					panic("must enter a user id")
				}
				exists, _ := check_request.Exists(user_id, vendor.Name, order_total, check_request.Date)
				if exists {
					return nil, errors.New("check request already created")
				}
				check_request.Create(user_id)
				return check_request, nil
			},
		},
	},
})
