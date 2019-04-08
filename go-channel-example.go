package main

import (
	"fmt"
	
)
 
func mysum(s []int, c chan int) {
     
      sum  := 0
      for _, v := range s {
        sum = sum + v
    }
  //fmt.Println(sum)
    c <- sum
    
}

func main() {
            a := [5]int{1,2,3,4,5}
            c := make(chan int)
           
            fmt.Println(a[0:3],a[3:5])
            
            go mysum(a[:len(a)/2],c)
            go mysum(a[len(a)/2:],c)
	    x := <- c 
	    y := <- c
	  
    fmt.Println("Finally: ",x+y)
   
   	
}
