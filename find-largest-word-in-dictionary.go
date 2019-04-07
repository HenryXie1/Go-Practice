package main

import (
	"fmt"
)


func main() {
        dict := []string{"ale", "apple", "monkey", "plea"} 
        str := "abpcplea"  
	fmt.Println(dict)
	findstr(dict, str)
	
	
}

func findstr(dict []string, str string)  int {
      bstr := []byte(str)
      inx :=0
      for i:=0; i < len(dict); i++ {

             bdict := []byte(dict[i])
             bstr1 := bstr
             res := false
             lenstr := 0
             n := len(bstr1)
             m := len(bdict)
              k := 0
            for j:=0; j < n && j < m; j++ {

                  if bstr1[j] == bdict[k] { 
                          k++ 
                          res = true
                       } else { res = false }
                 }
                 

       if res == true && lenstr < len(bdict) {
            lenstr = len(bdict)
            inx = i
             }

   }
   
   fmt.Println(dict[inx])
   return inx
}
