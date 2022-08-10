package user_api

import (
	"github.com/graphql-go/graphql"
)

// move email check to resolver
// emailCheck, _ := regexp.MatchString("[a-z0-9!#$%&'*+/=?^_{|}~-]*@norainc.org", value)
// if !emailCheck {
// 	return nil
// }
// return value

var UserMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "User Mutations",
	Fields: graphql.Fields{
		"create": &graphql.Field{
			Type:        UserType,
			Description: "Create a new user on initial sign up",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"type": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.EnumValueType),
				},
				"role": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(RoleType),
				},
			},
		},
		"login": &graphql.Field{
			Type:        UserType,
			Description: "Login a user and gather basic information about them",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		},
		"add_vehicle": &graphql.Field{
			Type:        VehicleType,
			Description: "Allow a user to add a vehicle to their account",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
		},
		"remove_vehicle": &graphql.Field{
			Type:        VehicleType,
			Description: "Allow a user to remove a vehicle from their account",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"vehicle_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
		},
	},
})
