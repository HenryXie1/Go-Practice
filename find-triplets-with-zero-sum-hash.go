package main

import (
      "fmt"
     
)

func main () {
 
      test1 := []int{0, -1, 2, -3, 1}
      test2 := []int{1,2,3}
      findtiplet(test1,5)
      findtiplet(test2,3)
     
  }

func findtiplet (t []int, n int )   {

     fmt.Println(t)
     found := false
     for i := 0 ; i < n -1 ; i ++ {
          s := []int{}
          k :=0
          for j := i + 1; j < n  ; j++ {
               x := -(t[i]+t[j])
              if containsx(s,x) {
              fmt.Printf("%d,%d,%d\n",x,t[i],t[j])
              found = true 
            } else {
              s = append(s,t[j])
       //    fmt.Printf("s is %v\n",s)
             k++
              }
         }
     
      }
   if found == false { 
     fmt.Println("Not found triplet")  
  }
 
}

func containsx(t []int, x int) bool {
       for _, n := range t {
        if n == x {
          return true
       } 
     }
   return false
}
