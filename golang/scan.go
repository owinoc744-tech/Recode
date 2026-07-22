package main
import "fmt"
func main(){

var name string
var id int
fmt.Println("Enter your name")
fmt.Scan(&name)
fmt.Println("Enter your id number")
fmt.Scan(&id)
fmt.Printf("Welcome to techinsigt %s of idnumber %v \n",name,id)
}
