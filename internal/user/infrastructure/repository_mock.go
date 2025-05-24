package infrastructure

import (
	"fmt"

	"go-poc-fx/internal/user/domain"
)

// MockUserRepository é um mock do UserRepository para testes
type MockUserRepository struct {
	users        map[int]*domain.User
	nextID       int
	shouldError  bool
	errorMsg     string
	findByIDFunc func(id int) (*domain.User, error)
	findAllFunc  func() ([]*domain.User, error)
	createFunc   func(user *domain.User) error
}

// NewMockUserRepository cria uma nova instância do mock repository
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users:  make(map[int]*domain.User),
		nextID: 1,
	}
}

// WithError configura o mock para retornar erro
func (m *MockUserRepository) WithError(msg string) *MockUserRepository {
	m.shouldError = true
	m.errorMsg = msg
	return m
}

// WithoutError remove a configuração de erro
func (m *MockUserRepository) WithoutError() *MockUserRepository {
	m.shouldError = false
	m.errorMsg = ""
	return m
}

// WithFindByIDFunc permite configurar um comportamento customizado para FindByID
func (m *MockUserRepository) WithFindByIDFunc(fn func(id int) (*domain.User, error)) *MockUserRepository {
	m.findByIDFunc = fn
	return m
}

// WithFindAllFunc permite configurar um comportamento customizado para FindAll
func (m *MockUserRepository) WithFindAllFunc(fn func() ([]*domain.User, error)) *MockUserRepository {
	m.findAllFunc = fn
	return m
}

// WithCreateFunc permite configurar um comportamento customizado para Create
func (m *MockUserRepository) WithCreateFunc(fn func(user *domain.User) error) *MockUserRepository {
	m.createFunc = fn
	return m
}

// AddUser adiciona um usuário pré-configurado no mock (útil para setup de testes)
func (m *MockUserRepository) AddUser(user *domain.User) *MockUserRepository {
	if user.ID == 0 {
		user.ID = m.nextID
		m.nextID++
	}
	m.users[user.ID] = user
	return m
}

// Reset limpa todos os dados do mock
func (m *MockUserRepository) Reset() {
	m.users = make(map[int]*domain.User)
	m.nextID = 1
	m.shouldError = false
	m.errorMsg = ""
	m.findByIDFunc = nil
	m.findAllFunc = nil
	m.createFunc = nil
}

// FindByID implementa domain.UserRepository
func (m *MockUserRepository) FindByID(id int) (*domain.User, error) {
	if m.findByIDFunc != nil {
		return m.findByIDFunc(id)
	}

	if m.shouldError {
		return nil, fmt.Errorf(m.errorMsg)
	}

	user, exists := m.users[id]
	if !exists {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	return user, nil
}

// FindAll implementa domain.UserRepository
func (m *MockUserRepository) FindAll() ([]*domain.User, error) {
	if m.findAllFunc != nil {
		return m.findAllFunc()
	}

	if m.shouldError {
		return nil, fmt.Errorf(m.errorMsg)
	}

	users := make([]*domain.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}

	return users, nil
}

// Create implementa domain.UserRepository
func (m *MockUserRepository) Create(user *domain.User) error {
	if m.createFunc != nil {
		return m.createFunc(user)
	}

	if m.shouldError {
		return fmt.Errorf(m.errorMsg)
	}

	user.ID = m.nextID
	m.users[user.ID] = user
	m.nextID++

	return nil
}

// GetUserCount retorna o número de usuários no mock (método auxiliar para testes)
func (m *MockUserRepository) GetUserCount() int {
	return len(m.users)
}

// GetUser retorna um usuário específico pelo ID (método auxiliar para testes)
func (m *MockUserRepository) GetUser(id int) *domain.User {
	return m.users[id]
}
