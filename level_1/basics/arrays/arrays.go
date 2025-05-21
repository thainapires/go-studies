package arrays

import (
	"fmt"
)

func PrintArrays() {
	arr := [3]int{} //Array com os três valores zerados
	fmt.Println(arr)
	arr2 := [3]int{1, 2, 3} //Array com os três valores preenchidos
	fmt.Println(arr2)
	arr3 := [3]int{1, 2} //Array com os dois primeiros valores preenchidos
	fmt.Println(arr3)
	//arr4 := [3]int{1,2,3,4} dessa forma não pode

	arr4 := [10]int{5: 400, 7: 300} //Especificar o índice
	fmt.Println(arr4)

	arr5 := [10]string{4: "Hello, World!"} // Outros tipos tbm são válidos
	fmt.Println(arr5)

	const x = 10
	arr6 := [x]string{4: "Hello, World!"} // Só aceita x se for uma constante
	fmt.Println(arr6)

	//Nunca pode mudar o tamanho do array
	//Porém, tem arrays dinâmicos
}
