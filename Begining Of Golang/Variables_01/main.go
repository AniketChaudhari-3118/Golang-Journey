package main

import "fmt"

func main() {
	var smsSendingLimit int
	var costPerSMS float64
	var hasPersmission bool
	var username string

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPersmission, username)

	congrats := "happy birthday!"
	fmt.Println(congrats)

	// penniesPerText := 2 //int value
	penniesPerText := 2.0
	fmt.Printf("The type of penniesPerText is %T\n", penniesPerText)

	averageOpenRate, displayMessage := .23, "is the average open rate of your message"
	fmt.Println(averageOpenRate, displayMessage)

	accountAge := 2.6

	//create a new "accountAgeInt" here
	accountAgeInt := int(accountAge)
	fmt.Println("your ccount has existed for", accountAgeInt, "years")

	// //constants:
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	fmt.Println("number of seconds in one hour:", secondsInHour)

	//printf and sprintf
	const name = "aniket chaudhari"
	const openrate = 30.5

	msg := fmt.Sprintf("Hi %s, your open rate is.%f percent", name, openrate)
	fmt.Println(msg)

	//conditional statements
	messagelen := 10
	maxMessageLen := 20
	if messagelen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not Sent")
	}

}
