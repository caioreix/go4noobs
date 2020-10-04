[← Pagina Inicial](../README.md#go4noobs)

# Cheat sheet

# Índice

1. [Sintaxe Básica](#sintaxe-básica)
2. [Operadores](#operadores)
    * [Aritméticos](#aritméticos)
    * [Comparação](#comparação)
    * [Lógico](#lógico)
    * [Outros](#outros)
3. [Declarações](#declarações)
4. [Funções](#funções)
    * [Funções como valores e closures](#funções-como-valores-e-closures)
    * [Funções Variadicas](#funções-variadicas)
5. [Tipos integrados](#tipos-integrados)
6. [Conversões de tipo](#conversões-de-tipo)
7. [Pacotes](#pacotes)
8. [Estruturas de controle](#estruturas-de-controle)
    * [If](#if)
    * [Loops](#loops)
    * [Switch](#switch)
9. [Arrays, Slices, Ranges](#arrays-slices-ranges)
    * [Arrays](#arrays)
    * [Slices](#slices)
    * [Operações em Arrays e Slices](#operações-em-arrays-e-slices)
10. [Maps](#maps)
11. [Structs](#structs)
12. [Ponteiros](#ponteiros)
13. [Interfaces](#interfaces)
14. [Incorporação](#incorporação)
15. [Erros](#erros)
16. [Concorrencia](#concorrencia)
    * [Goroutines](#goroutines)
    * [Channels](#channels)
    * [Channel Axioms](#channel-axioms)

## Créditos

Esse módulo foi criado e modificado a partir do repositório [golang-cheat-sheet](https://github.com/a8m/golang-cheat-sheet). A maioria dos códigos de exemplo tirados de [Um tour por Go](https://go-tour-br.appspot.com), que é uma excelente introdução ao Go.
Se você é novo no Go, faça esse tour. Seriamente.

## Go em poucas palavras

* Linguagem imperativa
* Estáticamente tipada
* Tokens de sintaxe semelhantes a C (mas, menos parênteses e sem ponto e vírgula) e a estrutura para Oberon-2
* Compila para código nativo(sem JVM)
* Sem classes, mas estruturas com métodos
* interfaces
* Sem herança de implementação. Apesar disso, há ['type embedding'](http://golang.org/doc/effective%5Fgo.html#embedding).
* Funções são cidadãos de primeira classe
* As funções podem retornar vários valores
* Possue closures
* Ponteiros, mas não possui aritmética de ponteiros
* Built-in primitivas de concorrencia: Goroutines e Channels

# Sintaxe Básica

## Hello World

File `hello.go`:

```go
package main

import "fmt"

func main() {
  fmt.Println("Olá Mundo!")
}
```

`$ go run hello.go`

## Operadores

### Aritméticos

|Operator|Description|
|--------|-----------|
|`+`|adição|
|`-`|subtração|
|`*`|multiplicação|
|`/`|divisão|
|`%`|resto|
|`&`|e(and)|
|`\|`|ou(or)|
|`^`|ou exclusivo(xor)|
|`&^`|e não(and not)|
|`<<`|desvio à esquerda|
|`>>`|desvio à direita|

### Comparação

|Operator|Description|
|--------|-----------|
|`==`|igual|
|`!=`|diferente|
|`<`|menor que|
|`<=`|menor ou igual que|
|`>`|maior que|
|`>=`|maio ou igual que|

### Lógico

|Operador|Descrição|
|--------|-----------|
|`&&`|e(and) lógico|
|`\|\|`|ou(or) lógico|
|`!`|não(not) lógico|

### Outros

|Operador|Descrição|
|--------|-----------|
|`&`|endereço de / criar ponteiro|
|`*`|ponteiro de desreferência|
|`<-`|enviar / receber operador (veja [Channels](#channels) abaixo)|

## Declarações

O tipo vai depois do identificador!

```go
var foo int                 // declaração sem inicialização
var foo int = 42            // declaração com inicialização
var foo, bar int = 42, 1302 // declara e inicia variáveis múltiplas de uma vez
var foo = 42                // tipo omitido, será inferido
foo := 42                   // abreviadamente, apenas em corpos de funções, omite
                            // a palavra-chave var, o tipo está sempre implícito
const constant = "Essa é uma constante"

// iota pode ser usado para incrementar números, começando do 0
const (
  _ = iota
  a
  b
  c = 1 << iota
  d
)
  fmt.Println(a, b) // 1 2 (0 foi pulado)
  fmt.Println(c, d) // 8 16 (2^3, 2^4)
```

## Funções

```go
// uma função simples
func functionName() {}

// função com parâmetros (novamente, os tipos vêm depois dos identificadores)
func functionName(param1 string, param2 int) {}

// vários parâmetros do mesmo tipo
func functionName(param1, param2 int) {}

// retorno do tipo de declaração
func functionName() int {
  return 42
}

// Pode retornar vários valores de uma vez
func returnMulti() (int, string) {
  return 42, "foobar"
}
var x, str = returnMulti()

// Retorne vários resultados nomeados simplesmente usando return
func returnMulti2() (n int, s string) {
  n = 42
  s = "foobar"

  return // n e s serão retornados
}
var x, str = returnMulti2()

```

### Funções Como Valores E Closures

```go
func main() {
  // atribuindo uma função a um nome
  add := func(a, b int) int {
    return a + b
  }
  // use o nome para chamar a função
  fmt.Println(add(3, 4))
}

// Closures, com escopo léxico: As funções podem acessar valores
// que foram declarados no escopo ao definir a função
func scope() func() int{
  outerVar := 2
  foo := func() int { return outerVar}
  return foo
}

func anotherScope() func() int{
  // não compilará porque outerVar e foo não estão definidos neste escopo
  outerVar = 444
  return foo
}


// Closures
func outer() (func() int, int) {
  outerVar := 2
  inner := func() int {
    outerVar += 99 // outerVar do escopo externo é mutada.
    return outerVar
  }
  inner()
  return inner, outerVar // retorno da função inner e variável mutada outerVar 101
}
```

### Funções Variadicas

```go
// Usando ... antes do nome do tipo do último parâmetro,
// você pode indicar que esse parâmetro leva zero ou mais desses parâmetros.
// A função é invocada como qualquer outra função,
// exceto que podemos passar quantos argumentos quisermos.
func adder(args ...int) int {
  total := 0
  for _, v := range args { // Repete os argumentos, seja qual for o número.
    total += v
  }
  return total
}

func main() {
  fmt.Println(adder(1, 2, 3)) // 6
  fmt.Println(adder(9, 9))  // 18

  nums := []int{10, 20, 30}
  fmt.Println(adder(nums...)) // 60
}
```

## Tipos integrados

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // apelido para uint8

rune // apelido para int32 ~= um caractere (ponto de código Unicode) - muito Viking

float32 float64

complex64 complex128
```

## Conversões de tipo

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// sintaxe alternativa
i := 42
f := float64(i)
u := uint(f)
```

## Pacotes

* Declaração de pacote no topo de cada arquivo fonte
* Os executáveis ​​estão no pacote `main`
* Convenção: nome do pacote == sobrenome do caminho de importação (caminho de importação `math/rand` => pacote `rand`)
* Identificador de maiúsculas: exportado (visível de outros pacotes)
* Identificador em minúsculas: privado (não é visível em outros pacotes)

## Estruturas de controle

### If

```go
func main() {
  // Mais básico
  if x > 10 {
    return x
  } else if x == 10 {
    return 10
  } else {
    return -x
  }

  // Você pode colocar uma declaração antes da condição
  if a := b + c; a < 42 {
    return a
  } else {
    return a - 42
  }

  // Afirmação de tipo dentro do if
  var val interface{}
  val = "foo"
  if str, ok := val.(string); ok {
    fmt.Println(str)
  }
}
```

### Loops

```go
// Há apenas `for`, não `while`, não `until`
for i := 1; i < 10; i++ {
}
for ; i < 10;  { // while - loop
}
for i < 10  { // você pode omitir ponto-e-vírgula se houver apenas uma condição
}
for { // você pode omitir a condição == while(true){}
}

// Use break/continue no loop atual
// Use break/continue com etiqueta no laço externo
here:
for i := 0; i < 2; i++ {
  for j := i + 1; j < 3; j++ {
    if i == 0 {
      continue here
    }
    fmt.Println(j)
    if j == 2 {
      break
    }
  }
}

there:
for i := 0; i < 2; i++ {
  for j := i + 1; j < 3; j++ {
    if j == 1 {
      continue
    }
    fmt.Println(j)
    if j == 2 {
      break there
    }
  }
}
```

### Switch

```go
  // switch statement
switch operatingSystem {
  case "darwin":
    fmt.Println("Mac OS Hipster")
    // cases quebrão automaticamente, sem fallthrough por padrão
  case "linux":
    fmt.Println("Linux Geek")
  default:
    // Windows, BSD, ...
    fmt.Println("Other")
}

// assim como no for e no if, você pode ter uma declaração de atribuição antes do switch
switch os := runtime.GOOS; os {
  case "darwin": ...
}

  // você também pode fazer comparações em casos de troca
number := 42
switch {
  case number < 42:
    fmt.Println("Smaller")
  case number == 42:
    fmt.Println("Equal")
  case number > 42:
    fmt.Println("Greater")
}

  // cases podem ser apresentados em listas separadas por vírgulas
var char byte = '?'
switch char {
  case ' ', '?', '&', '=', '#', '+', '%':
    fmt.Println("Should escape")
}

// Um switch de tipo é como uma instrução switch regular, mas os casos
// em um switch de tipo especificam tipos (não valores), e esses tipos são
// comparados com o tipo do valor mantido pelo valor de interface fornecido.
func do(i interface{}) {
  switch v := i.(type) {
    case int:
      fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
      fmt.Printf("%q is %v bytes long\n", v, len(v))
    default:
      fmt.Printf("I don't know about type %T!\n", v)
  }
}

func main() {
  do(21)
  do("hello")
  do(true)
}
```

## Arrays, Slices, Ranges

### Arrays

```go
var a [10]int // declara um array int com 10 de comprimento. O comprimento do array faz parte do tipo!
a[3] = 42     // definine elementos
i := a[3]     // lê elementos

// declara e inicializa
var a = [2]int{1, 2}
a := [2]int{1, 2}   // forma abreviada
a := [...]int{1, 2} // reticências -> Compilador calcula o comprimento do array
```

### Slices

```go
var a []int                // declara um slice - semelhante a um array, mas o comprimento não é especificado
var a = []int {1, 2, 3, 4} // declara e inicializa um slice (apoiado pelo array dado implicitamente)
a := []int{1, 2, 3, 4}     // forma abreviada
chars := []string{0:"a", 2:"c", 1: "b"} // ["a", "b", "c"]

var b = a[lo:hi]    // cria um slice (fatia de um array) do índice lo para hi-1
var b = a[1:4]      // slice do índice 1 ao 3
var b = a[:3]       // falta de índice baixo implica 0
var b = a[3:]       // falta de índice alto implica len(a)
a =  append(a,17,3) // acrescenta itens para o slice a
c := append(a,b...) // concatena os slices a e b

// crie uma fatia com make
a = make([]byte, 5, 5) // segundo argumento comprimento, terceiro capacidade
a = make([]byte, 5)  // capacidade é opcional

// crie uma slice de um array
x := [3]string{"Лайка", "Белка", "Стрелка"}
s := x[:] // um slice referenciando o armazenamento de x
```

### Operações em Arrays e Slices

`len(a)` dá a você o comprimento de um array/slice. É uma função integrada, não um atributo/método de um array.

```go
// loop sobre um array/slice
for i, e := range a {
  // 'i' é o índice, e 'e' o elemento
}

// se você só precisa do elemento:
for _, e := range a {
  // 'e' é o elemento
}

// ...e se você só precisa do índice
for i := range a {
}

// Em Go pre-1.4, você obterá um erro do compilador se não estiver usando 'i' e 'e'.
// Go 1.4 introduziu uma forma livre de variável, para que você possa fazer isso
for range time.Tick(time.Second) {
  // faça isso uma vez por segundo
}
```

## Maps

```go
var m map[string]int
m = make(map[string]int)
m["key"] = 42
fmt.Println(m["key"])

delete(m, "key")

elem, ok := m["key"] // testa se a chave "key" está presente e retorna isso, caso esteja

// map literal
var m = map[string]Vertex{
  "Bell Labs": {40.68433, -74.39967},
  "Google":  {37.42202, -122.08408},
}

// iterar sobre o conteúdo do map
for key, value := range m {
}
```

## Structs

Não há classes, apenas structs. As structs podem ter métodos.

```go
// Uma struct é um tipo. É também uma coleção de campos

// Declaração
type Vertex struct {
  X, Y int
}

// Criando
var v = Vertex{1, 2}
var v = Vertex{X: 1, Y: 2}          // Cria uma struct definindo valores com chaves
var v = []Vertex{{1,2},{5,2},{5,5}} // Inicializa um slice de structs

// Acessando membros
v.X = 4

// Você pode declarar métodos em structs. A struct que você quer declarar o
// método (o tipo de recebimento) fica entre a palavra-chave func e
// o nome do método. A struct é copiado em cada método de chamada(!)
func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Método de chamada
v.Abs()

// Para métodos mutantes, você precisa usar um ponteiro (olhe abaixo) A Struct
// como o tipo.Com isso, o valor da struct não é copiado para a chamada do método.
func (v *Vertex) add(n float64) {
  v.X += n
  v.Y += n
}

// Structs anônimas:
// Mais barato e seguro do que usar `map[string]interface{}`.
point := struct {
  X, Y int
}{1, 2}
```

## Ponteiros

```go
p := Vertex{1, 2}  // p é um Vertex
q := &p            // q é um ponteiro para um Vertex
r := &Vertex{1, 2} // r também é um ponteiro para um Vertex

// O tipo de um ponteiro para um Vertex é *Vertex

var s *Vertex = new(Vertex) // new cria um ponteiro para uma nova instância de struct
```

## Interfaces

```go
// declaração de interface
type Awesomizer interface {
  Awesomize() string
}

// tipos *não* declaram para implementar interfaces
type Foo struct {}

// em vez disso, os tipos satisfazem implicitamente uma interface
// se implementarem todos os métodos necessários
func (foo Foo) Awesomize() string {
  return "Awesome!"
}
```

## Incorporação

Não há subclasses no Go. Em vez disso, há incorporação de interface e estrutura.

```go
// ReadWriter sua implementação devem satisfazer ambos Reader e Writer
type ReadWriter interface {
  Reader
  Writer
}

// Server expõe todos os métodos que o Logger tem
type Server struct {
  Host string
  Port int
  *log.Logger
}

// inicialize o tipo incorporado da maneira usual
server := &Server{"localhost", 80, log.New(...)}

// métodos implementados na estrutura incorporada são passados ​​por
server.Log(...) // chama server.Logger.Log(...)

// o nome do campo do tipo incorporado é o seu nome de tipo (neste caso Logger)
var logger *log.Logger = server.Logger
```

## Erros

Não há tratamento de exceção. Funções que podem produzir um erro apenas declaram um valor de retorno adicional do tipo `Error`.

```go
// Esta é a interface `Error`:
type error interface {
  Error() string
}

// Uma função que pode retornar um erro:
func doStuff() (int, error) {
}

func main() {
  result, err := doStuff()
  if err != nil {
    // lidar com erro
  } else {
    // tudo está bem, use o resultado
  }
}
```

# Concorrencia

## Goroutines

Goroutines são threads leves (gerenciados pelo Go, e não threads de sistema operacional). `go f(a, b)` inicia um novo goroutine que executa `f` (dado `f` é uma função).

```go
// apenas uma função (que pode ser iniciada posteriormente como uma goroutine)
func doStuff(s string) {
}

func main() {
  // usando uma função nomeada em um goroutine
  go doStuff("foobar")

  // usando uma função interna anônima em uma goroutine
  go func (x int) {
    // o corpo da função vai aqui
  }(42)// executa a função apos ser criada
}
```

## Channels

```go
// Blocos de channel sem buffer.
// Lê os blocos quando nenhum valor estiver disponível,
// escreve os blocos até que haja uma leitura.
ch := make(chan int) // crie um channel do tipo int
ch <- 42             // envia um valor para o channel ch.
v := <-ch            // recebe um valor de ch


// Cria um channel com buffer.
// A gravação em channels com buffer não bloqueia a menos que
// o <tamanho do buffer> seja maior que valores não lidos forem gravados.
ch := make(chan int, 100)

close(ch) // fecha o channel (apenas o sender deve fechar)

// lê o channel e testa se ele foi fechado
v, ok := <-ch

// se ok for falso, o channel foi fechado

// Lê o channel até que seja fechado
for i := range ch {
  fmt.Println(i)
}

// seleciona blocos em operações de múltiplos canais, se um desbloquear, o caso correspondente é executado
func doStuff(channelOut, channelIn chan int) {
  select {
    case channelOut <- 42:
      fmt.Println("Poderíamos escrever para channelOut!")
    case x := <- channelIn:
      fmt.Println("Nós poderíamos ler de channelIn")
    case <-time.After(time.Second * 1):
      fmt.Println("tempo esgotado")
  }
}
```

### Channel Axioms

```go
// Um envio para um canal nulo bloqueia ele para sempre
var c chan string
c <- "Hello, World!"
// fatal error: all goroutines are asleep - deadlock!

// Uma recepção de um canal nulo bloqueia para sempre
var c chan string
fmt.Println(<-c)
// fatal error: all goroutines are asleep - deadlock!

// Um envio para um canal fechado entra em pânico
var c = make(chan string, 1)
c <- "Hello, World!"
close(c)
c <- "Hello, Panic!"
// panic: send on closed channel

// Uma recepção de um canal fechado retorna o valor zero imediatamente
var c = make(chan int, 2)
c <- 1
c <- 2
close(c)
for i := 0; i < 3; i++ {
  fmt.Printf("%d ", <-c)
}
// 1 2 0
```

[← Pagina Inicial](../README.md#go4noobs)
