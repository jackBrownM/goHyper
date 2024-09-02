//go:build wireinject
// +build wireinject

//go:generate go run github.com/google/wire/cmd/wire

package wire

import (
	"github.com/google/wire"
	"goHyper/core/svc"
	"goHyper/internal/api"
	"goHyper/internal/controller"
	"goHyper/internal/logic"
	"goHyper/internal/model"
)

func InitializeSvc() (*svc.Init, error) {
	wire.Build(
		svc.NewInit,
		svc.ProvideSet,
		model.ProvideSet,
		logic.ProvideSet,
		controller.ProvideSet,
		api.ProvideSet,
	)
	return &svc.Init{}, nil
}
