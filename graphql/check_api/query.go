package check_api

import (
	"context"
	conn "financial-api/m/db"
	r "financial-api/m/models/requests"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

var CheckQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "CheckQueries",
	Fields: graphql.Fields{
		"overview": &graphql.Field{
			Type:        graphql.NewList(CheckReqOverviewType),
			Description: "Gather overview information for all check requests, and basic info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var check_request r.Check_Request_Overview
				results, err := check_request.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		// "user_requests": &graphql.Field{
		// 	Type:        AggUserCheckReq,
		// 	Description: "Aggregate and gather all check requests for a given user",
		// 	Args: graphql.FieldConfigArgument{
		// 		"user_id": &graphql.ArgumentConfig{
		// 			Type: graphql.NewNonNull(graphql.ID),
		// 		},
		// 		"start_date": &graphql.ArgumentConfig{
		// 			Type:         graphql.DateTime,
		// 			DefaultValue: "",
		// 		},
		// 		"end_date": &graphql.ArgumentConfig{
		// 			Type:         graphql.DateTime,
		// 			DefaultValue: "",
		// 		},
		// 	},
		// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// 		var check_request r.User_Check_Overview
		// 		user_id, isOk := p.Args["user_id"].(string)
		// 		if !isOk {
		// 			panic("must enter a valid user id")
		// 		}
		// 		start_date := p.Args["start_date"].(string)
		// 		end_date := p.Args["end_date"].(string)
		// 		results, err := check_request.FindByUser(user_id, start_date, end_date)
		// 		if err != nil {
		// 			panic(err)
		// 		}
		// 		return results, nil
		// 	},
		// },
		"grant_requests": &graphql.Field{
			Type:        AggGrantCheckReq,
			Description: "Aggregate and gather all check requests for a given grant",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"start_date": &graphql.ArgumentConfig{
					Type:         graphql.DateTime,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.DateTime,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var check_request r.Grant_Check_Overview
				grant_id, isOk := p.Args["grant_id"].(string)
				if !isOk {
					panic("must enter a valid grant id")
				}
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := check_request.FindByGrant(grant_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				request_id, isOk := p.Args["id"].(string)
				if !isOk {
					panic("must enter a valid check request id")
				}
				var check_request r.Check_Request
				collection := conn.Db.Collection("check_requests")
				filter := bson.D{{Key: "_id", Value: request_id}}
				err := collection.FindOne(context.TODO(), filter).Decode(&check_request)
				if err != nil {
					panic(err)
				}
				return check_request, nil
			},
		},
	},
})
