[← Pagina Inicial](../README.md#go4noobs)

# Básico

## Índice

- [Pacotes](#pacotes)

- [Importações](#importações)

- [Exportação de nomes](#exportação-de-nomes)

- [Variáveis](#variáveis)

	- [Variável simples](#variável-simples)

	- [Variáveis múltiplas](#variáveis-múltiplas)

	- [Variáveis com inicializadores](#variáveis-com-inicializadores)

	- [Variáveis curtas](#variáveis-curtas)

	- [Tipos básicos de variáveis](#tipos-básicos-de-variáveis)

	- [Variáveis com valores zero](#variáveis-com-valores-zero)

- [Constantes](#constantes)

	- [Constante simples](#constante-simples)

  - [Constante consignada](#constante-consignada)

- [Funções](#funções)

	- [Função sem argumentos](#função-sem-argumentos)

	- [Função com argumentos](#função-com-argumentos)

	- [Função com tipos de argumentos iguais](#função-com-tipos-de-argumentos-iguais)

	- [Função com retorno](#função-com-retorno)

	- [Função com múltiplos retornos](#função-com-múltiplos-retornos)

	- [Função com retorno limpo](#função-com-retorno-limpo)

- [Condicionais](#condicionais)

	- [If](#if)

	- [If com declaração](#if-com-declaração)

	- [Else](#else)

	- [Else If](#else-if)

- [Loopings](#loopings)

	- [For simples](#for-simples)

	- [While (for)](#while-for)

	- [Loop infinito](#loop-infinito)

	- [Foreach (range)](#foreach-range)

## Pacotes

Todo programa em Go é composto por pacotes.

O início de um programa é feito no pacote ```main```.

Input:

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

Output:

```bash
Meu número da sorte é 42
```

> O número 42 é aleatório!

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

As formas de tornar uma variável, constante ou função Global ou não em Go é feita pela forma que seu nome é escrito.

Caso seja escrito com a primeira letra maiúscula poderá ser usada em qualquer outro pacote. Por outro lado, se o nome começa com a inicial minúscula poderá ser usado apenas dentro do pacote.

Input:

```go
package main

import "fmt"

func main(){
	fmt.Println("Forma correta!")

	fmt.println("Forma errada!")
}
```

Output:

```bash
.\main.go:8:2: cannot refer to unexported name fmt.println
.\main.go:8:2: undefined: fmt.println
```

Ao executar, irá gerar um erro durante a compilação na oitava linha devido à mesma estar iniciando com uma letra minúscula.

## Variáveis

### Variável simples

A instrução ```var``` declara uma variável o seu tipo é colocado apos seu nome que pode ou não ser atribuído um valor. 

Input:

```go
package main

import "fmt"

var name string

func main() {
	var age int = 19
	name = "Caio"

	fmt.Println(name, age, "anos")
}
```

Output:

```bash
Caio 19 anos
```

### Variáveis múltiplas

Quando uma variável possui o mesmo tipo que outra, seu tipo pode ser omisso, com exceção da última.

Input:

```go
package main

import "fmt"

var firstName, lastName string

func main() {
	var age, cpf int

	firstName = "Caio Alexandre"
	lastName = "Reis de Almeida"
	age = 19
	cpf = 12345678912

	fmt.Println(firstName, lastName, age, "anos / cpf:", cpf)
}
```

Output:

```bash
Caio Alexandre Reis de Almeida 19 anos / cpf: 12345678912
```

### Variáveis com inicializadores

Quando uma variável possui um inicializador, não é necessário declarar seu tipo, pois ele é auto definido.

Input:

```go
package main

import "fmt"

var name, age, balance = "Caio Alexandre", 19, 1.99

func main() {
	fmt.Println(name, "de", age, "anos possui exatos R$", balance, "em sua conta!")
}
```

Output:

```bash
Caio Alexandre de 19 anos possui exatos R$ 1.99 em sua conta!
```

### Variáveis curtas

Esse tipo de declaração não necessita da instrução ```var``` apenas o uso da expressão ```:=```.

Input:

```go
package main

import "fmt"

func main() {
	name, age, balance := "Caio Alexandre", 19, 1.99

	fmt.Println(name, "de", age, "anos possui exatos R$", balance, "em sua conta!")
}
```

Output:

```bash
Caio Alexandre de 19 anos possui exatos R$ 1.99 em sua conta!
```

> Nota: declarações curtas não funcionam fora de funções, nesse caso é necessário o uso da instrução ```var```

### Tipos básicos de variáveis

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // pseudônimo para uint8

rune // pseudônimo para int32
     // representa um ponto de código Unicode

float32 float64

complex64 complex128
```

### Variáveis com valores zero

Variáveis sem um inicializador são chamadas de variáveis zero.

Input:

```go
package main

import "fmt"

var name string  // Igual a "" string vazia
var age int      // Igual a 0
var marryed bool // Igual a false

func main() {

	fmt.Println("Nome:", name, "Idade:", age, "Casado:", marryed)

}
```

Output:

```bash
Nome:  Idade: 0 Casado: false
```

> ```0```  para tipos numéricos
>
> ```""``` (string vazia) para strings
>
> ```false``` para boleanos

### Variáveis consignadas

Da mesma forma que as importações as variáveis podem ser declaradas de forma consignada.

Input:

```go
package main

import "fmt"

var (
	name string
	age  = 19
)

func main() {
	name = "Caio Alexandre"

	fmt.Println(name, age, "anos")

}
```

Output:

```bash
Caio Alexandre 19 anos
```

## Constantes

### Constante simples

Assim como as variáveis as constantes, precisão de uma instrução, mas com a palavra-chave ```const```.

Input:

```go
package main

import "fmt"

const name = "Caio Alexandre"

func main() {

	fmt.Println(name)

}
```

Output:

```bash
Caio Alexandre
```

### Constante consignada

As constantes também podem ser declaradas de forma consignada.

Input:

```go
package main

import "fmt"

const (
	firstName, lastName = "Caio", "Almeida"
	cpf                 = 12345678912
)

func main() {

	fmt.Println(firstName, lastName, "cpf:", cpf)

}
```

Output:

```bash
Caio Almeida cpf: 12345678912
```

## Funções

### Função sem argumentos

Input:

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

Output:

```bash
Olá!
```

### Função com argumentos

Note que os tipos dos argumentos são colocados após as variáveis.

Input:

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

Output:

```bash
Caio Alexandre Reis de Almeida
```

### Função com tipos de argumentos iguais

Quando variáveis compartilham o mesmo tipo as mesmas podem omitir seus tipos, com exceção da última.

Input:

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

Output:

```bash
Caio Alexandre Reis de Almeida - 19 anos.
```

### Função com retorno

Para fazer uma função com retorno basta adicionar o tipo ao lado direito dos ```()```.

Input:

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

Output:

```bash
5
```

### Função com múltiplos retornos

As funções podem ter inúmeros retornos, quando há mais de um é necessário colocá-los em ```()```

Input:

```go
package main

import "fmt"

func sumAndSubCalc(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	sum, sub := sumAndSubCalc(3, 2)
	fmt.Println("Soma =", sum, "Subtração =", sub)
}
```

```bash
Soma = 5 Subtração = 1
```

Nas funções de múltiplo retorno, se não for necessário usar todos os valores retornados pela função, você pode usar o identificador ```_``` para evitar erros de compilação.

Input:

```go
package main

import "fmt"

func sumAndSubCalc(x, y int) (int,int) {
	return x + y, x - y
}

func main() {
	_, sub := sumAndSubCalc(3,2)
	fmt.Println("Subtração =", sub)
}
```

Output:

```bash
Subtração = 1
```

### Função com retorno limpo


As funções de retorno limpo agem como variáveis, e devem ser nomeadas para documentar o retorno.

Input:

```go
package main

import "fmt"

func invert(second, first int) (x, y int) {
	y = second
	x = first
	return
}

func main() {
	fmt.Println(invert(2, 3))
}
```

Output:

```bash
3 2
```

> Funções de retorno limpo apenas devem ser utilizadas em pequenas funções, seu uso em funções maiores podem dificultar na leitura do código.

## Condicionais

### If

Para fazer uma condicional if é necessário o uso da instrução ```if```.

Input:

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)

	fmt.Println("x =", x)

	if x > 50 {
		fmt.Println("x é maior que 50")
	}

}
```

Output:

Caso "y" seja maior que 50

```bash
x = 61
x é maior que 50
```

Ou

```bash
x = 9
```

> Os valores 61 e 9 são aleatórios!

> Diferente de outras linguagens como C, Java ou javascript, não há ```()``` parênteses rodeando os componentes do ```if``` mas as ```{}``` chaves são obrigatórias.

### If com declaração

A intrução ```if``` pode ter uma pequena declaração antes da condição. 

Input:

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	if y := randNumber(); y > 50 {
		fmt.Println("y =", y)
		fmt.Println("y é maior que 50")
	}
}

func randNumber() (x int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(100)

	return
}
```

Output:

Caso "y" seja maior que 50

```bash
y = 61
y é maior que 50
```

> O valor 61 é aleatório!

> Variáveis declaradas dentro da intrução são válidas somente nesse escopo (até o final do ```if```, ```else if``` ou ```else```) 

### Else

A instrução ```else``` tem que ser chamada logo após o ```if``` ou ```else if```, caso contrário irá causar um erro.

Input:

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)

	fmt.Println("x =", x)

	if x > 50 {
		fmt.Println("x é maior que 50")
	} else {
		fmt.Println("x é menor que 50")
	}

}
```

Output:

Caso "x" seja maior que 50

```bash
x = 91
x é maior que 50
```

Ou

```bash
x = 15
x é menor que 50
```

> Os valores 91 e 15 são aleatórios!

### Else If

Assim como o ```else``` a instrução ```else if``` também necessita iniciar após o fechamento de um ```if``` ou outro ```else if```.

Input:

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)

	fmt.Println("x =", x)

	if x > 50 {
		fmt.Println("x é maior que 50")
	} else if x == 50 {
		fmt.Println("x é igual a 50")
	} else {
		fmt.Println("x é menor que 50")
	}

}
```

Output:

Caso "x" seja maior que 50

```bash
x = 78
x é maior que 50
```

Caso "x" seja igual a 50

```bash
x = 50
x é igual a 50
```

Ou

```bash
x = 26
x é menor que 50
```

> Os valores 78, 50 e 26 são aleatórios!

## Loopings

### For simples

Assim como na maioria das linguagens o looping ```for``` possui 3 componentes:

- Instrução inicial (Executada antes da primeira interação)

- Expressão da condição (Avaliada antes de cada interação)

- Pós-instrução (executada depois de cada interação)

Input:

```go
package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}

}
```

Output:

```bash
01234
```

> Diferente de outras linguagens como C, Java ou javascript, não há ```()``` parênteses rodeando os componentes do ```for``` mas as ```{}``` chaves são obrigatórias.

### While (for)

Em Go a Instrução inicial e a Pós-instrução são opcionais.

Input:

```go
package main

import (
	"fmt"
)

func main() {

	i := 0
	for i < 5 {
		fmt.Print(i)
		i++
	}

	fmt.Println("")

	for i := 0; i <= 10; {
		fmt.Print(i)
		i++
	}

}
```

Output:

```bash
01234       
012345678910
```

> Nota: Em go existe apenas uma estrutura de looping o ```for``` mas elá é bem flexível podendo fazer qualquer um deles.

### Loop infinito

Caso a condição do ```for``` seja omissa, o looping se tornará infinito até que haja um ```break```.

Input:

```go
package main

import (
	"fmt"
)

func main() {

	i := 0
	for {
		fmt.Print(i)

		if i == 10 {
			break
		}

		i++
	}

}
```

Output:

```bash
012345678910
```

### Foreach (range)

Em Go o foreach é feito com a instrução ```range```.

Input:

```go
package main

import "fmt"

func main() {

	people := []string{"Caio", "paochapado", "abidinhs", "Vasco14pt"}

	for i, person := range people {
		fmt.Println("index:", i, "value:", person)
	}

}
```

Output:

```bash
index: 0 value: Caio
index: 1 value: paochapado
index: 2 value: abidinhs
index: 3 value: Vasco14pt
```

> O ```range``` é utilizado para correr sobre um [array](#array-simples) ou [slice](#slice).

[← Pagina Inicial](../README.md#go4noobs)
