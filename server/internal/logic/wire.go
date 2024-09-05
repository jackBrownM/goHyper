package logic

import "github.com/google/wire"

var ProvideSet = wire.NewSet(
	NewSystem)
