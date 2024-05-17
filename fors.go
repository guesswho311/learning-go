package main

import "fmt"

func main() {
  i := 1
  for i <= 3 {
     fmt.Println(i)
     i = i + 1
  }

  for j:=0; j<3; j++ {
    fmt.Println(j)
  }

  ii := 1
  for {
    fmt.Println("loop until break")
    ii = ii + 1
    if ii > 3 {
      break
    }
  }

  for n := range 6 {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}