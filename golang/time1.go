package main

import ("fmt"
"time"
)
func main(){
fmt.Println(time.Now())
clo := time.Now()
cloc := clo.Add(2*time.Hour)
fmt.Println(cloc)
}
