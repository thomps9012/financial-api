package mileage_api

import (
	"errors"
	r "financial-api/models/requests"
	"time"

	"github.com/graphql-go/graphql"
)

var MileageMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "MileageMutations",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        MileageType,
			Description: "Creates a new mileage request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"request": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(MileageInputType),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userID, isOK := p.Args["user_id"].(string)
				if !isOK {
					panic("must enter a valid user id")
				}
				mileageArgs := p.Args["request"].(map[string]interface{})
				date, dateisOK := mileageArgs["date"].(time.Time)
				if !dateisOK {
					panic("must enter a valid date")
				}
				start, startisOK := mileageArgs["starting_location"].(string)
				if !startisOK {
					panic("must enter a valid starting location")
				}
				destination, destinationisOK := mileageArgs["destination"].(string)
				if !destinationisOK {
					panic("must enter a valid destination")
				}
				purpose, purposeisOK := mileageArgs["trip_purpose"].(string)
				if !purposeisOK {
					panic("must enter a valid trip purpose")
				}
				start_odo, start_odoisOK := mileageArgs["start_odometer"].(int)
				if !start_odoisOK {
					panic("must enter a valid starting odometer")
				}
				end_odo, end_odoisOK := mileageArgs["end_odometer"].(int)
				if !end_odoisOK {
					panic("must enter a valid end odometer")
				}
				tolls, tollsisOK := mileageArgs["tolls"].(float64)
				if !tollsisOK {
					panic("must enter a valid tolls amount")
				}
				parking, parkingisOK := mileageArgs["parking"].(float64)
				if !parkingisOK {
					panic("must enter a valid parking amount")
				}
				mileage_req := &r.Mileage_Request{
					Date:              date,
					Starting_Location: start,
					Destination:       destination,
					Trip_Purpose:      purpose,
					Start_Odometer:    start_odo,
					End_Odometer:      end_odo,
					Tolls:             tolls,
					Parking:           parking,
				}
				exists, _ := mileage_req.Exists(userID, date, start_odo, end_odo)
				if exists {
					return nil, errors.New("mileage request already created")
				}
				mileage_req.Create(userID)
				return mileage_req, nil
			},
		},
	},
})
