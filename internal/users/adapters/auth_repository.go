package adapters

import (
	"context"
	"fmt"

	"github.com/migerick/cookie-authentication-go/internal/users/domain"
)

type AuthRepository struct {
}

func (a AuthRepository) Login(ctx context.Context, email string, password string) error {
	fmt.Printf("Login and Password: %s, %s\n", email, password)
	return nil
}

func (a AuthRepository) LoginQuery(ctx context.Context, email string, password string) (string, error) {
	fmt.Printf("Login and Password: %s, %s\n", email, password)
	return "Login success", nil
}

var _ domain.Repository = (*AuthRepository)(nil)

func NewAuthRepository() AuthRepository {
	return AuthRepository{}
}
