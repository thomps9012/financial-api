package petty_api

import "github.com/graphql-go/graphql"

var PettyCashQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Petty Cash Request Queries",
	Fields: graphql.Fields{
		"overview": &graphql.Field{
			Type:        PettyCashOverviewType,
			Description: "Gather overview information for all petty cash requests, and basic info",
		},
		"user_requests": &graphql.Field{
			Type:        AggUserPettyCash,
			Description: "Aggregate and gather all petty cash requests for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
		},
		"grant_requests": &graphql.Field{
			Type:        AggGrantPettyCashReq,
			Description: "Aggregate and gather all petty cash requests for a given grant",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
		},
		"detail": &graphql.Field{
			Type:        PettyCashType,
			Description: "Detailed information for a single petty cash request by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
		},
	},
})
