# Documentação do Projeto Fx

## O que é este projeto?

Este é um projeto de exemplo que demonstra o uso do Uber-go Fx para injeção de dependência em Go. O projeto implementa um simples serviço de usuário com uma arquitetura de camadas (domain, application, infrastructure).

## Requisitos

- Go 1.21 ou superior
- Fx (`go get go.uber.org/fx`)

## Como executar o projeto

1. Clone o repositório ou crie os arquivos:

```bash
go mod tidy
```

2. Execute a aplicação:

```bash
go run .
```

## Por que usar o Fx?

O Fx é um framework de injeção de dependência para Go desenvolvido pela Uber. Algumas vantagens de usar o Fx:

1. **Injeção de dependência em tempo de execução**: O Fx resolve dependências durante a inicialização da aplicação de forma automática e segura.
2. **Ciclo de vida gerenciado**: O Fx gerencia automaticamente o ciclo de vida dos componentes, incluindo inicialização e finalização.
3. **Facilidade de uso**: Com o Fx, você define providers e o framework cuida da resolução de dependências.
4. **Testabilidade**: Facilita a substituição de implementações reais por mocks durante testes.
5. **Graceful shutdown**: Suporte nativo para desligamento gracioso da aplicação.

## Estrutura do projeto

```
go-poc-fx/
├── cmd/
│   └── main.go              # Ponto de entrada da aplicação
├── internal/
│   └── user/
│       ├── domain/
│       │   └── entity.go    # Entidades e interfaces de domínio
│       ├── application/
│       │   └── service.go   # Serviços da aplicação
│       └── infrastructure/
│           └── repository.go # Implementações de repositórios
├── go.mod
└── README.md
```

## Como funciona?

1. O arquivo `main.go` define os "providers" - funções que criam instâncias dos componentes necessários.
2. O Fx automaticamente resolve as dependências entre os componentes com base nos tipos de entrada e saída dos providers.
3. A aplicação é iniciada com `fx.New()` que cria o container de dependências e executa os hooks de inicialização.
4. O framework gerencia o ciclo de vida completo da aplicação, incluindo o shutdown gracioso.

Esta abordagem permite que você mantenha um código limpo, organizado e fácil de testar, ao mesmo tempo que evita a complexidade de gerenciar dependências manualmente. 