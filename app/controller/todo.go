package controller

import (
	"go-gin-init-v2/app/model"
	"go-gin-init-v2/app/service"
	"net/http"
	"strconv"

	openapi "go-gin-init-v2/out/go/go"

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

// GetAll
func (d *TodoImpl) GetAll(c *gin.Context) {
	todos, err := d.service.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"type":    500,
			"message": "internal server error",
		})
	}
	c.JSON(http.StatusOK, &openapi.V1GetTodoAll200Response{
		List: convertOasTodoList(todos),
	})
}

// GetById
func (d *TodoImpl) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := d.service.GetById(id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"type":    500,
			"message": "internal server error",
		})
	}
	c.JSON(http.StatusOK, &openapi.V1GetTodoById200Response{convertOasTodo(*todo)})
}

// Post
func (d *TodoImpl) Post(c *gin.Context) {

	c.Status(http.StatusCreated)
}

// Update
func (d *TodoImpl) Update(c *gin.Context) {
	c.Status(http.StatusCreated)
}

// Delete
func (d *TodoImpl) Delete(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func convertOasTodo(todo model.Todo) openapi.Todo {
	return openapi.Todo{
		Id: int32(todo.Id()),
		Status: int32(todo.Status()),
		Text: todo.Text(),
	}
}

func convertOasTodoList(todos []*model.Todo) []openapi.Todo {
	var todoList []openapi.Todo
	for _, v := range todos {
		todoList = append(todoList, convertOasTodo(*v))
	}
	return todoList
}
