package main

import (
	"fmt"
)

func main() {
        
        matrix1 := []int{0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 2, 3 }
        xlen :=3
        ylen :=4
        x := 0
        y :=1
        k:=5
        c := matrix1[x*xlen + y]
	fmt.Println(matrix1,xlen,ylen,x,y,k,c)
	fillit(matrix1,xlen,ylen,x,y,k,c) 
	 fmt.Println(matrix1)
}

func fillit(matrix1 []int, xlen, ylen, x, y, k, c int )  {

  
        if x < 0 || y < 0 || x >= xlen || y >= ylen { return }
        if matrix1[x*xlen + y ] != c { return }
              
        matrix1[x*xlen + y ] = k
    
         fillit(matrix1,xlen,ylen,x,y-1,k,c )         
         fillit(matrix1,xlen,ylen,x,y+1,k,c )         
         fillit(matrix1,xlen,ylen,x-1,y,k,c )         
         fillit(matrix1,xlen,ylen,x+1,y,k,c )         

}
