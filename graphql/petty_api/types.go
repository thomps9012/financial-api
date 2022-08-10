package petty_api

import (
	u "financial-api/m/graphql/user_api"
	g "financial-api/m/models/grants"

	"github.com/graphql-go/graphql"
)

var PettyCashType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Petty Cash Request",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.ID,
			},
			"user": &graphql.Field{
				Type: u.UserType,
			},
			"grant_id": &graphql.Field{
				Type: graphql.ID,
			},
			"grant": &graphql.Field{
				Type: g.GrantType,
			},
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"amount": &graphql.Field{
				Type: graphql.Float,
			},
			"receipts": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"action_history": &graphql.Field{
				Type: graphql.NewList(u.ActionType),
			},
			"current_status": &graphql.Field{
				Type: u.StatusType,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var PettyCashOverviewType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Petty Cash Request Overview",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.ID,
			},
			"user": &graphql.Field{
				Type: u.UserType,
			},
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"amount": &graphql.Field{
				Type: graphql.Float,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"current_status": &graphql.Field{
				Type: u.StatusType,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var AggUserPettyCash = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Aggregate User Petty Cash Requests",
		Fields: graphql.Fields{
			"user_id": &graphql.Field{
				Type: graphql.ID,
			},
			"user": &graphql.Field{
				Type: u.UserType,
			},
			"last_request": &graphql.Field{
				Type: graphql.DateTime,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var AggGrantPettyCashReq = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Aggregate Grant Petty Cash Requests",
		Fields: graphql.Fields{
			"grant_id": &graphql.Field{
				Type: graphql.ID,
			},
			"grant": &graphql.Field{
				Type: g.GrantType,
			},
			"last_request": &graphql.Field{
				Type: graphql.DateTime,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
