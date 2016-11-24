package main

import (
  "time";
)

// Main structure for saving starting point of time
type Time struct {
  Time time.Time
}

// Format HH:MM:SS
func(t Time) timeFormat() string {
  return t.Time.Format("15:04:05")
}

// Format YYYY:MM:DD
func(t Time) dateFormat() string {
  return t.Time.Format("2006:01:02")
}

// First second of the date
func(t Time) getFirstSecond() int {
  return t.Time.Second()
}

// Structure for time queries
type TimeQuery struct {
  Time string `json:"time"`
  Timestamp int64 `json:"timestamp"`
}

// Structure for date queries
type DateQuery struct {
  Date string `json:"date"`
  Timestamp int64 `json:"timestamp"`
}

// getters/setters
var currentTimeAndDate = &Time{
	Time: time.Now(),
}

var timeQuery = TimeQuery{
  currentTimeAndDate.timeFormat(),
  int64(currentTimeAndDate.Time.Unix()),
}

var dateQuery = DateQuery{
  currentTimeAndDate.dateFormat(),
  int64(currentTimeAndDate.getFirstSecond()),
}


func GetTimeQuery() TimeQuery {
  return timeQuery
}

func GetDateQuery() DateQuery {
  return dateQuery
}
