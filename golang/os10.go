package main
import (
"fmt"
"os"
//"strconv"
)
func main(){

n,_ := os.ReadFile("mybio.txt")
fmt.Println(n)
fmt.Println(string(n))

}
