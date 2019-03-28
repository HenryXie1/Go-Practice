package main

import (
	"fmt"
)

func main() {
  
        replacestr("aababc")
}

func replacestr(str string) string {
   bstr := []byte(str)
   res  := make ([]byte,len(bstr) )
   res1  := make ([]byte,len(bstr) )
   for i:=0; i < len(bstr) ; i++ {
      if bstr[i] != []byte("b")[0] {
            res[i] = bstr[i]
            
     }
}

//fmt.Println(res)
res1 = res

for i:=0; i < len(res) - 1 ;i++ {
      if res[i] == []byte("a")[0] && res[i+1] == []byte("c")[0]  {
             res1[i]=0
             res1[i+1]=0
     } else {
         res1[i] = res[i] 
         
      }
}

 // fmt.Println(res1)
finalres:= string(res1[:len(res1)])
fmt.Println(finalres )

return finalres

}
