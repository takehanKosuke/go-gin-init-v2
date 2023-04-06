package handler

import (
	"go-gin-init-v2/app/service"

	"github.com/gin-gonic/gin"
)

// Todo Interface
type Todo interface {
	Ping(c *gin.Context)
}

// TodoImpl TodoImpl struct
type TodoImpl struct {
	service service.Todo
}

// NewTodo 新規Todo structを作成
func NewTodo(service service.Todo) Todo {
	return &TodoImpl{
		service: service,
	}
}

// Ping 初期デフォルトエンドポイント
func (d *TodoImpl) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
