//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package repository

import (
	"gorm.io/gorm"
)

type Todo interface {
	Ping()
}

type TodoImpl struct {
	db *gorm.DB
}

func NewTodo(db *gorm.DB) Todo {
	return &TodoImpl{db: db}
}

func (d *TodoImpl) Ping() {}
