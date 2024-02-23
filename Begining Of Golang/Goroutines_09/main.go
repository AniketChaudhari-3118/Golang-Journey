package main

import "fmt"

func generator() <-chan int {
	isOldChan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			isOldChan <- i
		}
		close(isOldChan)
	}()
	return isOldChan
}

func receiver(isOldChan <-chan int) {
	for v := range isOldChan {
		fmt.Println(v)
	}
}

func main() {

	fmt.Println("hello")
	c := generator()
	receiver(c)

}
