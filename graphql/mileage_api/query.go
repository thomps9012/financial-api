package mileage_api

import "github.com/graphql-go/graphql"

var MileageQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mileage Request Queries",
	Fields: graphql.Fields{
		"overview": &graphql.Field{
			Type:        MileageOverviewType,
			Description: "Gather overview information for all mileage requests, and basic info",
		},
		"monthly_mileage": &graphql.Field{
			Type:        AggMonthlyMileageType,
			Description: "Aggregate and gather all mileage requests for a given month and year",
			Args: graphql.FieldConfigArgument{
				"month": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"year": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
		},
		"detail": &graphql.Field{
			Type:        MileageType,
			Description: "Detailed information for a single mileage request by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
		},
	},
})
