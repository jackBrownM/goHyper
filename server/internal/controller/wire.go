package controller

import (
	"github.com/google/wire"
	"goHyper/internal/controller/admin"
	"goHyper/internal/controller/example"
)

var ProvideSet = wire.NewSet(
	example.NewExample,
	admin.NewAdmin,
)
