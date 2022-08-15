package check_api

import (
	u "financial-api/m/graphql/user_api"
	g "financial-api/m/models/grants"

	"github.com/graphql-go/graphql"
)

var AddressInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "VendorAddressInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"website": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"street": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"city": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"state": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"zip": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
	},
)

var AddressType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "VendorAddress",
		Fields: graphql.Fields{
			"website": &graphql.Field{
				Type: graphql.String,
			},
			"street": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"zip": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var CheckRequestInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CheckRequestInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"date": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"grant_id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"purchases": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(&graphql.List{OfType: PurchaseInputType}),
			},
			"receipts": &graphql.InputObjectFieldConfig{
				Type: &graphql.List{OfType: graphql.String},
			},
			"credit_card": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var VendorInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "VendorInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"address": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(AddressInputType),
			},
		},
	},
)
var VendorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Vendor",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: AddressType,
			},
		},
	},
)

type PurchaseInputStruct struct {
	grant_line_item string
	description     string
	amount          float64
}

var PurchaseInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "PurchaseInput",
		Fields: graphql.InputObjectConfigFieldMap{
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
		Name: "Purchase",
		Fields: graphql.Fields{
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
		Name: "CheckRequest",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.ID,
			},
			"grant_id": &graphql.Field{
				Type: graphql.ID,
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
		Name: "CheckRequestOverview",
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

var AggGrantCheckReq = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GrantCheckRequests",
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
			"vendors": &graphql.Field{
				Type: graphql.NewList(VendorType),
			},
			"credit_cards": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"last_request": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
