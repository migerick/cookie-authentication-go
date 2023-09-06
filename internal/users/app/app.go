package app

import (
	"github.com/migerick/cookie-authentication-go/internal/users/app/command"
	"github.com/migerick/cookie-authentication-go/internal/users/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

// Commands : RegisterCommands registers all commands
type Commands struct {
	Login command.LoginCommandHandler
}

// Queries : RegisterQueries registers all queries
type Queries struct {
	GetUsers query.GetUsersQueryHandler
}
