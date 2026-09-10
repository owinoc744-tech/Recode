package main

import "fmt"

func main(){

na := "letty" 
nb :=  &na
fmt.Println(&na)
//fmt.Println(na)
fmt.Println(*nb)

}
