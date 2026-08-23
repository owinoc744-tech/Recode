package main

import (
	"fmt"
	"weareunique/piscine"
)

func main() {
	fmt.Println(piscine.WeAreUnique("foo", "boo"))
	fmt.Println(piscine.WeAreUnique("", ""))
	fmt.Println(piscine.WeAreUnique("abc", "def"))
}
