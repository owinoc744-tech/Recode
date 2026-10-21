package main

import (
"fmt"
//"math/rand"
)
func main(){
nu := 59
count := 0
for count>=0 && count<=100{
   count++
fmt.Println("count",count,"num=",nu)
if count == nu{
break
}
}
}
