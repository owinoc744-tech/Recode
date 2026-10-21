package main

import (
"fmt"
"strings"
)
func main(){
fmt.Println("After leaving the cave you enter a dark forest")
var forest = "Monsters live here"
var hide = strings.Contains(forest,"Monsters")
fmt.Println("Hide as you run",hide )
}
