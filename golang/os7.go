package main
import (
"os"
"fmt"
)
func main(){
name := []byte("mary abi nona \n peter mike")
nam := os.WriteFile("n.txt", name,0744)
fmt.Println(nam)

}
