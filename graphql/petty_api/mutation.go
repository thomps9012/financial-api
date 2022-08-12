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
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(PettyCashInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_id, isOk := p.Args["user_id"].(string)
				if !isOk {
					panic("must enter a valid user id")
				}
				println(user_id)
				grant_id, grantOK := p.Args["grant_id"].(string)
				if !grantOK {
					panic("must enter a valid grant")
				}
				requestArgs := p.Args["request"].(map[string]interface{})
				amount, okAmount := requestArgs["amount"].(float64)
				if !okAmount {
					panic("must enter a valid amount")
				}
				date, okdate := requestArgs["date"].(time.Time)
				if !okdate {
					panic("must enter a valid date")
				}
				receipts, receiptsOK := requestArgs["receipts"].([]string)
				if !receiptsOK {
					panic("must enter a valid receipt item")
				}
				description, descriptionOK := requestArgs["description"].(string)
				if !descriptionOK {
					panic("must enter a valid description")
				}
				petty_cash_req := &r.Petty_Cash_Request{
					Date:        date,
					Grant_ID:    grant_id,
					Amount:      amount,
					Description: description,
					Receipts:    receipts,
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
