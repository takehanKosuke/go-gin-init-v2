//+build wireinject

package main

import (
	"go-gin-init-v2/app/router"
	"go-gin-init-v2/app/config"
	"go-gin-init-v2/app/controller"
	"go-gin-init-v2/app/service"
	"go-gin-init-v2/app/infrastructure/db"

	"github.com/google/wire"
)

func InitializeApplication() (APIApplication, error) {
	wire.Build(
		NewAPIApplication,
		config.Load,
		router.SetupRouter,
		config.ConnectDB,

		// controller
		controller.NewTodo,

		// service
		service.NewTodo,

		// infrastructure
		infrastructure.NewTodo,
	)
	return APIApplication{}, nil
}
