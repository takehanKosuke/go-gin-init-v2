package service

import "go-gin-init-v2/app/repository"

type Todo interface {
	Ping()
}

type TodoImpl struct {
	repository repository.Todo
}

func NewTodo(repository repository.Todo) Todo {
	return &TodoImpl{repository: repository}
}

func (d *TodoImpl) Ping() {}
