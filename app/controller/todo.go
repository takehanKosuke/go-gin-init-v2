package controller

import (
	"go-gin-init-v2/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo Interface
type Todo interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Post(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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

// GetAll 初期デフォルトエンドポイント
func (d *TodoImpl) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
// GetById 初期デフォルトエンドポイント
func (d *TodoImpl) GetById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
// Post 初期デフォルトエンドポイント
func (d *TodoImpl) Post(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "pong",
	})
}
// Update 初期デフォルトエンドポイント
func (d *TodoImpl) Update(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "pong",
	})
}
// Delete 初期デフォルトエンドポイント
func (d *TodoImpl) Delete(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"message": "pong",
	})
}
