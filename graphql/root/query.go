package root

import (
	"context"
	conn "financial-api/db"
	r "financial-api/models/requests"
	u "financial-api/models/user"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

var RootQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQueries",
	Fields: graphql.Fields{
		// user queries
		"me": &graphql.Field{
			Type:        UserDetailType,
			Description: "Gather Information for a specific user on login",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				user_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				filter := bson.D{{Key: "user_id", Value: user_id}}
				userCollection := conn.Db.Collection("users")
				err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
				if err != nil {
					panic(err)
				}
				petty_cash_reqs, petty_err := user.FindPettyCash(user_id)
				if petty_err != nil {
					panic(err)
				}
				mileage_reqs, mileage_err := user.FindMileage(user_id)
				if mileage_err != nil {
					panic(err)
				}
				check_reqs, check_err := user.AggregateChecks(user_id, "", "")
				if check_err != nil {
					panic(err)
				}
				return &u.User_Detail{
					ID:                      user.ID,
					Name:                    user.Name,
					Manager_ID:              user.Manager_ID,
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
		"all_users": &graphql.Field{
			Type:        graphql.NewList(UserType),
			Description: "Gather basic information for all users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				results, err := user.Findall()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"user_overview": &graphql.Field{
			Type:        UserOverviewType,
			Description: "Gather overview information for a user (all requests, and basic info)",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
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
				pettyCash, pettyCashErr := user.FindPettyCash(user_id)
				if pettyCashErr != nil {
					panic(pettyCashErr)
				}
				return u.User_Overview{
					ID:                      user_id,
					Name:                    user.Name,
					Last_Login:              user.Last_Login,
					Is_Active:               user.Is_Active,
					Role:                    user.Role,
					Manager_ID:              user.Manager_ID,
					Incomplete_Action_Count: len(user.InComplete_Actions),
					Check_Requests:          check_requests,
					Mileage_Requests:        mileage,
					Petty_Cash_Requests:     pettyCash,
				}, nil
			},
		},
		"user_monthly_mileage": &graphql.Field{
			Type:        UserMonthlyMileageType,
			Description: "Aggregate and gather all mileage requests for a user for a given month and year",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"month": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"year": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				user_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				month, validMo := p.Args["month"].(int)
				if !validMo {
					panic("must enter a valid month")
				}
				year, validYear := p.Args["year"].(int)
				if !validYear {
					panic("must enter a valid year")
				}
				results, err := user.MonthlyMileage(user_id, month, year)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"user_check_requests": &graphql.Field{
			Type:        UserCheckRequests,
			Description: "Aggregate and gather all check requests for a user over a given time period",
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
				var user u.User
				user_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := user.AggregateChecks(user_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		// mileage queries
		"mileage_overview": &graphql.Field{
			Type:        graphql.NewList(MileageOverviewType),
			Description: "Gather overview information for all mileage requests, and basic info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var mileage_req r.Mileage_Overview
				results, err := mileage_req.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"mileage_monthly_report": &graphql.Field{
			Type:        graphql.NewList(AggMonthlyMileageType),
			Description: "Aggregate and gather all mileage requests for a given month and year",
			Args: graphql.FieldConfigArgument{
				"month": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"year": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
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
				var records []r.Monthly_Mileage_Overview
				for users.Next(context.TODO()) {
					var user u.User
					decode_err := users.Decode(&user)
					if decode_err != nil {
						panic(decode_err)
					}
					user_mileage, err := user.MonthlyMileage(user.ID, month, year)
					if err != nil {
						panic(err)
					}
					user_record := &r.Monthly_Mileage_Overview{
						User_ID:       user.ID,
						Name:          user.Name,
						Month:         time.Month(month),
						Year:          year,
						Mileage:       user_mileage.Mileage,
						Tolls:         user_mileage.Tolls,
						Parking:       user_mileage.Parking,
						Reimbursement: user_mileage.Reimbursement,
						Request_IDS:   user_mileage.Request_IDS,
					}
					// possible to exclude null records
					records = append(records, *user_record)
				}
				return records, nil
			},
		},
		"mileage_detail": &graphql.Field{
			Type:        MileageType,
			Description: "Detailed information for a single mileage request by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var milage_req r.Mileage_Request
				mileage_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid request id")
				}
				results, err := milage_req.FindByID(mileage_id)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		// petty cash queries
		"petty_cash_overview": &graphql.Field{
			Type:        graphql.NewList(PettyCashOverviewType),
			Description: "Gather overview information for all petty cash requests, and basic info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var petty_cash_overview r.Petty_Cash_Overview
				results, err := petty_cash_overview.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"petty_cash_grant_requests": &graphql.Field{
			Type:        AggGrantPettyCashReq,
			Description: "Aggregate and gather all petty cash requests for a given grant",
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
				grant_id, isOK := p.Args["grant_id"].(string)
				if !isOK {
					panic("need to enter a valid grant id")
				}
				var grant_petty_cash r.Grant_Petty_Cash
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := grant_petty_cash.FindByGrant(grant_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"petty_cash_user_requests": &graphql.Field{
			Type:        AggUserPettyCash,
			Description: "Aggregate and gather all petty cash requests for a given user",
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
				user_id, isOK := p.Args["user_id"].(string)
				if !isOK {
					panic("need to enter a valid user id")
				}
				var user_petty_cash r.Petty_Cash_Request
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := user_petty_cash.FindByUser(user_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"petty_cash_detail": &graphql.Field{
			Type:        PettyCashType,
			Description: "Detailed information for a single petty cash request by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				request_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid check request id")
				}
				var petty_cash_req r.Petty_Cash_Request
				collection := conn.Db.Collection("petty_cash_requests")
				filter := bson.D{{Key: "_id", Value: request_id}}
				err := collection.FindOne(context.TODO(), filter).Decode(&petty_cash_req)
				if err != nil {
					panic(err)
				}
				return petty_cash_req, nil
			},
		},
		// check request queries
		"check_request_overview": &graphql.Field{
			Type:        graphql.NewList(CheckReqOverviewType),
			Description: "Gather overview information for all check requests, and basic info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var check_request r.Check_Request_Overview
				results, err := check_request.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"check_request_grant_requests": &graphql.Field{
			Type:        AggGrantCheckReq,
			Description: "Aggregate and gather all check requests for a given grant",
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
				var check_request r.Grant_Check_Overview
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
			Type:        CheckRequestType,
			Description: "Detailed information for a single check request by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				request_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid check request id")
				}
				var check_request r.Check_Request
				collection := conn.Db.Collection("check_requests")
				filter := bson.D{{Key: "_id", Value: request_id}}
				err := collection.FindOne(context.TODO(), filter).Decode(&check_request)
				if err != nil {
					panic(err)
				}
				return check_request, nil
			},
		},
	},
})