package main 
import (
"fmt"
"os"
//"io"
)
func main(){
file ,_ := os.Create("no.txt",file,0755)
fil,_ := os.OpenFile("no.txt")
defer file.Close()
fi :=os.WriteFile("no ooh noo")

fmt.Println(file,fil)
fmt.Println(fi)
}
