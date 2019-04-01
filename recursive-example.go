package main

import (
	"fmt"
)

func main() {

         fmt.Println(factcheck(3))
}

func factcheck (n int) int {

   if n == 0 {  return 1  }

   return n * factcheck(n-1)
}
