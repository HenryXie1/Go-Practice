package main

import (
      "fmt"
      "sort"
)

func main () {
 
      test1 := []int{0, -1, 2, -3, 1}
      test2 := []int{1,2,3}
      findtiplet(test1,5)
      findtiplet(test2,3)
     
  }
type myint []int

func (a myint) Len() int { return len(a) }
func (a myint) Swap(i,j int) { a[i],a[j] = a[j],a[i] }
func(a myint) Less(i,j int) bool { return a[i] < a [j] }
 
func findtiplet (t []int, n int ) bool  {

     sort.Sort(myint(t))
     fmt.Println(t)
     var sum int
     for i := 0 ; i < n -2 ; i ++ {
          for j := i + 1; j < n -1 ; j++ {
               
              if sum = t[i] + t[j] + t[j+1] ; sum == 0 {
         fmt.Println("true")
         return true 
         } 
    }
    
  }
fmt.Println("false")   
return false
}
