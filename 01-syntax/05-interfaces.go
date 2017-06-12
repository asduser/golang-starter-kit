package main

import "fmt"

type Person interface {
    greeting(text string)
}

type Man struct {
    Person
}

func (p *Man) greeting() {
    fmt.Println("Hello, I'm a man!")
}

func main() {
    p := new(Man)
    p.greeting()
}