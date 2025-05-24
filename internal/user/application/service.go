package application

import (
	"go-poc-fx/internal/user/domain"
)

// UserApplicationService implementa domain.UserService
type UserApplicationService struct {
	repository domain.UserRepository
}

// NewUserApplicationService cria uma nova instância do serviço de usuário
func NewUserApplicationService(repository domain.UserRepository) domain.UserService {
	return &UserApplicationService{
		repository: repository,
	}
}

// GetUser retorna um usuário pelo ID
func (s *UserApplicationService) GetUser(id int) (*domain.User, error) {
	return s.repository.FindByID(id)
}

// GetAllUsers retorna todos os usuários
func (s *UserApplicationService) GetAllUsers() ([]*domain.User, error) {
	return s.repository.FindAll()
}

// CreateUser cria um novo usuário
func (s *UserApplicationService) CreateUser(name, email string) (*domain.User, error) {
	user := &domain.User{
		Name:  name,
		Email: email,
	}

	err := s.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
