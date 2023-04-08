package router

import (
	"go-gin-init-v2/app/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	// すべてのmiddlewareを記述する

	// すべてのhandlerを記述する
	todoController controller.Todo,
) *gin.Engine {
	r := gin.Default()

	// middleware周りの実装
	// authMiddleware := middleware.NewFirebaseAuth(newMockFirebaseAuthClient())
	// authMiddleware := middleware.NewFirebaseAuth(firebaseAuthClient())
	// r.Use(authMiddleware.Authentication)

	// NewHandler
	todo := todoController

	// Routing
	r.GET("/todos", todo.GetAll)
	r.POST("/todos", todo.Post)
	r.GET("/todos/:id", todo.GetById)
	r.DELETE("/todos/:id", todo.Delete)
	r.PUT("/todos/:id", todo.Update)

	return r
}
