// Cada tipo de número no Go possui um tamanho diferente e ocupa uma quantidade diferente de recursos na máquina.
// Levando em conta a documentação existente aqui (https://golang.org/ref/spec#Numeric_types) crie
// o programa a seguir da forma mais otimizada possível em termos de tipos de variáveis.
//
// 1. Crie um programa que possua 8 variáveis com escopo de pacote, sendo elas:
// var1 = 12
// var2 = 125
// var3 = -15
// var4 = 256
// var5 = 65434
// var6 = -31746
// var7 = 1947483647
// var8 = -7523372036854775808
//
// 2. A sua função main deverá imprimir o valor de cada uma destas variáveis, bem como o tipo de cada uma delas.
//
// 3. Crie uma função chamada matematica. Esta função deverá ser chamada pela função main.
// A função matematica devera imprimir os valores das somas a seguir:
// var1 + var2
// var1 + var6
// var1 - var4
// var3 + var8
// var6 - var2
// var7 + var6