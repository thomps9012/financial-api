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
		// update functionality here after user collection has been seeded
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

				result, err := user.Login(id, name, email)
				if err != nil {
					panic(err)
				}
				token, tokenErr := auth.GenerateToken(result.ID, result.Role)
				if tokenErr != nil {
					panic(tokenErr)
				}
				return token, nil
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
				var user u.User
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				user_id, idOK := p.Args["user_id"].(string)
				if !idOK {
					panic("you must enter a valid user id")
				}
				isAdmin := user.CheckAdmin(p.Context)
				contextuser, _ := user.FindContextID(p.Context)
				if !isAdmin {
					if contextuser.ID != user_id {
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
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				user, userErr := user.FindContextID(p.Context)
				if userErr != nil {
					panic(userErr)
				}
				name, nameOK := p.Args["name"].(string)
				if !nameOK {
					panic("you must enter a valid vehicle name")
				}
				description, descriptionOK := p.Args["description"].(string)
				if !descriptionOK {
					panic("you must enter a valid vehicle description")
				}
				result, err := user.AddVehicle(user.ID, name, description)
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
				"vehicle_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				user, userErr := user.FindContextID(p.Context)
				if userErr != nil {
					panic(userErr)
				}
				vehicle_id, vehicle_idOK := p.Args["vehicle_id"].(string)
				if !vehicle_idOK {
					panic("you must enter a valid vehicle id")
				}
				result, err := user.RemoveVehicle(user.ID, vehicle_id)
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
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(MileageInputType),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				requestor, userErr := user.FindContextID(p.Context)
				if userErr != nil {
					panic(userErr)
				}
				grant_id, grantOK := p.Args["grant_id"].(string)
				if !grantOK {
					panic("must enter a valid grant id")
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
					Grant_ID:          grant_id,
					Starting_Location: start,
					Destination:       destination,
					Trip_Purpose:      purpose,
					Start_Odometer:    start_odo,
					End_Odometer:      end_odo,
					Tolls:             tolls,
					Parking:           parking,
				}
				exists, _ := mileage_req.Exists(requestor.ID, date, start_odo, end_odo)
				if exists {
					return nil, errors.New("mileage request already created")
				}
				mileage_req.Create(requestor)
				return mileage_req, nil
			},
		},
		"edit_mileage": &graphql.Field{
			Type:        MileageType,
			Description: "Allows a user to edit one of their rejected or pending mileage requests",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
				"request": &graphql.ArgumentConfig{
					Type: MileageInputType,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				request_id, idOK := p.Args["request_id"].(string)
				if !idOK {
					panic("must enter a valid mileage id")
				}
				var milage_req r.Mileage_Request
				result, err := milage_req.FindByID(request_id)
				contextuser, _ := user.FindContextID(p.Context)
				if contextuser.ID != result.User_ID {
					panic("you are unauthorized to edit this record")
				}
				if err != nil {
					panic(err)
				}
				if result.Current_Status != "PENDING" && result.Current_Status != "REJECTED" {
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
					result.Date = date
					result.Starting_Location = start
					result.Destination = destination
					result.Trip_Purpose = purpose
					result.Start_Odometer = start_odo
					result.End_Odometer = end_odo
					result.Tolls = tolls
					result.Parking = parking
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
			Type:        PettyCashType,
			Description: "Creates a new petty cash request for a given user",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(PettyCashInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				requestor, userErr := user.FindContextID(p.Context)
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
				petty_cash_req := &r.Petty_Cash_Request{
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
			Type:        PettyCashType,
			Description: "Allows a user to edit one of their rejected or pending petty cash requests",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
				"request": &graphql.ArgumentConfig{
					Type: PettyCashInput,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				var petty_cash r.Petty_Cash_Request
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				contextuser, _ := user.FindContextID(p.Context)
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
				if result.Current_Status != "PENDING" && result.Current_Status != "REJECTED" {
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
			Type:        CheckRequestType,
			Description: "Creates a new check request for a given user",
			Args: graphql.FieldConfigArgument{
				"vendor": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(VendorInput),
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CheckRequestInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				requestor, userErr := user.FindContextID(p.Context)
				if userErr != nil {
					panic(userErr)
				}
				grant_id, grantOK := p.Args["grant_id"].(string)
				if !grantOK {
					panic("must enter a valid grant")
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
					Grant_ID:    grant_id,
					Purchases:   purchases,
					Order_Total: order_total,
					Receipts:    receipts,
					Credit_Card: checkReqArgs["credit_card"].(string),
				}
				exists, _ := check_request.Exists(requestor.ID, vendor.Name, order_total, check_request.Date)
				if exists {
					return nil, errors.New("check request already created")
				}
				check_request.Create(requestor)
				return check_request, nil
			},
		},
		"edit_check_request": &graphql.Field{
			Type:        CheckRequestType,
			Description: "Allows a user to edit one of their rejected or pending check requests",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"vendor": &graphql.ArgumentConfig{
					Type: VendorInput,
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.ID,
				},
				"request": &graphql.ArgumentConfig{
					Type: CheckRequestInput,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				request_id, idOK := p.Args["request_id"].(string)
				if !idOK {
					panic("must enter a valid mileage id")
				}
				var check_req r.Check_Request
				result, err := check_req.FindByID(request_id)
				contextuser, _ := user.FindContextID(p.Context)
				if contextuser.ID != result.User_ID {
					panic("you are unauthorized to edit this record")
				}
				if err != nil {
					panic(err)
				}
				if result.Current_Status != "PENDING" && result.Current_Status != "REJECTED" {
					panic("this request is already being processed")
				}
				if p.Args["grant_id"] != nil {
					result.Grant_ID = p.Args["grant_id"].(string)
				}
				if p.Args["request"] != nil {
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
					result.Date = checkReqArgs["date"].(time.Time)
					result.Description = checkReqArgs["description"].(string)
					result.Order_Total = order_total
					result.Receipts = receipts
					result.Credit_Card = checkReqArgs["credit_card"].(string)
				}
				if p.Args["vendor"] != nil {
					vendor_input := p.Args["vendor"].(map[string]interface{})
					vendor_address_input := vendor_input["address"].(map[string]interface{})
					vendor_address := &r.Address{
						Website:  vendor_address_input["website"].(string),
						Street:   vendor_address_input["street"].(string),
						City:     vendor_address_input["city"].(string),
						State:    vendor_address_input["state"].(string),
						Zip_Code: vendor_address_input["zip"].(int),
					}
					vendor := r.Vendor{
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
			Description: "A method for a manager to approve a financial request",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request_type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				request_id, idOK := p.Args["request_id"].(string)
				if !idOK {
					panic("must enter a valid request id")
				}
				request_type, typeOk := p.Args["request_type"].(string)
				if !typeOk {
					panic("must enter a valid request type")
				}
				var user u.User
				var action r.Action
				// get request user id from request
				request, err := action.FindOne(request_id, request_type)
				if err != nil {
					panic(err)
				}
				// get manager id from context
				// get manager role from context
				loggedIn := user.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				manager, managerErr := user.FindContextID(p.Context)
				if managerErr != nil {
					panic(managerErr)
				}
				approveReq, approveErr := action.Approve(request_id, request, manager, request_type)
				if approveErr != nil {
					panic(approveErr)
				}
				if approveReq == false {
					panic("error approving request")
				}
				return approveReq, nil
			},
		},
		"reject_request": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "A method for a manager to reject a financial request",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request_type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				request_id, idOK := p.Args["request_id"].(string)
				if !idOK {
					panic("must enter a valid request id")
				}
				request_type, typeOk := p.Args["request_type"].(string)
				if !typeOk {
					panic("must enter a valid request type")
				}
				var action r.Action
				var user u.User
				// get request user id from request
				request, err := action.FindOne(request_id, request_type)
				if err != nil {
					panic(err)
				}
				// get manager id from context
				manager, managerErr := user.FindContextID(p.Context)
				if managerErr != nil {
					panic(managerErr)
				}
				rejectReq, rejectErr := action.Reject(request_id, request, manager.ID, request_type)
				if rejectErr != nil {
					panic(rejectErr)
				}
				if !rejectReq {
					panic("error rejecting request")
				}
				return rejectReq, nil
			},
		},
		"archive_request": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "A method for a manager to reject a financial request",
			Args: graphql.FieldConfigArgument{
				"request_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request_type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				request_id, idOK := p.Args["request_id"].(string)
				if !idOK {
					panic("must enter a valid request id")
				}
				request_type, typeOk := p.Args["request_type"].(string)
				if !typeOk {
					panic("must enter a valid request type")
				}
				var action r.Action
				var user u.User
				user, userErr := user.FindContextID(p.Context)
				if userErr != nil {
					panic(userErr)
				}
				archiveReq, archiveErr := action.Archive(request_id, request_type, user.ID)
				if archiveErr != nil {
					panic(archiveErr)
				}
				return archiveReq, nil
			},
		},
		"clear_notification": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "A method for a user to clear a notification that has been dealt with",
			Args: graphql.FieldConfigArgument{
				"item_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				item_id, idOK := p.Args["item_id"].(string)
				if !idOK {
					panic("must enter a valid action id")
				}
				var user u.User
				userInfo, idErr := user.FindContextID(p.Context)
				if idErr != nil {
					panic(idErr)
				}
				notificationClear, clearErr := user.ClearNotification(item_id, userInfo.ID)
				if clearErr != nil {
					panic(clearErr)
				}
				return notificationClear, nil
			},
		},
		"clear_all_notifications": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "A method for a user to clear all of their notifications",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				userInfo, idErr := user.FindContextID(p.Context)
				if idErr != nil {
					panic(idErr)
				}
				notificationClear, clearErr := user.ClearNotifications(userInfo.ID)
				if clearErr != nil {
					panic(clearErr)
				}
				return notificationClear, nil
			},
		},
	},
})
