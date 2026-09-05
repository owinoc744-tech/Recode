package main
import "fmt"
func main(){
na := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
fmt.Println("printing alphabet using range")
for _,n := range na{
 fmt.Print(n)
}
fmt.Println()
}
