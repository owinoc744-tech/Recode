package main 

import (
"fmt"
)
func main(){
var red, green, blue uint8 = 0, 141, 213
var red1, green1, blue1 uint8 = 0x00, 0x8d, 0xd5
fmt.Printf("%x %x %x \n", red, green, blue)
fmt.Printf("%x %x %x \n", red1, green1, blue1)
fmt.Printf("color: #%02x%02x%02x; \n", red, green, blue)
}
