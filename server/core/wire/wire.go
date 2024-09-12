//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire

package wire

import (
	"github.com/google/wire"
	"goHyper/internal/api"
	"goHyper/internal/controller"
	"goHyper/internal/dao"
	"goHyper/internal/logic"
	"goHyper/internal/svc"
)

func InitializeSvc() (*svc.Init, error) {
	wire.Build(
		svc.NewInit,
		svc.ProvideSet,
		dao.ProvideSet,
		logic.ProvideSet,
		controller.ProvideSet,
		api.ProvideSet,
	)
	return &svc.Init{}, nil
}
