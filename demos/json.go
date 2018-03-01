package main

import (
	"encoding/json"
	"log"
)

func main() {
	b := []byte(`{"name": "Wednesday", "NNN": "xxx", "age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		log.Printf("err: %v", err)
	} else {
		log.Printf("f: %v", f)
	}

	type FamilyMember struct {
		Name    string // 首字母必须大写
		Aa      int    `json:"age"` // 使用tag
		Parents []string
		NNN     string
	}
	var m FamilyMember
	err1 := json.Unmarshal(b, &m)
	if err1 != nil {
		log.Printf("err: %v", err1)
	} else {
		log.Printf("m: %v", m)
	}
}
