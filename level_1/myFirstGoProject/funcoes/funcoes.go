package funcoes

import (
	"fmt"
	"myFirstGoProject/pacote"
)

func digaOi() {
	fmt.Println("Oi")
}

func somar(a int, b int) int {
	return a + b
}

func somar2(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

// Naked return não é uma prática muito recomendada pois deixa o código mais difícil de ser lido, mas podem ser úteis
func dividirNakedReturn(a, b int) (res1 int, rem1 int) {
	res1 = a / b
	rem1 = a % b
	return
}

// Obs: Todo closure é uma função anônima, mas nem toda função anônima é uma closure
func somarHighOrder(a int) func(int) int {
	//Retorna uma função anônima
	return func(b int) int {
		//Essa função func(b int) int é uma closure, está capturando uma variável que não faz parte do escopo dela. Está pegando uma variável que está no escopo acima.
		return a + b
	}
}

// Função que pode receber um número variável de argumentos
// Argumento variático tem que ser sempre o último argumento da função
func somarVariatica(nums ...int) int {
	var out int
	for _, num := range nums {
		out += num
	}
	return out
}

func Print_funcoes() {
	fmt.Println("Hello, World!") //Print do Hello, World!

	pacote.PrintMinha() //Função de pacote

	digaOi() //Função

	fmt.Println(somar(1, 2)) //Função com parâmetros
	fmt.Println(somar2(1, 2))

	a, b := swap(10, 20) //Função com retorno múltiplo
	fmt.Println(a, b)

	res, rem := dividir(10, 3) //Função com parâmetro nomeado
	fmt.Println(res, rem)

	res1, rem1 := dividirNakedReturn(10, 3) //Função com retorno nomeado e que utiliza naked return
	fmt.Println(res1, rem1)

	result := somarHighOrder(2)(1) //Função high order
	fmt.Println(result)

	f := somarHighOrder(2) //Função high order - segunda forma
	x := f(1)
	fmt.Println(x)

	foo := func() { //Função anônima
		fmt.Println("Olá")
	}
	foo()

	fmt.Println(somarVariatica(10, 10, 10)) //Função variática
}
