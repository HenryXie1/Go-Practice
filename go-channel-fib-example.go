package main

import (
	"fmt"
)

func worker(input chan int, result chan int) {
     
     for n := range input {
     r := fib(n)
     fmt.Println("fib result of ",n," is ",r)
     result <- r
   }
}

func fib(n int) int {
   if n == 0 || n == 1 { return n }

   return n + fib(n -1 )
}


func main() {
         jobs := make(chan int,50)
         result := make(chan int,50)
 
         go worker(jobs,result)
         go worker(jobs,result)
         go worker(jobs,result)
         go worker(jobs,result)

         for i := 0;i < 50;i++ {
          jobs <- i
         }
         close(jobs)
         for j := 0;j < 50;j++ {
            fmt.Println( <- result )
         }
	
}
