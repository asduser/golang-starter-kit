package main

import "fmt"
import "sync"

func main() {
  waitGroup := sync.WaitGroup{}
  waitGroup.Add(1)
  go func() {
    fmt.Println("Hello, playground")
    waitGroup.Done()
  }()
  waitGroup.Wait()
}