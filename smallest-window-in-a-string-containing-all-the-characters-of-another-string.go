package main

import (
	"fmt"
)

func main() {
       i := 0
       s := "timetopractice"
       t := "toc"
       bs := []byte(s)
       bt := []byte(t)
        
        lenbt := len(bt)
        lenbs := len(bs)
        res := lenbs
       if lenbt > lenbs  { 
       fmt.Println(" Subtext is longer than string")
       return
 }
       
       for i=0;i<lenbs;i++ {
              
              res = min(findnum(bs,bt,i,lenbt,lenbs),res)
	   
	}
      fmt.Println(res)	
}

func min (x,y int) int {
   if x > y { 
   return y
 } else { return x }
}

func findnum(bs []byte,bt []byte, i int,lenbt int ,lenbs int ) int{
    j:=0
    k:=0
    ends := 0
    flagbt := make([]int, lenbt)

   if i > lenbs - lenbt  { 
    //   fmt.Println(" not found !")
       return lenbs  
 }
   for k=i;k<lenbs;k++{
    for j=0;j<lenbt;j++  {
       if bt[j]==bs[k] && flagbt[j] == 0 { 
          // fmt.Println(i,k,j,bt[j])
           flagbt[j]= 1
           ends = k
           break
        }
     }
 }

for j=0;j<lenbt;j++ {
    if flagbt[j] == 0 { 
 //    fmt.Println("Not found")
     return  lenbs
    } 
   
}
fmt.Println("found string is" , string(bs[i:ends+1]) )
return ends-i+1
}
