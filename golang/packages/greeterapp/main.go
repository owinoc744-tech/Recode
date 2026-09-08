package main
import(
    "fmt"
   "greeterapp/greeting"
)
func main(){
msg := greeting.Hello("Nairobi")

fmt.Println(msg)
}
