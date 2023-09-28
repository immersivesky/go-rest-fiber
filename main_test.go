package main_test

import (
  "fmt"
  "context"
  "testing"

  "net/http"
  "encoding/json"
)

var client = http.DefaultClient

func Benchmark_Fiber(b *testing.B) {
  for i := 0; i < b.N; i++ {
    request, err := http.NewRequestWithContext(context.Background(), "GET", "http://localhost:2300", nil)
    if err != nil {
      panic(err)
    }

    resp, err := client.Do(request)
    if err != nil {
      panic(err)
    }

    data := struct{Message string `json:"message"`}{}
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
      panic(err)
    }
    
    fmt.Println(data.Message)
  }
}
