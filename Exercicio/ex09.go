// Você possui uma turma com 4 alunos cujos nomes são Pedro, Ricardo, Paula e Marcia.
// Cada um fez 3 exames e tiveram notas diversas, porém a média final de cada deverá ser igual ou maior a 7.0 para que sejam aprovados.
// As notas de cada um deles são as seguintes:
// - Pedro: 4.5, 7.9 e 8.1
// - Ricardo: 5.7, 6.0 e 9.1
// - Paula: 6.8, 7.6 e 8.7
// - Marcia: 7.5, 5.9 e 7.6

// a. Crie uma variável que receba o nome de cada aluno. Estas variáveis devem possuir escopo geral.

// b. Para cada nota de cada aluno, crie uma variável. Estas variáveis de nota precisam ser com escopo de função.

// c. Em sua função main imprima:
// c.1- "Os alunos da turma são:"
// c.2- Na linha seguinte imprima o valor das variáveis com os nomes dos 4 alunos.
// c.3- Chame uma nova função chamada media

// d. a função media deverá imprimir:
// d.1- A média final de cada aluno indicando o nome do mesmo;
// d.2- Um resultado booleano que represente "true" para cada aluno que foi aprovado (cuja média seja igual ou maior a 7.0)

// e. Crie uma nova função chamada diplomas
// e.1- A função main deverá chamar também a função diplomas.
// e.2- A função diplomas irá imprimir a seguinte linha para cada alunoq ue foi aprovado:
// "Parabéns pela sua aprovação. Seu diploma estará disponível na próxima semana <NOME>."

// f. Crie uma função chamada reprovados.
// f.1- A função main também deverá chamar a função reprovados
// f.2- A função reprovados deverá imprimir o seguinte para cada aluno que foi reprovado:
// "Não sabe, não sabe, vai ter que aprender.. orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima <ALUNO>."

package main

import (
	"fmt"
)

var alunos [4]string

func main() {
	alunos[0] = "Pedro"
	alunos[1] = "Ricardo"
	alunos[2] = "Paula"
	alunos[3] = "Marcia"
	var exames1 [4]float64
	var exames2 [4]float64
	var exames3 [4]float64

	exames1[0] = 4.5
	exames1[1] = 5.7
	exames1[2] = 6.8
	exames1[3] = 7.5

	exames2[0] = 7.9
	exames2[1] = 6.0
	exames2[2] = 7.6
	exames2[3] = 5.9

	exames3[0] = 8.1
	exames3[1] = 9.1
	exames3[2] = 8.7
	exames3[3] = 7.6

	fmt.Printf("\nOs alinos da turma são:")
	for i := 0; i < len(alunos); i++ {
		fmt.Printf("\n%q", alunos[i])
	}
	fmt.Println("\n")
	for z := 0; z < len(alunos); z++ {
		if media(alunos[z], exames1[z], exames2[z], exames3[z]) >= float64(7.0) {
			diplomas(alunos[z])
		} else {
			reprovados(alunos[z])
		}
	}
}

func media(aluno string, nota1, nota2, nota3 float64) float64 {
	var notas [3]float64
	notas[0] = nota1
	notas[1] = nota2
	notas[2] = nota3
	media := (notas[0] + notas[1] + notas[2]) / float64(len(notas))
	var j bool = true
	fmt.Printf("\nAluno: %q, Media: %f\n", aluno, media)
	if float64(media) >= float64(7.0) {
		fmt.Printf("%t\n", j)
	}
	return media
}

func diplomas(aluno string) {
	fmt.Printf("Parabés pela sua aprovação. Seu diploma estará disponível na próxima semana %q.\n", aluno)
}

func reprovados(aluno string) {
	fmt.Printf("Não sabe, não sabe, vai ter que aprender... orelha, de burro, cabeça de et. Brincadeirinha, mas melhor sorte na próxima %q.\n", aluno)

}
