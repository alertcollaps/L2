package main

import (
	"testing"
	time "time"
)

func TestGetTime(t *testing.T) {
	tm := time.Now()
	tmTest, err := getTime()
	if err != nil {
		t.Error(err)
	}
	latency := 1
	secundesTm := tm.Second() + 60*tm.Minute() + 3600*tm.Hour()
	secundesTest := tmTest.Second() + 60*tmTest.Minute() + 3600*tmTest.Hour()
	if secundesTm-latency >= secundesTest || secundesTest >= secundesTm+latency {
		t.Error("Latency more then 1 second")
	}
}
