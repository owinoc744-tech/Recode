package main
import (
"fmt"
"time"
)
func main(){
n := "kenya"
fmt.Println(n)
clock := time.Now()
fmt.Println(clock)
tomorrow := clock.Add(24*time.Hour)
fmt.Println(tomorrow)
fmt.Println(clock.Year())
k:= time.Now()
fmt.Println(k)
c := k.Add(-200*time.Hour)
fmt.Println(c)
}

