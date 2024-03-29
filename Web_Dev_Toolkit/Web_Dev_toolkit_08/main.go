package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //make sure all paths cancelthe context to avoid the context leak

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
	time.Sleep(1 * time.Minute)
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return // avoid leaking goroutine when ctx is done
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}
