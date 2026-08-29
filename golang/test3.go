package main
import "fmt"
func main(){
var name string
fmt.Println("Enter your name")
fmt.Scan(&name)
fmt.Println(name)
var email string
fmt.Println("Enter your email")
fmt.Scan(&email)
fmt.Println(email)
 var age int
fmt.Println("Enter your age")
fmt.Scan(&age)
if age >= 18 {
 fmt.Println("qualified")
}else if age < 18{
  fmt.Println("you age does not qualify")

}
var phone int
fmt.Println("Enter your phone number")
fmt.Scan(&phone)
fmt.Println(phone)


}
