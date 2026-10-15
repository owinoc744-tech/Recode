package main

import (
"fmt"
"math"
)
func main(){
//degrees 5o radians
de := 90.0
ra := de *math.Pi/180
fmt.Println("90° to radians:",ra)
rad := 1.5707963267948966
deg := rad*180/math.Pi
fmt.Println("1.5707.. to degrees:",deg)

}
