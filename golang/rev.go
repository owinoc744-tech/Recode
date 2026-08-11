package main
import (
       "strings"
       "sort"
       "fmt"
      "time"
)
func main(){
var greet string = "welcome to the review"
fmt.Println(greet)
  time.Sleep(2*time.Second)
  var name string
 fmt.Println("Enter your name")
fmt.Scan(&name)
  time.Sleep(1*time.Second)
fmt.Printf("welcome %s \n",name)
 var num [8]int =[8]int{1,2,4,5,6,7,8,3} 
fmt.Print("here is your array")
fmt.Println(num)
fmt.Println("below is a slice containg your friends names" )
  time.Sleep(2*time.Second) 
 friend := [] string {"Berry","Aaron","Deen","Carl"}
fmt.Println(friend)
fmt.Println("and here it is sorted in alphabetical order")
sort.Strings(friend)
fmt.Println(friend)
fmt.Println("now we we want to search the index of Aaron in the sorted slice")
time.Sleep(2*time.Second)
//search.Index(friend)
//fmt.Println(Search.Index(friend,"Aaron"))
//fmt.Println("this is the index of Aaron ")
//fmt.Println(Strings.ReplaceAll("Aaron" ,"Abel"))
friends := append (friend,"Erick")
fmt.Println(friends)
f := "Erick"
fmt.Println(strings.Contains(f,"Erick"))

}
