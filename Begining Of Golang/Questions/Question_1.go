package main

//average of three numbers

import "fmt"

func main() {
	var a uint
	var b uint
	var c uint

	fmt.Println("Enter the value of a:")
	fmt.Scan(&a)

	fmt.Println("Enter the value of b:")
	fmt.Scan(&b)

	fmt.Println("Enter the value of c:")
	fmt.Scan(&c)

	var add uint = a + b + c
	var Average uint = add / 3

	fmt.Printf("The average of the three numbers is: %v", Average)
}
