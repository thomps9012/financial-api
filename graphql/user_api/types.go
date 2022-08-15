package user_api

import (
	g "financial-api/models/grants"

	"github.com/graphql-go/graphql"
)

var RoleType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "Role",
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
		Name: "Status",
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
		Name: "RequestAction",
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
		Name: "VehicleInformation",
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

var UserVendorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserVendor",
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
var UserPurchaseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserPurchase",
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

var UserOverviewType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserOverviewType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"manager_id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"role": &graphql.Field{
				Type: RoleType,
			},
			"incomplete_action_count": &graphql.Field{
				Type: graphql.Int,
			},
			"last_login": &graphql.Field{
				Type: graphql.DateTime,
			},
			"mileage_requests": &graphql.Field{
				Type: AggUserMileage,
			},
			"check_requests": &graphql.Field{
				Type: AggUserChecks,
			},
			"petty_cash_requests": &graphql.Field{
				Type: AggUserPettyCash,
			},
		},
	},
)

var AggUserMileage = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AggUserMileage",
		Fields: graphql.Fields{
			"vehicles": &graphql.Field{
				Type: graphql.NewList(VehicleType),
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
			"request_ids": &graphql.Field{
				Type: graphql.NewList(graphql.ID),
			},
		},
	},
)
var AggUserChecks = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AggUserChecks",
		Fields: graphql.Fields{
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"receipts": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"vendors": &graphql.Field{
				Type: graphql.NewList(UserVendorType),
			},
			"purchases": &graphql.Field{
				Type: graphql.NewList(UserPurchaseType),
			},
			"request_ids": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
var AggUserPettyCash = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AggUserPettyCash",
		Fields: graphql.Fields{
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"receipts": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"request_ids": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)

var UserMileageOverview = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserMileageOverview",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"trip_mileage": &graphql.Field{
				Type: graphql.Int,
			},
			"reimbursement": &graphql.Field{
				Type: graphql.Float,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"current_status": &graphql.Field{
				Type: StatusType,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var UserPettyCashOverview = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserPettyCashOverview",
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
				Type: StatusType,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var UserCheckReqOverview = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserCheckRequestOverview",
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
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"vendor": &graphql.Field{
				Type: UserVendorType,
			},
			"order_total": &graphql.Field{
				Type: graphql.Float,
			},
			"created_at": &graphql.Field{
				Type: graphql.DateTime,
			},
			"current_status": &graphql.Field{
				Type: StatusType,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
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
		Name: "UserMonthlyMileageRequests",
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
				Type: graphql.String,
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
			"request_ids": &graphql.Field{
				Type: graphql.NewList(graphql.ID),
			},
		},
	},
)
var UserMonthlyPettyCash = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserMonthlyPettyCashRequests",
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
			"total_amount": &graphql.Field{
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
		Name: "UserTotalCheckRequests",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"start_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"end_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"vendors": &graphql.Field{
				Type: graphql.NewList(UserVendorType),
			},
			"purchases": &graphql.Field{
				Type: graphql.NewList(UserPurchaseType),
			},
			"receipts": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"request_ids": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
		},
	},
)
