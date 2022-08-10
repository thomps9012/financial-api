package user_api

import (
	"github.com/graphql-go/graphql"
)

var UserQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "User Queries",
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
				// return , nil
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
				// return , nil
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
				// return , nil
			},
		},
		"check_requests": &graphql.Field{
			Type:        UserCheckRequests,
			Description: "Aggregate and gather all check requests for a user",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// return , nil
			},
		},
	},
})
