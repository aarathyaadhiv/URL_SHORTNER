//go:build wireinject
// +build wireinject

package dependency

import (
	"github.com/google/wire"
	"github.com/url-shortner/pkg/api"
	"github.com/url-shortner/pkg/api/handler"
	"github.com/url-shortner/pkg/config"
	"github.com/url-shortner/pkg/db"
	"github.com/url-shortner/pkg/repository"
	"github.com/url-shortner/pkg/usecase"
)

func InitializeAPI(c config.Config) (*api.ServerHTTP, error) {
	wire.Build(
		db.ConnectDB,
		repository.NewRepository,
		usecase.NewUseCase,
		handler.NewHandler,
		api.NewServerHTTP,
	)
	return &api.ServerHTTP{}, nil
}
