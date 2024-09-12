package main

import (
	"github.com/pkg/errors"
	"goHyper/svc/wire"
	"log"
)

// 入口文件
func main() {
	svcSvc, err := wire.InitializeSvc() // 初始化
	defer svcSvc.Stop()
	if err != nil {
		log.Fatal(errors.Wrap(err, "initialize failed"))
	}
	svcSvc.Start()
}
