package main

import (
	"fmt"
)

var y int = 43

var z int
var w string
var a bool
var b float64

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)

	foo()
	fmt.Println("Encerrei foo...")
}

func foo() {
	fmt.Println("Novamente", y)
	fmt.Println("O valor de z é: ", z)
	fmt.Println("O valor de w é: ", w)
	fmt.Println("O valor de a é: ", a)
	fmt.Println("O valor de b é :", b)
}
