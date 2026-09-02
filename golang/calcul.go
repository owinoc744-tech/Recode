package main

import "fmt"

func main() {
    var num float64
    var op string
    var num2 float64

 for {
 fmt.Println("Enter your calculation (or 'q' to quit):")
        fmt.Scan(&num, &op, &num2)

        if op == "q" {
            break
        }

        switch op {
        case "+":
            fmt.Println("Result =", num+num2)
        case "-":
            fmt.Println("Result =", num-num2)
        case "*":
            fmt.Println("Result =", num*num2)
        case "/":
            if num2 != 0 {
                fmt.Println("Result =", num/num2)
            } else {
                fmt.Println("Error: division by zero")
            }
        default:
            fmt.Println("Invalid operator")
        }
    }
}
