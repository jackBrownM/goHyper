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

// HelloHay 服务层只做判断与逻辑处理
func (h *helloService) HelloHay(hello string) string {
	return model.HelloModel().HelloHay(hello)
}
