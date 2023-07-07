package service

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/migerick/cookie-authentication-go/internal/common/metrics"
	"github.com/migerick/cookie-authentication-go/internal/users/app/command"

	"github.com/migerick/cookie-authentication-go/internal/users/adapters"
	"github.com/migerick/cookie-authentication-go/internal/users/app"
	"github.com/migerick/cookie-authentication-go/internal/users/app/query"
)

func NewApplication(ctx context.Context) app.Application {
	//	TODO: Connection to database
	fmt.Printf("context %s", ctx)

	userRepository := adapters.NewAuthRepository()

	logger := logrus.NewEntry(logrus.StandardLogger())

	metricsClient := metrics.NoOp{}

	return app.Application{
		Commands: app.Commands{
			Login: command.NewLoginCommandHandler(&userRepository, logger, metricsClient),
		},
		Queries: app.Queries{
			GetUsers: query.NewGetUsersQueryHandler(&userRepository, logger, metricsClient),
		},
	}
}
