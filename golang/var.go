package main
// package main maens this is an executable program 
import "fmt"
// import  allows us to utilize go packages 
func main(){
// funcmain this is where we start to execute  our code
var name string = "Collins"
// var store data of different types
fmt.Printf("My name is %v i am a and i am a dev \n ",name)
fmt.Println(&name)
// &name this lets us print the address where name is stored
}
