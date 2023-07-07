package domain

import "context"

type Repository interface {
	Login(context.Context, string, string) error
	GetUsers(context.Context, string) (string, error)
}
