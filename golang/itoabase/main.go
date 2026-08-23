package main

import (
	"fmt"
	"itoabase/piscine"
)

func main() {
	fmt.Println(piscine.ItoaBase(42, 10))
	fmt.Println(piscine.ItoaBase(42, 2))
	fmt.Println(piscine.ItoaBase(255, 16))
	fmt.Println(piscine.ItoaBase(-42, 10))
	fmt.Println(piscine.ItoaBase(63, 4))
}
