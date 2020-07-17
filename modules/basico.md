[← Pagina Inicial](../README.md#go4noobs)

# Básico

## Índice

- [Pacotes](./basico.md#pacotes)

- [Importações](./basico.md#importações)

- [Exportação de nomes](./basico.md#exportação-de-nomes)

- [Funções](./basico.md#funções)

	- [Função sem argumentos](./basico.md#função-sem-argumentos)

	- [Função com argumentos](./basico.md#função-com-argumentos)

	- [Função com tipos de argumentos iguais](./basico.md#função-com-tipos-de-argumentos-iguais)

	- [Função com retorno](./basico.md#função-com-retorno)

	- [Função com multiplos retornos](./basico.md#função-com-múltiplos-retornos)

	- [Função com retorno limpo](./basico.md#função-com-retorno-limpo)

## Pacotes

Todo programa em Go é composto por pacotes.

O início de um programa é feito no pacote ```main```.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Meu número da sorte é", rand.Intn(100))

}
```

O exemplo a cima está usando os pacotes com caminho de importação "```fmt```", "```math/rand```" e "```time```".

Por convenção, o nome do pacote é o mesmo que o último elemento do caminho de importação. Por exemplo, o pacote "```math/rand```" compreende arquivos que começão com ```package rand```.

## Importações

O grupo de códigos em parênteses da importação, é a declaração de uma importação "consignada":

```go
import (
  "fmt"
  "math/rand"
  "time"
)
```

Você também pode escrever várias importações dessa forma:

```go
import  "fmt"
import  "math/rand"
import  "time"
```

Entretanto, o mais recomendado é utilizar o estilo de importação consignada.

## Exportação de nomes

As formas de tornar uma variável Global ou não em Go é feita pela forma que seu nome é escrito.

Caso ele seja escrito com a primeira letra maiúscula poderá ser usada em qualquer outro pacote. Por outro lado, os que começam com a inicial minúscula poderá ser usado apenas dentro do pacote.

```go
package main

import "fmt"

func main(){
	fmt.Println("Forma correta!")

	fmt.println("Forma errada!")
}
```

Ao executar, irá gerar um erro durante a compilação na oitava linha devido à mesma estar iniciando com uma letra minúscula.

## Funções

### Função sem argumentos

```go
package main

import "fmt"

func sayHello(){
	fmt.Println("Olá!")
}

func main(){
	sayHello()
}
```

### Função com argumentos

Note que os tipos dos argumentos são colocados após as variáveis.

```go
package main

import "fmt"

func fullName(firstName string, lastName string) {
	fmt.Println(firstName, lastName)
}

func main() {
	fullName("Caio Alexandre", "Reis de Almeida")
}
```

### Função com tipos de argumentos iguais

Quando variáveis compartilham o mesmo tipo as mesmas podem omitir seus tipos, com exceção da última.

```go
package main

import "fmt"

func fullName(firstName, lastName string, age int) {
	fmt.Println(firstName, lastName, "-", age, "anos.")
}

func main() {
	fullName("Caio Alexandre", "Reis de Almeida", 19)
}
```

### Função com retorno

Para fazer uma função com retorno basta adicionar o tipo ao lado direito dos ```()```.

```go
package main

import "fmt"

func sumCalc(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(sumCalc(3, 2))
}
```

### Função com múltiplos retornos

As funções podem ter inúmeros retornos, quando há mais de um é necessário colocá-los em ```()```

```go
package main

import "fmt"

func sumAndSubCalc(x, y int) (int,int) {
	return x + y,x-y
}

func main() {
	sum, sub := sumAndSubCalc(3,2)
	fmt.Println(sum,sub)
}
```

### Função com retorno limpo


As funções de retorno limpo agem como variáveis, e devem ser nomeadas para documentar o retorno.

```go
package main

import "fmt"

func split(value int) (x, y int) {
	x = value * 4 / 9
	y = value - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

> Funções de retorno limpo apenas devem ser utilizadas em pequenas funções, seu uso em funções maiores podem dificultar na leitura do código.
