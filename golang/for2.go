package main 
import (
"fmt"
"strings"
"strconv"

)
func main(){
var na [5] string = [5]string{"a","b","c","d","e"}
fmt.Println("original array",na,len(na))
n := "allanwalker"
fmt.Println("original string",n)
fmt.Println(strings.ToUpper(n))
fmt.Println(strings.Index(n,"l"))
fmt.Println(strconv.Atoi(n))
}
