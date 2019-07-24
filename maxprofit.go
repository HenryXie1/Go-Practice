package main

import (
	"fmt"
	
)

func main() {
        test := []int{2333,9445,667,8887,8886}
       
	//fmt.Println(len(test))
	fmt.Println(Solution(test))
}

func Solution(A []int) int {
    minprice  := 1<<32
    maxprofit := 0

    for _,value := range(A){
       minprice = min(minprice,value)
       maxprofit = max(value-minprice, maxprofit)
       
     }
     return maxprofit
}


func min(x,y int) int {
  if x > y { 
    return y 
}
  return x
}


func max(x,y int) int {
  if x < y { 
    return y 
}
  return x
}
