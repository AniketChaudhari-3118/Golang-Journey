package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime time.Time
	recipienName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %v :", bm.recipienName, bm.birthdayTime)
}

type sendingReprt struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReprt) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v number of times`, sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
}

func main() {
	//as test function takes an interface so we can pass any struct that implemnts the intrface
	test(sendingReprt{
		reportName:    "First Report",
		numberOfSends: 10,
	})

	test(birthdayMessage{
		recipienName: "ANiket",
		birthdayTime: time.Date(2001, 06, 19, 0, 0, 0, 0, time.UTC),
	})
}
