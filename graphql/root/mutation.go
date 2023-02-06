package root

import (
	"encoding/json"
	"errors"
	"financial-api/middleware"
	"financial-api/models"
	"fmt"
	"math"
	"regexp"
	"time"

	"github.com/graphql-go/graphql"
)

var RootMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutations",
	Fields: graphql.Fields{
		"login": &graphql.Field{
			Type:        graphql.String,
			Description: "A mutation that either creates, or logs a user in, dependent on whether or not they have an account in the database.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
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
				id, okid := p.Args["id"].(string)
				if !okid {
					panic(okid)
				}
				var user models.User
				exists, _ := user.Exists(id)
				if exists {
					result, err := user.Login(id)
					if err != nil {
						panic(err)
					}
					return result, nil
				} else {
					user.ID = id
					user.Email = email
					user.Name = name
					result, err := user.Create()
					if err != nil {
						panic(err)
					}
					return result, nil
				}
			},
		},
		"deactivate_user": &graphql.Field{
			Type:        user_detail,
			Description: "Deactivates a user by setting the status to inactive",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				user_id, idOK := p.Args["user_id"].(string)
				if !idOK {
					panic("you must enter a valid user id")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				context_user := middleware.ForID(p.Context)
				if !isAdmin {
					if context_user != user_id {
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
			Type:        user_vehicle,
			Description: "Allows a user to add a vehicle to their account",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				user_id := middleware.ForID(p.Context)
				name, nameOK := p.Args["name"].(string)
				if !nameOK {
					panic("you must enter a valid vehicle name")
				}
				description, descriptionOK := p.Args["description"].(string)
				if !descriptionOK {
					panic("you must enter a valid vehicle description")
				}
				result, err := user.AddVehicle(user_id, name, description)
				println(result)
				if err != nil {
					panic(err)
				}
				return &models.Vehicle{
					ID:          result,
					Name:        name,
					Description: description,
				}, nil
			},
		},
		"remove_vehicle": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Removes a specific vehicle from a user's account.",
			Args: graphql.FieldConfigArgument{
				"vehicle_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				user_id := middleware.ForID(p.Context)
				vehicle_id, vehicle_idOK := p.Args["vehicle_id"].(string)
				if !vehicle_idOK {
					panic("you must enter a valid vehicle id")
				}
				result, err := user.RemoveVehicle(user_id, vehicle_id)
				if err != nil {
					panic(err)
				}
				return result, nil
			},
		},
		// mileage mutations
		"create_mileage": &graphql.Field{
			Type:        mileage_request,
			Description: "Creates a new mileage request for a given user based on the logged in user context.",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(mileage_input),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				requestor_id := middleware.ForID(p.Context)
				grant_id, grantOK := p.Args["grant_id"].(string)
				if !grantOK {
					panic("must enter a valid grant id")
				}
				mileageArgs := p.Args["request"].(map[string]interface{})
				date, isOK := mileageArgs["date"].(time.Time)
				if !isOK {
					panic("must enter a valid date")
				}
				start, isOK := mileageArgs["starting_location"].(string)
				if !isOK {
					panic("must enter a valid starting location")
				}
				destination, isOK := mileageArgs["destination"].(string)
				if !isOK {
					panic("must enter a valid destination")
				}
				purpose, isOK := mileageArgs["trip_purpose"].(string)
				if !isOK {
					panic("must enter a valid trip purpose")
				}
				start_odo, isOK := mileageArgs["start_odometer"].(int)
				if !isOK {
					panic("must enter a valid starting odometer")
				}
				end_odo, isOK := mileageArgs["end_odometer"].(int)
				if !isOK {
					panic("must enter a valid end odometer")
				}
				tolls, isOK := mileageArgs["tolls"].(float64)
				if !isOK {
					panic("must enter a valid tolls amount")
				}
				parking, isOK := mileageArgs["parking"].(float64)
				if !isOK {
					panic("must enter a valid parking amount")
				}
				category, isOK := mileageArgs["category"].(string)
				if !isOK {
					panic("must enter a valid category type")
				}
				// add in mileage variance response here
				mileage_req := &models.Mileage_Request{
					Date:              date,
					Category:          models.Category(category),
					Grant_ID:          grant_id,
					Starting_Location: start,
					Destination:       destination,
					Trip_Purpose:      purpose,
					Start_Odometer:    start_odo,
					End_Odometer:      end_odo,
					Tolls:             tolls,
					Parking:           parking,
				}
				exists, _ := mileage_req.Exists(requestor_id, date, start_odo, end_odo)
				if exists {
					return nil, errors.New("mileage request already created")
				}
				var user models.User
				requestor, err := user.FindByID(requestor_id)
				if err != nil {
					panic("error finding user with that id")
				}
				mileage_req.Create(requestor)
				return mileage_req, nil
			},
		},
		"test_create_mileage": &graphql.Field{
			Type:        test_mileage_request,
			Description: "Creates a new mileage request for a given user based on the logged in user context.",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(test_mileage_input),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				requestor_id := middleware.ForID(p.Context)
				grant_id, grantOK := p.Args["grant_id"].(string)
				if !grantOK {
					panic("must enter a valid grant id")
				}
				mileageArgs := p.Args["request"].(map[string]interface{})
				date, isOK := mileageArgs["date"].(time.Time)
				if !isOK {
					panic("must enter a valid date")
				}
				start, isOK := mileageArgs["starting_location"].(map[string]interface{})
				if !isOK {
					panic("must enter a valid starting location")
				}
				starting_location := models.Location{
					Latitude:  start["latitude"].(float64),
					Longitude: start["longitude"].(float64),
				}
				destination, isOK := mileageArgs["destination"].(map[string]interface{})
				if !isOK {
					panic("must enter a valid destination")
				}
				destination_location := models.Location{
					Latitude:  destination["latitude"].(float64),
					Longitude: destination["longitude"].(float64),
				}
				purpose, isOK := mileageArgs["trip_purpose"].(string)
				if !isOK {
					panic("must enter a valid trip purpose")
				}
				tolls, isOK := mileageArgs["tolls"].(float64)
				if !isOK {
					panic("must enter a valid tolls amount")
				}
				parking, isOK := mileageArgs["parking"].(float64)
				if !isOK {
					panic("must enter a valid parking amount")
				}
				category, isOK := mileageArgs["category"].(string)
				if !isOK {
					panic("must enter a valid category type")
				}
				mileage_variance, isOK := mileageArgs["request_variance"].(map[string]interface{})
				if !isOK {
					panic("must enter a valid mileage variance")
				}
				location_points, isOK := mileageArgs["location_points"].([]interface{})
				if !isOK {
					panic("must enter valid location points")
				}
				var locations []models.Location
				marshalled, err := json.Marshal(location_points)
				if err != nil {
					panic(err)
				}
				unmarshal_err := json.Unmarshal(marshalled, &locations)
				if unmarshal_err != nil {
					panic(unmarshal_err)
				}
				// add in mileage variance response here
				mileage_req := &models.New_Mileage_Request{
					Date:              date,
					Category:          models.Category(category),
					Grant_ID:          grant_id,
					Starting_Location: starting_location,
					Destination:       destination_location,
					LocationPoints:    locations,
					Trip_Purpose:      purpose,
					Tolls:             tolls,
					Parking:           parking,
					Request_Variance: models.ResponseCompare{
						Matrix_Distance:   mileage_variance["matrix_distance"].(float64),
						Traveled_Distance: mileage_variance["traveled_distance"].(float64),
						Difference:        mileage_variance["difference"].(float64),
						Variance:          models.Variance_Level(mileage_variance["variance"].(string)),
					},
				}
				exists, _ := mileage_req.NewExists(requestor_id, date, starting_location, destination_location)
				if exists {
					return nil, errors.New("mileage request already created")
				}
				var user models.User
				requestor, err := user.FindByID(requestor_id)
				if err != nil {
					panic("error finding user with that id")
				}
				mileage_req.NewCreate(requestor)
				return mileage_req, nil
			},
		},
		"edit_mileage": &graphql.Field{
			Type:        mileage_request,
			Description: "Allows a user to edit one of their rejected or pending mileage requests, requests that have been approved (at any stage) may not be edited.",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
				"request": &graphql.ArgumentConfig{
					Type: mileage_input,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				request_id, idOK := p.Args["request_id"].(string)
				if !idOK {
					panic("must enter a valid mileage id")
				}
				var milage_req models.Mileage_Request
				result, err := milage_req.FindByID(request_id)
				user_id := middleware.ForID(p.Context)
				contextuser, _ := user.FindByID(user_id)
				if contextuser.ID != result.User_ID {
					panic("you are unauthorized to edit this record")
				}
				if err != nil {
					panic(err)
				}
				if result.Current_Status != models.PENDING && result.Current_Status != models.REJECTED {
					panic("this request is already being processed")
				}
				// add in conditional update of fields based on input
				if p.Args["grant_id"] != nil {
					result.Grant_ID = p.Args["grant_id"].(string)
				}
				if p.Args["request"] != nil {
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
					category, category_ok := mileageArgs["category"].(string)
					if !category_ok {
						panic("must enter a valid request category")
					}
					result.Date = date
					result.Starting_Location = start
					result.Destination = destination
					result.Trip_Purpose = purpose
					result.Start_Odometer = start_odo
					result.End_Odometer = end_odo
					result.Tolls = tolls
					result.Parking = parking
					result.Category = models.Category(category)
				}
				updatedDoc, updateErr := milage_req.Update(result, contextuser)
				if updateErr != nil {
					panic(updateErr)
				}
				return updatedDoc, nil
			},
		},
		// petty cash mutations
		"create_petty_cash": &graphql.Field{
			Type:        petty_cash_request,
			Description: "Creates a new petty cash request for a given user, based on the logged in user context.",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(petty_cash_input),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				requestor_id := middleware.ForID(p.Context)
				requestor, userErr := user.FindByID(requestor_id)
				if userErr != nil {
					panic(userErr)
				}
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
				category, category_ok := requestArgs["category"].(string)
				if !category_ok {
					panic("must enter a valid request category")
				}
				petty_cash_req := &models.Petty_Cash_Request{
					Category:    models.Category(category),
					Date:        date,
					Grant_ID:    grant_id,
					Amount:      amount,
					Description: description,
					Receipts:    receipts,
				}
				exists, _ := petty_cash_req.Exists(requestor.ID, amount, date)
				if exists {
					return nil, errors.New("duplicate petty cash request")
				}
				petty_cash_req.Create(requestor)
				return petty_cash_req, nil
			},
		},
		"edit_petty_cash": &graphql.Field{
			Type:        petty_cash_request,
			Description: "Allows a user to edit one of their rejected or pending petty cash requests, requests that have been approved (at any stage) may not be edited.",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
				"request": &graphql.ArgumentConfig{
					Type: petty_cash_input,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				var petty_cash models.Petty_Cash_Request
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				user_id := middleware.ForID(p.Context)
				contextuser, _ := user.FindByID(user_id)
				request_id, idOK := p.Args["request_id"].(string)
				if !idOK {
					panic("must enter a valid request id")
				}
				result, findErr := petty_cash.FindByID(request_id)
				if findErr != nil {
					panic(findErr)
				}
				if contextuser.ID != result.User_ID {
					panic("you are unauthorized to edit this record")
				}
				if result.Current_Status != models.PENDING && result.Current_Status != models.REJECTED {
					panic("this request is currently being processed")
				}
				if p.Args["request"] != nil {
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
					category, category_ok := requestArgs["category"].(string)
					if !category_ok {
						panic("must enter a valid request category")
					}
					result.Category = models.Category(category)
					result.Amount = amount
					result.Date = date
					result.Receipts = receipts
					result.Description = description
				}
				if p.Args["grant_id"] != nil {
					result.Grant_ID = p.Args["grant_id"].(string)
				}
				updatedDoc, updateErr := petty_cash.Update(result, contextuser)
				if updateErr != nil {
					panic(updateErr)
				}
				return updatedDoc, nil
			},
		},
		// check request mutations
		"create_check_request": &graphql.Field{
			Type:        check_request,
			Description: "Creates a new check request for a given user, based on the logged in user context.",
			Args: graphql.FieldConfigArgument{
				"vendor": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(vendor_input),
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(check_request_input),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				requestor_id := middleware.ForID(p.Context)
				requestor, userErr := user.FindByID(requestor_id)
				if userErr != nil {
					panic(userErr)
				}
				fmt.Printf("%v\n", p.Args)
				grant_id, grantOK := p.Args["grant_id"].(string)
				if !grantOK {
					panic("must enter a valid grant")
				}
				vendor_input := p.Args["vendor"].(map[string]interface{})
				vendor_address_input := vendor_input["address"].(map[string]interface{})
				vendor_address := &models.Address{
					Website:  vendor_address_input["website"].(string),
					Street:   vendor_address_input["street"].(string),
					City:     vendor_address_input["city"].(string),
					State:    vendor_address_input["state"].(string),
					Zip_Code: vendor_address_input["zip"].(int),
				}
				vendor := &models.Vendor{
					Name:    vendor_input["name"].(string),
					Address: *vendor_address,
				}
				checkReqArgs := p.Args["request"].(map[string]interface{})
				purchases_input := checkReqArgs["purchases"].([]interface{})
				var purchases []models.Purchase
				var order_total = 0.0
				for _, purchase_item_obj := range purchases_input {
					fmt.Printf("%s\n", purchase_item_obj)
					purchase_item := purchase_item_obj.(map[string]interface{})
					amount := purchase_item["amount"].(float64)
					description := purchase_item["description"].(string)
					grant_line_item := purchase_item["grant_line_item"].(string)
					purchase := &models.Purchase{
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
				category, category_ok := checkReqArgs["category"].(string)
				if !category_ok {
					panic("must enter a valid request category")
				}
				var check_request models.Check_Request
				check_request.Category = models.Category(category)
				check_request.Date = checkReqArgs["date"].(time.Time)
				check_request.Vendor = *vendor
				check_request.Description = checkReqArgs["description"].(string)
				check_request.Grant_ID = grant_id
				check_request.Purchases = purchases
				check_request.Order_Total = order_total
				check_request.Receipts = receipts
				check_request.Credit_Card = checkReqArgs["credit_card"].(string)
				exists := check_request.Exists(requestor.ID, vendor.Name, order_total, check_request.Date)
				if exists {
					return nil, errors.New("check request already created")
				}
				fmt.Println("exists checker", exists)
				check_request.Create(requestor)
				return check_request, nil
			},
		},
		"edit_check_request": &graphql.Field{
			Type:        check_request,
			Description: "Allows a user to edit one of their rejected or pending check requests, requests that have been approved (at any stage) may not be edited.",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"vendor": &graphql.ArgumentConfig{
					Type: vendor_input,
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
				"request": &graphql.ArgumentConfig{
					Type: check_request_input,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				request_id, idOK := p.Args["request_id"].(string)
				if !idOK {
					panic("must enter a valid mileage id")
				}
				var check_req models.Check_Request
				result, err := check_req.FindByID(request_id)
				user_id := middleware.ForID(p.Context)
				contextuser, _ := user.FindByID(user_id)
				if contextuser.ID != result.User_ID {
					panic("you are unauthorized to edit this record")
				}
				if err != nil {
					panic(err)
				}
				if result.Current_Status != models.PENDING && result.Current_Status != models.REJECTED {
					panic("this request is already being processed")
				}
				if p.Args["grant_id"] != nil {
					result.Grant_ID = p.Args["grant_id"].(string)
				}
				if p.Args["request"] != nil {
					checkReqArgs := p.Args["request"].(map[string]interface{})
					purchases_input := checkReqArgs["purchases"].([]interface{})
					var purchases []models.Purchase
					var order_total = 0.0
					for _, purchase_item_obj := range purchases_input {
						fmt.Printf("%s\n", purchase_item_obj)
						purchase_item := purchase_item_obj.(map[string]interface{})
						amount := purchase_item["amount"].(float64)
						description := purchase_item["description"].(string)
						grant_line_item := purchase_item["grant_line_item"].(string)
						purchase := &models.Purchase{
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
					category, category_ok := checkReqArgs["category"].(string)
					if !category_ok {
						panic("must enter a valid request category")
					}
					result.Category = models.Category(category)
					result.Date = checkReqArgs["date"].(time.Time)
					result.Description = checkReqArgs["description"].(string)
					result.Order_Total = order_total
					result.Receipts = receipts
					result.Purchases = purchases
					result.Credit_Card = checkReqArgs["credit_card"].(string)
				}
				if p.Args["vendor"] != nil {
					vendor_input := p.Args["vendor"].(map[string]interface{})
					vendor_address_input := vendor_input["address"].(map[string]interface{})
					vendor_address := &models.Address{
						Website:  vendor_address_input["website"].(string),
						Street:   vendor_address_input["street"].(string),
						City:     vendor_address_input["city"].(string),
						State:    vendor_address_input["state"].(string),
						Zip_Code: vendor_address_input["zip"].(int),
					}
					vendor := models.Vendor{
						Name:    vendor_input["name"].(string),
						Address: *vendor_address,
					}
					result.Vendor = vendor
				}
				updatedDoc, updateErr := check_req.Update(result, contextuser)
				if updateErr != nil {
					panic(updateErr)
				}
				return updatedDoc, nil
			},
		},
		// action mutations
		"approve_request": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Allows an administrator or manager to approve a specific financial request.",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request_type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(request_type),
				},
				"request_category": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(request_category),
				},
				"new_status": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(request_status),
				},
				"exec_review": &graphql.ArgumentConfig{
					Type:         graphql.NewNonNull(graphql.Boolean),
					DefaultValue: false,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				request_id, isOK := p.Args["request_id"].(string)
				if !isOK {
					panic("must enter a valid request id")
				}
				request_type, isOK := p.Args["request_type"].(string)
				if !isOK {
					panic("must enter a valid request type")
				}
				request_category, isOK := p.Args["request_category"].(string)
				if !isOK {
					panic("must enter a valid request category")
				}
				new_status, isOK := p.Args["new_status"].(string)
				if !isOK {
					panic("must enter a valid request status")
				}
				exec_review, isOK := p.Args["exec_review"].(bool)
				if !isOK {
					panic("must enter a valid option for executive review")
				}
				admin := middleware.ForAdmin(p.Context)
				if !admin {
					panic("you are unable to approve requests")
				}
				var action models.Action
				request_info, err := action.Get(request_id, models.Request_Type(request_type))
				if err != nil {
					panic(err)
				}
				if !request_info.CheckStatus(models.Status(new_status)) {
					panic("current action has already been taken")
				}
				new_action := action.Create(models.Status(new_status), request_info)
				current_user_email := models.UserEmailHandler(models.Category(request_category), models.Status(new_status), exec_review)
				user_id, err := models.DetermineUserID(current_user_email, request_info)
				if err != nil {
					panic(err)
				}
				var user models.User
				prev_user_clear_notification, err := user.ClearNotification(request_id, request_info.Current_User)
				if err != nil {
					panic(err)
				}
				if !prev_user_clear_notification {
					panic("error clearing the previous reviewer's notifications")
				}
				updated_doc, err := models.UpdateRequest(new_action, user_id)
				if err != nil || !updated_doc {
					panic("error updating the request")
				}
				current_user_notified, err := user.AddNotification(new_action, user_id)
				if err != nil {
					panic(err)
				}
				if !current_user_notified {
					panic("error notifying new request reviewer")
				}
				if request_info.User_ID != user_id {
					requestor_notified, err := user.AddNotification(new_action, request_info.User_ID)
					if err != nil {
						panic(err)
					}
					return requestor_notified, nil
				} else {
					return current_user_notified, nil
				}
			},
		},
		"reject_request": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Allows a manager or administrator to reject a financial request and kicks it back to the user for further revisions.",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request_type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(request_type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				request_id, isOK := p.Args["request_id"].(string)
				if !isOK {
					panic("must enter a valid request id")
				}
				request_type, isOK := p.Args["request_type"].(string)
				if !isOK {
					panic("must enter a valid request type")
				}
				admin := middleware.ForAdmin(p.Context)
				if !admin {
					panic("you are unable to reject finance requests")
				}
				var action models.Action
				var new_status = models.REJECTED
				request_info, err := action.Get(request_id, models.Request_Type(request_type))
				if err != nil {
					panic(err)
				}
				if !request_info.CheckStatus(models.Status(new_status)) {
					panic("current action has already been taken")
				}
				new_action := action.Create(models.Status(new_status), request_info)
				var user models.User
				prev_user_clear_notification, err := user.ClearNotification(request_id, request_info.Current_User)
				if err != nil {
					panic(err)
				}
				if !prev_user_clear_notification {
					panic("error clearing previous user's notification")
				}
				updated_doc, err := models.UpdateRequest(new_action, request_info.User_ID)
				if err != nil || !updated_doc {
					panic("error updating the request")
				}
				requestor_notified, err := user.AddNotification(new_action, request_info.User_ID)
				if err != nil {
					panic(err)
				}
				return requestor_notified, nil
			},
		},
		"archive_request": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Allows a user or manager to archive a specified financial request, taking it out of the approval process.",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request_type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(request_type),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you must be logged in")
				}
				request_id, idOK := p.Args["request_id"].(string)
				if !idOK {
					panic("must enter a valid request id")
				}
				request_type, typeOk := p.Args["request_type"].(string)
				if !typeOk {
					panic("must enter a valid request type")
				}
				var action models.Action
				user_id := middleware.ForID(p.Context)
				admin := middleware.ForAdmin(p.Context)
				var new_status = models.ARCHIVED
				request_info, err := action.Get(request_id, models.Request_Type(request_type))
				if err != nil {
					panic(err)
				}
				if !admin && request_info.User_ID != user_id {
					panic("you are unauthorized to archive this request")
				}
				if !request_info.CheckStatus(models.Status(new_status)) {
					panic("current action has already been taken")
				}
				new_action := action.Create(models.Status(new_status), request_info)
				var user models.User
				prev_user_clear_notification, err := user.ClearNotification(request_id, request_info.Current_User)
				if err != nil {
					panic(err)
				}
				if !prev_user_clear_notification {
					panic("error clearing previous user's notification")
				}
				updated_doc, err := models.UpdateRequest(new_action, request_info.User_ID)
				if err != nil || !updated_doc {
					panic("error updating the request")
				}
				requestor_notified, err := user.AddNotification(new_action, request_info.User_ID)
				if err != nil {
					panic(err)
				}
				return requestor_notified, nil
			},
		},
		"clear_notification": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Clears a logged in user's notification based on the notification id.",
			Args: graphql.FieldConfigArgument{
				"notification_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you must be logged in")
				}
				notification_id, idOK := p.Args["notification_id"].(string)
				if !idOK {
					panic("must enter a valid action id")
				}
				var user models.User
				user_id := middleware.ForID(p.Context)
				notificationClear, clearErr := user.ClearNotificationByID(notification_id, user_id)
				if clearErr != nil {
					panic(clearErr)
				}
				return notificationClear, nil
			},
		},
		"clear_all_notifications": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Clears all of a logged in user's notifications.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you must be logged in")
				}
				var user models.User
				user_id := middleware.ForID(p.Context)
				notificationClear, clearErr := user.ClearNotifications(user_id)
				if clearErr != nil {
					panic(clearErr)
				}
				return notificationClear, nil
			},
		},
	},
})
