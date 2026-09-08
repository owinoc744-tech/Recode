package main

import "fmt"

var language = "Go"

func main() {
    name := "Collins"

    if true {
        age := 19

        fmt.Println(language)
        fmt.Println(name)
        fmt.Println(age)
    }

    fmt.Println(language)
    fmt.Println(name)

    // fmt.Println(age) // ERROR
}
