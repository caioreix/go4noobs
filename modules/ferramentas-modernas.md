[← Pagina Inicial](../README.md#go4noobs)

# Ferramentas modernas

## Índice

- [Módulos com go mod](#módulos-com-go-mod)
- [Executando com go run](#executando-com-go-run)
- [Testando com go test](#testando-com-go-test)
- [Formatando e revisando com go fmt e go vet](#formatando-e-revisando-com-go-fmt-e-go-vet)
- [Lendo documentação com go doc e pkg.go.dev](#lendo-documentação-com-go-doc-e-pkggodev)
- [Ajuda rápida](#ajuda-rápida)

## Módulos com go mod

Hoje, praticamente todo projeto Go começa com um módulo:

```bash
go mod init github.com/seu-usuario/seu-projeto
go mod tidy
```

O arquivo `go.mod` registra o nome do módulo e a versão mínima do Go esperada pelo projeto. Já o `go mod tidy` limpa dependências não utilizadas e adiciona as que estiverem faltando.

## Executando com go run

Durante o estudo, o comando mais comum é:

```bash
go run .
```

## Testando com go test

Go já vem com suíte de testes na biblioteca padrão. Para rodar todos os testes do projeto:

```bash
go test ./...
```

Sempre que você adicionar exemplos, helpers ou pequenos exercícios, vale a pena acompanhá-los com testes.

## Formatando e revisando com go fmt e go vet

Esses dois comandos cobrem uma boa parte da higiene básica de um projeto Go:

```bash
go fmt ./...
go vet ./...
```

- `go fmt` mantém o código no estilo padrão da linguagem.
- `go vet` aponta problemas suspeitos que não são exatamente erros de compilação.

## Lendo documentação com go doc e pkg.go.dev

Para explorar a documentação sem sair do terminal:

```bash
go doc fmt.Println
go help test
```

Na web, a referência principal fica em [pkg.go.dev](https://pkg.go.dev/), especialmente a documentação da biblioteca padrão em [pkg.go.dev/std](https://pkg.go.dev/std).

## Ajuda rápida

Alguns links oficiais úteis para quem está estudando:

- [Download e instalação](https://go.dev/doc/install)
- [Tutorial inicial](https://go.dev/doc/tutorial/getting-started)
- [Effective Go](https://go.dev/doc/effective_go)
- [Especificação da linguagem](https://go.dev/ref/spec)

[← Pagina Inicial](../README.md#go4noobs)
