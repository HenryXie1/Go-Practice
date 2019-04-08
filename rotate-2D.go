package main

import (
	"fmt"
)

func main() {
         array1 := []int{1,2,3,4,5,6,7,8,9}
         n := 3
         array1_2d := [3][3]int{}
         array2_2d := [3][3]int{}
         
         for i := 0; i < n ; i++ {
            for j:=0; j < n; j++ {
               array1_2d[i][j] = array1[i*n + j]
   }
  }

//fmt.Println(array1,array1_2d)

         count :=0
         for i := n-1; i > -1; i--{
              for j := 0; j < n ; j++ {
                  //fmt.Println(count,j," : ", j,i)
                  array2_2d[j][i]=array1_2d[count][j]
                  //fmt.Println(array2_2d)
                 
                              
 } 
     count++
     
}
	fmt.Println(array1_2d," vs ",array2_2d)

}
