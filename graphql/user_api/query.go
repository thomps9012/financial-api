package user_api

import (
	u "financial-api/models/user"

	"github.com/graphql-go/graphql"
)

var UserQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"all": &graphql.Field{
			Type:        graphql.NewList(UserType),
			Description: "Gather basic information for all users",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var user u.User
				results, err := user.Findall(p.Context)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"overview": &graphql.Field{
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
		"monthly_mileage": &graphql.Field{
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
		"check_requests": &graphql.Field{
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
	},
},
)
