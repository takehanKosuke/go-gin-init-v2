// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"go-gin-init-v2/app/config"
	"go-gin-init-v2/app/controller"
	"go-gin-init-v2/app/infrastructure/db"
	"go-gin-init-v2/app/router"
	"go-gin-init-v2/app/service"
)

// Injectors from wire.go:

func InitializeApplication() (APIApplication, error) {
	configConfig, err := config.Load()
	if err != nil {
		return APIApplication{}, err
	}
	db := config.ConnectDB(configConfig)
	todo := infrastructure.NewTodo(db)
	serviceTodo := service.NewTodo(todo)
	controllerTodo := controller.NewTodo(serviceTodo)
	engine := router.SetupRouter(controllerTodo)
	apiApplication := NewAPIApplication(configConfig, engine, db)
	return apiApplication, nil
}
