package main

import "fmt"

// func concat(s1 string, s2 string) string {
// 	return s1 + s2
// }

func main() {
	// fmt.Println(concat("Aniket,", " happy birthday"))
	// fmt.Println(concat("yor are", " a great man"))

	// sendsSoFar := 430
	// const sendsToAdd = 25
	// sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	// fmt.Println("You have sent", sendsSoFar, "message")

	//IGNORING RETURN VALUE
	//  firstName, lastName := getNames()
	firstName, _ := getNames()
	fmt.Println("Welcome to Pune,", firstName)
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

func getNames() (string, string) {
	return "Aniket", "Chaudhari"
}
