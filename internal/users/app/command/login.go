package command

import (
	"context"
	"github.com/migerick/cookie-authentication-go/internal/common/decorator"
	"github.com/migerick/cookie-authentication-go/internal/users/domain"
	"github.com/sirupsen/logrus"
)

type Login struct {
	Email    string
	Password string
}
type LoginCommandHandler decorator.CommandHandler[Login]

type loginHandler struct {
	authRepo domain.Repository
}

func NewLoginCommandHandler(
	authRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) LoginCommandHandler {
	if authRepo == nil {
		panic("hourRepo is nil")
	}

	return decorator.ApplyCommandDecorators[Login](
		loginHandler{authRepo: authRepo},
		logger,
		metricsClient,
	)
}

func (c loginHandler) Handle(ctx context.Context, cmd Login) error {
	return c.authRepo.Login(ctx, cmd.Email, cmd.Password)
}
