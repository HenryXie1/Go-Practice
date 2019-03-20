package main

import (
	"fmt"
)

func main() {

        test1 :=  []int{4,5,1,2,3}
        var result  [5]int
	//fmt.Println("Hello, playground")
	fmt.Println(len(test1))
	for i :=0; i < len(test1)  ; i++ {
	       num := 0
	     for j := i + 1;  j < len(test1)  ; j++ {
	       //fmt.Println(test1[i],test1[j])
  	        if test1[j] > test1[i] {
	         //fmt.Println(test1[i],test1[j])
	         num++
	      }
	   result[i] = num 
	  }
	}
	
	fmt.Println(result)
	
}
