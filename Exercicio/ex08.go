// Utilizando o código do exemplo anterior:
// 1. no escopo de pacote, utilizando a palavra chave "var",
// crie uma VARIÁVEL com o IDENTIFICADOR "y". A variável
// deverá ser do tipo original ou que dá origem ao SEU TIPO
// criado anteriormente.
//
// 2. na função main:
// 	a. isto já deverá estar feito:
// 		i. imprimir o valor de "x"
// 		ii. imprimir o tipo de "x"
// 		iii. atribuir o VALOR à variável "x" utilizando "="
// 		iv. imprimir o valor de "x"
// 	b. Agora faça:
// 		i. utilize CONVERSÃO para converter o TIPO do valor
// 		de "x" para o tipo original de seu TIPO.
// 		ii. Atribuir este valor a "y".
// 		iii. Imprima o tipo de "y".

package main

import (
	"fmt"
)

type sml int

var x sml
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%T\n", y)
}
