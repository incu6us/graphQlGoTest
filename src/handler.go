package main

import (
  "github.com/graphql-go/graphql"
  "github.com/graphql-go/handler"
  "net/http";
)

// Run HTTP multiplexor with graphql's schema
func Handler()  {

  var Schema graphql.Schema = GraphSchema()

  handler := handler.New(&handler.Config{
    Schema: &Schema,
    Pretty: true,
  })

  // serve HTTP
  serverMux := http.NewServeMux()
  serverMux.Handle("/graphql", handler)
  http.ListenAndServe(":8081", serverMux)
}
