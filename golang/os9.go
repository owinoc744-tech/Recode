package main
import (
"fmt"
"os"
)
func main(){
data,_ := os.ReadFile("hello.txt")
fmt.Println(string(data))

}
