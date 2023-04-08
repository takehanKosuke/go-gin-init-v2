package model

import "fmt"

type Status int

const (
	Yet Status = iota + 1
	Doing
	Done
)
type Todo interface {
	Status() Status
	Text() string
	SetStatus(status Status) error
	SetText(text string) error
}

type Id int

type todoImpl struct {
	id Id
	status Status
	text string
}

func NewTodo (id int, status Status, text string) (Todo, error) {
	todo := &todoImpl{id: Id(id), status: status, text: ""}
	if err := todo.SetText(text); err != nil {
		return nil, err
	}

	return todo, nil
}

func (m *todoImpl) Status() Status {
	return m.status;
}

func (m *todoImpl) Text() string {
	return m.text;
}

func (m *todoImpl) SetStatus(status Status) error {
	m.status = status;
	return nil
}

func (m *todoImpl) SetText(text string) error {
	if len(text) < 100 {
		return fmt.Errorf("todo text is lather then 100, this text is %d", len(text))
	}
	m.text = text;
	return nil
}
