package command

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/migerick/cookie-authentication-go/internal/common/decorator"
	"github.com/migerick/cookie-authentication-go/internal/users/domain"
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
		panic("repository is nil")
	}

	return decorator.ApplyCommandDecorators[Login](
		loginHandler{authRepo: authRepo},
		logger,
		metricsClient,
	)
}

func (h loginHandler) Handle(ctx context.Context, cmd Login) error {
	return h.authRepo.Login(ctx, cmd.Email, cmd.Password)
}
