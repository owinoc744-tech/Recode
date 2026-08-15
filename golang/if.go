package main
import(
"fmt"
)
func main(){
size := 56
fmt.Println(size == 45)
fmt.Println(size > 56)
fmt.Println(size < 56)
fmt.Println(size == 56)
if size ==56 {
fmt.Println("it fits")
}else if size < 56{
  fmt.Println("too tight")
}/* else size > 56{
fmt.Println("too big")
}*/
}
