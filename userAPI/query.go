package userAPI

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
			// Type:        UserType,
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
			},
		},
	},
})
