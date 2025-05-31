package main

import (
	"go-poc-fx/internal/server"
	"go-poc-fx/internal/user/application"
	"go-poc-fx/internal/user/infrastructure"

	"go.uber.org/fx"
)

// providePort provides the server port as a dependency
func providePort() string {
	return "8080"
}

func main() {
	fx.New(
		// Providers - funções que criam instâncias dos componentes
		fx.Provide(
			// User domain dependencies
			infrastructure.NewInMemoryUserRepository,
			application.NewUserApplicationService,

			// Server dependencies
			providePort,
			server.New,
		),

		// Invoke - executa a lógica principal da aplicação
		fx.Invoke(server.StartServer),
	).Run()
}
