package main

import (
  "fmt"
  "math/rand"
  "time"
)

type User struct {
  name string
}

type Message struct {
  user *User
  body string
}

func (m Message) log() {
  fmt.Printf("%s: %q\n", m.user.name, m.body)
}


func anyOf(input1, input2 <-chan Message) <-chan Message {
  c := make(chan Message)

  go func() {
    for {
      select {
      case in := <-input1: c <- in
      case in := <-input2: c <- in
      case <-time.After(1 * time.Second):
        fmt.Println(" ↑this guy's taking too slow!")
      }
    }
  }()

  return c
}

func chatter(name string) <-chan Message {
  c := make(chan Message)
  i := 0

  go func() {
    for {
      c <- Message{
        user: &User{name: name},
        body: fmt.Sprintf("%d", i),
      }

      i++
      time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
    }
  }()

  return c
}

func main() {
  c := anyOf(chatter("Alice"), chatter("Bob"))

  for i := 0; i < 10; i++ {
    (<-c).log()
  }

  fmt.Println("-- All conversation done, the room is closing.")
}
