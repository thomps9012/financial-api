package check_api

import (
	. "financial-api/m/models/requests"
	"time"

	"github.com/graphql-go/graphql"
)

var CheckRequestMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Check Request Mutations",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        CheckRequestType,
			Description: "Creates a new check request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"vendor": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(VendorInputType),
				},
				"date": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"purchases": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(&graphql.List{OfType: PurchaseInputType}),
				},
				"receipts": &graphql.ArgumentConfig{
					Type: &graphql.List{OfType: graphql.String},
				},
				"credit_card": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				vendor_input := p.Args["vendor"].(map[string]interface{})
				vendor_address_input := vendor_input["address"].(map[string]interface{})
				vendor_address := &Address{
					Website:  vendor_address_input["website"].(string),
					Street:   vendor_address_input["street"].(string),
					City:     vendor_address_input["city"].(string),
					State:    vendor_address_input["state"].(string),
					Zip_Code: vendor_address_input["zip"].(int64),
				}
				vendor := &Vendor{
					Name:    vendor_input["name"].(string),
					Address: *vendor_address,
				}
				purchases_input := p.Args["purchases"].(map[string]interface{})
				var purchases []Purchase
				for range purchases_input {
					purchase := &Purchase{
						Grant_ID:        purchases_input["grant_id"].(string),
						Grant_Line_Item: purchases_input["line_item"].(string),
						Description:     purchases_input["description"].(string),
						Amount:          purchases_input["amount"].(float64),
					}
					purchases = append(purchases, *purchase)
				}
				check_request := &Check_Request{
					Date:        p.Args["date"].(time.Time),
					Vendor:      *vendor,
					Description: p.Args["description"].(string),
					Purchases:   purchases,
					Receipts:    p.Args["receipts"].([]string),
					Credit_Card: p.Args["credit_card"].(string),
				}
				user_id, isOk := p.Args["user_id"].(string)
				if !isOk {
					panic("must enter a user id")
				}
				check_request.Create(user_id)
				return check_request, nil
			},
		},
	},
})
