package config

import (
	"go-gin-init-v2/app/handler"
	"go-gin-init-v2/app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	// すべてのmiddlewareを記述する

	// すべてのhandlerを記述する
	defaultHandler handler.Default,
) *gin.Engine {
	r := gin.Default()

	authMiddleware := middleware.NewFirebaseAuth(newMockFirebaseAuthClient())
	// authMiddleware := middleware.NewFirebaseAuth(firebaseAuthClient())
	r.Use(authMiddleware.Authentication)

	// NewHandler
	defo := defaultHandler

	// Routing
	r.GET("/", defo.Ping)

	return r
}
