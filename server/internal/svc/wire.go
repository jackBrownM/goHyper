package svc

import (
	"github.com/google/wire"
	"goHyper/core/router"
	base2 "goHyper/internal/svc/base"
)

var ProvideSet = wire.NewSet(
	base2.NewConfig,
	base2.NewLogger,
	NewHttpServ,
	base2.NewRedis,
	base2.NewMysql,
	router.NewRoute,
)
