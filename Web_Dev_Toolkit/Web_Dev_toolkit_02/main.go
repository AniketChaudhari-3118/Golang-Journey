package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "I am a fresher with a degree of B.E in the field of Computer Science and seeking entry-level Software Engineer job at your esteemed company. Enthusiastic about cybernetics and eager to contribute my key skills. Available for immediate interview"
	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("I'm giving her all she's got Caption!", err)
	}
	fmt.Println(string(bs))
}
