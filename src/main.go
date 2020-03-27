package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
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
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Entries []Entry `json:"entries"`
}

type Start struct {
	Log Log
}

func MatchString(s string) (matched bool) {
	matched, err := regexp.MatchString("^http://", s)
	if err != nil {
		fmt.Print(err)
	}
	return matched
}

func main() {

	data, err := ioutil.ReadFile("data/test.json")
	if err != nil {
		fmt.Print(err)
	}

	var resources Start

	err = json.Unmarshal(data, &resources)
	// fmt.Println(" URL: ", resources.Log.Entries[0].Request.URL)
	// fmt.Println(" results: ", reflect.TypeOf(resources))
	for i := range resources.Log.Entries {
		thisURL := resources.Log.Entries[i].Request.URL
		thisURLNoSSL := MatchString(thisURL)
		fmt.Println("thisURL ", thisURL)
		fmt.Println("thisURLNoSSL ", thisURLNoSSL)
		if thisURLNoSSL == true {
			fmt.Println(" ", thisURL)
		}
	}
}
