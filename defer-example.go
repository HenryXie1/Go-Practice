package main

import (
	"fmt"
)

func main() {
  
         defer print2()
         print1()
         fmt.Println("heello")
         fmt.Println("heello again")
}

func print1() { fmt.Println(1) }
func print2() { fmt.Println(2) }
