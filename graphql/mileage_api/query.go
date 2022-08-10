package mileage_api

import (
	"context"
	conn "financial-api/m/db"
	. "financial-api/m/models/requests"
	"financial-api/m/models/user"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

var MileageQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mileage Request Queries",
	Fields: graphql.Fields{
		"overview": &graphql.Field{
			Type:        MileageOverviewType,
			Description: "Gather overview information for all mileage requests, and basic info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var mileage_req Mileage_Overview
				results, err := mileage_req.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"monthly_mileage": &graphql.Field{
			Type:        AggMonthlyMileageType,
			Description: "Aggregate and gather all mileage requests for a given month and year",
			Args: graphql.FieldConfigArgument{
				"month": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"year": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				month, validMo := p.Args["month"].(int)
				if !validMo {
					panic("must enter a valid month")
				}
				year, validYear := p.Args["year"].(int)
				if !validYear {
					panic("must enter a valid year")
				}
				users, err := conn.DB.Collection("users").Find(context.TODO(), bson.D{})
				if err != nil {
					panic(err)
				}
				var records []Monthly_Mileage_Overview
				for users.Next(context.TODO()) {
					var user user.User
					decode_err := users.Decode(&user)
					if decode_err != nil {
						panic(decode_err)
					}
					user_mileage, err := user.MonthlyMileage(user.ID, month, year)
					if err != nil {
						panic(err)
					}
					user_record := &Monthly_Mileage_Overview{
						User_ID:       user.ID,
						Name:          user.Name,
						Month:         time.Month(month),
						Year:          year,
						Mileage:       user_mileage.Mileage,
						Tolls:         user_mileage.Tolls,
						Parking:       user_mileage.Parking,
						Reimbursement: user_mileage.Reimbursement,
						Request_IDS:   user_mileage.Request_IDS,
					}
					records = append(records, *user_record)
				}
				return records, nil
			},
		},
		"detail": &graphql.Field{
			Type:        MileageType,
			Description: "Detailed information for a single mileage request by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var milage_req Mileage_Request
				mileage_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid request id")
				}
				results, err := milage_req.FindByID(mileage_id)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
	},
})
