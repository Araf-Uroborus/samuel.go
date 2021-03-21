package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println("Olá Mundo")
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Printf("%b\n", y)  //em binario
	fmt.Printf("%x\n", y)  //em hexadecimal
	fmt.Printf("%#x\n", y) // em hexadecimal 0x na frente
	fmt.Printf("%X\n", y)  // hexadecimal em maiusculo
}
