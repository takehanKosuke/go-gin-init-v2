//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package repository

import (
	"go-gin-init-v2/app/model"

	"gorm.io/gorm"
)

// Todo Interface
type Todo interface {
	GetAll() ([]*model.Todo, error)
	GetById(id int) (*model.Todo, error)
	Post(todo *model.Todo) (*model.Todo, error)
	Update(todo *model.Todo) (*model.Todo, error)
	Delete(id int) error
}


type TodoImpl struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) Todo {
	return &TodoImpl{db: db}
}

// GetAll
func (i *TodoImpl) GetAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	i.db.Limit(20).Find(&todos)
	return todos, nil
}

// GetById
func (i *TodoImpl) GetById(id int) (*model.Todo, error) {
	var todo *model.Todo
	i.db.First(&todo, id)
	return todo, nil
}

// Post
func (i *TodoImpl) Post(todo *model.Todo) (*model.Todo, error) {
	i.db.Create(todo)
	return todo, nil
}

// Update
func (i *TodoImpl) Update(todo *model.Todo) (*model.Todo, error) {
	i.db.Updates(todo)
	return todo, nil
}

// Delete
func (i *TodoImpl) Delete(id int) error {
	var todo *model.Todo
	i.db.Delete(todo, id)
	return nil
}
