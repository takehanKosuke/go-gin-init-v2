package service

import "go-gin-init-v2/app/repository"

type Template interface {
	Ping()
}

type TemplateImpl struct {
	repository repository.Template
}

func NewTemplate(repository repository.Template) Template {
	return &TemplateImpl{repository: repository}
}

func (d *TemplateImpl) Ping() {}
