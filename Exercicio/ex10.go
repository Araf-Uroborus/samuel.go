// O código a seguir está quebrado e não consegue ser executado.
// De forma curta e grossa: Te vira para corrigir.
package main

import (
	"fmt"
)

var pessoa1 string = "Chico"
var pessoa2 string = "Chiquinho"
var pessoa3 string = "Chiquinha"
var pessoa1Idade int = 15
var pessoa2Idade int = 17
var pessoa3Idade int = 16

func main() {
	fmt.Println("Este é o início do nosso programa")
	totalIdade := pessoa1Idade + pessoa2Idade + pessoa3Idade
	fmt.Println("Se a ciência genética estivesse evoluída e pudéssemos fundir corspos...")
	fmt.Println("Será que", pessoa1, "+", pessoa2, "+", pessoa3, "resultaria em Chiconhoquinha?")
	fmt.Println("Qual será o tipo da variável totalIdade?")
	fmt.Printf("%T\n", totalIdade)

	possivelNome(totalIdade)
	fmt.Println("A Chiquinha é maior de idade?")
	maiorDeIdade(totalIdade)
}

func possivelNome(totalIdade int) {
	fmt.Println("Mas calma... Algo está errado...")
	fmt.Println("Porque a fusão dos nomes", pessoa1, "+", pessoa2, "+", pessoa3, "seria Chiconhoquinha?")
	fmt.Printf("Vejamos os nomes novamente:\n")
	fmt.Println(pessoa1, pessoa2, pessoa3)
	fmt.Println("Se pegarmos a primeira sílaba do primeiro nome, a segunda sílaba do segundo nome, e a terceira do terceiro nome...")
	fmt.Println("Teremos o CHI de", pessoa1, "+ o QUI de", pessoa2, "+ o NHA de", pessoa3)
	fmt.Println("Logo o resultado seria CHI-QUI-NHA.")
	fmt.Println("Isso significa que Chico e Chiquinho já não existiam desde o início? Já que a fusão dos 3 resulta em Chiquinha?")
	fmt.Println("Apenas Chiquinha era real o tempo todo?")
	fmt.Println("E se isso é verdade, então a idade de Chiquinha seria a soma dos 3?")
	fmt.Println("Os outros dois podem ser apenas alter-egos da Chiquinha e a sua idade real seria", totalIdade)
}

func maiorDeIdade(totalIdade int) {
	fmt.Println(totalIdade >= 18)
}
