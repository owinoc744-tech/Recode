package main
import "fmt"
func main(){
// float64 stands for whole numbers with very many decimals places as compared to float32 
var num float64
var op = ("*,+,-,/")
var num2 float64
fmt.Println("Enter your calculation:")
fmt.Scan(&num,&op,&num2)
fmt.Println("addition=",num+num2)
fmt.Println("substraction=",num-num2)
fmt.Println("Multiplication=",num*num2)
fmt.Println("Division=",num/num2)
}
