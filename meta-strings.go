package main

import (
	"fmt"
)

func main() {
	mstra := "ageek"
	mstrb := "gaeek"
//	fmt.Println(mstra,mstrb)
	fmt.Println(metastrcheck(mstra,mstrb))
}


func metastrcheck(mstra string, mstrb string)   int {

    if mstra == mstrb { return  0 }

    if len(mstra) != len(mstrb) {  return 0 }
    
    stra := []byte(mstra)
    strb := []byte(mstrb)
    strc := make([]byte,len(mstra)) 

    for i:=0; i < len(stra); i++ {
             for j:=0;j < len(strb);j++ {

             if stra[i] == strb[j] { 
                 strc[i] = strb[j]
                 strb[j] = 0
                 break

                   } 

       }
    }
 //   fmt.Println(stra,strb,strc) 
    mstrc := string(strc[:len(strc)])
    fmt.Println(mstra,mstrc) 
    if mstrc == mstra { 
      return 1
 }  else { 
    return 0 
   }

}
