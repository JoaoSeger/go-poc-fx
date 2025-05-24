package main

import (
	"fmt"

	"go.uber.org/fx"

	"go-poc-fx/internal/user/application"
	"go-poc-fx/internal/user/domain"
	"go-poc-fx/internal/user/infrastructure"
)

// demonstrateUserService demonstra o uso do serviço de usuário
func demonstrateUserService(userService domain.UserService) error {
	fmt.Println("🚀 Aplicação iniciada com sucesso!")
	fmt.Println("\n📝 Demonstrando o serviço de usuário:")

	// Criar usuários
	user1, err := userService.CreateUser("João Silva", "joao@example.com")
	if err != nil {
		return fmt.Errorf("erro ao criar usuário 1: %w", err)
	}
	fmt.Printf("✅ Usuário criado: %+v\n", user1)

	user2, err := userService.CreateUser("Maria Santos", "maria@example.com")
	if err != nil {
		return fmt.Errorf("erro ao criar usuário 2: %w", err)
	}
	fmt.Printf("✅ Usuário criado: %+v\n", user2)

	// Buscar usuário por ID
	foundUser, err := userService.GetUser(1)
	if err != nil {
		return fmt.Errorf("erro ao buscar usuário: %w", err)
	}
	fmt.Printf("🔍 Usuário encontrado: %+v\n", foundUser)

	// Listar todos os usuários
	allUsers, err := userService.GetAllUsers()
	if err != nil {
		return fmt.Errorf("erro ao listar usuários: %w", err)
	}
	fmt.Printf("📋 Todos os usuários (%d):\n", len(allUsers))
	for _, user := range allUsers {
		fmt.Printf("  - %+v\n", user)
	}

	fmt.Println("\n🎉 Demonstração concluída com sucesso!")
	return nil
}

func main() {
	fx.New(
		// Providers - funções que criam instâncias dos componentes
		fx.Provide(
			infrastructure.NewInMemoryUserRepository,
			application.NewUserApplicationService,
		),

		// Invoke - executa a lógica principal da aplicação
		fx.Invoke(demonstrateUserService),
	).Run()
}
