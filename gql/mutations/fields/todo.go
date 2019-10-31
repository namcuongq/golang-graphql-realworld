package fields

import (
	"fmt"
	"golang-graphql-realworld/common/constant"
	"golang-graphql-realworld/container"
	"golang-graphql-realworld/model"

	"github.com/graphql-go/graphql"

	types "golang-graphql-realworld/gql/types"
)

var CreateTodo = &graphql.Field{
	Type:        types.Todo,
	Description: "Create a Todo",

	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},

	Resolve: createToDo,
}

func createToDo(params graphql.ResolveParams) (interface{}, error) {
	name := params.Args["name"]
	if name == nil {
		return nil, fmt.Errorf("name not null")
	}

	description := params.Args["description"]
	if description == nil {
		return nil, fmt.Errorf("description not null")
	}

	var todo = model.Todo{fmt.Sprintf("%v", name), fmt.Sprintf("%v", description)}

	collection := container.Get().MongoClient.C(constant.MONGO_COLLECTION_TODO)
	err := collection.Insert(todo)

	return todo, err

}
