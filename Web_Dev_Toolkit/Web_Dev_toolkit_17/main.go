package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//string
	var data1 string
	rcvd1 := `"Todd"`
	err1 := json.Unmarshal([]byte(rcvd1), &data1)
	if err1 != nil {

		log.Fatalln(err1)
	}
	fmt.Println(data1)

	//int
	var data2 int
	rcvd2 := `42`
	err2 := json.Unmarshal([]byte(rcvd2), &data2)
	if err2 != nil {
		log.Fatalln(err2)
	}
	fmt.Println(data2)

	//bool
	var data3 bool
	rcvd3 := `true`
	err3 := json.Unmarshal([]byte(rcvd3), &data3)
	if err3 != nil {
		log.Fatalln(err3)
	}

	fmt.Println(data3)

	//null
	var a []string

	rcvd := `null`
	err := json.Unmarshal([]byte(rcvd), &a)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

}
