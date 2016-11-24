package main

import (
  "fmt";
)

func main() {
  // short debug(the values must be the same in the API)
  print := fmt.Println
  print(GetTimeQuery())
  print(GetDateQuery())

  // Run HTTP with graphql's schema
  Handler()
}
