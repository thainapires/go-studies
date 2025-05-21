package defers

import (
	"database/sql"
	"fmt"
	"os"
)

// Adia uma chamada de uma função até que a função principal retorne
func PrintDefer() {
	x := printDefer()
	fmt.Println(x)

	printDeferMultiple()

	printDefer2() // 50 10
	printDefer3() // 50 50
	printDefer4()

	printDeferEscopos()
}

func printDefer() int {
	defer fmt.Println("World")
	println("Hello")
	return 10
}

func printDeferMultiple() {
	// Last In First Out
	// O último a ser adicionado é o primeiro a ser executado
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)
}

func printDefer2() {
	x := 10
	defer func(y int) {
		fmt.Println(y)
	}(x)

	x = 50
	fmt.Println(x)
}

func printDefer3() {
	x := 10
	defer func() {
		fmt.Println(x)
	}()

	x = 50
	fmt.Println(x)
}

func printDefer4() {
	file, err := os.Open("foo.txt")
	if err != nil {
		panic(err)
	}

	//Escopo do defer sempre responderá ao escopo da func ao redor
	defer func() {
		fmt.Println("Closing file")
		file.Close()
	}()

	if err := customError("First error"); err != nil {
		panic(err)
	}

	if err := customError("Second error"); err != nil {
		panic(err)
	}
}

func customError(message string) error {
	fmt.Println("Error:", message)
	return fmt.Errorf(message)
}

var arquivos = []string{"foo.txt", "bar.txt", "baz.txt"}

func printDeferEscopos() {
	for _, f := range arquivos {
		func() {
			file, err := os.Open(f)
			if err != nil {
			}
			defer file.Close()
		}()
	}
}

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
	}

	defer db.Close() //Cuidado! Será fechado antes de retornar, não no final do código
	return db, nil
}
