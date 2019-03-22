package main

import (
	"fmt"
)

func main() {
        n := 4
        res := make([][2][3]int,n+1)
 //     var res [][2][3]int
	fmt.Println( "heeelo",mycount(res,n,1,2))
	 
}

func  mycount(res [][2][3]int, n int, bcount int,ccount int) int {
     if bcount < 0  { return 0 }
     if ccount < 0 { return 0 }
     if n == 0 { return 1 } 
     if bcount == 0 && ccount == 0 { return 1}
     if res[n][bcount][ccount] != 0 {
         return res[n][bcount][ccount]
      }

    count1 := mycount(res,n-1,bcount,ccount)
    count1 += mycount(res,n-1,bcount-1,ccount)
    count1 += mycount(res,n-1,bcount,ccount-1)
    res[n][bcount][ccount] = count1

   return count1
}
