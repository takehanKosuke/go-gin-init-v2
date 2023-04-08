//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package infrastructure

import (
	"go-gin-init-v2/app/model"

	"gorm.io/gorm"
)

// Todo Interface
type Todo interface {
	GetAll() []*model.Todo
	GetById(id int) *model.Todo
	Post(todo model.Todo) *model.Todo
	Update(todo model.Todo) *model.Todo
	Delete()
}


type TodoImpl struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) Todo {
	return &TodoImpl{db: db}
}

// GetAll
func (d *TodoImpl) GetAll() []*model.Todo{
	var todos []*model.Todo
	return todos
}

// GetById
func (d *TodoImpl) GetById(id int) *model.Todo {
	var todo *model.Todo
	return todo
}

// Post
func (d *TodoImpl) Post(todo model.Todo) *model.Todo{
	return &todo
}

// Update
func (d *TodoImpl) Update(todo model.Todo) *model.Todo{
	return &todo
}

// Delete
func (d *TodoImpl) Delete() {}
