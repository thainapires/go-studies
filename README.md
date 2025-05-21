<h3 align="center">
  ⚡ Projeto: Go Essentials
</h3>

## :rocket: Sobre o projeto

Go Essentials é um repositório com exemplos de código Go com os fundamentos básicos de Go.

## :computer: Tecnologias utilizadas

<td><img src="https://img.shields.io/badge/go-%23777BB4.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go"></td>

### 0. Instalar Go

```wget https://go.dev/dl/go1.24.3.linux-amd64.tar.gz```

```sudo rm -rf /usr/local/go```

```sudo tar -C /usr/local -xzf go1.24.3.linux-amd64.tar.gz ```

add this to the path:

```
export PATH=$HOME/bin:/usr/local/go/bin:$PATH
export PATH=$HOME/go/bin:$PATH
```
### 1. Comandos

##### Build
```go build main.go```

##### Execução
```go run main.go```

#### Compilação espeficando o nome do arquivo executável
```go build -o "foo2" main.go```

#### Compilação pra windows via linux
```GOOS=windows GOARCH=amd64 go build main.go```

### 2. Primeiros Passos:

#### 2.1. Nomes

Exemplos:

```var Foo string

const Bar string = "bar"

type MeuTipo struct{}

func minhaFuncao() {}
```

##### Nomes privados e públicos

```
func foo() // Nome privado, pois primeira letra é minuscula
func Foo() // Nome público, pois primeira letra é maiúscula
```

#### 2.2. Internal

Apenas pacotes dentro do mesmo pacote podem importar pacotes internal. 

Exemplo b.go:

```
package pacote

import (
	"fmt"
	"myFirstGoProject/pacote/internal/foo"
)

var Bar string = "Hello, Bar!"

func PrintMinha() {
	fmt.Println(foo.Minha)
}
```

#### 2.3. Funções

Função sem parâmetros:

```
func digaOi() {
	fmt.Println("Oi")
}
```

Função com parâmetros:

```
func somar(a int, b int) int {
	return a + b
}
```

Funções com tipos dos parâmetros iguais: podemos omitir uma declaração de tipo:

```
func somar(a, b int) int {
    return a + b
}
```

Função com retorno de múltiplos valores:

```
func swap(a, b int) (int, int) {
	return b, a
}
```

Função com parâmetro nomeado:

```
func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}
```

Função com Naked Return:


```
func dividir(a, b int) (res1 int, rem1 int) {
	res1 = a / b
	rem1 = a % b
	return
}
```

Função High Order 

```
func somar(a int) func (int) int {
	//Retorna uma função anônima
	return func(b int) int {
		//Essa função func(b int) int é uma closure, ela está capturando uma variável que não faz parte do escopo dela. Está pegando uma variável que está no escopo acima.
		return a + b
	}
}
```

Função variática

```

func somar(nums ...int) int {
	var out int
	for _, num := range nums {
		out += num
	}
	return out
}

```

#### 2.4. Variáveis

Variável global:

```
var idade int
var idade2 int = 30 //definir valor
var idade3 = 30     //omitir tipo

```

Variável de escopo:

```
var nome, sobrenome string //Declaração
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

var foo, bar = "foo", 50 //tipos diferentes em mesma linha
fmt.Println(foo, bar)
```

#### 2.5. Tipos e constantes

Mais comuns:

- bool
- int int8 int16 int32 int64 - Inteiros positivos e negativos
- uint uint8 uint16 uint32 uint64 uintptr - Inteiros somente positivos
- byte (Mesmo que um uint8)
- rune (Mesmo que um int32) - Normalmente representa caractere
- float32 float64 - Ponto flutuante
- complex64 complex128 - Complexos
- string

Conversão

```
var x int = 65
f := float64(x)
fmt.Println(f)
```
Conversão rara de acontecer

```
var y int = 10084
s := string(y)
fmt.Println(s)
```

Conversão de int para string

```
s2 := strconv.FormatInt(int64(y), 10)
fmt.Println(s2)
```

Constantes

```
const g int = 10 // Não precisa ser utilizada (pode ser runa, byte, string, bool, numéricos)
//g := 10 não pode encurtar
const h = 10
//takeInt32(g) //Não funciona
takeInt64(h) //vai assumir o tipo necessário pelo contexto
```