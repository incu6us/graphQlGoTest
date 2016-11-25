package main

import (
  "github.com/graphql-go/graphql"
)

// Create graphQl Schema
func GraphSchema() graphql.Schema {
  var Schema graphql.Schema

  var timeType = graphql.NewObject(graphql.ObjectConfig{
    Name: "time",
    Fields: graphql.Fields{
      "time": &graphql.Field{
        Type: graphql.String,
      },
      "timestamp": &graphql.Field{
        Type: graphql.String,
      },
    },
  })

  var dateType = graphql.NewObject(graphql.ObjectConfig{
    Name: "date",
    Fields: graphql.Fields{
      "date": &graphql.Field{
        Type: graphql.String,
      },
      "timestamp": &graphql.Field{
        Type: graphql.String,
      },
    },
  })

  var queryObject = graphql.NewObject(graphql.ObjectConfig{
    Name: "Query",
    Fields: graphql.Fields{
      "timeQuery": &graphql.Field{
        Type: timeType,
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
          return GetTimeQuery(), nil
        },
      },
      "dateQuery": &graphql.Field{
          Type: dateType,
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
