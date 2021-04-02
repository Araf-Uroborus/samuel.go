// 1. Crie um programa que, no escopo de pacote crie as seguintes variáveis:
// - nome (string)
// - idade (int)
// - peso (float64)
// - pesopizza (int)
//
// 	a. para nome, atribua seu nome
// 	b. para idade, atribua sua idade
// 	c. para peso, atribua seu peso
//
// 2. Na função main:
//
// Imprima o seguinte parágrafo da forma que você desejar,
// respeitando as seguintes regras:
//
// 	a. Onde você encontrar <nome>,
// 	deverá ser impresso o valor da variável nome. Onde encontrar <idade>, deverá
// 	ser impresso o valor da variável idade. E assim por diante.
// 	b. Onde você encontrar <tipo_variável_nome>, deverá ser impresso o tipo da variável
// 	nome e assim por diante.
//
//
// Meu nome é <nome>. Estou aprendendo o básico de Go e git.
// Possuo <idade> anos e estou pesando aproximadamente <peso>.
// Utilizei variáveis para lhe passar estas informações pessoais com Go.
// Por exemplo, utilizei a variável nome, com o tipo <tipo_variável_nome>.
// Utilizei outras duas variáveis, para idade e peso, com os tipos <tipo_variável_idade> e <tipo_variável_peso>.
// Gosto muito de pizza, portanto fui em um rodízio de pizzas ontem e acabei engordando 1 kg.
//
// 3. Após imprimir o texto acima, atribua o valor da variável pesopizza de forma que ele seja o seu peso + 1kg.
// Em seguida imprima o valor e o tipo da variável pesopizza

package main

import (
	"fmt"
)

var name string = "Samuel"
var idade int = 20
var peso float64 = 80
var pesopizza int

func main() {
	fmt.Printf(`Meu nome é %q. Estou aprendendo o básico de Go e git.
Possuo %d anos e estou pesando aproximadamente %0.2f.
Utilizei variáveis para lhe passar estas informações pessoais com Go.
por exemplo, utilizei a variável nome, com o tipo %T.
Utilizei outras duas variáveis, para idade e peso, com os tipos %T e %T.
Gosto muito de pizza, portanto fui em um rodízio de pizzas ontem e acabei engordando 1kg.`, name, idade, peso, name, idade, peso)

	pesopizza = 1
	if pesopizza == 1 {
		pesopizza = int(peso) + 1
		fmt.Printf("\n%d ", pesopizza)
		fmt.Printf("%T\n", pesopizza)
	} else {
		fmt.Println("A balança estava quebrada. Não engordei.")
	}
}
