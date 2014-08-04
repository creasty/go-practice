package main

import (
  "fmt"
  "math/rand"
  "time"
)

type Result string
type Engine func(query string) Result

func server(category string) Engine {
  return func(query string) Result {
    time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
    return Result(fmt.Sprintf("<%s result of %q>", category, query))
  }
}

func firstOf(query string, engines ...Engine) Result {
  c := make(chan Result)

  for _, engine := range engines {
    go func() { c <- engine(query) }()
  }

  return <-c
}

var (
  webSearchServer1   = server("web")
  webSearchServer2   = server("web")
  imageSearchServer1 = server("image")
  imageSearchServer2 = server("image")
  videoSearchServer1 = server("video")
  videoSearchServer2 = server("video")
)

func search(query string) (results []Result) {
  c := make(chan Result)

  acrossAllEngines := func(engines ...Engine) {
    go func() { c <- firstOf(query, engines...) }()
  }

  acrossAllEngines(webSearchServer1, webSearchServer2)
  acrossAllEngines(imageSearchServer1, imageSearchServer2)
  acrossAllEngines(videoSearchServer1, videoSearchServer2)

  timeout := time.After(600 * time.Millisecond)

  for i := 0; i < 3; i++ {
    select {
    case result := <-c:
      results = append(results, result)
    case <-timeout:
      fmt.Println("timeout!")
      return
    }
  }

  return
}

func main() {
  rand.Seed(time.Now().UnixNano())

  start := time.Now()
  results := search("golang")
  elapsed := time.Since(start)

  fmt.Println(results)
  fmt.Println(elapsed)
}
