package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"golang-graphql-realworld/common/constant"
	"golang-graphql-realworld/container"
	"golang-graphql-realworld/gql/mutations"
	"golang-graphql-realworld/gql/queries"
	"golang-graphql-realworld/server"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

type reqBody struct {
	Query string `json:"query"`
}

var (
	schema graphql.Schema
	config string
)

func init() {
	flag.StringVar(&config, "config", constant.DEFAULT_CONFIG, "path of file config")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	flag.Parse()
	err := container.Init(config)
	if err != nil {
		panic(err)
	}

	schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queries.RootQuery,
		Mutation: mutations.RootMutation,
	})
	if err != nil {
		panic(err)
	}

	router := server.New()
	router.POST("/graphql", gin.WrapH(gqlHandler()))

	err = router.Run(container.Get().Config.Listen)
	if err != nil {
		panic(err)
	}
}

func gqlHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "No query data", 400)
			return
		}

		var rBody reqBody
		err := json.NewDecoder(r.Body).Decode(&rBody)
		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		params := graphql.Params{Schema: schema, RequestString: rBody.Query}
		result := graphql.Do(params)
		if len(result.Errors) > 0 {
			fmt.Printf("failed to execute graphql operation, errors: %+v", result.Errors)
		}

		rJSON, err := json.Marshal(result)
		if err != nil {
			fmt.Printf("failed to execute graphql operation, errors: %v", err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, "%s", rJSON)

	})
}
