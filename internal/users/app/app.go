package app

import (
	"github.com/migerick/cookie-authentication-go/internal/users/app/command"
	"github.com/migerick/cookie-authentication-go/internal/users/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	Login command.LoginCommandHandler
}

type Queries struct {
	GetUsers query.GetUsersQueryHandler
}
