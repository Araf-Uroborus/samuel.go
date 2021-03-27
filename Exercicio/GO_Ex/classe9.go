package main

import (
	"fmt"
)

var y = 32

func main() {
	fmt.Printf("Em hexadecimal com 0x %#X, em binario %b, e em hexadecimal %X\n", y, y, y)
	fmt.Printf("o valor de y é %v\t, já os demais são: %#X\t%b\t%x\n", y, y, y, y)
}
