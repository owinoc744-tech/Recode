package main 
import "fmt"
func main(){
 num := []int{0,1,2,3,4,5,6,7,8,9}
for i,n := range num {
   fmt.Println(i,n)
}
}
