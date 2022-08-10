package mileage_api

import (
	"github.com/graphql-go/graphql"
)

var MileageMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mileage Mutations",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        MileageType,
			Description: "Creates a new mileage request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"date": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"starting_location": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"destination": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"trip_purpose": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"start_odometer": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"end_odometer": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"tolls": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"parking": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
			},
		},
	},
})
