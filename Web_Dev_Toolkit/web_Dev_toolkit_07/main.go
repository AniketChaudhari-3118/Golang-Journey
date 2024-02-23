package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results := dbAccess(ctx)

	fmt.Fprintln(res, results)
}

func dbAccess(ctx context.Context) int {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	ch := make(chan int)

	go func() {
		//ridiculous long running task
		uid := ctx.Value("userID").(int)
		time.Sleep(4 * time.Second)

		//check to make sure we're not running in vain
		//if ctx.Done() has
		if ctx.Err() != nil {
			return
		}

		ch <- uid
	}()

	select {
	case <-ctx.Done():
		return 0
	case i := <-ch:
		return i
	}
}

func bar(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	log.Println(ctx)
	fmt.Fprintln(res, ctx)
}
