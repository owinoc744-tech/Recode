package main
import "fmt"
func add(){
a := 4
b:= 5
fmt.Println("add",a+b)
}
func substract(){
a:= 28
b := 12
fmt.Println("subtract",a-b)
}
func multiply(){
a := 4
b := 7
fmt.Println("multiply",a*b)
}

func main(){
add()
substract()
multiply()
devide()
modulo()
}
func devide(){
a := 56
b:= 4
fmt.Println("devision",a/b)
}
func modulo(){

a := 45
b := 20
fmt.Println("modulus",a%b)
}
