//+build wireinject

package main

import (
	"go-gin-init-v2/app/config"
	"go-gin-init-v2/app/handler"
	"go-gin-init-v2/app/repository"
	"go-gin-init-v2/app/service"

	"github.com/google/wire"
)

func InitializeApplication() (APIApplication, error) {
	wire.Build(
		NewAPIApplication,
		config.Load,
		config.SetupRouter,
		config.ConnectDB,

		// handler
		handler.NewTodo,

		// service
		service.NewTodo,

		// repository
		repository.NewTodo,
	)
	return APIApplication{}, nil
}
