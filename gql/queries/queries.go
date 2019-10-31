package queries

import (
	fields "golang-graphql-realworld/gql/queries/fields"

	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"todos": fields.GetTodos,
	},
})
