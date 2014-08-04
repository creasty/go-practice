package main

import (
  "fmt"
  "math/rand"
  "time"
)

func routine(id int, semaphore chan bool, done chan<- int) {
  go func() {
    semaphore <- true

    time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

    fmt.Printf("Done #%d\n", id)

    <-semaphore
    done <- id
  }()
}

func main() {
  const (
    concurrency = 5
    routines    = 10
  )

  semaphore := make(chan bool, concurrency)
  done := make(chan int, routines)

  for i := 0; i < routines; i++ {
    routine(i, semaphore, done)
  }

  for i := 0; i < routines; i++ {
    <-done
  }

  fmt.Println("All routines are done")
}
