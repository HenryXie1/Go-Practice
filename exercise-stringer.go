package main

import "fmt"
//import "strconv"


type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func (a IPAddr) String() string {
      var b string
      for i:=0;i < 4;i++ {
	      if i ==3 {
		    b = b + fmt.Sprintf("%v",a[i])
			} else {  
		     b = b + fmt.Sprintf("%v",a[i]) + "."
		 }
		}
 //      fmt.Println(b)
	   return b
}
