package main

import (
	"fmt"
)

var a int

type cerveja int

var b cerveja

func main() {
	a = 42
	b = 12
	fmt.Println(a)
	fmt.Printf("%T", a)
	fmt.Println(b)
	fmt.Printf("%T", b)

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T", a)
}
