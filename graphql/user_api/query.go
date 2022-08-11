package user_api

import (
	u "financial-api/m/models/user"

	"github.com/graphql-go/graphql"
)

var UserQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"overview": &graphql.Field{
			Type:        UserType,
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
				userRes, err := user.FindByID(user_id)
				if err != nil {
					panic(err)
				}
				return userRes, nil
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
		"monthly_petty_cash": &graphql.Field{
			Type:        UserMonthlyPettyCash,
			Description: "Aggregate and gather all petty cash requests for a user for a given month and year",
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
				results, err := user.MonthlyPettyCash(user_id, month, year)
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
					Type:         graphql.DateTime,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.DateTime,
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
