package main

import (
	//"bufio"
	"bufio"
	"fmt"
	// "io"
	"log"
	"net"
	"time"
	//"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {	
		conn, err := li.Accept()
		if err != nil {
			log.Println()
		}
	    go handle(conn)
		// io.WriteString(conn, "\nHello from TCP server\n")
		// fmt.Fprintln(conn, "How is your day?")
		// fmt.Fprintf(conn, "%v", "Well, I hope!")

		// conn.Close()
	}
}
	
func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		fmt.Println("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Code got here")

}
