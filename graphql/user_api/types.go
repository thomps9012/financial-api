package user_api

import "github.com/graphql-go/graphql"

var RoleType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "User Organization Role",
		Values: graphql.EnumValueConfigMap{
			"EMPLOYEE": &graphql.EnumValueConfig{
				Value: "EMPLOYEE",
			},
			"MANAGER": &graphql.EnumValueConfig{
				Value: "MANAGER",
			},
			"FINANCE": &graphql.EnumValueConfig{
				Value: "FINANCE",
			},
			"EXECUTIVE": &graphql.EnumValueConfig{
				Value: "EXECUTIVE",
			},
		},
	},
)

var StatusType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "Request Status",
		Values: graphql.EnumValueConfigMap{
			"PENDING": &graphql.EnumValueConfig{
				Value: "PENDING",
			},
			"MANAGER_APPROVED": &graphql.EnumValueConfig{
				Value: "MANAGER_APPROVED",
			},
			"FINANCE_APPROVED": &graphql.EnumValueConfig{
				Value: "FINANCE_APPROVED",
			},
			"ORG_APPROVED": &graphql.EnumValueConfig{
				Value: "ORG_APPROVED",
			},
			"REJECTED": &graphql.EnumValueConfig{
				Value: "REJECTED",
			},
			"ARCHIVED": &graphql.EnumValueConfig{
				Value: "ARCHIVED",
			},
		},
	},
)

var ActionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Request Action",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: StatusType,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

var VehicleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Vehicle Information",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
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

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"manager_id": &graphql.Field{
				Type: graphql.ID,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: RoleType,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
			"last_login": &graphql.Field{
				Type: graphql.DateTime,
			},
			"incomplete_actions": &graphql.Field{
				Type: graphql.NewList(graphql.ID),
			},
			"vehicles": &graphql.Field{
				Type: graphql.NewList(VehicleType),
			},
		},
	},
)

var UserMonthlyMileageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User Monthly Mileage Requests",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"vehicles": &graphql.Field{
				Type: graphql.NewList(VehicleType),
			},
			"month": &graphql.Field{
				Type: graphql.Int,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"mileage": &graphql.Field{
				Type: graphql.Int,
			},
			"tolls": &graphql.Field{
				Type: graphql.Float,
			},
			"parking": &graphql.Field{
				Type: graphql.Float,
			},
			"reimbursement": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
var UserMonthlyPettyCash = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User Monthly Petty Cash Requests",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"month": &graphql.Field{
				Type: graphql.Int,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"amount": &graphql.Field{
				Type: graphql.Float,
			},
			"receipts": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
var UserCheckRequests = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User Monthly Check Requests",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"month": &graphql.Field{
				Type: graphql.Int,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"amount": &graphql.Field{
				Type: graphql.Float,
			},
			"vendors": &graphql.Field{
				Type: graphql.NewList(VendorType),
			},
			"purchases": &graphql.Field{
				Type: graphql.NewList(PurchaseType),
			},
			"receipts": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
