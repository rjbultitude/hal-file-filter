package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	// read file
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Print(err)
	}

	type Entry struct {
		URL string
	}
	var entries []Entry

	// unmarshall it
	err = json.Unmarshal(data, &entries)
	fmt.Println("Urls ", entries)
}
