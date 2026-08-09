package main
import(
       "fmt"
       "strings"
       "sort"
)
func main(){

var name = "allan berry carlos  derick ethan ford"

fmt.Println("original string:", name)

fmt.Println(strings.Contains( name,"berry"))
//fmt.Println(strings.SearchStrings(name,"ethan"))
fmt.Println(strings.ToUpper(name))
fmt.Println(strings.ReplaceAll( name,"ford","ferdy"))
fmt.Println("the index of ethan")
fmt.Println(strings.Index(name,"ethan"))
//  var num int = 9 3 7 2 8 4 2 5
   na := "berry allan carl Dan Eric" 

fmt.Println("original int:",na)
  fmt.Println(strings.Split(na," "))
words := []string {"berry allan carl Dan Eric"}
 //fmt.Println(sort.Strings(na))
fmt.Println(words,len(words)
 x :=sort.String(words)
}
