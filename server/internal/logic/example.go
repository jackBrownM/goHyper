package logic

import "goHyper/internal/model"

type Example struct {
	example *model.Example
}

func NewExample(example *model.Example) *Example {
	return &Example{
		example: example,
	}
}

func (l *Example) Example() string {
	return l.example.Example()
}
