package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{}
	fmt.Println(m)

	bs, err := json.Marshal(m)
	if err != nil {
		log.Println("error:", err)
	}

	fmt.Println(string(bs))

}
