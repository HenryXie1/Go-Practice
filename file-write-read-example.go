package main

import (
	"fmt"
	 "os"
	 "log"
	 "io/ioutil"
	 
)

func main() {

        file, err := os.Create("test.txt")
   
           if err != nil {
            log.Fatal(err)
          }

       file.WriteString("this is radom test string")
       file.Close()
       stream,err := ioutil.ReadFile("test.txt")
         if err != nil {
            log.Fatal(err)
          }

       fmt.Printf("Type stream is %T, convert to string to be : %v",stream,string(stream))

}
