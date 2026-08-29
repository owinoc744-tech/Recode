package main
import "fmt"
func main(){
 var score int
fmt.Println("Enter score")
 fmt.Scan(&score)

if score >= 50 {
  fmt.Println("above average")
}else{
   fmt.Println("below avarage")
}
}
