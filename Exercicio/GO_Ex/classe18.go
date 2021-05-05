package main

import (
	"fmt"
)

func main() {
	s := "Olá Mundo"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	s2 := `Olá,
	mundo`
	fmt.Println(s2)
	fmt.Printf("%T\n", s2)

	s3 := "Ola, Mundo"
	fmt.Println(s3)
	fmt.Printf("%T\n", s3)

	bs := []byte(s3)
	fmt.Println(bs)

	fmt.Println(bs[2])
	fmt.Println(bs[0])
	fmt.Println(string(s3[5]))

}
