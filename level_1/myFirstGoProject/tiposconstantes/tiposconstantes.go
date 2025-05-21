package tiposconstantes

import (
	"fmt"
	"strconv"
)

// Únicos tipos que a documentação de go faz garantia do tamanho deles em memória (tirando int e uuint)
// bool
// int int8 int16 int32 int64 - Inteiros positivos e negativos
// uint uint8 uint16 uint32 uint64 uintptr - Inteiros somente positivos
// byte (Mesmo que um uint8)
// rune (Mesmo que um int32)
// float32 float64 - Ponto flutuante
// complex64 complex128 - Complexos
// string
func PrintTiposConstantes() {
	var b uint8 = 10
	takeByte(b)

	var r int32 = 3
	takeRune(r)

	//Conversão
	var x int = 65
	f := float64(x)
	fmt.Println(f)

	//Conversão rara de acontecer
	var y int = 10084
	s := string(y)
	fmt.Println(s)

	s2 := strconv.FormatInt(int64(y), 10)
	fmt.Println(s2)
}

func takeByte(b byte) {
	fmt.Println(b)
}

func takeRune(r rune) {
	fmt.Println(r)
}
