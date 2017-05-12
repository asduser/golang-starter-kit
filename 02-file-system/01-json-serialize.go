package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	const jsonSerialized = `
		{"Name": "Ed", "Age": "20"}
		{"Name": "Sam", "Age": "35"}
	`
	type User struct {
		Name string
		Age string
	}
	dec := json.NewDecoder(strings.NewReader(jsonSerialized))
	for {
		var user User
		if err := dec.Decode(&user); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", user.Name, user.Age)
	}
}