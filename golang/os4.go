package main
import (
"fmt"
"os"
"strings"

)
func main(){
na := "name"
fmt.Println(na)
fmt.Println(&na)
fmt.Println(strings.ToUpper(na))
   os.Setenv("penguin","true")
fmt.Println(os.Getenv)
fmt.Println(os.Setenv)
fmt.Println(os.LookupEnv)
fmt.Println(os.LookupEnv(na))

all := os.Environ()
fmt.Println(all)
}
