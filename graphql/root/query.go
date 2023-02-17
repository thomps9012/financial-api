package root

import (
	"context"
	"encoding/json"
	conn "financial-api/db"
	"financial-api/middleware"
	"financial-api/models"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

var RootQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQueries",
	Fields: graphql.Fields{
		// user queries
		"me": &graphql.Field{
			Type:        user_overview,
			Description: "Gathers overview information for a specific logged in user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user models.User
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				user_id := middleware.ForID(p.Context)
				user, userErr := user.FindByID(user_id)
				if userErr != nil {
					panic(userErr)
				}
				petty_cash_reqs, petty_err := user.AggregatePettyCash(user.ID)
				if petty_err != nil {
					panic(petty_err)
				}
				mileage_reqs, mileage_err := user.FindMileage(user.ID)
				if mileage_err != nil {
					panic(mileage_err)
				}
				check_reqs, check_err := user.AggregateChecks(user.ID, "", "")
				if check_err != nil {
					panic(check_err)
				}
				return &models.User_Detail{
					ID:                      user.ID,
					Name:                    user.Name,
					Admin:                   user.Admin,
					Permissions:             user.Permissions,
					Incomplete_Actions:      user.InComplete_Actions,
					Incomplete_Action_Count: len(user.InComplete_Actions),
					Last_Login:              user.Last_Login,
					Vehicles:                user.Vehicles,
					Mileage_Requests:        mileage_reqs,
					Check_Requests:          check_reqs,
					Petty_Cash_Requests:     petty_cash_reqs,
				}, nil
			},
		},
		"user_name": &graphql.Field{
			Type:        user_name,
			Description: "Gathers a user's name for populating inbox requests",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				var user models.User
				user_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				user, userErr := user.FindByID(user_id)
				if userErr != nil {
					panic(userErr)
				}
				return models.UserName{
					ID:   user.ID,
					Name: user.Name,
				}, nil
			},
		},
		"all_users": &graphql.Field{
			Type:        graphql.NewList(user_detail),
			Description: "Gather basic information for all users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				var user models.User
				isAdmin := middleware.ForAdmin(p.Context)
				if !isAdmin {
					panic("you are unauthorized to view this page")
				}
				results, err := user.Findall()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"user_overview": &graphql.Field{
			Type:        user_overview,
			Description: "Overview information for a user (all requests, and basic info), used in administrative queries and views.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				var user models.User
				isAdmin := middleware.ForAdmin(p.Context)
				if !isAdmin {
					panic("you are unauthorized to view this page")
				}
				user_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				user, userErr := user.FindByID(user_id)
				if userErr != nil {
					panic(userErr)
				}
				check_requests, err := user.AggregateChecks(user_id, "", "")
				if err != nil {
					panic(err)
				}
				mileage, mileageErr := user.FindMileage(user_id)
				if mileageErr != nil {
					panic(mileageErr)
				}
				pettyCash, pettyCashErr := user.AggregatePettyCash(user_id)
				if pettyCashErr != nil {
					panic(pettyCashErr)
				}
				return models.User_Overview{
					ID:                      user_id,
					Name:                    user.Name,
					Last_Login:              user.Last_Login,
					Is_Active:               user.Is_Active,
					Permissions:             user.Permissions,
					Admin:                   user.Admin,
					Incomplete_Action_Count: len(user.InComplete_Actions),
					Check_Requests:          check_requests,
					Mileage_Requests:        mileage,
					Petty_Cash_Requests:     pettyCash,
				}, nil
			},
		},
		"user_mileage": &graphql.Field{
			Type:        total_user_mileage,
			Description: "All mileage requests and information pertaining to those requests, for a specified user, over a specified time period.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"start_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				var user models.User
				user_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				isAdmin := middleware.ForAdmin(p.Context)
				contextuser := middleware.ForID(p.Context)
				if !isAdmin {
					if user_id != contextuser {
						panic("you are unauthorized to view this page")
					}
				}
				results, err := user.AggregateMileage(user_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"user_check_requests": &graphql.Field{
			Type:        total_user_check_requests,
			Description: "All check requests and information pertaining to those requests, for a specified user, over a specified time period.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"start_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				var user models.User
				user_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				contextuser := middleware.ForID(p.Context)
				if !isAdmin {
					if user_id != contextuser {
						panic("you are unauthorized to view this page")
					}
				}
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := user.FindCheckReqs(user_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"user_petty_cash_requests": &graphql.Field{
			Type:        total_user_petty_cash,
			Description: "All petty cash requests and information pertaining to those requests, for a specified user, over a specified time period.",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"start_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				user_id, isOK := p.Args["user_id"].(string)
				if !isOK {
					panic("need to enter a valid user id")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				contextuser := middleware.ForID(p.Context)
				if !isAdmin {
					if user_id != contextuser {
						panic("you are unauthorized to view this page")
					}
				}
				var user_petty_cash models.Petty_Cash_Request
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := user_petty_cash.FindByUser(user_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		// mileage queries
		"mileage_overview": &graphql.Field{
			Type:        graphql.NewList(mileage_request_overview),
			Description: "Overview information for a mileage request, used in administrative views.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				if !isAdmin {
					panic("you are unauthorized to view this page")
				}
				var mileage_req models.Mileage_Overview
				results, err := mileage_req.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"mileage_monthly_report": &graphql.Field{
			Type:        graphql.NewList(monthly_mileage_report),
			Description: "Aggregate and gather all mileage requests for a given month and year.",
			Args: graphql.FieldConfigArgument{
				"month": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"year": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				if !isAdmin {
					panic("you are unauthorized to view this page")
				}
				month, validMo := p.Args["month"].(int)
				if !validMo {
					panic("must enter a valid month")
				}
				year, validYear := p.Args["year"].(int)
				if !validYear {
					panic("must enter a valid year")
				}
				users, err := conn.Db.Collection("users").Find(context.TODO(), bson.D{})
				if err != nil {
					panic(err)
				}
				var records []models.Monthly_Mileage_Overview
				for users.Next(context.TODO()) {
					var user models.User
					decode_err := users.Decode(&user)
					if decode_err != nil {
						panic(decode_err)
					}
					user_mileage, err := user.MonthlyMileage(user.ID, month, year)
					if err != nil {
						panic(err)
					}
					user_record := &models.Monthly_Mileage_Overview{
						Grant_IDS:     user_mileage.Grant_IDS,
						User_ID:       user.ID,
						Name:          user.Name,
						Month:         time.Month(month),
						Year:          year,
						Mileage:       user_mileage.Mileage,
						Tolls:         user_mileage.Tolls,
						Parking:       user_mileage.Parking,
						Reimbursement: user_mileage.Reimbursement,
						Requests:      user_mileage.Requests,
					}
					// possible to exclude null records
					records = append(records, *user_record)
				}
				return records, nil
			},
		},
		"mileage_detail": &graphql.Field{
			Type:        mileage_request,
			Description: "Detailed information for a single mileage request, specified by id.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				var milage_req models.Mileage_Request
				mileage_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid request id")
				}
				results, err := milage_req.FindByID(mileage_id)
				contextuser := middleware.ForID(p.Context)
				if !isAdmin {
					if contextuser != results.User_ID {
						panic("you are unauthorized to view this record")
					}
				}
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"grant_mileage_report": &graphql.Field{
			Type:        grant_mileage,
			Description: "Aggregate and gather all mileage requests for a given grant.",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"start_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				isAdmin := middleware.ForAdmin(p.Context)
				if !isAdmin {
					panic("you are unauthorized to view this page")
				}
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				var mileage_request models.Grant_Mileage_Overview
				grant_id, isOk := p.Args["grant_id"].(string)
				if !isOk {
					panic("must enter a valid grant id")
				}
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := mileage_request.FindByGrant(grant_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"mileage_request_variance": &graphql.Field{
			Type:        mileage_request_variance,
			Description: "A query that calculates the variance between a user's actual mileage and recorded mileage",
			Args: graphql.FieldConfigArgument{
				"starting_point": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(location_point_input),
				},
				"ending_point": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(location_point_input),
				},
				"location_points": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(&graphql.List{OfType: location_point_input}),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				start, isOk := p.Args["starting_point"].(map[string]interface{})
				if !isOk {
					panic("missing a valid starting point")
				}
				destination, isOk := p.Args["ending_point"].(map[string]interface{})
				if !isOk {
					panic("missing a valid destination point")
				}
				location_points, isOk := p.Args["location_points"].([]interface{})
				if !isOk {
					panic("missing valid location points")
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
				fmt.Print(locations)
				mileage_points := models.Mileage_Points{
					LocationPoints: locations,
					Starting_Point: models.Location{
						Latitude:  start["latitude"].(float64),
						Longitude: start["longitude"].(float64),
					},
					Destination: models.Location{
						Latitude:  destination["latitude"].(float64),
						Longitude: destination["longitude"].(float64),
					},
				}
				matrix_res, err := mileage_points.CallMatrixAPI()
				if err != nil {
					panic(err)
				}
				calculated_dist := mileage_points.CalculatePreSnapDistance()
				response, err := matrix_res.CompareToMatrix(calculated_dist)
				if err != nil {
					panic(err)
				}
				return response, nil
			},
		},
		// petty cash queries
		"petty_cash_overview": &graphql.Field{
			Type:        graphql.NewList(petty_cash_overview),
			Description: "Gather overview information for all petty cash requests, and basic info, used in administrative views and queries.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				var petty_cash_overview models.Petty_Cash_Overview
				isAdmin := middleware.ForAdmin(p.Context)
				if !isAdmin {
					panic("you are unauthorized to view this page")
				}
				results, err := petty_cash_overview.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"grant_petty_cash_requests": &graphql.Field{
			Type:        grant_petty_cash,
			Description: "Aggregate and gather all petty cash requests for a given grant.",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"start_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				if !isAdmin {
					panic("you are unauthorized to view this page")
				}
				grant_id, isOK := p.Args["grant_id"].(string)
				if !isOK {
					panic("need to enter a valid grant id")
				}
				var grant_petty_cash models.Grant_Petty_Cash
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := grant_petty_cash.FindByGrant(grant_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"petty_cash_detail": &graphql.Field{
			Type:        petty_cash_request,
			Description: "Detailed information for a single petty cash request, specified by id.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				request_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid check request id")
				}
				var petty_cash_req models.Petty_Cash_Request
				collection := conn.Db.Collection("petty_cash_requests")
				filter := bson.D{{Key: "_id", Value: request_id}}
				err := collection.FindOne(context.TODO(), filter).Decode(&petty_cash_req)
				contextuser := middleware.ForID(p.Context)
				if !isAdmin {
					if contextuser != petty_cash_req.User_ID {
						panic("you are unauthorized to view this record")
					}
				}
				if err != nil {
					panic(err)
				}
				return petty_cash_req, nil
			},
		},
		// check request queries
		"check_request_overview": &graphql.Field{
			Type:        graphql.NewList(check_request_overview),
			Description: "Gather overview information for all check requests, and basic info, used in administrative views and queries.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				if !isAdmin {
					panic("you are unauthorized to view this page")
				}
				var check_request models.Check_Request_Overview
				results, err := check_request.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"grant_check_requests": &graphql.Field{
			Type:        grant_check_requests,
			Description: "Aggregate and gather all check requests for a given grant.",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"start_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				if !isAdmin {
					panic("you are unauthorized to view this page")
				}
				var check_request models.Grant_Check_Overview
				grant_id, isOk := p.Args["grant_id"].(string)
				if !isOk {
					panic("must enter a valid grant id")
				}
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := check_request.FindByGrant(grant_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"check_request_detail": &graphql.Field{
			Type:        check_request,
			Description: "Detailed information for a single check request, identified by id.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				isAdmin := middleware.ForAdmin(p.Context)
				request_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid check request id")
				}
				var check_request models.Check_Request
				collection := conn.Db.Collection("check_requests")
				filter := bson.D{{Key: "_id", Value: request_id}}
				err := collection.FindOne(context.TODO(), filter).Decode(&check_request)
				contextuser := middleware.ForID(p.Context)
				if !isAdmin {
					if contextuser != check_request.User_ID {
						panic("you are unauthorized to view this record")
					}
				}
				if err != nil {
					panic(err)
				}
				return check_request, nil
			},
		},
		// grant queries
		"all_grants": &graphql.Field{
			Type:        graphql.NewList(grant),
			Description: "Returns all grant information in the database.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var grant models.Grant
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				results, err := grant.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"single_grant": &graphql.Field{
			Type:        grant,
			Description: "Returns information for a single grant, queried by id.",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var grant models.Grant
				loggedIn := middleware.LoggedIn(p.Context)
				if !loggedIn {
					panic("you are not logged in")
				}
				grant_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid grant id")
				}
				result, err := grant.Find(grant_id)
				if err != nil {
					panic(err)
				}
				return result, nil
			},
		},
	},
})
