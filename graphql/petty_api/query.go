package petty_api

import (
	"context"
	conn "financial-api/db"
	r "financial-api/models/requests"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
)

var PettyCashQueries = graphql.NewObject(graphql.ObjectConfig{
	Name: "PettyCashQueries",
	Fields: graphql.Fields{
		"overview": &graphql.Field{
			Type:        graphql.NewList(PettyCashOverviewType),
			Description: "Gather overview information for all petty cash requests, and basic info",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var petty_cash_overview r.Petty_Cash_Overview
				results, err := petty_cash_overview.FindAll()
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"user_requests": &graphql.Field{
			Type:        AggUserPettyCash,
			Description: "Aggregate and gather all petty cash requests for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"start_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_id, isOK := p.Args["user_id"].(string)
				if !isOK {
					panic("need to enter a valid user id")
				}
				var user_petty_cash r.Petty_Cash_Request
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := user_petty_cash.FindByUser(user_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"grant_requests": &graphql.Field{
			Type:        AggGrantPettyCashReq,
			Description: "Aggregate and gather all petty cash requests for a given grant",
			Args: graphql.FieldConfigArgument{
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"start_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
				"end_date": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				grant_id, isOK := p.Args["grant_id"].(string)
				if !isOK {
					panic("need to enter a valid grant id")
				}
				var grant_petty_cash r.Grant_Petty_Cash
				start_date := p.Args["start_date"].(string)
				end_date := p.Args["end_date"].(string)
				results, err := grant_petty_cash.FindByGrant(grant_id, start_date, end_date)
				if err != nil {
					panic(err)
				}
				return results, nil
			},
		},
		"detail": &graphql.Field{
			Type:        PettyCashType,
			Description: "Detailed information for a single petty cash request by id",
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
				var petty_cash_req r.Petty_Cash_Request
				collection := conn.Db.Collection("petty_cash_requests")
				filter := bson.D{{Key: "_id", Value: request_id}}
				err := collection.FindOne(context.TODO(), filter).Decode(&petty_cash_req)
				if err != nil {
					panic(err)
				}
				return petty_cash_req, nil
			},
		},
	},
},
)
