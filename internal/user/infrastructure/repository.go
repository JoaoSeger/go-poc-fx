package infrastructure

import (
	"errors"
	"sync"

	"go-poc-fx/internal/user/domain"
)

// InMemoryUserRepository implementa UserRepository usando memória
type InMemoryUserRepository struct {
	users  map[int]*domain.User
	nextID int
	mu     sync.RWMutex
}

// NewInMemoryUserRepository cria uma nova instância do repositório em memória
func NewInMemoryUserRepository() domain.UserRepository {
	return &InMemoryUserRepository{
		users:  make(map[int]*domain.User),
		nextID: 1,
	}
}

// FindByID busca um usuário pelo ID
func (r *InMemoryUserRepository) FindByID(id int) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("usuário não encontrado")
	}

	return user, nil
}

// FindAll retorna todos os usuários
func (r *InMemoryUserRepository) FindAll() ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]*domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

// Create cria um novo usuário
func (r *InMemoryUserRepository) Create(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++

	return nil
}
