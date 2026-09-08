package main

import "fmt"

func main() {
    name := "Collins"

    {
        name := "John"

        fmt.Println(name)
    }

    fmt.Println(name)
}
