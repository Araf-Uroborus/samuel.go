package main

import (
	"fmt"
)

var x int
var y = 64
var z = "Sou uma variavel string"

func main() {
	fmt.Println(y)      // ln = line
	fmt.Printf("%T", y) // f = format (formatação)
	fmt.Println("E agora?")
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
