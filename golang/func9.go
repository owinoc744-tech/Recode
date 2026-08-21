//func with parameteres
package main
import "fmt"
func minus(a int,b int)int{
res := a-b
fmt.Println(res)
return res
}
func add(a int,b int) int{
res := a+b
fmt.Println(res)
return res
}
func multiply(a int,b int) int{
res :=  a*b
return res
}
func main(){
  add(5,6)
 minus(50,41)
multiply(2,2)
}
