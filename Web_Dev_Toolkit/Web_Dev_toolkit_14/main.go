package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

type thumbnail struct {
	URl            string
	Height, Weight int
}

func main() {
	var data img
	rcvd := `{"Width":800,"Height":600,"Title":"View from 15th Floor","Thumbnail":{"Url":"http://www.example.com/image/481989943","Height":125,"Width":100},"Animated":false,"IDs":[116,943,234,38793]}`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln("err unmarshaling", err)
	}

	fmt.Println(data)

	for i, v := range data.IDs {
		fmt.Println(i, v)
	}

	fmt.Println(data.Thumbnail.URl)
}
