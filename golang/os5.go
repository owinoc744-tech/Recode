package main

import (
"fmt"
"os"
//""
)
func main(){
 data := []byte("Hello from Go\n Line2")
  dat:= os.WriteFile("hello.txt",data , 0744)
fmt.Println(data)
fmt.Println(dat)
}

