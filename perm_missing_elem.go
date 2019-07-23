package main

import (
	"fmt"
	"sort"
)

func main() {
        test := []int{2,3,1,5}
       
	//fmt.Println(len(test))
	fmt.Println(Solution(test))
}

func Solution(A []int) int {
    sort.Ints(A)
   fmt.Println(A)
   for i := 0 ; i < len(A);i++ {
     if A[i+1] - 1 != A[i] {return A[i+1] - 1 }
    }

return -1
}
