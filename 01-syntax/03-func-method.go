package main

import "fmt"

type Country struct {
	Name string
}

type User struct {
    Age int
    Name string
    Money int
    Country Country
}

func (u User) greeting() string {
    return fmt.Sprintf("Hi, my name is %s and I'm %d years old. I am from %s.", u.Name, u.Age, u.Country.Name)
}

func getInfo(u User) string {
	return fmt.Sprintf("Hi, my name is %s. I am from %s.", u.Name, u.Country.Name)
}

func getAccess(u User) bool {
	if u.Age > 18 {
		return true
	} else {
		return false
	}
}

func main() {
	country := Country { "USA" }
    user := User{ 30, "Bob", 100, country }
    fmt.Println(user.greeting())
    fmt.Println(getInfo(user))
    fmt.Println(getAccess(user))
}