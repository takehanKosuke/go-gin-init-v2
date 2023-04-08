//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package infrastructure

import (
	"go-gin-init-v2/app/model"

	"gorm.io/gorm"
)

// Todo Interface
type Todo interface {
	GetAll() ([]*model.Todo, error)
	GetById(id int) (*model.Todo, error)
	Post(todo model.Todo) (*model.Todo, error)
	Update(todo model.Todo) (*model.Todo, error)
	Delete(id int) error
}


type TodoImpl struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) Todo {
	return &TodoImpl{db: db}
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
