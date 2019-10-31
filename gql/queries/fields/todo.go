package queries

import (
	"golang-graphql-realworld/common/constant"
	"golang-graphql-realworld/container"
	types "golang-graphql-realworld/gql/types"
	"golang-graphql-realworld/model"

	"github.com/graphql-go/graphql"
)

var GetTodos = &graphql.Field{
	Type:        graphql.NewList(types.Todo),
	Description: "Get list todos",
	Resolve:     getListToDo,
}

func getListToDo(params graphql.ResolveParams) (interface{}, error) {
	notTodoCollection := container.Get().MongoClient.C(constant.MONGO_COLLECTION_TODO)
	var todosList []model.Todo
	err := notTodoCollection.Find(nil).All(&todosList)
	return todosList, err
}
