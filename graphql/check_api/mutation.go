package check_api

import (
	"github.com/graphql-go/graphql"
)

var CheckRequestMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Check Request Mutations",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        CheckRequestType,
			Description: "Creates a new check request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"vendor": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(VendorInputType),
				},
				"date": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"purchases": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(&graphql.List{OfType: PurchaseInputType}),
				},
				"receipts": &graphql.ArgumentConfig{
					Type: &graphql.List{OfType: graphql.String},
				},
				"credit_card": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
		},
	},
})
