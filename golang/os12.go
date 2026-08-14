package main
import (
"fmt"
"os"
)
func main(){
l,_ := os.Create("l.txt")

 l.WriteString("never have i ever")

fmt.Println(l)
}
