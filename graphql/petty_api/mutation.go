package petty_api

import (
	"github.com/graphql-go/graphql"
)

var PettyCashMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Petty Cash Mutations",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        PettyCashType,
			Description: "Creates a new petty cash request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"amount": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"date": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"receipts": &graphql.ArgumentConfig{
					Type: &graphql.List{OfType: graphql.String},
				},
			},
		},
	},
})
