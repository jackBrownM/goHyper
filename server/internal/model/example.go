package model

type Example struct{}

func NewExample() *Example {
	return &Example{}
}

func (m *Example) Example() string {
	return "hello world"
}
