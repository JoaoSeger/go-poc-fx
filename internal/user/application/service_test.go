package application

import (
	"testing"

	"go-poc-fx/internal/user/domain"
	"go-poc-fx/internal/user/infrastructure"
)

func TestUserApplicationService_CreateUser(t *testing.T) {
	// Arrange
	mockRepo := infrastructure.NewMockUserRepository()
	service := NewUserApplicationService(mockRepo)

	// Act
	user, err := service.CreateUser("João Silva", "joao@example.com")

	// Assert
	if err != nil {
		t.Fatalf("CreateUser() retornou erro inesperado: %v", err)
	}

	if user == nil {
		t.Fatal("CreateUser() retornou usuário nil")
	}

	if user.Name != "João Silva" {
		t.Errorf("CreateUser() nome esperado 'João Silva', obtido '%s'", user.Name)
	}

	if user.Email != "joao@example.com" {
		t.Errorf("CreateUser() email esperado 'joao@example.com', obtido '%s'", user.Email)
	}

	if user.ID == 0 {
		t.Error("CreateUser() ID não foi atribuído")
	}

	// Verificar se o usuário foi salvo no repositório
	if mockRepo.GetUserCount() != 1 {
		t.Errorf("CreateUser() esperado 1 usuário no repositório, obtido %d", mockRepo.GetUserCount())
	}
}

func TestUserApplicationService_CreateUser_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := infrastructure.NewMockUserRepository().
		WithError("erro de conexão com banco")
	service := NewUserApplicationService(mockRepo)

	// Act
	user, err := service.CreateUser("João Silva", "joao@example.com")

	// Assert
	if err == nil {
		t.Fatal("CreateUser() deveria ter retornado erro")
	}

	if user != nil {
		t.Fatal("CreateUser() deveria ter retornado usuário nil em caso de erro")
	}

	if err.Error() != "erro de conexão com banco" {
		t.Errorf("CreateUser() erro esperado 'erro de conexão com banco', obtido '%s'", err.Error())
	}
}

func TestUserApplicationService_GetUser(t *testing.T) {
	// Arrange
	mockRepo := infrastructure.NewMockUserRepository()
	existingUser := &domain.User{
		ID:    1,
		Name:  "Maria Santos",
		Email: "maria@example.com",
	}
	mockRepo.AddUser(existingUser)
	service := NewUserApplicationService(mockRepo)

	// Act
	user, err := service.GetUser(1)

	// Assert
	if err != nil {
		t.Fatalf("GetUser() retornou erro inesperado: %v", err)
	}

	if user == nil {
		t.Fatal("GetUser() retornou usuário nil")
	}

	if user.ID != 1 {
		t.Errorf("GetUser() ID esperado 1, obtido %d", user.ID)
	}

	if user.Name != "Maria Santos" {
		t.Errorf("GetUser() nome esperado 'Maria Santos', obtido '%s'", user.Name)
	}
}

func TestUserApplicationService_GetUser_NotFound(t *testing.T) {
	// Arrange
	mockRepo := infrastructure.NewMockUserRepository()
	service := NewUserApplicationService(mockRepo)

	// Act
	user, err := service.GetUser(999)

	// Assert
	if err == nil {
		t.Fatal("GetUser() deveria ter retornado erro para usuário inexistente")
	}

	if user != nil {
		t.Fatal("GetUser() deveria ter retornado usuário nil para usuário inexistente")
	}
}

func TestUserApplicationService_GetAllUsers(t *testing.T) {
	// Arrange
	mockRepo := infrastructure.NewMockUserRepository()

	// Adicionar usuários de teste
	user1 := &domain.User{Name: "João Silva", Email: "joao@example.com"}
	user2 := &domain.User{Name: "Maria Santos", Email: "maria@example.com"}

	mockRepo.AddUser(user1).AddUser(user2)
	service := NewUserApplicationService(mockRepo)

	// Act
	users, err := service.GetAllUsers()

	// Assert
	if err != nil {
		t.Fatalf("GetAllUsers() retornou erro inesperado: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("GetAllUsers() esperado 2 usuários, obtido %d", len(users))
	}
}

func TestUserApplicationService_GetAllUsers_Empty(t *testing.T) {
	// Arrange
	mockRepo := infrastructure.NewMockUserRepository()
	service := NewUserApplicationService(mockRepo)

	// Act
	users, err := service.GetAllUsers()

	// Assert
	if err != nil {
		t.Fatalf("GetAllUsers() retornou erro inesperado: %v", err)
	}

	if len(users) != 0 {
		t.Errorf("GetAllUsers() esperado 0 usuários, obtido %d", len(users))
	}
}

func TestUserApplicationService_GetAllUsers_RepositoryError(t *testing.T) {
	// Arrange
	mockRepo := infrastructure.NewMockUserRepository().
		WithError("erro de conexão")
	service := NewUserApplicationService(mockRepo)

	// Act
	users, err := service.GetAllUsers()

	// Assert
	if err == nil {
		t.Fatal("GetAllUsers() deveria ter retornado erro")
	}

	if users != nil {
		t.Fatal("GetAllUsers() deveria ter retornado slice nil em caso de erro")
	}
}

// Exemplo de teste usando função customizada no mock
func TestUserApplicationService_CustomMockBehavior(t *testing.T) {
	// Arrange
	mockRepo := infrastructure.NewMockUserRepository()

	// Configurar comportamento customizado: sempre retornar o mesmo usuário
	expectedUser := &domain.User{ID: 42, Name: "Usuário Fixo", Email: "fixo@example.com"}
	mockRepo.WithFindByIDFunc(func(id int) (*domain.User, error) {
		return expectedUser, nil
	})

	service := NewUserApplicationService(mockRepo)

	// Act
	user, err := service.GetUser(999) // Qualquer ID

	// Assert
	if err != nil {
		t.Fatalf("GetUser() retornou erro inesperado: %v", err)
	}

	if user.ID != 42 {
		t.Errorf("GetUser() ID esperado 42, obtido %d", user.ID)
	}

	if user.Name != "Usuário Fixo" {
		t.Errorf("GetUser() nome esperado 'Usuário Fixo', obtido '%s'", user.Name)
	}
}
