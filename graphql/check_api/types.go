package check_api

import (
	u "financial-api/m/graphql/user_api"
	g "financial-api/m/models/grants"

	"github.com/graphql-go/graphql"
)

var VendorInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Vendor Input",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"address": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)
var VendorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Vendor Information",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var PurchaseInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "Purchase Input",
		Fields: graphql.InputObjectConfigFieldMap{
			"grant_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"grant_line_item": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"amount": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},
	},
)

var PurchaseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Vendor Information",
		Fields: graphql.Fields{
			"grant_id": &graphql.Field{
				Type: graphql.String,
			},
			"grant_line_item": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"amount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

var CheckRequestType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Check Request",
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
			"vendor": &graphql.Field{
				Type: VendorType,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"purchases": &graphql.Field{
				Type: graphql.NewList(PurchaseType),
			},
			"receipts": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"order_total": &graphql.Field{
				Type: graphql.Float,
			},
			"credit_card": &graphql.Field{
				Type: graphql.String,
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

var CheckReqOverviewType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Check Request Overview",
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
			"grant": &graphql.Field{
				Type: g.GrantType,
			},
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"vendor": &graphql.Field{
				Type: VendorType,
			},
			"order_total": &graphql.Field{
				Type: graphql.Float,
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

var AggUserCheckReq = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User Check Requests",
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
			"vendor": &graphql.Field{
				Type: VendorType,
			},
			"order_total": &graphql.Field{
				Type: graphql.Float,
			},
			"credit_card": &graphql.Field{
				Type: graphql.String,
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

var AggGrantCheckReq = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Grant Check Requests",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"grant_id": &graphql.Field{
				Type: graphql.ID,
			},
			"grant": &graphql.Field{
				Type: g.GrantType,
			},
			"vendor": &graphql.Field{
				Type: VendorType,
			},
			"order_total": &graphql.Field{
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
