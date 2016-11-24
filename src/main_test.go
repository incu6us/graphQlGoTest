package main

import (
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

  data0, err0 := HttpCheck("timeQuery")
  if err0 != nil {
    t.Fail()
  }

  // Check Time
  if strings.Split(data0, " ")[0] != GetTimeQuery().Time {
    t.Fail()
  }

  // Check Timestamp
  if strings.Split(data0, " ")[1] != strconv.Itoa(int(GetTimeQuery().Timestamp)) {
    t.Fail()
  }

  data1, err1 := HttpCheck("dateQuery")
  if err1 != nil {
    t.Fail()
  }

  // Chaeck Date
  if strings.Split(data1, " ")[0] != GetDateQuery().Date {
    t.Fail()
  }

  // Chaeck Date Timestamp
  if strings.Split(data1, " ")[1] != strconv.Itoa(int(GetDateQuery().Timestamp)) {
    t.Fail()
  }
}

// Query to server with "timeQuery" or "dateQuery"
func HttpCheck(typeQuery string) (data string, err error){
  client := &http.Client{}

  url := "http://localhost:8081/graphql?query={"+typeQuery+"}"
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
    json.Unmarshal(body, &result)
    convF := strings.Replace(result.Data.TimeQuery, "{", "", -1)
    convL := strings.Replace(convF, "}", "", -1)
    data = convL
  } else if typeQuery == "dateQuery" {
    var result DateDataStruct
    json.Unmarshal(body, &result)
    convF := strings.Replace(result.Data.DateQuery, "{", "", -1)
    convL := strings.Replace(convF, "}", "", -1)
    data = convL
  }
  return
}

// Time query
type TimeDataStruct struct {
  Data TimeQueryStruct `json:"data"`
}

type TimeQueryStruct struct {
  TimeQuery string `json:"timeQuery"`
}

// Date query
type DateDataStruct struct {
  Data DateQueryStruct `json:"data"`
}

type DateQueryStruct struct {
  DateQuery string `json:"dateQuery"`
}
