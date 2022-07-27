package main

import (
  "fmt"
  "github.com/google/uuid"
  "flag"
)

func main() {
  
  var count int
  flag.IntVar(&count, "n", 1, "number of UUIDs to generate")
  flag.Parse()

  for i:=0; i < count; i++ {
    fmt.Println(uuid.NewString());
  }
}
