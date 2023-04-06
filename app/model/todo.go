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


type TodoImpl struct {
	id *int
	status Status
	text string
}

func NewTodo (id int, status *Status, text string) *TodoImpl {
	if status == nil {
		status = Yet
	}
	todo := &TodoImpl{id: id, status: *status, text: ""}
	todo.SetText(text)
	return todo
}

func (m *TodoImpl) Status() Status {
	return m.status;
}

func (m *TodoImpl) Text() string {
	return m.text;
}

func (m *TodoImpl) SetStatus(status Status) error {
	m.status = status;
	return nil
}

func (m *TodoImpl) SetText(text string) error {
	if len(text) < 100 {
		return fmt.Errorf("todo text is lather then 100, this text is %d", len(text))
	}
	m.text = text;
	return nil
}
