//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE
package repository

import (
	"gorm.io/gorm"
)

type Template interface {
	Ping()
}

type TemplateImpl struct {
	db *gorm.DB
}

func NewTemplate(db *gorm.DB) Template {
	return &TemplateImpl{db: db}
}

func (d *TemplateImpl) Ping() {}
