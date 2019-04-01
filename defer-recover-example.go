package main

import (
	"fmt"
)

func main() {
  
         
         fmt.Println(safediv(3,0))
         fmt.Println(safediv(3,2))
}

func safediv(x, y int) int {

    defer func() {
      fmt.Println(recover())
}()

  return x / y
}
