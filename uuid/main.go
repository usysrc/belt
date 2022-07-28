package main

import (
  "fmt"
  "github.com/google/uuid"
  "flag"
)

func main() {
  
  count := flag.Int("n", 1, "number of UUIDs to generate")
  verify := flag.String("v", "", "verify a UUID")

  flag.Parse()

  if *verify != "" {
    _, err := uuid.Parse(*verify)
    if err != nil {
      fmt.Println("invalid")
      return
    }
    fmt.Println("valid")
    return
  }

  for i:=0; i < *count; i++ {
    fmt.Println(uuid.NewString());
  }
}
