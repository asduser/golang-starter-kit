package main

import (
  "fmt"
)

type User struct {
    name string
    age int
}

func main() {

  users := []User { 
    User{ "Bob", 20 },
    User{ "Thomas", 30 },
  }

  fmt.Println(users[0]);
  fmt.Println(users[1].name);

}