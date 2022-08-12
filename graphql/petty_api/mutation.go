package petty_api

import (
	"errors"
	r "financial-api/m/models/requests"
	"time"

	"github.com/graphql-go/graphql"
)

var PettyCashMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "PettyCashMutations",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        PettyCashType,
			Description: "Creates a new petty cash request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"grant_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"amount": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"date": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"receipts": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				petty_cash_req := &r.Petty_Cash_Request{
					Date:        p.Args["date"].(time.Time),
					Grant_ID:    p.Args["grant_id"].(string),
					Amount:      p.Args["amount"].(float64),
					Description: p.Args["description"].(string),
					Receipts:    p.Args["receipts"].([]string),
				}
				user_id, isOk := p.Args["user_id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				amount, okAmount := p.Args["amount"].(float64)
				if !okAmount {
					panic("must enter a valid amount")
				}
				date, okdate := p.Args["date"].(time.Time)
				if !okdate {
					panic("must enter a valid date")
				}
				exists, _ := petty_cash_req.Exists(user_id, amount, date)
				if exists {
					return nil, errors.New("duplicate petty cash request")
				}
				petty_cash_req.Create(user_id)
				return petty_cash_req, nil
			},
		},
	},
})
