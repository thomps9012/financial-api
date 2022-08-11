package mileage_api

import (
	. "financial-api/m/models/requests"
	"time"

	"github.com/graphql-go/graphql"
)

var MileageMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutations",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        MileageType,
			Description: "Creates a new mileage request for a given user",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"date": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.DateTime),
				},
				"starting_location": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"destination": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"trip_purpose": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"start_odometer": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"end_odometer": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"tolls": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"parking": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				userID, isOK := p.Args["user_id"].(string)
				if !isOK {
					panic("must enter a valid user id")
				}
				date, dateisOK := p.Args["date"].(time.Time)
				if !dateisOK {
					panic("must enter a valid date")
				}
				start, startisOK := p.Args["starting_location"].(string)
				if !startisOK {
					panic("must enter a valid starting location")
				}
				destination, destinationisOK := p.Args["destination"].(string)
				if !destinationisOK {
					panic("must enter a valid destination")
				}
				purpose, purposeisOK := p.Args["trip_purpose"].(string)
				if !purposeisOK {
					panic("must enter a valid trip purpose")
				}
				start_odo, start_odoisOK := p.Args["start_odometer"].(int64)
				if !start_odoisOK {
					panic("must enter a valid starting odometer")
				}
				end_odo, end_odoisOK := p.Args["end_odometer"].(int64)
				if !end_odoisOK {
					panic("must enter a valid end odometer")
				}
				tolls, tollsisOK := p.Args["tolls"].(float64)
				if !tollsisOK {
					panic("must enter a valid tolls amount")
				}
				parking, parkingisOK := p.Args["parking"].(float64)
				if !parkingisOK {
					panic("must enter a valid parking amount")
				}
				mileage_req := &Mileage_Request{
					Date:              date,
					Starting_Location: start,
					Destination:       destination,
					Trip_Purpose:      purpose,
					Start_Odometer:    start_odo,
					End_Odometer:      end_odo,
					Tolls:             tolls,
					Parking:           parking,
				}
				mileage_req.Create(userID)
				return mileage_req, nil
			},
		},
	},
})
