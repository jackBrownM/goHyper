package example

import (
	"goHyper/internal/model"
)

type helloService struct{}

func HelloService() *helloService {
	return &helloService{}
}

// HelloHay 服务层只做判断与逻辑处理
func (h *helloService) HelloHay(hello string) string {
	return model.HelloModel().HelloHay(hello)
}
