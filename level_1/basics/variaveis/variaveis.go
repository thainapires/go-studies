package variaveis

import (
	"fmt"
)

// Variável global, pode não ser utilizada, diferente das de escopo de função
var idade int
var idade2 int = 30 //definir valor
var idade3 = 30     //omitir tipo
//nomeDoMeio := "Henrique" //Isso não pode ser feito

func Print_variaveis() {
	var nome, sobrenome string
	fmt.Println(nome, sobrenome, idade) //zeradas por padrão

	var nome2, sobrenome2 string = "John", "Doe"
	fmt.Println(nome2, sobrenome2)

	//da pra retirar o tipo, e vai ser inferido
	var (
		nome3      = "Jane"
		sobrenome3 = "Doe"
	)

	fmt.Println(nome3, sobrenome3)

	nome4 := "Fernando"
	sobrenome4 := "Pessoa"
	fmt.Println(nome4, sobrenome4)

	//foo string := foo // não pode definir tipo, é sempre inferido

	var foo, bar = "foo", 50 //tipos diferentes em mesma linha
	fmt.Println(foo, bar)
	//foo, bar := "foo", 50
	//var baz, qux int, string //Não pode
}
