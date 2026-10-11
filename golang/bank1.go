package main
import "fmt"
var cho uint
var mon uint = 0
func CheckBal(){
fmt.Println("your balance is Ksh",mon)
}
func Depo(){
var Dep uint
fmt.Println("Enter amount to deposit")
fmt.Scan(&Dep)
mon += Dep
fmt.Println("your new balance is Ksh",mon)
}
func Withdraw(){
 var with uint 
fmt.Println("Enter amount to withdraw")
fmt.Scan(&with)
mon -= with
fmt.Println("your new balance is",mon)
}
func main(){
for {
fmt.Println("Enter 1 to checkbalance")
fmt.Println("Enter 2 to deposit")
fmt.Println("Enter 3 to withdraw")
fmt.Println("Enter 4 to exit")
fmt.Println("Enter choice")
fmt.Scan(&cho)
if cho ==1 {
  CheckBal()
   }else if cho == 2{
Depo()
     }else if cho == 3{
    Withdraw()
}else if cho == 4 {
    break
  }else{
  fmt.Println("Invalid Choice")
    }
  }
}
