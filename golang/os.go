package main
import (
     "fmt"
      "os"
)
func main(){
//na := os.Args[1]
//fmt.Println(na)
os.Exit(1)
fmt.Println("my pid",os.Getpid)
fmt.Println("my ppid",os.Getppid)
os.Exit(0)
}
