package handler

import (
	"github.com/gin-gonic/gin"
)

// Template Interface
type Template interface {
	Ping(c *gin.Context)
}

// TemplateImpl TemplateImpl struct
type TemplateImpl struct {
	service service.Template
}

// NewTemplate 新規Template structを作成
func NewTemplate(service service.Template) Template {
	return &TemplateImpl{
		service: service,
	}
}

// Ping 初期デフォルトエンドポイント
func (d *TemplateImpl) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
