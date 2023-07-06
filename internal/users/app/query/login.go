package query

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

type LoginQueryHandler decorator.QueryHandler[Login, string]

type loginHandler struct {
	userRepo domain.Repository
}

func NewLoginQueryHandler(
	userRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,

) LoginQueryHandler {
	if userRepo == nil {
		panic("nil userRepo")
	}

	return decorator.ApplyQueryDecorators[Login, string](
		loginHandler{userRepo: userRepo},
		logger,
		metricsClient,
	)
}

func (l loginHandler) Handle(ctx context.Context, query Login) (string, error) {
	return l.userRepo.LoginQuery(ctx, query.Email, query.Password)
}
