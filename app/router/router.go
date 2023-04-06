package config

import (
	"go-gin-init-v2/app/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	// すべてのmiddlewareを記述する

	// すべてのhandlerを記述する
	defaultHandler handler.Default,
) *gin.Engine {
	r := gin.Default()

	// middleware周りの実装
	// authMiddleware := middleware.NewFirebaseAuth(newMockFirebaseAuthClient())
	// authMiddleware := middleware.NewFirebaseAuth(firebaseAuthClient())
	// r.Use(authMiddleware.Authentication)

	// NewHandler
	defo := defaultHandler

	// Routing
	r.GET("/", defo.Ping)

	return r
}
