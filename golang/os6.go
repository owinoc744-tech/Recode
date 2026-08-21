package main
import (
"fmt"
"os"

)
func main(){

data := os.ReadFile("hello.txt")
fmt.Println(os.ReadFile("hello.txt"))

}
