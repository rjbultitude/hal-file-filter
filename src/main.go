package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Request struct {
	Method string
	URL    string
}

type Entry struct {
	StartedDateTime string
	Time            float64
	Request         Request
}

type Entries struct {
	Entry Entry
}

type Creator struct {
	Name    string
	Version string
}

type Log struct {
	Version string
	Creator Creator
	Entries Entries
}

func main() {
	data, err := ioutil.ReadFile("data/homepage.json")
	if err != nil {
		fmt.Print(err)
	}

	var result []Log
	json.Unmarshal(data, &result)
	fmt.Println(" Result: ", result)

	var resources []Entries

	err = json.Unmarshal(data, &resources)
	fmt.Println(" Entries: ", resources)
}
