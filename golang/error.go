package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("divide by zero")
    }
    return a / b, nil
}

func main() {
    res, err := divide(5, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println(res)
    }
}

