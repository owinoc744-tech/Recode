package main
import (
"fmt"
"os"

)
func main(){
   fish := []byte("tilapia \n dagaa \n salamander")
   fis := os.WriteFile("fish.txt",fish ,0744)
fmt.Println(fis)
fmt.Println(fish)
fmt.Println("data written in fish.txt")
}
