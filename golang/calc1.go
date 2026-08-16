package main
import(
"fmt"
)
func main(){
fmt.Println("=====CALCULATOR=====")
fmt.Println("Enter yuor calculation")
var num  float64
fmt.Println("Enter first number")
fmt.Scanln(&num)
fmt.Println("Enter operator")
var op string =("+,-,/,*")
fmt.Scanln(&op)
var num1 float64
fmt.Println("Enter 2nd number")
fmt.Scanln(&num1)

switch op {
case "+":fmt.Println(num+num1)
case "-" :fmt.Println(num-num1)
 case "/" :fmt.Println(num/num1)
case "*" :fmt.Println(num*num1)

if num == 0 || num1 == 0 {
  fmt.Println("deviding by 0 is infinity")
}
}}
