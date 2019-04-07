package main

import (
	"fmt"
	//"strings"
	//"strconv"
)

func main() {
        str := "1001??0?101"
	mstr := []byte(str)
	q := []byte("?")[0]
	a := []byte("0")[0]
	b := []byte("1")[0]
	fmt.Println(mstr)
        allstr(mstr,0, len(mstr),q,a,b)
	
}

func allstr(mstr []byte, index int, size int, q,a,b byte) {

       if index >= size  { 
          fmt.Println(string(mstr) )
          return
       }
       if mstr[index] == q {
           mstr[index] = a
           allstr(mstr,index + 1,size,q,a,b)

           mstr[index] = b
           allstr(mstr,index + 1,size,q,a,b)

           mstr[index] = q   //back track

   } else {
  
    allstr(mstr,index + 1,size,q,a,b)   
  }

  
} 
