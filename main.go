package main

import (
	"fmt"
)

var y string
var z int

func main() {
	y = "Bond, James Bond"
	fmt.Println("Imprimindo o valor de y ", y, " e pronto.")
	fmt.Printf("%T\n", y)

	fmt.Println("Imprimindo o valor de z ", z, " e pront.")
	fmt.Printf("%T\n", z)
}
