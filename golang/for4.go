package main
import (
"fmt"

)
func main(){

name := []string{"roy","toy","ray","buma"}
fmt.Println("original slice",name)

for Index ,value := range name{

fmt.Println("the index of",Index,value)

}
}
