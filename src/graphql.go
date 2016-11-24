package main

import (
  "github.com/graphql-go/graphql"
)

// Create graphQl Schema
func GraphSchema() graphql.Schema {
  var Schema graphql.Schema

  var queryObject = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
      "timeQuery": &graphql.Field{
        Type: graphql.String,
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetTimeQuery(), nil
        },
      },
      "dateQuery": &graphql.Field{
          Type: graphql.String,
          Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            return GetDateQuery(), nil
          },
      },
    },
  })

  Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryObject,
  })

  return Schema
}
