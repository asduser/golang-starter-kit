package main

import (
	"fmt"
	"errors"
)

func main() {
	err := errors.New("custom exception")
	fmt.Println(err)
}