package main

import(
"fmt"
)
func main(){
num := make([]int,5)
nu := make([][]int,7*4)
for i:=0;i<5;i++{
num[i]= i
   }
fmt.Println(num)
fmt.Println(nu)
}
