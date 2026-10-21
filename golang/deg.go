package main

import (
"math/rand"
"fmt")

func main(){
var degrees = 0
for {
 fmt.Println(degrees)
 degrees++
 if degrees >= 360 {
 degrees = 0
 if rand.Intn(2) == 0 {
 break
 }
 }
}

}
