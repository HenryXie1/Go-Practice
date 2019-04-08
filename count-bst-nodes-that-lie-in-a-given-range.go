package main

import (
	"fmt"
)

type node struct {
    data int
    left *node
    right *node
}


func newnode(data int)  *node {
   
     var temp node
     temp.data = data
     temp.left = nil 
     temp.right = nil
     return &temp
         
}
func main() {
        myroot := newnode(10)
        myroot.left = newnode(5)
        myroot.right = newnode(50)
        myroot.left.left = newnode(1)
        myroot.right.left = newnode(40)
        myroot.right.right = newnode(100)
 //     fmt.Println(myroot.right.right.data)

        fmt.Println(getcount(myroot, 5,45))

}

func getcount(root *node, i int, h int) int {
   if root == nil { return 0 }
   if root.data >= i && root.data <= h  {
     return 1 + getcount(root.left,i,h) + getcount(root.right,i,h)
   }

  if root.data < i {
      return getcount(root.right,i,h)
   }
  if root.data > h {
   return getcount(root.left,i,h)
  }

  return 0
}


