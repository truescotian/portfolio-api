package models

type Model struct {
	ID   int
	Name string
}

func NewModel(id int, name string) (m *Model) {
	return &Model{
		ID:   id,
		Name: name,
	}
}
