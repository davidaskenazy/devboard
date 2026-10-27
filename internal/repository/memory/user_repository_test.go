package memory

import (
	"context"
	"errors"
	"testing"

	"github.com/davidaskenazy/devboard/internal/domain"
)

func TestUserRepository_Create(t *testing.T) {
	repo := NewUserRepository()
	ctx := context.Background()

	user := &domain.User{
		Name:  "Ana",
		Email: "ana@devboard.dev",
	}

	err := repo.Create(ctx, user)

	if err != nil {
		t.Fatalf("Create() devolvio un error inesperado: %v", err)
	}

	if user.ID == "" {
		t.Fatalf("Create() no establecio un ID al usuario")
	}

	if user.ID != "user_1" {
		t.Errorf("Create() establecio un id %q; se esperaba %q", user.ID, "user_1")
	}
}

func TestUserRepository_GetByID_NotFound(t *testing.T) {
	repo := NewUserRepository()
	ctx := context.Background()

	_, err := repo.GetByID(ctx, "user_999")

	if err == nil {
		t.Fatalf("GetByID() debia devolver un error")
	}

	if !errors.Is(err, domain.ErrNotFound) {
		t.Errorf("GetByID() devolvio %v; se esperaba ErrNotFound", err)
	}
}

//Test _GetByID

//Test _GetByEmail
