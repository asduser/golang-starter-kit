package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type User struct {
		Name string
		Age string
	}

func main() {
	// data
	const jsonSerialized = `
		{"Name": "Ed", "Age": "20"}
		{"Name": "Sam", "Age": "35"}
	`
    
    // json -> object
	dec := json.NewDecoder(strings.NewReader(jsonSerialized))

	// functionality
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