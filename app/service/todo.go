package service

import (
	"go-gin-init-v2/app/model"
	"go-gin-init-v2/app/repository"
)

type Todo interface {
	GetAll() ([]*model.Todo, error)
	GetById(id int) (*model.Todo, error)
	Post(todo *model.Todo) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
	Delete(id int) error
}

type TodoImpl struct {
	repository repository.Todo
}

func NewTodo(repository repository.Todo) Todo {
	return &TodoImpl{repository: repository}
}

// GetAll
func (d *TodoImpl) GetAll() ([]*model.Todo, error) {
	todos, err := d.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// GetById
func (d *TodoImpl) GetById(id int) (*model.Todo, error) {
	todo, err := d.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// Post
func (d *TodoImpl) Post(todo *model.Todo) (*model.Todo, error) {
	newTodo, err := d.repository.Post(todo)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}

// Update
func (d *TodoImpl) Update(todo *model.Todo) (*model.Todo, error) {
	updateTodo, err := d.repository.Update(todo)
	if err != nil {
		return nil, err
	}
	return updateTodo, nil
}

// Delete
func (d *TodoImpl) Delete(id int) error {
	err := d.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
