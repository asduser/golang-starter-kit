package main
import (
    "fmt"
)

func main() {
    outputParagraph()
    simplePointer()
    outputParagraph()
    dynamicValueChanging()
    outputParagraph()
}

func simplePointer() {
    value := 100
    linked := &value
    fmt.Println("address of 'value' is", linked)
    fmt.Println("value of 'value' is", *linked)
}

func dynamicValueChanging() {
    a := 100
    b := &a
    fmt.Println("new 'b' equals ", *b * 20)
}

func outputParagraph() {
    fmt.Println("------")
}