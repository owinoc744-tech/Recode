package main
import (
         "fmt"
         "os"
          )
func main(){
fmt.Fprintln(os.Stdout,"Hello developers an")
fmt.Fprintln(os.Stderr,"error occured")
 f,_:= os.Create("log.txt")
fmt.Fprintln(f,"Hello developers an error occured")

}
