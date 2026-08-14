package main
import (
"os"
"fmt"
)
func main(){
file,_ := os.Create("data.txt")
  defer file.Close()
  file.WriteString("Golang go rock baby")
fmt.Println(file)
}
