package main

import "fmt"


type Country struct {
	Name string
}

type User struct {
    Age int
    Name string
    Country Country
}

func (u User) greeting() string {
    return fmt.Sprintf("Hi, my name is %s and I'm %d years old. I am from %s.", u.Name, u.Age, u.Country.Name)
}

func main() {
	country := Country { "USA" }
    user := User{ 30, "Bob", country }
    fmt.Println(user.greeting());
}