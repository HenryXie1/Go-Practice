package main

import (
	"fmt"
)

func main() {

        test := []int{4,3,5,1,3}
	//fmt.Println(len(test))
	fmt.Println(Solution(test,3))
}

func Solution(A []int, k int) []int {
     B := make([]int,len(A)) 
     copy(B,A)
     fmt.Println(A)
    fmt.Println(B)
     for i:=0;i<len(A);i++ {
         if i + k < len(A) {
           B[i+k] = A[i]  
        fmt.Printf("less : %v\n",B)
    } else {
      B[(i+k) % len(A) ] = A[i]
       fmt.Printf("more: %v\n",B)
       fmt.Println((i+k) % len(A) )
       fmt.Printf("%v,%v\n",i,A[i])
}   
}
return B
}
