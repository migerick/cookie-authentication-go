package adapters

import (
	"context"
	"fmt"
	"log"

	"github.com/migerick/cookie-authentication-go/internal/users/domain"
)

type AuthRepository struct{}

func (a AuthRepository) Login(ctx context.Context, email string, password string) error {
	log.Printf("success login: %s %s\n", email, password)
	return nil
}

func (a AuthRepository) GetUsers(_ context.Context, search string) (string, error) {
	//TODO: implement connection to database
	fmt.Printf("Search: %s\n", search)
	return "User list", nil
}

var _ domain.Repository = (*AuthRepository)(nil)

func NewAuthRepository() AuthRepository {
	return AuthRepository{}
}
