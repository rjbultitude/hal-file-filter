package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
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

type Creator struct {
	Name    string
	Version string
}

type Log struct {
	Version string
	Creator Creator
	Entries []Entry
}

type Start struct {
	Log Log
}

func main() {

	data, err := ioutil.ReadFile("data/test.json")
	if err != nil {
		fmt.Print(err)
	}

	var resources Start

	err = json.Unmarshal(data, &resources)
	fmt.Println(" results: ", resources)
	fmt.Println(" results: ", reflect.TypeOf(resources))
}
