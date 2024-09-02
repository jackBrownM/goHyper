package logic

import "goHyper/internal/dao"

type Example struct {
	example *dao.Example
}

func NewExample(example *dao.Example) *Example {
	return &Example{
		example: example,
	}
}

func (l *Example) Example() string {
	return l.example.Example()
}
