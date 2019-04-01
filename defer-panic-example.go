package main

import (
	"fmt"
)

func main() {
  
        demopanic()
}

func demopanic() {

    defer func() {
      fmt.Println(recover())
}()
  fmt.Println("calm down, panic is coming")
  panic("This is panic, run away")
// fmt.Println("calm down, panic passed")
}
