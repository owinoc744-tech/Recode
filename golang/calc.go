package main

import "fmt"

func main(){
   fmt.Println("##Go Calculator ##")
    var num float64
     var op = ("+-*/")
       var num2 float64
        fmt.Println("Enter your calculation")
        fmt.Scan(&num,&op,&num2)
       switch op{
      case "+":fmt.Println(num+num2)
     case "-":fmt.Println(num-num2)
    case "/":fmt.Println(num/num2)
   case "*":fmt.Println(num*num2)
}}
