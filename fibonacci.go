package main

import "fmt"

func fibonacciFactory() <-chan int {
  c := make(chan int)
  i, j := 0, 1

  go func() {
    for {
      c <- j
      i, j = j, i+j
    }
  }()

  return c
}

func main() {
  fibonacci := fibonacciFactory()

  for i := 0; i < 20; i++ {
    fmt.Println(<-fibonacci)
  }
}
