package mutations

import (
	fields "golang-graphql-realworld/gql/mutations/fields"

	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createNotTodo": fields.CreateTodo,
	},
})
