// Arrays and Slices
package main

import "fmt"

func main(){
array1 :=  [4]string {" gray","brown" ,"yellow","rouge" }
 var array2 [5]int =[5] int{1,2,3,4,5}
  array1 [1] = "black"
   slice := [] int {9,8,7,6,5,4,3}
    slice = append(slice ,2) 
     var slice1 []float64 =[]float64{12.2,13.4,12.3,15.3}
      range1 := slice[0:5]
       city := []string{"newyork","newjersey","neworleans"}
        state := []string{"michigan","minnesotta",}
        fmt.Println(city,len(city))
       fmt.Println(slice1,len(slice1))
      fmt.Println(array1,len(array1))
     fmt.Println(array2,len(array2))
    fmt.Println(slice,len(slice))
   fmt.Println(range1,len(range1))

}
