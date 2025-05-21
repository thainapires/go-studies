package switches

import (
	"fmt"
	"math"
	"time"
)

//Apenas o caso selecionado roda, os abaixo não rodam
//Não precisa colocar break
//Performance do if ou switch é praticamente a mesma, escolher o que for melhor no seu código

func PrintSwitches() {
	printSwitch()
	printSwitchFallthrough()
	printSwitchSemCondicoes()
	printSwitchQueDeclaraVariavel()
	printSwitchMultiplosCasosEmUmaLinha()
	printSwitchDeUmTipo()
}

func printSwitch() {
	do(1)
	do(2)
	do(3)
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("outra coisa")
	}
}

func printSwitchFallthrough() {
	doFallthrough(1)
}

func doFallthrough(x int) {
	//Força a rodar o próximo caso, mesmo que não seja o caso
	switch x {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough
	default:
		fmt.Println("outra coisa")
	}
}

func printSwitchSemCondicoes() {
	fmt.Println(isWeekend(time.Now()))
}

func isWeekend(x time.Time) bool {
	//Casos não precisam ser constantes
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		return false
	default:
		return true
	}
}

func printSwitchQueDeclaraVariavel() {
	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("resultado é 2")
	default:
		fmt.Println("algo de errado")
	}
}

func printSwitchMultiplosCasosEmUmaLinha() {
	fmt.Println(isWeekend2(time.Now()))
}

func isWeekend2(x time.Time) bool {
	switch x.Weekday() {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}

func printSwitchDeUmTipo() {
	do3("str")
}

func do3(x any) {
	switch t := x.(type) {
	case int:
		fmt.Println("int")
	case string:
		takeString(t)
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("outro tipo")
	}
}

func takeString(s string) {
	fmt.Println(s)
}
