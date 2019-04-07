package main

import (
	"fmt"
)

func main() {

       str := "aabacbebebe"
        k := 3
	fmt.Println(str,k)
	bstr := []byte(str)
	count1 := 0
	u := make([]byte,3)
	i := 0
	j := 0
	res := 0
	 for i=0;i<len(bstr);i++ {
	        if count1 < k {  u[count1]=bstr[i] }
	         
	    for j=i;j+i<len(bstr);j++ {
	      if count1 >= k-1 { 
	        // fmt.Println(u,i, j) 
	         res = gmax(res,j)
	         count1 = 0
	         break
	       }
	      if bstr[i] != bstr[i+j]  { 
	          count1++
	         u[ count1 ] = bstr[i+j]  
	        // fmt.Println(u,count1,i, j)
	        }
	 }
	 
	}
	fmt.Printf("The longest length is : %v",res+1)
}

func gmax (x,y int) int {
   if x > y {
     return x 
  } else { return y }
}
