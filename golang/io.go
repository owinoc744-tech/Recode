package main
import (
"fmt"
"os"
"io"
)
func main(){
  file,_ :=os.Open("hello.txt")
  fil,_ := os.ReadFile("hello.txt")
  data,_ := io.ReadAll(file)
fmt.Println(file)
fmt.Println(fil)
fmt.Println(string(fil))
fmt.Println(data)
fmt.Println(string(data))
}
