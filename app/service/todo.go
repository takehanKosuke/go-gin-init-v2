package service

import infrastructure "go-gin-init-v2/app/infrastructure/db"

type Todo interface {
	Ping()
}

type TodoImpl struct {
	infrastructure infrastructure.Todo
}

func NewTodo(infrastructure infrastructure.Todo) Todo {
	return &TodoImpl{infrastructure: infrastructure}
}

func (d *TodoImpl) Ping() {}
