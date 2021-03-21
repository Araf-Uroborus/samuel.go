// Altere o código a seguir e inclua o seguinte:
// - Crie uma nova variável chamada nota1
// - Crie uma variável chamada nota2
// - Crie uma variável chamada nota3
// - Crie uma variável chamada media:
// -- A variável media deverá receber o valor da media entre nota1, nota2 e nota3
// - Imprima o resultado da media

package main

import (
	"fmt"
)

var nota1 float64
var nota2 float64
var nota3 float64
var notas [3]float64

func main() {
	notas[0] = nota1
	notas[1] = nota2
	notas[2] = nota3
	soma := nota1 + nota2 + nota3
	media := soma / float64(len(notas))
	fmt.Println("A media é:", media)
}
