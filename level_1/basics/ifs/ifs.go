package ifs

import (
	"errors"
	"fmt"
	"math"
)

func PrintIfs() {
	printIfVariavelNoEscopoDoIf()
	printIfWithElse()
	printTestError()
}

func printIfVariavelNoEscopoDoIf() {
	//Short statement
	if x := math.Sqrt(4); x < 10 {
		fmt.Println(x)
	}
	//fmt.Println(x) // x não está definido nesse escopo
}

func printIfWithElse() {
	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x < 1 {
		fmt.Println("x é maior que 0")
	} else {
		fmt.Println("x é menor que 0")
	}
}

func printTestError() {
	meuBool := true
	if meuBool {
		fmt.Println("meuBool é true")
	}
	if err := doError(); err != nil {
	}
}

func doError() error {
	return errors.New("error")
}
