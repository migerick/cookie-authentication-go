package query

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/migerick/cookie-authentication-go/internal/common/decorator"
	"github.com/migerick/cookie-authentication-go/internal/users/domain"
)

type User struct {
	Query string
}

type GetUsersQueryHandler decorator.QueryHandler[User, string]

type loginHandler struct {
	userRepo domain.Repository
}

func NewGetUsersQueryHandler(
	userRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,

) GetUsersQueryHandler {
	if userRepo == nil {
		panic("repository is nil")
	}

	return decorator.ApplyQueryDecorators[User, string](
		loginHandler{userRepo: userRepo},
		logger,
		metricsClient,
	)
}

func (h loginHandler) Handle(ctx context.Context, user User) (string, error) {
	return h.userRepo.GetUsers(ctx, user.Query)
}
