package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// You can choose to only unmarshal some of the json data
// Create a data structure that only has fields for some of the data
type city struct {
	Latitude, Longitude float64
	city                string
}

type cities []city

func main() {
	var data cities
	rcvd := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {

		log.Fatalln(err)
	}
	fmt.Println(data)
}

//Here we recieved the data in the variable 'rcvd' as the json data and after that,
//to unmarshal it we passed the slice of bytes to convert the whole data in bytes,
//and then with that we passed the address of the data structure that,  where we are going to 
//store the unmarshaled data. 