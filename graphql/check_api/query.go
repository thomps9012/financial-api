package check_api

import "github.com/graphql-go/graphql"

var CheckQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Check Request Queries",
	Fields: graphql.Fields{
		"overview": &graphql.Field{
			Type:        CheckReqOverviewType,
			Description: "Gather overview information for all check requests, and basic info",
		},
		"user_requests": &graphql.Field{
			Type:        AggUserCheckReq,
			Description: "Aggregate and gather all check requests for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
		},
		"grant_requests": &graphql.Field{
			Type:        AggGrantCheckReq,
			Description: "Aggregate and gather all check requests for a given grant",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
		},
		"detail": &graphql.Field{
			Type:        CheckRequestType,
			Description: "Detailed information for a single check request by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
		},
	},
})
