package loops

import (
	"fmt"
	"sync"
)

func PrintLoops() {
	loopSimples()
	//loopInfinito()
	loopRange()
	loopInt()
	loopTest()
	loopTest2()
}

func loopSimples() {
	// i := 0  - Init statement (Executa uma vez antes do loop)
	// i < 10 - Condition expression (Avaliado antes de cada iteração)
	// i++  - Post statement (Executa após cada iteração)

	var res int

	for i := 0; i < 10; i++ {
		res++
	}
	fmt.Println(res)
}

func loopInfinito() {
	var res int
	var i int
	for {
		res++
		i++
	}
}

func loopRange() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, elem := range arr {
		fmt.Println(i)  // Index
		fmt.Print(elem) // Element
	}

	for _, elem := range arr {
		fmt.Print(elem) // Element
	}
}

func loopInt() {
	for range 10 {
		fmt.Println("dentro")
	}

	//Só pode iterar em uma variavel
	for i := range 10 {
		fmt.Println(i)
	}
}

func loopTest() {
	arr := [3]int{1, 2, 3}
	for i, elem := range arr { //Mudança boa pois causava muito bug
		fmt.Println(&i, &elem) //a cada iteração, 2 variáveis são criadas
	}
}

func loopTest2() {
	const n = 10
	var wg sync.WaitGroup //grupo de espera
	wg.Add(10)            // Esperar 10 execuções
	for i := 0; i < n; i++ {
		// i := i //Solução antiga: shadow
		// go func(i int) //Solução antiga 2
		// Go routine
		go func() { //Código roda concorrentemente ou paralelamente dependendo da CPU
			defer wg.Done()
			fmt.Println(i)
		}() //(i) //Solução antiga 2
	}
	wg.Wait()
}
