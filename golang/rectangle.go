package main
import (
"fmt"

)
func LeftCorner(){
 fmt.Print("\n")
    fmt.Print("/")
}
    func UpperRow(){
      for i:=0;i<5;i++{
        fmt.Print("*")
}}
       func RightCorner(){
         fmt.Print("\\")
}
          func SecondRow(){
             fmt.Println()
                fmt.Print("*     *")
}
              func ThirdRow(){
               fmt.Println()
                fmt.Println("*     *")
}
              func FourthRow(){
            fmt.Print()
          fmt.Println("*     *")
}
      func LowerCorner(){
     fmt.Print("\\")
for i:=0;i<5;i++{
fmt.Print("*")
}
fmt.Print("/")
}
func LowerLine(){
  fmt.Print("\n")
}
func main(){
LeftCorner()
UpperRow()
RightCorner()
SecondRow()
ThirdRow()
FourthRow()
LowerCorner()
LowerLine()
}
