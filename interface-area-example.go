package main

import (
	"fmt"
	"math"
)

func main() {
  
        rect1 := rectangle{2,3,5,7}
        circ1 := circle{6} 
        fmt.Println(getarea(rect1))
        fmt.Println(getarea(circ1))
}


type rectangle struct {
  x float64 
  y float64 
  height float64 
  width float64 
  }

type circle struct {
   radius float64 
}

type shape interface {
  area() float64 
}

func (rect rectangle) area() float64 {
  return rect.height * rect.width
   
}

func(circ circle) area() float64 {
  return math.Pi * math.Pow(circ.radius,2)
}

func getarea(shap shape) float64 {
   return shap.area()
} 
