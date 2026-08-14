package main

import (
	"fmt"
	"your-module-name/piscine"
)

func main() {
	fmt.Println(piscine.Gcd(42, 10))
	fmt.Println(piscine.Gcd(42, 12))
	fmt.Println(piscine.Gcd(14, 77))
	fmt.Println(piscine.Gcd(17, 3))
	fmt.Println(piscine.Gcd(0, 10))
	fmt.Println(piscine.Gcd(10, 0))
}
