package main
import(
    "sort"
    "fmt"
)
func main(){

num := []int{9,6,7,3,2,1,4,8}
num = append(num,10)

fmt.Println(num)
sort.Ints(num)
fmt.Println(num)
fmt.Println(sort.SearchInts(num,5))
}
