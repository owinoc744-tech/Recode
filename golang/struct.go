package main

import "fmt"
type Student struct{

name string
age  int
course string

}
func main(){ 

  dev := Student{

name: "liam",
age : 21,
course : "java",
}
fmt.Println(dev.name)
fmt.Println(dev.age)
fmt.Println(dev.course)
}
