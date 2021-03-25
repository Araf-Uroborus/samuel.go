// 1. Utilize var para DECLARAR três variáveis. As variáveis
// devem possuir escopo de pacote. Não atribua VALORES a essas
// variáveis. Utilize os seguintes IDENTIFICADORES para as
// variáveis e certifique-se de que as variáveis serão dos tipos:
// a. identificador "x" do tipo int
// b. identificador "y" do tipo string
// c. identificador "z" do tipo bool
//
// 2. Na função main:
// a. Imprima os valores de cada variável
// b. O compilador atribuiu um valor para cada variável. Como esse
// valor se chama?

package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
