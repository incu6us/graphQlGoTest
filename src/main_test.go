package main

import (
  "fmt";
  "testing";
  "strings";
  "encoding/json";
  "net/http";
  "io/ioutil";
  "strconv";
)

// Testing setter & getter time is the same
func TestTime(t *testing.T){

  if len(strings.Split(currentTimeAndDate.timeFormat(), ":")) != 3 {
    t.Fail()
  }

  if currentTimeAndDate.getFirstSecond() > 60 {
    t.Fail()
  }

  if timeQuery != GetTimeQuery() {
    t.Fail()
  }
}

// Testing setter & getter date is the same
func TestDate(t *testing.T){

  if len(strings.Split(currentTimeAndDate.dateFormat(), ":")) != 3 {
    t.Fail()
  }

  if dateQuery != GetDateQuery() {
    t.Fail()
  }
}

// Testing HTTP with "timeQuery" or "dateQuery"
func TestHttp(t *testing.T)  {
  go Handler()

  // HttpCheck("timeQuery")
  // HttpCheck("dateQuery")
  time, _, err0 := HttpCheck("timeQuery")
  if err0 != nil {
    t.Fail()
  }

  // Check Time
  if time.Time != GetTimeQuery().Time {
    t.Fail()
  }

  // Check Timestamp
  if time.Timestamp != strconv.Itoa(int(GetTimeQuery().Timestamp)) {
    t.Fail()
  }

  _, date, err1 := HttpCheck("dateQuery")
  if err1 != nil {
    t.Fail()
  }

  // Chaeck Date
  if date.Date != GetDateQuery().Date {
    t.Fail()
  }

  // Chaeck Date Timestamp
  if date.Timestamp != strconv.Itoa(int(GetDateQuery().Timestamp)) {
    t.Fail()
  }
}

// Query to server with "timeQuery" or "dateQuery"
func HttpCheck(typeQuery string) (time TimeQueryTest, date DateQueryTest, err error){
  var url string
  client := &http.Client{}

  if typeQuery == "timeQuery" {
    url = "http://localhost:8081/graphql?query={"+typeQuery+"{time,timestamp}}"
  } else if typeQuery == "dateQuery"{
    url = "http://localhost:8081/graphql?query={"+typeQuery+"{date,timestamp}}"
  }

  req, e := http.NewRequest("GET", url, nil)
  req.Header.Set("Content-Type", "application/json")
  if e != nil {
    err = e
  }

  resp, _ := client.Do(req)

  defer resp.Body.Close()
  body, e := ioutil.ReadAll(resp.Body)
  if e != nil {
    err = e
  }

  if typeQuery == "timeQuery" {
    var result TimeDataStruct
    err := json.Unmarshal(body, &result)
    if err != nil {
      fmt.Println(err)
      fmt.Println(string(body))
    }
    time = result.Data.TQuery
  } else if typeQuery == "dateQuery" {
    var result DateDataStruct
    err := json.Unmarshal(body, &result)
    if err != nil {
      fmt.Println(err)
      fmt.Println(string(body))
    }
    date = result.Data.DQuery
  }
  return
}

// Time query
type TimeDataStruct struct {
  Data TimeQueryStruct `json:"data"`
}

type TimeQueryStruct struct {
  TQuery TimeQueryTest `json:"timeQuery"`
}

type TimeQueryTest struct {
  Time string `json:"time"`
  Timestamp string `json:"timestamp"`
}

// Date query
type DateDataStruct struct {
  Data DateQueryStruct `json:"data"`
}

type DateQueryStruct struct {
  DQuery DateQueryTest `json:"dateQuery"`
}

type DateQueryTest struct {
  Date string `json:"date"`
  Timestamp string `json:"timestamp"`
}
