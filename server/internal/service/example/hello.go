package example

import (
	"goHyper/internal/model"
)

var helloServiceInstance *helloService

type helloService struct{}

func init() {
	helloServiceInstance = &helloService{}
}

func HelloService() *helloService {
	return helloServiceInstance
}

func (h *helloService) HelloHay(hello string) string {
	return model.HelloModel().HelloHay(hello)
}
