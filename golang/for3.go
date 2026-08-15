package main
import (
"fmt"
)
func main(){

 name  := []string{"mario","yoshi","bowser","lenny","tom"} 

for c := 0;c < len(name);c++{

fmt.Println(name[c])
}

fmt.Println(name)
}
