package controller

import (
	"github.com/google/wire"
	ctr_admin "goHyper/internal/controller/admin"
)

var ProvideSet = wire.NewSet(
	ctr_admin.NewSystem,
	ctr_admin.NewRole,
	ctr_admin.NewMenu,
)
