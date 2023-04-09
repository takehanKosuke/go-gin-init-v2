package model

import "fmt"

type Status int

const (
	Yet Status = iota + 1
	Doing
	Done
)
type Todo interface {
	Id() int
	Status() Status
	Text() string
	SetStatus(status Status) error
	SetText(text string) error
}

type todoImpl struct {
	id int
	status Status
	text string
}

func NewTodo (status Status, text string) (Todo, error) {
	todo := &todoImpl{status: status, text: ""}
	if err := todo.SetText(text); err != nil {
		return nil, err
	}

	return todo, nil
}

func (m *todoImpl) Id() int {
	return m.id;
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
