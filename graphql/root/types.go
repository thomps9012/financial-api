package root

import (
	"github.com/graphql-go/graphql"
)

// user types
var user_detail = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "user_detail",
		Description: "Detailed information about a specific user, excluding request information",
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
				Type: graphql.NewList(request_action),
			},
			"vehicles": &graphql.Field{
				Type: graphql.NewList(user_vehicle),
			},
		},
	},
)
var user_overview = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "user_overview",
		Description: "Basic information about a user, coupled with overview information about their various requests.",
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
			"vehicles": &graphql.Field{
				Type: graphql.NewList(user_vehicle),
			},
			"incomplete_actions": &graphql.Field{
				Type: graphql.NewList(request_action),
			},
			"incomplete_action_count": &graphql.Field{
				Type: graphql.Int,
			},
			"last_login": &graphql.Field{
				Type: graphql.DateTime,
			},
			"mileage_requests": &graphql.Field{
				Type: aggregate_user_mileage,
			},
			"check_requests": &graphql.Field{
				Type: aggregate_user_check_requests,
			},
			"petty_cash_requests": &graphql.Field{
				Type: aggregate_user_petty_cash,
			},
		},
	},
)

// user aggregates
var aggregate_user_mileage = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "aggregate_user_mileage",
		Description: "The aggregate mileage for a specific user, including detailed information about their last request.",
		Fields: graphql.Fields{
			"vehicles": &graphql.Field{
				Type: graphql.NewList(user_vehicle),
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
				Type: mileage_request,
			},
		},
	},
)
var aggregate_user_check_requests = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "aggregate_user_check_requests",
		Description: "The aggregate total of check requests for a specific user, including detailed information about their last request.",
		Fields: graphql.Fields{
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"receipts": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"vendors": &graphql.Field{
				Type: graphql.NewList(vendor),
			},
			"purchases": &graphql.Field{
				Type: graphql.NewList(purchase_item),
			},
			"total_requests": &graphql.Field{
				Type: graphql.Int,
			},
			"last_request": &graphql.Field{
				Type: check_request,
			},
		},
	},
)
var aggregate_user_petty_cash = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "AggregateUserPettyCash",
		Description: "The total amount of petty cash requested by a given user, including detailed information about their most recent request.",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: user_detail,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"total_requests": &graphql.Field{
				Type: graphql.Int,
			},
			"last_request": &graphql.Field{
				Type: petty_cash_request,
			},
		},
	},
)

// user totals
var total_user_check_requests = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "total_user_check_requests",
		Description: "All of the check requests, including detailed information about each request, for a given user.",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: user_detail,
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
				Type: graphql.NewList(vendor),
			},
			"requests": &graphql.Field{
				Type: graphql.NewList(check_request),
			},
		},
	},
)
var total_user_mileage = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "total_user_mileage",
		Description: "All of the mileage requests, including detailed information about each request, for a given user.",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: user_detail,
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
				Type: graphql.NewList(mileage_request),
			},
		},
	},
)
var total_user_petty_cash = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "total_user_petty_cash",
		Description: "All of the petty cash requests, including detailed information about each request, for a given user.",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: user_detail,
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
				Type: graphql.NewList(petty_cash_request),
			},
		},
	},
)

// mileage types
var mileage_request = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "mileage_request",
		Description: "Detailed information about a specified mileage request.",
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
				Type: graphql.NewList(request_action),
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
var test_mileage_request = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "test_mileage_request",
		Description: "Detailed information about a specified mileage request.",
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
				Type: location_point,
			},
			"destination": &graphql.Field{
				Type: location_point,
			},
			"trip_purpose": &graphql.Field{
				Type: graphql.String,
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
				Type: graphql.NewList(request_action),
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
			"location_points": &graphql.Field{
				Type: graphql.NewList(location_point),
			},
			"request_variance": &graphql.Field{
				Type: mileage_request_variance,
			},
		},
	},
)
var mileage_request_overview = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "mileage_request_overview",
		Description: "Basic information about a mileage request, used for administrative views of the application.",
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
				Type: user_detail,
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
				Type: graphql.String,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)
var monthly_mileage_report = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "monthly_mileage_report",
		Description: "A general monthly mileage report which can be generated for a specified month and year.",
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
				Type: graphql.NewList(mileage_request),
			},
		},
	},
)
var grant_mileage = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "grant_mileage",
		Description: "All of the mileage requests associated with a specified grant.",
		Fields: graphql.Fields{
			"grant": &graphql.Field{
				Type: grant,
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
				Type: graphql.NewList(mileage_request),
			},
		},
	},
)
var mileage_request_variance = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "mileage_variance",
		Description: "Describes the calculated variance between a user's logged milage points and the calculated distance",
		Fields: graphql.Fields{
			"matrix_distance": &graphql.Field{
				Type: graphql.Float,
			},
			"traveled_distance": &graphql.Field{
				Type: graphql.Float,
			},
			"variance": &graphql.Field{
				Type: graphql.String,
			},
			"difference": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)

// petty cash types
var petty_cash_request = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "petty_cash_request",
		Description: "Detailed information about a specified petty cash request.",
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
				Type: graphql.NewList(request_action),
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
var petty_cash_overview = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "petty_cash_overview",
		Description: "Basic information about a petty cash request, used for administrative views and queries.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.ID,
			},
			"user": &graphql.Field{
				Type: user_detail,
			},
			"grant_id": &graphql.Field{
				Type: graphql.ID,
			},
			"grant": &graphql.Field{
				Type: grant,
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
var grant_petty_cash = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "grant_petty_cash",
		Description: "All of the petty cash requests for a specific grant.",
		Fields: graphql.Fields{
			"grant": &graphql.Field{
				Type: grant,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"total_requests": &graphql.Field{
				Type: graphql.Int,
			},
			"requests": &graphql.Field{
				Type: graphql.NewList(petty_cash_request),
			},
		},
	},
)

