// O código a seguir está quebrado e não consegue ser executado.
// De forma curta e grossa: Te vira para corrigir.
package mian

pessoa1 := "Chico"
pessoa2 := "Chiquinho"
pessoa3 := "Chiquinha"
pessoa1Idade := 15
pessoa2Idade := 17
pessoa3Idade := 16

func main() {
	fmt.println("Este é o início do nosso programa")
	totalIdade := pessoa1Idade + pessoa2Idade + pessoa3Idade
	fmt.println("Se a ciência genética estivesse evoluída e pudéssemos fundir corspos...")
	fmt.println("Será que", pessoa1, "+", pessoa2, "+", pessoa3, "resultaria em Chiconhoquinha?")
	fmt.Println("Qual será o tipo da variável totalIdade?")
	fmt.printf("%T\n", totalIdade)

	possivelNome()
	fmt.Println("A Chiquinha é maior de idade?")
	maiorDeIdade
}

possivelNome() {
	fmt.Println("Mas calma... Algo está errado...")
	fmt.Println("Porque a fusão dos nomes", pessoa1, "+", pessoa2, "+", pessoa3, "seria Chiconhoquinha?")
	fmt.Println(Vejamos os nomes novamente:)
	fmt.Println(pessoa1, pessoa2, pessoa3)
	fmt.Println("Se pegarmos a primeira sílaba do primeiro nome, a segunda sílaba do segundo nome, e a terceira do terceiro nome...")
	fmt.Println("Teremos o CHI de", pessoa1, "+ o QUI de", pessoa2, "+ o NHA de", pessoa3)
	fmt.Println("Logo o resultado seria CHI-QUI-NHA.")
	fmt.Println("Isso significa que Chico e Chiquinho já não existiam desde o início? Já que a fusão dos 3 resulta em Chiquinha?")
	fmt.Println("Apenas Chiquinha era real o tempo todo?")
	fmt.Println("E se isso é verdade, então a idade de Chiquinha seria a soma dos 3?")
	fmt.Println("Os outros dois podem ser apenas alter-egos da Chiquinha e a sua idade real seria", totalIdade)	
}

maiorDeIdade() {
	totalIdade >= 18
}