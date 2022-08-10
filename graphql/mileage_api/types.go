package mileage_api

import (
	u "financial-api/m/graphql/user_api"

	"github.com/graphql-go/graphql"
)

var MileageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mileage Request",
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
var MileageOverviewType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mileage Request Overview",
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
				Type: u.StatusType,
			},
			"is_active": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
	},
)

var AggMonthlyMileageType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Monthly Mileage Requests",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"month": &graphql.Field{
				Type: graphql.Int,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"user_id": &graphql.Field{
				Type: graphql.ID,
			},
			"user": &graphql.Field{
				Type: u.UserType,
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
