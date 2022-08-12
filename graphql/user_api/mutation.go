package user_api

import (
	u "financial-api/m/models/user"
	"regexp"

	"github.com/graphql-go/graphql"
)

var UserMutations = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutations",
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
				"role": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(RoleType),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email, isOk := p.Args["email"].(string)
				if !isOk {
					panic("must enter a valid email")
				}
				emailCheck, _ := regexp.MatchString("[a-z0-9!#$%&'*+/=?^_{|}~-]*@norainc.org", email)
				if !emailCheck {
					panic("must have a Northern Ohio Recovery Association Email to register")
				}
				user := &u.User{
					Name: p.Args["name"].(string),
				}
				role, roleOK := p.Args["role"].(string)
				if !roleOK {
					panic("user must have an active role")
				}
				result, err := user.Create(email, role)
				if err != nil {
					panic(err)
				}
				return result, nil
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				email, isOk := p.Args["email"].(string)
				if !isOk {
					panic("must enter a valid email")
				}
				emailCheck, _ := regexp.MatchString("[a-z0-9!#$%&'*+/=?^_{|}~-]*@norainc.org", email)
				if !emailCheck {
					panic("must have a Northern Ohio Recovery Association Email to register")
				}
				var user u.User
				result, err := user.Login(email)
				if err != nil {
					panic(err)
				}
				return result, nil
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
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_id, idOK := p.Args["user_id"].(string)
				if !idOK {
					panic("you must enter a valid user id")
				}
				name, nameOK := p.Args["name"].(string)
				if !nameOK {
					panic("you must enter a valid vehicle name")
				}
				description, descriptionOK := p.Args["description"].(string)
				if !descriptionOK {
					panic("you must enter a valid vehicle description")
				}
				var user u.User
				result, err := user.AddVehicle(user_id, name, description)
				if err != nil {
					panic(err)
				}
				return &u.Vehicle{
					ID:          result,
					Name:        name,
					Description: description,
				}, nil
			},
		},
		"remove_vehicle": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Allow a user to remove a vehicle from their account",
			Args: graphql.FieldConfigArgument{
				"user_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
				"vehicle_id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.ID),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user_id, idOK := p.Args["user_id"].(string)
				if !idOK {
					panic("you must enter a valid user id")
				}
				vehicle_id, vehicle_idOK := p.Args["vehicle_id"].(string)
				if !vehicle_idOK {
					panic("you must enter a valid vehicle id")
				}
				var user u.User
				result, err := user.RemoveVehicle(user_id, vehicle_id)
				if err != nil {
					panic(err)
				}
				return result, nil
			},
		},
	},
})
