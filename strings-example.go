package main

import (
	"fmt"
	 "strings"
	 "sort"
)

func main() {
     mystr := "3zdhe,yafbe,bee,0cfem"
     mystr4 := []string{"a","bd","ee"}
     fmt.Println(strings.Contains(mystr,"be"))
     fmt.Println(strings.Index(mystr,"be") ) 
	   mystr1 := strings.Replace(mystr,",","",777) 
	   fmt.Println(strings.Replace(mystr1," ","",777) )
	   mystr2 := strings.Split(mystr,",")
	   sort.Strings(mystr2)
	   fmt.Println(mystr2,mystr)
	   fmt.Println(mystr2[0] + " " + mystr)
	   fmt.Println(strings.Join(mystr2,":"))
	   fmt.Println(strings.Join(mystr4, ","))
	
}
