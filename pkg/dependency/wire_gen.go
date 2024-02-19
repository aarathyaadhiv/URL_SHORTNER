// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dependency

import (
	"github.com/url-shortner/pkg/api"
	"github.com/url-shortner/pkg/api/handler"
	"github.com/url-shortner/pkg/config"
	"github.com/url-shortner/pkg/db"
	"github.com/url-shortner/pkg/repository"
	"github.com/url-shortner/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(c config.Config) (*api.ServerHTTP, error) {
	gormDB, err := db.ConnectDB(c)
	if err != nil {
		return nil, err
	}
	interfaceRepoRepository := repository.NewRepository(gormDB)
	interfaceUsecaseUsecase := usecase.NewUseCase(interfaceRepoRepository, c)
	handlerInterfaceHandler := handler.NewHandler(interfaceUsecaseUsecase, c)
	serverHTTP := api.NewServerHTTP(handlerInterfaceHandler)
	return serverHTTP, nil
}