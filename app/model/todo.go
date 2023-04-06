package model

type Status int

const (
	Yet Status = iota
	Doing
	Done
)

type Todo struct {
	id int
	status Status
	text string
}

func NewTodo (id int, status Status, text string) *Todo {
	return &Todo{
		id: id,
		status: status,
		text: text,
	}
}

func (m *Todo) Status() Status {
	return m.status;
}

func (m *Todo) Text() string {
	return m.text;
}

func (m *Todo) ChangeStatus(status Status)  {
	m.status = status;
}