// check request types
var vendor = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "vendor",
		Description: "Vendor information that is associated with the check request model.",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: address,
			},
		},
	},
)
var purchase_item = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "purchase_item",
		Description: "A purchased item, associated with a specific grant line item.",
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
var check_request = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "check_request",
		Description: "Detailed information about a specific check request.",
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
				Type: vendor,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"purchases": &graphql.Field{
				Type: graphql.NewList(purchase_item),
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
				Type: graphql.NewList(request_action),
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
var check_request_overview = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "check_request_overview",
		Description: "High level information about a check request, used in administrative queries and views.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"user_id": &graphql.Field{
				Type: graphql.ID,
			},
			"user": &graphql.Field{
				Type: user_detail,
			},
			"grant_id": &graphql.Field{
				Type: graphql.ID,
			},
			"grant": &graphql.Field{
				Type: grant,
			},
			"date": &graphql.Field{
				Type: graphql.DateTime,
			},
			"vendor": &graphql.Field{
				Type: vendor,
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
var grant_check_requests = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "grant_check_requests",
		Description: "All of the check requests that are associated with a specific grant.",
		Fields: graphql.Fields{
			"grant": &graphql.Field{
				Type: grant,
			},
			"vendors": &graphql.Field{
				Type: graphql.NewList(vendor),
			},
			"total_requests": &graphql.Field{
				Type: graphql.Int,
			},
			"total_amount": &graphql.Field{
				Type: graphql.Float,
			},
			"requests": &graphql.Field{
				Type: graphql.NewList(check_request),
			},
		},
	},
)

// input types
var location_point_input = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "location_point_input",
		Description: "A singular location point",
		Fields: graphql.InputObjectConfigFieldMap{
			"longitude": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"latitude": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
		},
	},
)

var variance_input = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "mileage_variance_input",
		Description: "Details about the Mileage Request's Variance",
		Fields: graphql.InputObjectConfigFieldMap{
			"matrix_distance": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"traveled_distance": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"difference": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"variance": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(variance_level),
			},
		},
	},
)
var address_input = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "address_input",
		Description: "The input specifications for a check request vendor's address.",
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
var vendor_input = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "vendor_input",
		Description: "The input specifications for a check request vendor.",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"address": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(address_input),
			},
		},
	},
)
var purchase_input = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "purchase_input",
		Description: "The input specifications for a check request purchase item.",
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
var petty_cash_input = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "petty_cash_input",
		Description: "The input specifications for a petty cash request.",
		Fields: graphql.InputObjectConfigFieldMap{
			"category": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(request_category),
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
var mileage_input = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "mileage_input",
		Description: "The input specifications for a mileage request.",
		Fields: graphql.InputObjectConfigFieldMap{
			"date": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"category": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(request_category),
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
var test_mileage_input = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "test_mileage_input",
		Description: "The input specifications for a mileage request.",
		Fields: graphql.InputObjectConfigFieldMap{
			"date": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"category": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(request_category),
			},
			"starting_location": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(location_point_input),
			},
			"destination": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(location_point_input),
			},
			"location_points": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(&graphql.List{OfType: location_point_input}),
			},
			"trip_purpose": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"tolls": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"parking": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"request_variance": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(variance_input),
			},
		},
	},
)
var check_request_input = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "check_request_input",
		Description: "The input specifications for a check request.",
		Fields: graphql.InputObjectConfigFieldMap{
			"category": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(request_category),
			},
			"date": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.DateTime),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"purchases": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(&graphql.List{OfType: purchase_input}),
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
var request_info_error = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name:        "request_info_error",
		Description: "Basic info about a request that caused an error",
		Fields: graphql.InputObjectConfigFieldMap{
			"operation_name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"query": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"request": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(&graphql.InputObject{}),
			},
		},
	},
)

// util types
var address = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "address",
		Description: "Information about a vendor's address.",
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
var request_type = graphql.NewEnum(
	graphql.EnumConfig{
		Name:        "request_type",
		Description: "The type of request, used for action history and query specification.",
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
var request_category = graphql.NewEnum(
	graphql.EnumConfig{
		Name:        "request_category",
		Description: "The category for a specific request, used to determine which supervisors and managers are responsible for approving the request.",
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
var grant = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "grant",
		Description: "Basic information about a grant.",
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
var location_point = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "location_point",
		Description: "A singular location point",
		Fields: graphql.Fields{
			"longitude": &graphql.Field{
				Type: graphql.Float,
			},
			"latitude": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
var variance_level = graphql.NewEnum(
	graphql.EnumConfig{
		Name:        "variance_level",
		Description: "The differing levels of variance between the user's calculated and actual mileage request",
		Values: graphql.EnumValueConfigMap{
			"HIGH": &graphql.EnumValueConfig{
				Value: "HIGH",
			},
			"MEDIUM": &graphql.EnumValueConfig{
				Value: "MEDIUM",
			},
			"LOW": &graphql.EnumValueConfig{
				Value: "LOW",
			},
		},
	},
)
var request_status = graphql.NewEnum(
	graphql.EnumConfig{
		Name:        "request_status",
		Description: "The different levels that a request can maintain throughout the approval process.",
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
var request_action = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "request_action",
		Description: "Basic information about any action taken on a specific request.",
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
var user_vehicle = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "user_vehicle",
		Description: "Basic information about a user's vehicle.",
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
