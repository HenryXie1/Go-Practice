package main

import (
	"fmt"
)

func main() {
  
        rect1 := rectangle{2,3,5,7}
        fmt.Println(rect1.area())
}


type rectangle struct {
  x int
  y int
  height int
  width int
  }

func (rect *rectangle) area() int {
   return rect.height * rect.width
}
