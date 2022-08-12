package check_api

import (
	"errors"
	r "financial-api/m/models/requests"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
)

var CheckRequestMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "CheckMutations",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        CheckRequestType,
			Description: "Creates a new check request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"vendor": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(VendorInput),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(CheckRequestInput),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				vendor_input := p.Args["vendor"].(map[string]interface{})
				vendor_address_input := vendor_input["address"].(map[string]interface{})
				vendor_address := &r.Address{
					Website:  vendor_address_input["website"].(string),
					Street:   vendor_address_input["street"].(string),
					City:     vendor_address_input["city"].(string),
					State:    vendor_address_input["state"].(string),
					Zip_Code: vendor_address_input["zip"].(int),
				}
				vendor := &r.Vendor{
					Name:    vendor_input["name"].(string),
					Address: *vendor_address,
				}
				checkReqArgs := p.Args["request"].(map[string]interface{})
				purchases_input := checkReqArgs["purchases"].([]interface{})
				var purchases []r.Purchase
				var order_total = 0.0
				for _, purchase_item_obj := range purchases_input {
					fmt.Printf("%s\n", purchase_item_obj)
					purchase_item := purchase_item_obj.(map[string]interface{})
					amount := purchase_item["amount"].(float64)
					description := purchase_item["description"].(string)
					grant_line_item := purchase_item["grant_line_item"].(string)
					purchase := &r.Purchase{
						Grant_Line_Item: grant_line_item,
						Description:     description,
						Amount:          amount,
					}
					order_total += amount
					purchases = append(purchases, *purchase)
				}
				receiptArgs := checkReqArgs["receipts"].([]interface{})
				var receipts []string
				for item := range receiptArgs {
					receipts = append(receipts, receiptArgs[item].(string))
				}
				check_request := &r.Check_Request{
					Date:        checkReqArgs["date"].(time.Time),
					Vendor:      *vendor,
					Description: checkReqArgs["description"].(string),
					Grant_ID:    checkReqArgs["grant_id"].(string),
					Purchases:   purchases,
					Order_Total: order_total,
					Receipts:    receipts,
					Credit_Card: checkReqArgs["credit_card"].(string),
				}
				user_id, isOk := p.Args["user_id"].(string)
				if !isOk {
					panic("must enter a user id")
				}
				exists, _ := check_request.Exists(user_id, vendor.Name, order_total, check_request.Date)
				if exists {
					return nil, errors.New("check request already created")
				}
				check_request.Create(user_id)
				return check_request, nil
			},
		},
	},
})
