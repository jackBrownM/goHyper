package svc

import (
	"github.com/google/wire"
	"goHyper/core/router"
	"goHyper/svc/base"
)

var ProvideSet = wire.NewSet(
	base.NewConfig,
	base.NewLogger,
	NewHttpServ,
	base.NewRedis,
	base.NewMysql,
	router.NewRoute,
)
