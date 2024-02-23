package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "aniket", "******")
	messageVal = strings.ReplaceAll(messageVal, "Pranit", "******")
	messageVal = strings.ReplaceAll(messageVal, "Rutuja", "******")
	*message = messageVal
}

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	// myString := "hello"
	// myStringPtr := &myString
	// fmt.Println(myStringPtr)
	// fmt.Println(*myStringPtr)
	// *myStringPtr = "world"
	// fmt.Println(myString)
	var messages [3]string
	messages[0] = "aniket"
	messages[1] = "Pranit"
	messages[2] = "Rutuja"
	test(messages[:])

}
