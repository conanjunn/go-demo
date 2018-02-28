package main

import (
	"encoding/json"
	"log"
)

func main() {
	b := []byte(`{"Name": "Wednesday", "NNN": "xxx", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Printf("err: %v", err)
	} else {
		log.Printf("f: %v", f)
	}

	type FamilyMember struct {
		Name    string
		Age     int
		Parents []string
		NNN     string
	}
	var m FamilyMember
	err1 := json.Unmarshal(b, &m)
	if err != nil {
		log.Printf("err: %v", err1)
	} else {
		log.Printf("m: %v", m)
	}
}
