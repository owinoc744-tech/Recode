package main
import(
       "fmt"
       "sort"
   )
func main(){
ages := []int{12,14,13,51,17,19,29}
fmt.Println(ages)
sort.Ints(ages)
fmt.Println(ages)

index := sort.SearchInts(ages,51) 
fmt.Println("the index of 51is below")
fmt.Println(index)

names := []string{"dorothy","caroline","abigail","betty","esta"}
fmt.Println(names)
fmt.Println("The index of caroline is below from the original slice ")
fmt.Println(sort.SearchStrings(names,"caroline"))
sort.Strings(names)
fmt.Println(names)

}
