package model

var helloModelInstance *helloModel

type helloModel struct{}

func init() {
	helloModelInstance = &helloModel{}
}

func HelloModel() *helloModel {
	return helloModelInstance
}

// HelloHay 模型层只做数据的封装
func (h *helloModel) HelloHay(hello string) string {
	return hello
}
