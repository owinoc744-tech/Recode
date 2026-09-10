package main

import "fmt"

type Car struct { 
name string 
year int 
}
func main (){
voom:= Car{name: "toyota",year: 1959 }

fmt.Println(voom.name)
fmt.Println(voom.year)


}
