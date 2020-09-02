[← Pagina Inicial](../README.md#go4noobs)

# Instalando

## Índice

- [Playground](./instalando.md#playgroud)

- [Requisitos](./instalando.md#requisitos)

- [Downloads](./instalando.md#downloads)

## Playground

A [golang.org](https://play.golang.org) criou um serviço web chamado [playground](https://play.golang.org/) que compila, executa e mostra o output do código escrito pelo usuário. 
É um serviço útil para poder testar a linguagem sem ter a necessidade de fazer a instalação.
## Requisitos

As distribuições binárias Go estão disponíveis para esses sistemas operacionais e arquiteturas suportados. Verifique se o seu sistema atende a esses requisitos antes de continuar. Se o seu sistema operacional ou arquitetura não estiver na lista, você poderá instalar a partir da fonte ou usar o gccgo.

| Sistema Operacional               | Arquiteturas                           | Notas                                                                           |
|-----------------------------------|----------------------------------------|---------------------------------------------------------------------------------|
| FreeBSD 10.3 or later             | amd64, 386                             | Debian GNU / kFreeBSA not supported                                             |
| Linux 2.6.23 or later with glibc  | amd64, 386, arm, arm64, s390x, ppc64le | CentOS / RHEL 5.x not supported. Intall from source for other libc.             |
| macOS 10.11 or later              | amd64                                  | use the clang or gcc that comes with Xcode for ```cgo``` support.               |
| Windows 7, Server 2008R2 or later | amd64, 386                             | use MinGW(```386```) or MinGW-W64(```amd64```) gcc. No need for cygwin or msys. |

Um compilador C é necessário apenas se você planeja usar o cgo.

Você só precisa instalar as ferramentas de linha de comando do Xcode. Se você já instalou o Xcode 4.3+, poderá instalá-lo na guia Componentes do painel de preferências de Downloads.

## Downloads

- [Windows](https://golang.org/dl/)
- [Linux](https://golang.org/dl/)
- [macOS](https://golang.org/dl/)
- [Source](https://golang.org/dl/)

Para um passo a passo basta seguir a [documentação](https://golang.org/doc/install).

[← Pagina Inicial](../README.md#go4noobs)
