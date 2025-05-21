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
// rune (Mesmo que um int32) - Normalmente representa caractere
// float32 float64 - Ponto flutuante
// complex64 complex128 - Complexos
// string
func PrintTiposConstantes() {
	var xx bool
	fmt.Println(xx)

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

	//CONSTANTES

	const g int = 10 // Não precisa ser utilizada (caractere - runa byte, string, bool, numéricos)
	//g := 10 não pode encurtar
	const h = 10
	//takeInt32(g)
	takeInt32(h)
	takeInt64(h) //vai assumir o tipo necessário pelo contexto

	takeInt32(10)     // 10 é um literal
	takeString("foo") // foo é um literal

}

func takeByte(b byte) {
	fmt.Println(b)
}

func takeRune(r rune) {
	fmt.Println(r)
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}

func takeString(s string) {
	fmt.Println(s)
}
