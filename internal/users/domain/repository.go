package domain

import "context"

type Repository interface {
	Login(context.Context, string, string) error
	LoginQuery(context.Context, string, string) (string, error)
}
