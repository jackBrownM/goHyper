package api

import (
	"github.com/google/wire"
	route_admin "goHyper/internal/api/admin"
)

var ProvideSet = wire.NewSet(
	route_admin.NewAdmin,
)
