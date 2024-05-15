package model

var helloModelInstance *helloModel

type helloModel struct{}

func init() {
	helloModelInstance = &helloModel{}
}

func HelloModel() *helloModel {
	return helloModelInstance
}

func (h *helloModel) HelloHay(hello string) string {
	return hello
}
