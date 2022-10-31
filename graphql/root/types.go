package root

import (
	"financial-api/models"

	"github.com/graphql-go/graphql"
)

var UserDetailType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserDetail",
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
		"admin": &graphql.Field{
			Type: graphql.Boolean,
		},
		"permissions": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
		"vehicles": &graphql.Field{
			Type: graphql.NewList(VehicleType),
		},
		"incomplete_actions": &graphql.Field{
			Type: graphql.NewList(ActionType),
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
})

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

var RequestType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "RequestType",
		Values: graphql.EnumValueConfigMap{
			"CHECK": &graphql.EnumValueConfig{
				Value: "CHECK",
			},
			"MILEAGE": &graphql.EnumValueConfig{
				Value: "MILEAGE",
			},
			"PETTY_CASH": &graphql.EnumValueConfig{
				Value: "PETTY_CASH",
			},
		},
	},
)

var CategoryType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "Category",
		Values: graphql.EnumValueConfigMap{
			"IOP": &graphql.EnumValueConfig{
				Value: "IOP",
			},
			"INTAKE": &graphql.EnumValueConfig{
				Value: "INTAKE",
			},
			"PEERS": &graphql.EnumValueConfig{
				Value: "PEERS",
			},
			"ACT_TEAM": &graphql.EnumValueConfig{
				Value: "ACT_TEAM",
			},
			"IHBT": &graphql.EnumValueConfig{
				Value: "IHBT",
			},
			"PERKINS": &graphql.EnumValueConfig{
				Value: "PERKINS",
			},
			"MENS_HOUSE": &graphql.EnumValueConfig{
				Value: "MENS_HOUSE",
			},
			"NEXT_STEP": &graphql.EnumValueConfig{
				Value: "NEXT_STEP",
			},
			"LORAIN": &graphql.EnumValueConfig{
				Value: "LORAIN",
			},
			"PREVENTION": &graphql.EnumValueConfig{
				Value: "PREVENTION",
			},
			"ADMINISTRATIVE": &graphql.EnumValueConfig{
				Value: "ADMINISTRATIVE",
			},
			"FINANCE": &graphql.EnumValueConfig{
				Value: "FINANCE",
			},
		},
	},
)

var CheckRequestInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CheckRequestInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"category": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(CategoryType),
			},
			"date": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
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
			"category": &graphql.Field{
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
				Type: graphql.NewList(ActionType),
			},
			"current_status": &graphql.Field{
				Type: graphql.String,
			},
			"current_user": &graphql.Field{
				Type: graphql.ID,
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
				Type: UserType,
			},
			"grant_id": &graphql.Field{
				Type: graphql.ID,
			},
			"grant": &graphql.Field{
				Type: models.GrantType,
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
				Type: graphql.String,
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
			"grant": &graphql.Field{
				Type: models.GrantType,
			},
			"vendors": &graphql.Field{
				Type: graphql.NewList(VendorType),
			},
			"total_requests": &graphql.Field{
				Type: graphql.Int,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"requests": &graphql.Field{
				Type: graphql.NewList(CheckRequestType),
			},
		},
	},
)

var PettyCashType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PettyCashRequest",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
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
				Type: graphql.NewList(ActionType),
			},
			"current_status": &graphql.Field{
				Type: graphql.String,
			},
			"current_user": &graphql.Field{
				Type: graphql.ID,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var PettyCashOverviewType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PettyCashOverview",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.ID,
			},
			"user": &graphql.Field{
				Type: UserType,
			},
			"grant_id": &graphql.Field{
				Type: graphql.ID,
			},
			"grant": &graphql.Field{
				Type: models.GrantType,
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
				Type: graphql.String,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var UserAggPettyCash = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "UserAggPettyCash",
		Description: "the query for a specified users petty cash requests",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: UserType,
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
			"requests": &graphql.Field{
				Type: graphql.NewList(PettyCashType),
			},
			"last_request": &graphql.Field{
				Type: PettyCashType,
			},
		},
	},
)

var AggUserPettyCash = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AggregateUserPettyCash",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: UserType,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"total_requests": &graphql.Field{
				Type: graphql.Int,
			},
			"last_request": &graphql.Field{
				Type: PettyCashType,
			},
		},
	},
)

var PettyCashInput = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "PettyCashInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"category": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(CategoryType),
			},
			"amount": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"date": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"receipts": &graphql.InputObjectFieldConfig{
				Type: &graphql.List{OfType: graphql.String},
			},
		},
	},
)

var AggGrantPettyCashReq = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "AggregateGrantPettyCash",
		Fields: graphql.Fields{
			"grant": &graphql.Field{
				Type: models.GrantType,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"total_requests": &graphql.Field{
				Type: graphql.Int,
			},
			"requests": &graphql.Field{
				Type: graphql.NewList(PettyCashType),
			},
		},
	},
)

var MileageInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "MileageInputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"date": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"category": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(CategoryType),
			},
			"starting_location": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"destination": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"trip_purpose": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"start_odometer": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"end_odometer": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"tolls": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"parking": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},
	},
)

var MileageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MileageRequest",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"category": &graphql.Field{
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
			"starting_location": &graphql.Field{
				Type: graphql.String,
			},
			"destination": &graphql.Field{
				Type: graphql.String,
			},
			"trip_purpose": &graphql.Field{
				Type: graphql.String,
			},
			"start_odometer": &graphql.Field{
				Type: graphql.Int,
			},
			"end_odometer": &graphql.Field{
				Type: graphql.Int,
			},
			"tolls": &graphql.Field{
				Type: graphql.Float,
			},
			"parking": &graphql.Field{
				Type: graphql.Float,
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
			"action_history": &graphql.Field{
				Type: graphql.NewList(ActionType),
			},
			"current_user": &graphql.Field{
				Type: graphql.ID,
			},
			"current_status": &graphql.Field{
				Type: graphql.String,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
var MileageOverviewType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MileageRequestOverview",
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
			"user": &graphql.Field{
				Type: UserType,
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

var AggMonthlyMileageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "MonthlyMileageRequests",
		Fields: graphql.Fields{
			"user_id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
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
			"requests": &graphql.Field{
				Type: graphql.NewList(MileageType),
			},
		},
	},
)
var AggGrantMileage = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GrantMileageRequests",
		Fields: graphql.Fields{
			"grant": &graphql.Field{
				Type: models.GrantType,
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
			"requests": &graphql.Field{
				Type: graphql.NewList(MileageType),
			},
		},
	},
)

// change
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
			"SUPERVISOR_APPROVED": &graphql.EnumValueConfig{
				Value: "SUPERVISOR_APPROVED",
			},
			"FINANCE_APPROVED": &graphql.EnumValueConfig{
				Value: "FINANCE_APPROVED",
			},
			"EXECUTIVE_APPROVED": &graphql.EnumValueConfig{
				Value: "EXECUTIVE_APPROVED",
			},
			"ORGANIZATION_APPROVED": &graphql.EnumValueConfig{
				Value: "ORGANIZATION_APPROVED",
			},
			"REJECTED": &graphql.EnumValueConfig{
				Value: "REJECTED",
			},
			"REJECTED_EDIT": &graphql.EnumValueConfig{
				Value: "REJECTED_EDIT",
			},
			"ARCHIVED": &graphql.EnumValueConfig{
				Value: "ARCHIVED",
			},
		},
	},
)
var PermissionType = graphql.NewEnum(
	graphql.EnumConfig{
		Name: "Permission",
		Values: graphql.EnumValueConfigMap{
			"EMPLOYEE": &graphql.EnumValueConfig{
				Value: "EMPLOYEE",
			},
			"MANAGER": &graphql.EnumValueConfig{
				Value: "MANAGER",
			},
			"SUPERVISOR": &graphql.EnumValueConfig{
				Value: "SUPERVISOR",
			},
			"EXECUTIVE": &graphql.EnumValueConfig{
				Value: "EXECUTIVE",
			},
			"FINANCE_TEAM": &graphql.EnumValueConfig{
				Value: "FINANCE_TEAM",
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
			"user": &graphql.Field{
				Type: graphql.String,
			},
			"request_type": &graphql.Field{
				Type: graphql.String,
			},
			"request_id": &graphql.Field{
				Type: graphql.ID,
			},
			"status": &graphql.Field{
				Type: graphql.String,
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
				Type: AddressType,
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
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"admin": &graphql.Field{
				Type: graphql.Boolean,
			},
			"permissions": &graphql.Field{
				Type: graphql.NewList(graphql.String),
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
		Name:        "AggUserMileage",
		Description: "The aggregate user mileage for the user overview query",
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
			"total_requests": &graphql.Field{
				Type: graphql.Int,
			},
			"last_request": &graphql.Field{
				Type: MileageType,
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
			"total_requests": &graphql.Field{
				Type: graphql.Int,
			},
			"last_request": &graphql.Field{
				Type: CheckRequestType,
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
				Type: models.GrantType,
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
		Name: "UserCheckReqOverview",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"grant_id": &graphql.Field{
				Type: graphql.ID,
			},
			"grant": &graphql.Field{
				Type: models.GrantType,
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
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"admin": &graphql.Field{
				Type: graphql.Boolean,
			},
			"permissions": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
			"last_login": &graphql.Field{
				Type: graphql.DateTime,
			},
			"incomplete_actions": &graphql.Field{
				Type: graphql.NewList(ActionType),
			},
			"vehicles": &graphql.Field{
				Type: graphql.NewList(VehicleType),
			},
		},
	},
)

var UserAggCheckRequests = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "UserAggCheckRequests",
		Description: "the aggregate check requests for a specified user",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: UserType,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"start_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"end_date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"vendors": &graphql.Field{
				Type: graphql.NewList(VendorType),
			},
			"requests": &graphql.Field{
				Type: graphql.NewList(CheckRequestType),
			},
		},
	},
)

var UserAggMileage = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "UserAggMileage",
		Description: "the aggregate user mileage for the query that looks up mileage specific to a specified user",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: UserType,
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
			"requests": &graphql.Field{
				Type: graphql.NewList(MileageType),
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

var GrantType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "GrantType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
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
			"requests": &graphql.Field{
				Type: graphql.NewList(CheckRequestType),
			},
		},
	},
)
