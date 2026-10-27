package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/davidaskenazy/devboard/internal/domain"
)

// UserRepository es un adapter que implementa UserRepository
type UserRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

// NewUserRepository constructor
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]*domain.User),
	}
}

// Create metodo para crear usuario en memoria
func (r *UserRepository) Create(_ context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = fmt.Sprintf("user_%d", len(r.users)+1)
	r.users[user.ID] = user
	return nil
}
