<h3 align="center">
  ⚡ Projeto: Go Essentials
</h3>

## :rocket: Sobre o projeto

Go Essentials é um repositório com exemplos de código Go com os fundamentos básicos de Go.

## :computer: Tecnologias utilizadas

<td><img src="https://img.shields.io/badge/go-%23777BB4.svg?style=for-the-badge&logo=go&logoColor=white" alt="Go"></td>

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

