package main

import "fmt"


func cal() int {
	var sum int
	for	i := 0; i < 10; i++ {
		sum = sum + i
	}
	return sum
}


func main() {
    val := cal()
	fmt.Println(val)
}