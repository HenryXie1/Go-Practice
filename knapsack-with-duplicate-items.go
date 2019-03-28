package main

import (
	"fmt"
)
func main() {

          test1 := []int{4,8}
          test1v := []int{1,4,5,7}
          test1w := []int{1,3,4,5}
        
	fmt.Println(kanpsack(test1[0],test1[1],test1v,test1w))
}

func kanpsack(n int, w int, testv []int, testw []int ) int {

        
          res := make([]int,w+1)

       for i:=0 ; i <= w ; i++ {
         for j:=0;j < n ; j ++  {
             if i >= testw[j] {
           res[i] = max(res[i],res[i-testw[j]] + testv[j] )
            }
     }
}
  
   return res[w]    
}
func max(x ,y int) int {
  if x > y { 
   return x
  } else {
  return y

 } 
}
