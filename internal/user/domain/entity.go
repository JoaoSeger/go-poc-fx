package domain

// User representa a entidade de usuário
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserRepository define o contrato para operações de usuário
type UserRepository interface {
	FindByID(id int) (*User, error)
	FindAll() ([]*User, error)
	Create(user *User) error
}

// UserService define o contrato para serviços de usuário
type UserService interface {
	GetUser(id int) (*User, error)
	GetAllUsers() ([]*User, error)
	CreateUser(name, email string) (*User, error)
}
