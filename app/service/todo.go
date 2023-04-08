package service

import (
	infrastructure "go-gin-init-v2/app/infrastructure/db"
	"go-gin-init-v2/app/model"
)

type Todo interface {
	GetAll() ([]*model.Todo, error)
	GetById(id int) (*model.Todo, error)
	Post(todo model.Todo) (*model.Todo, error)
	Update(todo model.Todo) (*model.Todo, error)
	Delete(id int) error
}

type TodoImpl struct {
	infrastructure infrastructure.Todo
}

func NewTodo(infrastructure infrastructure.Todo) Todo {
	return &TodoImpl{infrastructure: infrastructure}
}

// GetAll
func (d *TodoImpl) GetAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	return todos, nil
}

// GetById
func (d *TodoImpl) GetById(id int) (*model.Todo, error) {
	var todo *model.Todo
	return todo, nil
}

// Post
func (d *TodoImpl) Post(todo model.Todo) (*model.Todo, error) {
	return &todo, nil
}

// Update
func (d *TodoImpl) Update(todo model.Todo) (*model.Todo, error) {
	return &todo, nil
}

// Delete
func (d *TodoImpl) Delete(id int) error {
	return nil
}
