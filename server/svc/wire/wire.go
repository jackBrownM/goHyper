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
	svc2 "goHyper/svc"
)

func InitializeSvc() (*svc2.Init, error) {
	wire.Build(
		svc2.NewInit,
		svc2.ProvideSet,
		dao.ProvideSet,
		logic.ProvideSet,
		controller.ProvideSet,
		api.ProvideSet,
	)
	return &svc2.Init{}, nil
}
